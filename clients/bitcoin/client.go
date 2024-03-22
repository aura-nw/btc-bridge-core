package bitcoin

import (
	"encoding/hex"
	"fmt"
	"log/slog"
	"strings"

	"github.com/aura-nw/btc-bridge-core/config"
	"github.com/aura-nw/btc-bridge-core/types"
	"github.com/btcsuite/btcd/btcjson"
	"github.com/btcsuite/btcd/rpcclient"
	"github.com/btcsuite/btcd/txscript"
)

// Client defines Bitcoin client.
type Client interface {
	GetBtcDeposits(height int64, filterAddr string, minConfirms int64) ([]types.BtcDeposit, error)
	GetTokenDeposits(height int64, filterAddr string, minConfirms int64) ([]types.InscriptionDeposit, error)
}

type clientImpl struct {
	info      config.BitcoinInfo
	logger    *slog.Logger
	rpcClient *rpcclient.Client
}

var _ Client = &clientImpl{}

func NewClient(logger *slog.Logger, info config.BitcoinInfo) (Client, error) {
	connCfg := &rpcclient.ConnConfig{
		Host:         info.Host,
		User:         info.User,
		Pass:         info.Password,
		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	rpcClient, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return nil, err
	}
	return &clientImpl{
		info:      info,
		logger:    logger,
		rpcClient: rpcClient,
	}, nil
}

func (c *clientImpl) GetBtcDeposits(height int64, filterAddr string, minConfirms int64) ([]types.BtcDeposit, error) {
	blockHash, err := c.rpcClient.GetBlockHash(height)
	if err != nil {
		return nil, err
	}
	block, err := c.rpcClient.GetBlockVerboseTx(blockHash)
	if err != nil {
		return nil, err
	}
	var results []types.BtcDeposit
	for _, tx := range block.Tx {
		for _, vout := range tx.Vout {
			if filterAddr == vout.ScriptPubKey.Address {
				memo, err := c.getMemo(&vout)
				if err != nil {
					c.logger.Error("get memo failed", "err", err)
					continue
				}
				results = append(results, types.BtcDeposit{
					TxHash:         tx.Hash,
					Height:         height,
					Memo:           memo,
					Receiver:       "",
					Sender:         "",
					MultisigWallet: filterAddr,
					Amount:         fmt.Sprintf("%f", vout.Value),
					Idx:            vout.N,
				})
			}
		}
	}
	if len(results) == 0 {
		return nil, fmt.Errorf("no txs found for height: %d, filter addr: %s", height, filterAddr)
	}
	return results, nil
}

func (c *clientImpl) GetTokenDeposits(height int64, filterAddr string, minConfirms int64) ([]types.InscriptionDeposit, error) {
	var results []types.InscriptionDeposit
	return results, nil
}

// getMemo returns memo for a btc tx, using vout OP_RETURN
func (c *clientImpl) getMemo(vOut *btcjson.Vout) (string, error) {
	var opReturns string
	buf, err := hex.DecodeString(vOut.ScriptPubKey.Hex)
	if err != nil {
		c.logger.Error("fail to hex decode scriptPubKey", "err", err)
	}

	asm, err := txscript.DisasmString(buf)

	if err != nil {
		c.logger.Error("fail to disasm script pubkey", "err", err)
	}
	opReturnFields := strings.Fields(asm)
	if len(opReturnFields) == 2 {
		var decoded []byte
		decoded, err = hex.DecodeString(opReturnFields[1])
		if err != nil {
			c.logger.Error("fail to decode OP_RETURN string: %s", "err", err)
		}
		opReturns += string(decoded)
	}

	return opReturns, nil
}
