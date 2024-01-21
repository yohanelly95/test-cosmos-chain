package razorchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "razorChain/testutil/keeper"
	"razorChain/testutil/nullify"
	"razorChain/x/razorchain/module"
	"razorChain/x/razorchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.RazorchainKeeper(t)
	razorchain.InitGenesis(ctx, k, genesisState)
	got := razorchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
