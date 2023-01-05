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

type TableHerokuAppsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableHerokuAppsGenerator{}

func (x *TableHerokuAppsGenerator) GetTableName() string {
	return "heroku_apps"
}

func (x *TableHerokuAppsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableHerokuAppsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableHerokuAppsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableHerokuAppsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c, _ := heroku_client.Connect(ctx, client.(*heroku_client.Client).Config)
			nextRange := &heroku.ListRange{
				Field: "id",
				Max:   1000,
			}

			for nextRange.Max != 0 {
				ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
				v, err := c.AppList(ctxWithRange, nextRange)
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

func (x *TableHerokuAppsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableHerokuAppsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internal_routing").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("released_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("team").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("git_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("GitURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("archived_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repo_size").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stack").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WebURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("acm").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("build_stack").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("buildpack_provided_description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance").ColumnType(schema.ColumnTypeBool).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("slug_size").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("space").ColumnType(schema.ColumnTypeJSON).Build(),
	}
}

func (x *TableHerokuAppsGenerator) GetSubTables() []*schema.Table {
	return nil
}
