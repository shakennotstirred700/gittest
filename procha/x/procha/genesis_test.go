package procha_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "procha/testutil/keeper"
	"procha/testutil/nullify"
	"procha/x/procha"
	"procha/x/procha/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.ProchaKeeper(t)
	procha.InitGenesis(ctx, *k, genesisState)
	got := procha.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
