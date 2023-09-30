package parser

import (
	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/plugin/parser/base"
	mysqlparser "github.com/bytebase/bytebase/backend/plugin/parser/mysql"
	pgparser "github.com/bytebase/bytebase/backend/plugin/parser/pg"
	plsqlparser "github.com/bytebase/bytebase/backend/plugin/parser/plsql"
	snowparser "github.com/bytebase/bytebase/backend/plugin/parser/snowflake"
	standardparser "github.com/bytebase/bytebase/backend/plugin/parser/standard"
	tidbparser "github.com/bytebase/bytebase/backend/plugin/parser/tidb"
	tsqlparser "github.com/bytebase/bytebase/backend/plugin/parser/tsql"
	storepb "github.com/bytebase/bytebase/proto/generated-go/store"
)

// SplitMultiSQL splits statement into a slice of the single SQL.
func SplitMultiSQL(engineType storepb.Engine, statement string) ([]base.SingleSQL, error) {
	switch engineType {
	case storepb.Engine_MYSQL, storepb.Engine_MARIADB, storepb.Engine_OCEANBASE:
		return mysqlparser.SplitMySQL(statement)
	case storepb.Engine_TIDB:
		return tidbparser.SplitSQL(statement)
	case storepb.Engine_POSTGRES, storepb.Engine_REDSHIFT, storepb.Engine_RISINGWAVE:
		return pgparser.SplitSQL(statement)
	case storepb.Engine_ORACLE, storepb.Engine_DM:
		return plsqlparser.SplitPLSQL(statement)
	case storepb.Engine_MSSQL:
		return tsqlparser.SplitSQL(statement)
	default:
		return standardparser.SplitSQL(statement)
	}
}

// ExtractDatabaseList extracts all databases from statement.
func ExtractDatabaseList(engineType storepb.Engine, statement string, currentDatabase string) ([]string, error) {
	switch engineType {
	case storepb.Engine_MYSQL, storepb.Engine_MARIADB, storepb.Engine_OCEANBASE:
		// TODO(d): use mysql parser.
		return tidbparser.ExtractDatabaseList(statement, currentDatabase)
	case storepb.Engine_TIDB:
		return tidbparser.ExtractDatabaseList(statement, currentDatabase)
	case storepb.Engine_SNOWFLAKE:
		return snowparser.ExtractDatabaseList(statement, currentDatabase)
	case storepb.Engine_MSSQL:
		return tsqlparser.ExtractDatabaseList(statement, currentDatabase)
	default:
		return nil, errors.Errorf("engine type is not supported: %s", engineType)
	}
}
