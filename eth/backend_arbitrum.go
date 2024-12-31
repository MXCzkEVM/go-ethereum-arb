package eth

import (
	"context"

	"github.com/mxczkevm/go-ethereum-arb/core"
	"github.com/mxczkevm/go-ethereum-arb/core/state"
	"github.com/mxczkevm/go-ethereum-arb/core/types"
	"github.com/mxczkevm/go-ethereum-arb/core/vm"
	"github.com/mxczkevm/go-ethereum-arb/eth/tracers"
	"github.com/mxczkevm/go-ethereum-arb/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*types.Transaction, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	return eth.stateAtTransaction(ctx, block, txIndex, reexec)
}
