package table_schema_generator_tables

import (
	"context"
	"github.com/selefra/selefra-provider-heroku/heroku_client"

	heroku "github.com/heroku/heroku-go/v5"
	"github.com/pkg/errors"
	"github.com/selefra/selefra-provider-heroku/table_schema_generator"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
)

type TableHerokuTeamAppPermissionsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableHerokuTeamAppPermissionsGenerator{}

func (x *TableHerokuTeamAppPermissionsGenerator) GetTableName() string {
	return "heroku_team_app_permissions"
}

func (x *TableHerokuTeamAppPermissionsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableHerokuTeamAppPermissionsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableHerokuTeamAppPermissionsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableHerokuTeamAppPermissionsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c, _ := heroku_client.Connect(ctx, client.(*heroku_client.Client).Config)
			nextRange := &heroku.ListRange{
				Field: "id",
				Max:   1000,
			}

			for nextRange.Max != 0 {
				ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
				v, err := c.TeamAppPermissionList(ctxWithRange, nextRange)
				if err != nil {
					maybeError := errors.WithStack(err)
					return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
				}
				resultChannel <- v
			}
			return nil
		},
	}
}

func (x *TableHerokuTeamAppPermissionsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableHerokuTeamAppPermissionsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableHerokuTeamAppPermissionsGenerator) GetSubTables() []*schema.Table {
	return nil
}
