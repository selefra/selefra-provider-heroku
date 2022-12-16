package provider

import (
	"context"
	"os"

	"github.com/selefra/selefra-provider-heroku/constants"

	"testing"

	"github.com/selefra/selefra-provider-sdk/env"
	"github.com/selefra/selefra-provider-sdk/grpc/shard"

	"github.com/selefra/selefra-provider-sdk/provider"

	"github.com/selefra/selefra-provider-sdk/provider/schema"

	"github.com/selefra/selefra-provider-sdk/storage/database_storage/postgresql_storage"

	"github.com/selefra/selefra-utils/pkg/json_util"

	"github.com/selefra/selefra-utils/pkg/pointer"
)

func TestProvider_PullTable(t *testing.T) {

	os.Setenv(constants.SELEFRADATABASEDSN, constants.Hostuserpostgrespasswordpassportdbnamepostgressslmodedisable)
	wk := constants.Constants_5

	config := `providers:
  email: gemkmdsafdm@outlook.com
  api_key: 11f36d92-c8cc-4217-b5a3-49058ba3377d`
	myProvider := GetProvider()

	Pull(myProvider, config, wk, constants.Constants_6)

}

func Pull(myProvider *provider.Provider, config, workspace string, pullTables ...string) {

	diagnostics := schema.NewDiagnostics()

	initProviderRequest := &shard.ProviderInitRequest{

		Storage: &shard.Storage{

			Type:           0,
			StorageOptions: json_util.ToJsonBytes(postgresql_storage.NewPostgresqlStorageOptions(env.GetDatabaseDsn())),
		},

		Workspace: &workspace,

		IsInstallInit: pointer.TruePointer(),

		ProviderConfig: &config,
	}

	response, err := myProvider.Init(context.Background(), initProviderRequest)

	if err != nil {
		panic(diagnostics.AddFatal(constants.Initprovidererrors, err.Error()).ToString())

	}
	if diagnostics.AddDiagnostics(response.Diagnostics).HasError() {
		panic(diagnostics.ToString())

	}

	err = myProvider.PullTables(context.Background(), &shard.PullTablesRequest{

		Tables: pullTables,

		MaxGoroutines: 100,
		Timeout:       1000 * 60 * 60,
	}, shard.NewFakeProviderServerSender())

	if err != nil {
		panic(diagnostics.AddFatal(constants.Providerpulltableerrors, err.Error()).ToString())
	}

}
