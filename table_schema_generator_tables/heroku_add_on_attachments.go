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

type TableHerokuAddOnAttachmentsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableHerokuAddOnAttachmentsGenerator{}

func (x *TableHerokuAddOnAttachmentsGenerator) GetTableName() string {
	return "heroku_add_on_attachments"
}

func (x *TableHerokuAddOnAttachmentsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableHerokuAddOnAttachmentsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableHerokuAddOnAttachmentsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableHerokuAddOnAttachmentsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c, _ := heroku_client.Connect(ctx, client.(*heroku_client.Client).Config)
			nextRange := &heroku.ListRange{
				Field: "id",
				Max:   1000,
			}

			for nextRange.Max != 0 {
				ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
				v, err := c.AddOnAttachmentList(ctxWithRange, nextRange)
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

func (x *TableHerokuAddOnAttachmentsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableHerokuAddOnAttachmentsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("log_input_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("LogInputURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("addon").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("app").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("namespace").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("web_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("WebURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
	}
}

func (x *TableHerokuAddOnAttachmentsGenerator) GetSubTables() []*schema.Table {
	return nil
}
