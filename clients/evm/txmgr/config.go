package txmgr

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type SignerFn func(context.Context, common.Address, *types.Transaction) (*types.Transaction, error)

type Config struct {
	Backend EthBackend
	// ResubmissionTimeout is the interval at which, if no previously
	// published transaction has been mined, the new tx with a bumped gas
	// price will be published. Only one publication at MaxGasPrice will be
	// attempted.
	ResubmissionTimeout time.Duration

	// The multiplier applied to fee suggestions to put a hard limit on fee increases.
	FeeLimitMultiplier uint64

	// Minimum threshold (in Wei) at which the FeeLimitMultiplier takes effect.
	// On low-fee networks, like test networks, this allows for arbitrary fee bumps
	// below this threshold.
	FeeLimitThreshold *big.Int

	// Minimum base fee (in Wei) to assume when determining tx fees.
	MinBaseFee *big.Int

	// Minimum tip cap (in Wei) to enforce when determining tx fees.
	MinTipCap *big.Int

	// ChainID is the chain ID of the L1 chain.
	ChainID *big.Int

	// TxSendTimeout is how long to wait for sending a transaction.
	// By default it is unbounded. If set, this is recommended to be at least 20 minutes.
	TxSendTimeout time.Duration

	// TxNotInMempoolTimeout is how long to wait before aborting a transaction send if the transaction does not
	// make it to the mempool. If the tx is in the mempool, TxSendTimeout is used instead.
	TxNotInMempoolTimeout time.Duration

	// NetworkTimeout is the allowed duration for a single network request.
	// This is intended to be used for network requests that can be replayed.
	NetworkTimeout time.Duration

	// RequireQueryInterval is the interval at which the tx manager will
	// query the backend to check for confirmations after a tx at a
	// specific gas price has been published.
	ReceiptQueryInterval time.Duration

	// NumConfirmations specifies how many blocks are need to consider a
	// transaction confirmed.
	NumConfirmations uint64

	// SafeAbortNonceTooLowCount specifies how many ErrNonceTooLow observations
	// are required to give up on a tx at a particular nonce without receiving
	// confirmation.
	SafeAbortNonceTooLowCount uint64

	// Signer is used to sign transactions when the gas price is increased.
	Signer SignerFn
	From   common.Address
}

func (m Config) Check() error {
	if m.Backend == nil {
		return errors.New("must provide the Backend")
	}
	if m.NumConfirmations == 0 {
		return errors.New("NumConfirmations must not be 0")
	}
	if m.NetworkTimeout == 0 {
		return errors.New("must provide NetworkTimeout")
	}
	if m.FeeLimitMultiplier == 0 {
		return errors.New("must provide FeeLimitMultiplier")
	}
	if m.MinBaseFee != nil && m.MinTipCap != nil && m.MinBaseFee.Cmp(m.MinTipCap) == -1 {
		return fmt.Errorf("minBaseFee smaller than minTipCap, have %v < %v",
			m.MinBaseFee, m.MinTipCap)
	}
	if m.ResubmissionTimeout == 0 {
		return errors.New("must provide ResubmissionTimeout")
	}
	if m.ReceiptQueryInterval == 0 {
		return errors.New("must provide ReceiptQueryInterval")
	}
	if m.TxNotInMempoolTimeout == 0 {
		return errors.New("must provide TxNotInMempoolTimeout")
	}
	if m.SafeAbortNonceTooLowCount == 0 {
		return errors.New("SafeAbortNonceTooLowCount must not be 0")
	}
	if m.Signer == nil {
		return errors.New("must provide the Signer")
	}
	if m.ChainID == nil {
		return errors.New("must provide the ChainID")
	}
	return nil
}
