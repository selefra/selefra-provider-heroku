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

type TableHerokuPeeringsGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableHerokuPeeringsGenerator{}

func (x *TableHerokuPeeringsGenerator) GetTableName() string {
	return "heroku_peerings"
}

func (x *TableHerokuPeeringsGenerator) GetTableDescription() string {
	return ""
}

func (x *TableHerokuPeeringsGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableHerokuPeeringsGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableHerokuPeeringsGenerator) GetDataSource() *schema.DataSource {
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
					v, err := c.PeeringList(ctxWithRange, it.ID, nextRange)
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

func (x *TableHerokuPeeringsGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, client any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableHerokuPeeringsGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("aws_account_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AwsAccountID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_region").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("aws_vpc_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("AwsVpcID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("status").ColumnType(schema.ColumnTypeString).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("selefra_id").ColumnType(schema.ColumnTypeString).SetUnique().Description("random id").
			Extractor(column_value_extractor.UUID()).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("cidr_blocks").ColumnType(schema.ColumnTypeStringArray).
			Extractor(column_value_extractor.StructSelector("CIDRBlocks")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("expires").ColumnType(schema.ColumnTypeTimestamp).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("pcx_id").ColumnType(schema.ColumnTypeString).
			Extractor(column_value_extractor.StructSelector("PcxID")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("type").ColumnType(schema.ColumnTypeString).Build(),
	}
}

func (x *TableHerokuPeeringsGenerator) GetSubTables() []*schema.Table {
	return nil
}
