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

type TableHerokuPipelineReleasesGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableHerokuPipelineReleasesGenerator{}

func (x *TableHerokuPipelineReleasesGenerator) GetTableName() string {
	return "heroku_pipeline_releases"
}

func (x *TableHerokuPipelineReleasesGenerator) GetTableDescription() string {
	return ""
}

func (x *TableHerokuPipelineReleasesGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableHerokuPipelineReleasesGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableHerokuPipelineReleasesGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c, _ := heroku_client.Connect(ctx, client.(*heroku_client.Client).Config)
			nextRange := &heroku.ListRange{
				Field: "id",
				Max:   1000,
			}
			items := make([]heroku.Pipeline, 0, 10)

			for nextRange.Max != 0 {
				ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
				v, err := c.PipelineList(ctxWithRange, nextRange)
				if err != nil {
					maybeError := errors.WithStack(err)
					return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
				}
				items = append(items, v...)
			}

			for _, it := range items {
				nextRange = &heroku.ListRange{
					Field: "id",
					Max:   1000,
				}

				for nextRange.Max != 0 {
					ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
					v, err := c.PipelineReleaseList(ctxWithRange, it.ID, nextRange)
					if err != nil {
						maybeError := errors.WithStack(err)
						return schema.NewDiagnosticsErrorPullTable(task.Table, maybeError)
					}
					resultChannel <- v
				}
			}
			return nil
		},
	}
}

func (x *TableHerokuPipelineReleasesGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableHerokuPipelineReleasesGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("version").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("description").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("user").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("output_stream_url").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("OutputStreamURL")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("slug").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("addon_plan_names").ColumnType(schema.ColumnTypeStringArray).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("app").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("current").ColumnType(schema.ColumnTypeBool).Build(),
	}
}

func (x *TableHerokuPipelineReleasesGenerator) GetSubTables() []*schema.Table {
	return nil
}
