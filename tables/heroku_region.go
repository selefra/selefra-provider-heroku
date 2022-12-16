package tables

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/selefra/selefra-provider-heroku/heroku_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableHerokuRegionGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableHerokuRegionGenerator{}

func (x *TableHerokuRegionGenerator) GetTableName() string {
	return "heroku_region"
}

func (x *TableHerokuRegionGenerator) GetTableDescription() string {
	return ""
}

func (x *TableHerokuRegionGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableHerokuRegionGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableHerokuRegionGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := heroku_client.Connect(ctx, taskClient.(*heroku_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			opts := heroku.ListRange{Field: "id", Max: 1000}

			items, err := conn.RegionList(ctx, &opts)
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

func (x *TableHerokuRegionGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableHerokuRegionGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("locale").ColumnType(schema.ColumnTypeString).Description("Area in the country where the region exists.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("When region was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Unique name of region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Description("Description of region.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("private_capable").ColumnType(schema.ColumnTypeBool).Description("Whether or not region is available for creating a Private Space.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provider").ColumnType(schema.ColumnTypeJSON).Description("Provider of underlying substrate.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("country").ColumnType(schema.ColumnTypeString).Description("Country where the region exists.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When region was created.").Build(),
	}
}

func (x *TableHerokuRegionGenerator) GetSubTables() []*schema.Table {
	return nil
}
