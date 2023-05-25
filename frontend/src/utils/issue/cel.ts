import { cloneDeep, last } from "lodash-es";
import { getDatabaseIdByName, getDatabaseNameById } from "./expr";
import { celServiceClient } from "@/grpcweb";
import { SimpleExpr, resolveCELExpr } from "@/plugins/cel";
import { useDatabaseStore } from "@/store";
import { DatabaseResource } from "@/components/Issue/form/SelectDatabaseResourceForm/common";

interface DatabaseLevelCondition {
  database: string[];
}

interface SchemaLevelCondition {
  database: string;
  schema: string[];
}

interface TableLevelCondition {
  database: string;
  schema: string;
  table: string[];
}

type DatabaseResourceCondition =
  | DatabaseLevelCondition
  | SchemaLevelCondition
  | TableLevelCondition;

export const stringifyDatabaseResources = (resources: DatabaseResource[]) => {
  const conditionList: DatabaseResourceCondition[] = [];

  for (const resource of resources) {
    const database = useDatabaseStore().getDatabaseById(resource.databaseId);
    const databaseName = getDatabaseNameById(database.id);
    if (resource.table === undefined && resource.schema === undefined) {
      // Database level
      conditionList.push({
        database: [databaseName],
      });
    } else if (resource.schema !== undefined && resource.table === undefined) {
      // Schema level
      conditionList.push({
        database: databaseName,
        schema: [resource.schema],
      });
    } else if (resource.schema !== undefined && resource.table !== undefined) {
      // Table level
      conditionList.push({
        database: databaseName,
        schema: resource.schema,
        table: [resource.table],
      });
    } else {
      throw new Error("Invalid database resource");
    }
  }

  const databaseLevelConditionList = mergeDatabaseLevelConditions(
    conditionList.filter((condition): condition is DatabaseLevelCondition =>
      Array.isArray(condition.database)
    )
  ).filter((condition) => condition.database.length > 0);
  const schemaLevelConditionList = mergeSchemaLevelConditions(
    conditionList.filter(
      (condition): condition is SchemaLevelCondition =>
        typeof condition.database === "string" &&
        Array.isArray((condition as SchemaLevelCondition).schema)
    )
  ).filter((condition) => condition.schema.length > 0);
  const tableLevelConditionList = mergeTableLevelConditions(
    conditionList.filter(
      (condition): condition is TableLevelCondition =>
        typeof condition.database === "string" &&
        typeof (condition as TableLevelCondition).schema === "string" &&
        Array.isArray((condition as TableLevelCondition).table)
    )
  ).filter((condition) => condition.table.length > 0);

  const cel = convertToCEL([
    ...databaseLevelConditionList,
    ...schemaLevelConditionList,
    ...tableLevelConditionList,
  ]);
  return cel;
};

const convertToCEL = (
  conditions: (
    | DatabaseLevelCondition
    | SchemaLevelCondition
    | TableLevelCondition
  )[]
): string => {
  function buildCondition(
    condition:
      | DatabaseLevelCondition
      | SchemaLevelCondition
      | TableLevelCondition
  ): string {
    if ("table" in condition) {
      return `resource.database == "${
        condition.database
      }" && resource.schema == "${
        condition.schema
      }" && resource.table in ${JSON.stringify(condition.table)}`;
    } else if ("schema" in condition) {
      return `resource.database == "${
        condition.database
      }" && resource.schema in ${JSON.stringify(condition.schema)}`;
    } else {
      return `resource.database in ${JSON.stringify(condition.database)}`;
    }
  }

  function buildGroup(
    conditions: (
      | DatabaseLevelCondition
      | SchemaLevelCondition
      | TableLevelCondition
    )[]
  ): string {
    if (conditions.length === 1) {
      return buildCondition(conditions[0]);
    } else {
      const conditionStrings = conditions.map(buildCondition);
      return `${conditionStrings.map((s) => `(${s})`).join(" || ")}`;
    }
  }

  const topLevelCondition = buildGroup(conditions);

  return `(${topLevelCondition})`;
};

interface ConditionExpression {
  databases?: string[];
  databaseResources?: DatabaseResource[];
  expiredTime?: string;
  statement?: string;
  rowLimit?: number;
  exportFormat?: string;
}

export const converFromCEL = async (
  cel: string
): Promise<ConditionExpression> => {
  const { expression: celExpr } = await celServiceClient.parse({
    expression: cel,
  });

  if (!celExpr || !celExpr.expr) {
    return {};
  }

  const simpleExpr = resolveCELExpr(celExpr.expr);
  const conditionExpression: ConditionExpression = {
    databaseResources: [],
  };

  async function processCondition(expr: SimpleExpr) {
    if (expr.operator === "_&&_" || expr.operator === "_||_") {
      for (const arg of expr.args) {
        await processCondition(arg);
      }
    } else if (expr.operator === "@in") {
      const [property, values] = expr.args;
      if (typeof property === "string" && Array.isArray(values)) {
        if (property === "resource.database") {
          for (const value of values) {
            const databaseId = await getDatabaseIdByName(value as string);
            const databaseResource: DatabaseResource = {
              databaseId: databaseId,
            };
            conditionExpression.databaseResources!.push(databaseResource);
          }
        } else if (property === "resource.schema") {
          const databaseResource = conditionExpression.databaseResources?.pop();
          if (databaseResource) {
            for (const value of values) {
              const temp: DatabaseResource = cloneDeep(
                databaseResource
              ) as DatabaseResource;
              temp.schema = value as string;
              conditionExpression.databaseResources!.push(temp);
            }
          }
        } else if (property === "resource.table") {
          const databaseResource = conditionExpression.databaseResources?.pop();
          if (databaseResource) {
            for (const value of values) {
              const temp: DatabaseResource = cloneDeep(
                databaseResource
              ) as DatabaseResource;
              temp.table = value as string;
              conditionExpression.databaseResources!.push(temp);
            }
          }
        }
      }
    } else if (expr.operator === "_==_") {
      const [left, right] = expr.args;
      if (typeof left === "string") {
        if (typeof right === "string") {
          if (left === "resource.database") {
            const databaseId = await getDatabaseIdByName(right);
            const databaseResource: DatabaseResource = {
              databaseId: databaseId,
            };
            conditionExpression.databaseResources!.push(databaseResource);
          } else if (left === "resource.schema") {
            const databaseResource = last(
              conditionExpression.databaseResources
            );
            if (databaseResource) {
              databaseResource.schema = right;
            }
          } else if (left === "request.statement") {
            const statement = atob(right);
            conditionExpression.statement = statement;
          } else if (left === "request.export_format") {
            conditionExpression.exportFormat = right;
          }
        } else if (typeof right === "number") {
          if (left === "request.row_limit") {
            conditionExpression.rowLimit = right;
          }
        }
      }
    } else if (expr.operator === "_<_") {
      const [left, right] = expr.args;
      if (left === "request.time") {
        conditionExpression.expiredTime = (right as Date).toISOString();
      }
    }
  }
  await processCondition(simpleExpr);
  return conditionExpression;
};

const mergeDatabaseLevelConditions = (
  conditions: DatabaseLevelCondition[]
): DatabaseLevelCondition[] => {
  return [
    {
      database: conditions.map((condition) => condition.database).flat(),
    },
  ];
};

const mergeSchemaLevelConditions = (
  conditions: SchemaLevelCondition[]
): SchemaLevelCondition[] => {
  const groupedConditions: Record<string, string[]> = {};

  for (const condition of conditions) {
    const { database, schema } = condition;

    if (groupedConditions[database]) {
      groupedConditions[database] = [...groupedConditions[database], ...schema];
    } else {
      groupedConditions[database] = [...schema];
    }
  }

  const mergedConditions: SchemaLevelCondition[] = [];

  for (const database in groupedConditions) {
    const condition: SchemaLevelCondition = {
      database,
      schema: groupedConditions[database],
    };
    mergedConditions.push(condition);
  }

  return mergedConditions;
};

const mergeTableLevelConditions = (
  conditions: TableLevelCondition[]
): TableLevelCondition[] => {
  const groupedConditions: Record<string, string[]> = {};

  for (const condition of conditions) {
    const { database, schema, table } = condition;
    const key = `${database}:${schema}`;

    if (groupedConditions[key]) {
      groupedConditions[key] = [...groupedConditions[key], ...table];
    } else {
      groupedConditions[key] = [...table];
    }
  }

  const mergedConditions: TableLevelCondition[] = [];

  for (const key in groupedConditions) {
    const [database, schema] = key.split(":");
    const condition: TableLevelCondition = {
      database,
      schema,
      table: groupedConditions[key],
    };
    mergedConditions.push(condition);
  }

  return mergedConditions;
};
