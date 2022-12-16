package tables

import (
	"context"

	"github.com/selefra/selefra-provider-heroku/heroku_client"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

type TableHerokuAccountGenerator struct {
}

var _ table_schema_generator.TableSchemaGenerator = &TableHerokuAccountGenerator{}

func (x *TableHerokuAccountGenerator) GetTableName() string {
	return "heroku_account"
}

func (x *TableHerokuAccountGenerator) GetTableDescription() string {
	return ""
}

func (x *TableHerokuAccountGenerator) GetVersion() uint64 {
	return 0
}

func (x *TableHerokuAccountGenerator) GetOptions() *schema.TableOptions {
	return &schema.TableOptions{}
}

func (x *TableHerokuAccountGenerator) GetDataSource() *schema.DataSource {
	return &schema.DataSource{
		Pull: func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask, resultChannel chan<- any) *schema.Diagnostics {

			conn, err := heroku_client.Connect(ctx, taskClient.(*heroku_client.Client).Config)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			item, err := conn.AccountInfo(ctx)
			if err != nil {

				return schema.NewDiagnosticsErrorPullTable(task.Table, err)
			}
			resultChannel <- item
			return schema.NewDiagnosticsErrorPullTable(task.Table, nil)

		},
	}
}

func (x *TableHerokuAccountGenerator) GetExpandClientTask() func(ctx context.Context, clientMeta *schema.ClientMeta, taskClient any, task *schema.DataSourcePullTask) []*schema.ClientTaskContext {
	return nil
}

func (x *TableHerokuAccountGenerator) GetColumns() []*schema.Column {
	return []*schema.Column{
		table_schema_generator.NewColumnBuilder().ColumnName("last_login").ColumnType(schema.ColumnTypeTimestamp).Description("When account last authorized with Heroku.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("updated_at").ColumnType(schema.ColumnTypeTimestamp).Description("When account was updated.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_team").ColumnType(schema.ColumnTypeJSON).Description("Team selected by default.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("delinquent_at").ColumnType(schema.ColumnTypeTimestamp).Description("When account became delinquent.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("id").ColumnType(schema.ColumnTypeString).Description("Unique identifier of an account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("identity_provider").ColumnType(schema.ColumnTypeJSON).Description("Identity Provider details for federated users.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("italian_partner_terms").ColumnType(schema.ColumnTypeString).Description("Whether account has acknowledged the Italian provider terms of service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("sms_number").ColumnType(schema.ColumnTypeString).Description("SMS number of account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("name").ColumnType(schema.ColumnTypeString).Description("Full name of the account owner.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("acknowledged_msa").ColumnType(schema.ColumnTypeBool).Description("Whether account has acknowledged the MSA terms of service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("allow_tracking").ColumnType(schema.ColumnTypeBool).Description("Whether to allow third party web activity tracking.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("created_at").ColumnType(schema.ColumnTypeTimestamp).Description("When account was created.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("default_organization").ColumnType(schema.ColumnTypeJSON).Description("Team selected by default.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("acknowledged_msa_at").ColumnType(schema.ColumnTypeTimestamp).Description("When account has acknowledged the MSA terms of service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("beta").ColumnType(schema.ColumnTypeBool).Description("Whether allowed to utilize beta Heroku features.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("country_of_residence").ColumnType(schema.ColumnTypeString).Description("Country where account owner resides.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("italian_customer_terms").ColumnType(schema.ColumnTypeString).Description("Whether account has acknowledged the Italian customer terms of service.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("two_factor_authentication").ColumnType(schema.ColumnTypeBool).Description("Whether two-factor auth is enabled on the account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("email").ColumnType(schema.ColumnTypeString).Description("Unique email address of account.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("federated").ColumnType(schema.ColumnTypeBool).Description("Whether the user is federated and belongs to an Identity Provider.").Build(),
		table_schema_generator.NewColumnBuilder().ColumnName("suspended_at").ColumnType(schema.ColumnTypeTimestamp).Description("When account was suspended.").Build(),
	}
}

func (x *TableHerokuAccountGenerator) GetSubTables() []*schema.Table {
	return nil
}
