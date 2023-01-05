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

type TableHerokuBuildpackInstallationsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableHerokuBuildpackInstallationsGenerator{}

func (x *TableHerokuBuildpackInstallationsGenerator) GetTableName() string {
	return "heroku_buildpack_installations"
}

func (x *TableHerokuBuildpackInstallationsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableHerokuBuildpackInstallationsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableHerokuBuildpackInstallationsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableHerokuBuildpackInstallationsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c, _ := heroku_client.Connect(ctx, client.(*heroku_client.Client).Config)
			nextRange := &heroku.ListRange{
				Field: "id",
				Max:   1000,
			}
			items := make([]heroku.App, 0, 10)

			for nextRange.Max != 0 {
				ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
				v, err := c.AppList(ctxWithRange, nextRange)
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
					v, err := c.BuildpackInstallationList(ctxWithRange, it.ID, nextRange)
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

func (x *TableHerokuBuildpackInstallationsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableHerokuBuildpackInstallationsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("buildpack").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("ordinal").ColumnType(schema.ColumnTypeInt).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
	}
}

func (x *TableHerokuBuildpackInstallationsGenerator) GetSubTables() []*schema.Table {
	return nil
}
