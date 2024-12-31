package arbitrum

import (
	"context"

	"github.com/mxczkevm/go-ethereum-arb/arbitrum_types"
	"github.com/mxczkevm/go-ethereum-arb/core"
	"github.com/mxczkevm/go-ethereum-arb/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction, options *arbitrum_types.ConditionalOptions) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}
