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

type TableHerokuOutboundRulesetsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableHerokuOutboundRulesetsGenerator{}

func (x *TableHerokuOutboundRulesetsGenerator) GetTableName() string {
	return "heroku_outbound_rulesets"
}

func (x *TableHerokuOutboundRulesetsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableHerokuOutboundRulesetsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableHerokuOutboundRulesetsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{
		PrimaryKeys: []string{
			"id",
		},
	}
}

func (x *TableHerokuOutboundRulesetsGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {
			c, _ := heroku_client.Connect(ctx, client.(*heroku_client.Client).Config)
			nextRange := &heroku.ListRange{
				Field: "id",
				Max:   1000,
			}
			items := make([]heroku.Space, 0, 10)

			for nextRange.Max != 0 {
				ctxWithRange := context.WithValue(ctx, "nextRange", nextRange)
				v, err := c.SpaceList(ctxWithRange, nextRange)
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
					v, err := c.OutboundRulesetList(ctxWithRange, it.ID, nextRange)
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

func (x *TableHerokuOutboundRulesetsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableHerokuOutboundRulesetsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_by").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("ID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("rules").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("space").ColumnType(schema.ColumnTypeJSON).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("primary keys value md5").
			Extractor(column_value_extractor.PrimaryKeysID()).Build(),
	}
}

func (x *TableHerokuOutboundRulesetsGenerator) GetSubTables() []*schema.Table {
	return nil
}
