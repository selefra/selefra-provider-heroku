package tables

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/selefra/selefra-provider-heroku/heroku_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableHerokuAppGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableHerokuAppGenerator{}

func (x *TableHerokuAppGenerator) GetTableName() string {
	return "heroku_app"
}

func (x *TableHerokuAppGenerator) GetTableDescription() string {
	return ""
}

func (x *TableHerokuAppGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableHerokuAppGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableHerokuAppGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := heroku_client.Connect(ctx, taskClient.(*heroku_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			opts := heroku.ListRange{Field: "id", Max: 1000}

			items, err := conn.AppList(ctx, &opts)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			for _, i := range items {
				resultChannel <- i

				if heroku_client.IsCancelled(ctx) {
					return schema.NewDiagnosticsErrorPullTable(task.Table, nil)
				}
			}
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableHerokuAppGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableHerokuAppGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("git_url").ColumnType(schema.ColumnTypeString).Description("Git repo URL of app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("owner").ColumnType(schema.ColumnTypeJSON).Description("Identity of app owner.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("slug_size").ColumnType(schema.ColumnTypeInt).Description("Slug size in bytes of app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("stack").ColumnType(schema.ColumnTypeJSON).Description("Identity of app stack.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Unique name of app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("buildpack_provided_description").ColumnType(schema.ColumnTypeString).Description("Description from buildpack of app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("maintenance").ColumnType(schema.ColumnTypeBool).Description("Maintenance status of app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("released_at").ColumnType(schema.ColumnTypeTimestamp).Description("When app was released.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("When app was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("archived_at").ColumnType(schema.ColumnTypeTimestamp).Description("When app was archived.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When app was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("organization").ColumnType(schema.ColumnTypeJSON).Description("Identity of team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_url").ColumnType(schema.ColumnTypeString).Description("Web URL of app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("acm").ColumnType(schema.ColumnTypeString).Description("ACM status of this app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("build_stack").ColumnType(schema.ColumnTypeJSON).Description("Identity of the stack that will be used for new builds.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("repo_size").ColumnType(schema.ColumnTypeInt).Description("Git repo size in bytes of app.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("space").ColumnType(schema.ColumnTypeJSON).Description("Identity of space.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("team").ColumnType(schema.ColumnTypeJSON).Description("identity of team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("internal_routing").ColumnType(schema.ColumnTypeBool).Description("Describes whether a Private Spaces app is externally routable or not.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("region").ColumnType(schema.ColumnTypeJSON).Description("Identity of app region.").Build(),
	}
}

func (x *TableHerokuAppGenerator) GetSubTables() []*schema.Table {
	return nil
}
