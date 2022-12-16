package tables

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/selefra/selefra-provider-heroku/heroku_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableHerokuKeyGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableHerokuKeyGenerator{}

func (x *TableHerokuKeyGenerator) GetTableName() string {
	return "heroku_key"
}

func (x *TableHerokuKeyGenerator) GetTableDescription() string {
	return ""
}

func (x *TableHerokuKeyGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableHerokuKeyGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableHerokuKeyGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := heroku_client.Connect(ctx, taskClient.(*heroku_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			opts := heroku.ListRange{Field: "id", Max: 1000}

			items, err := conn.KeyList(ctx, &opts)
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

func (x *TableHerokuKeyGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableHerokuKeyGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("comment").ColumnType(schema.ColumnTypeString).Description("Comment on the key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When key was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("Deprecated. Please refer to 'comment' instead.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("fingerprint").ColumnType(schema.ColumnTypeString).Description("A unique identifying string based on contents.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of this key.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("public_key").ColumnType(schema.ColumnTypeString).Description("Full public_key as uploaded.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("When key was updated.").Build(),
	}
}

func (x *TableHerokuKeyGenerator) GetSubTables() []*schema.Table {
	return nil
}
