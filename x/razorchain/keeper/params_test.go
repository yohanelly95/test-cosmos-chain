package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "razorChain/testutil/keeper"
	"razorChain/x/razorchain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.RazorchainKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
