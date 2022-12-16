package tables

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/selefra/selefra-provider-heroku/heroku_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableHerokuAddonGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableHerokuAddonGenerator{}

func (x *TableHerokuAddonGenerator) GetTableName() string {
	return "heroku_addon"
}

func (x *TableHerokuAddonGenerator) GetTableDescription() string {
	return ""
}

func (x *TableHerokuAddonGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableHerokuAddonGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableHerokuAddonGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := heroku_client.Connect(ctx, taskClient.(*heroku_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			opts := heroku.ListRange{Field: "id", Max: 1000}

			items, err := conn.AddOnList(ctx, &opts)
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

func (x *TableHerokuAddonGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableHerokuAddonGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("addon_service").ColumnType(schema.ColumnTypeJSON).Description("Identity of add-on service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("state").ColumnType(schema.ColumnTypeString).Description("State in the add-on's lifecycle.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("plan").ColumnType(schema.ColumnTypeJSON).Description("Identity of add-on plan.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("actions").ColumnType(schema.ColumnTypeJSON).Description("Provider actions for this specific add-on.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When add-on was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("When add-on was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_url").ColumnType(schema.ColumnTypeString).Description("URL for logging into web interface of add-on (e.g. a dashboard).").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Globally unique name of the add-on.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of add-on.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provider_id").ColumnType(schema.ColumnTypeString).Description("Id of this add-on with its provider.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("app").ColumnType(schema.ColumnTypeJSON).Description("Billing application associated with this add-on.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billed_price").ColumnType(schema.ColumnTypeJSON).Description("Billed price.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("billing_entity").ColumnType(schema.ColumnTypeJSON).Description("Billing entity associated with this add-on.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("config_vars").ColumnType(schema.ColumnTypeJSON).Description("Config vars exposed to the owning app by this add-on.").Build(),
	}
}

func (x *TableHerokuAddonGenerator) GetSubTables() []*schema.Table {
	return nil
}
