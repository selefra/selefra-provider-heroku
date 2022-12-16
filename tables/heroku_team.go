package tables

import (
	"context"

	heroku "github.com/heroku/heroku-go/v5"

	"github.com/selefra/selefra-provider-heroku/heroku_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/provider/transformer/column_value_extractor"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableHerokuTeamGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableHerokuTeamGenerator{}

func (x *TableHerokuTeamGenerator) GetTableName() string {
	return "heroku_team"
}

func (x *TableHerokuTeamGenerator) GetTableDescription() string {
	return ""
}

func (x *TableHerokuTeamGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableHerokuTeamGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableHerokuTeamGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := heroku_client.Connect(ctx, taskClient.(*heroku_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			opts := heroku.ListRange{Field: "id", Max: 1000}

			items, err := conn.TeamList(ctx, &opts)
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

func (x *TableHerokuTeamGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableHerokuTeamGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the team was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("role").ColumnType(schema.ColumnTypeString).Description("Role in the team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("When the team was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("membership_limit").ColumnType(schema.ColumnTypeFloat).Description("Upper limit of members allowed in a team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("provisioned_licenses").ColumnType(schema.ColumnTypeBool).Description("Whether the team is provisioned licenses by salesforce.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("team_type").ColumnType(schema.ColumnTypeString).Description("Type of team.").
			Extractor(column_value_extractor.StructSelector("Type")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Unique name of team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("credit_card_collections").ColumnType(schema.ColumnTypeBool).Description("Whether charges incurred by the team are paid by credit card.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("is_default").ColumnType(schema.ColumnTypeBool).Description("Whether to use this team when none is specified.").
			Extractor(column_value_extractor.StructSelector("Default")).Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("enterprise_account").ColumnType(schema.ColumnTypeJSON).Description("Enterprise Account associated with the team.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_provider").ColumnType(schema.ColumnTypeJSON).Description("Identity Provider associated with the team.").Build(),
	}
}

func (x *TableHerokuTeamGenerator) GetSubTables() []*schema.Table {
	return nil
}
