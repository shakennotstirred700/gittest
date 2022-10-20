package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "procha/testutil/keeper"
	"procha/x/procha/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ProchaKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
