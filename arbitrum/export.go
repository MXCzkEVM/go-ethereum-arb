package arbitrum

import (
	"context"

	"github.com/mxczkevm/go-ethereum-arb/common/hexutil"
	"github.com/mxczkevm/go-ethereum-arb/core"
	"github.com/mxczkevm/go-ethereum-arb/internal/ethapi"
	"github.com/mxczkevm/go-ethereum-arb/rpc"
)

type TransactionArgs = ethapi.TransactionArgs

func EstimateGas(ctx context.Context, b ethapi.Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, overrides *ethapi.StateOverride, gasCap uint64) (hexutil.Uint64, error) {
	return ethapi.DoEstimateGas(ctx, b, args, blockNrOrHash, overrides, gasCap)
}

func NewRevertReason(result *core.ExecutionResult) error {
	return ethapi.NewRevertError(result)
}
