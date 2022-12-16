package provider

import (
	"github.com/selefra/selefra-provider-heroku/tables"
	"github.com/selefra/selefra-provider-sdk/provider/schema"
	"github.com/selefra/selefra-provider-sdk/table_schema_generator"
)

func GenTables() []*schema.Table {
	return []*schema.Table{
		table_schema_generator.GenTableSchema(&tables.TableHerokuAccountGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableHerokuRegionGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableHerokuAddonGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableHerokuKeyGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableHerokuAppGenerator{}),
		table_schema_generator.GenTableSchema(&tables.TableHerokuTeamGenerator{}),
	}
}
