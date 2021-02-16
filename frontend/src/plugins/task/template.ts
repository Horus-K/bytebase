import { TaskTemplate, TemplateContext } from "../types";
import { TaskNew } from "../../types";

export const taskTemplateList: TaskTemplate[] = [
  {
    type: "bytebase.general",
    buildTask: (ctx: TemplateContext): TaskNew => {
      return {
        type: "task",
        attributes: {
          name: "New General Task",
          type: "bytebase.general",
          content: "",
          stageProgressList: [
            {
              id: "1",
              name: "Request",
              type: "SIMPLE",
              status: "PENDING",
            },
          ],
          creator: {
            id: ctx.currentUser.id,
            name: ctx.currentUser.attributes.name,
          },
          payload: {},
        },
      };
    },
  },
  {
    type: "bytebase.datasource.create",
    buildTask: (ctx: TemplateContext): TaskNew => {
      return {
        type: "task",
        attributes: {
          name: "New Data Source",
          type: "bytebase.datasource.create",
          content: "Estimated QPS: 10",
          stageProgressList: [
            {
              id: "1",
              name: "Request Data Source",
              type: "ENVIRONMENT",
              status: "PENDING",
            },
          ],
          creator: {
            id: ctx.currentUser.id,
            name: ctx.currentUser.attributes.name,
          },
          payload: {},
        },
      };
    },
    outputFieldList: [
      {
        id: 1,
        name: "Data Source URL1",
        required: true,
      },
      {
        id: 2,
        name: "Hello world",
        required: true,
      },
      {
        id: 3,
        name: "Data Source URL3",
        required: true,
      },
    ],
    fieldList: [
      {
        id: 1,
        slug: "db",
        name: "Database Name",
        type: "String",
        required: true,
        preprocessor: (name: string): string => {
          // In case caller passes corrupted data.
          // Handled here instead of the caller, because it's
          // preprocessor specific behavior to handle fallback.
          return name?.toLowerCase();
        },
      },
    ],
  },
];
