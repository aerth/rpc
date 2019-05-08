package rpc

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	//	"github.com/ethereum/go-ethereum/common"
	//	"github.com/ethereum/go-ethereum/core/types"
	//	ethclient "github.com/ethereum/go-ethereum/ethclient"
	//	"github.com/ethereum/go-ethereum/params"
	//	rpcclient "github.com/ethereum/go-ethereum/rpc"

	// easier to use pure-go aquachain
	"gitlab.com/aquachain/aquachain/common"
	"gitlab.com/aquachain/aquachain/core/types"
	ethclient "gitlab.com/aquachain/aquachain/opt/aquaclient"
	"gitlab.com/aquachain/aquachain/params"
	rpcclient "gitlab.com/aquachain/aquachain/rpc/rpcclient"
)

var OneAqua = params.Aqua
var OneEther = OneAqua

// Ethereum implements the RPCClient interface
var _ RPCClient = &Ethereum{}

// Ethereum implements the RPCClient interface.
type Ethereum struct {
	rpcclient *rpcclient.Client
}

// NewEthereumF force create new aquachain
func NewEthereumF(rpchost string) *Ethereum {
	r, err := NewEthereum(rpchost)
	if err != nil {
		fmt.Println("err+forced=panic")
		panic(err)
	}
	return r
}

// NewEthereum returns new RPCClient for connecting to Ethereum node with network/chain config
func NewEthereum(rpchost string) (*Ethereum, error) {
	tghc, err := tgun.HTTPClient()
	if err != nil {
		return nil, err
	}

	rc, err := rpcclient.DialHTTPWithClient(rpchost, tghc)
	return &Ethereum{
		rpcclient: rc,
	}, err
}

// Balance returns total coins (eg: 1.23456789) as a string
func (a *Ethereum) Balance(addr string) (string, error) {
	var (
		ctx, _ = context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
	)
	var address = common.HexToAddress(addr)
	var bignum, err = ethclient.NewClient(a.rpcclient).BalanceAt(ctx, address, nil)
	if err != nil {
		log.Printf("aqua: %v", err)
		return "ERR", err
	}
	dots := new(big.Float).Quo(new(big.Float).SetInt(bignum), big.NewFloat(OneEther))
	return fmt.Sprintf("%.8f", dots), nil
}

// BlockByNumber returns Block
func (a *Ethereum) BlockByNumber(num int64) (Block, error) {
	var (
		ctx, _   = context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
		bignum   = big.NewInt(num)
		blk, err = ethclient.NewClient(a.rpcclient).BlockByNumber(ctx, bignum)
	)
	if err != nil {
		return nil, err
	}
	return EtherBlock(*blk), nil
}

// BlockByHash returns Block
func (a *Ethereum) BlockByHash(hashstr string) (Block, error) {
	var (
		hash     = common.HexToHash(hashstr)
		ctx, _   = context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
		blk, err = ethclient.NewClient(a.rpcclient).BlockByHash(ctx, hash)
	)
	if err != nil {
		return nil, err
	}
	return EtherBlock(*blk), nil
}

// Tx returns Tx if exists
func (a *Ethereum) Tx(hashstr string) (Tx, error) {
	var (
		hash               = common.HexToHash(hashstr)
		ctx, _             = context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
		tx, isPending, err = ethclient.NewClient(a.rpcclient).TransactionByHash(ctx, hash)
	)
	if err != nil {
		return nil, err
	}
	if isPending {

	}
	return EtherTx(*tx), errors.New("is pending")
}

// Call custom method
func (a *Ethereum) Call(resultptr interface{}, method string, args ...interface{}) error {
	return a.rpcclient.Call(resultptr, method, args...)
}

func EtherHash(h common.Hash) Hash {
	return h
}

func EtherTx(tx types.Transaction) Tx {
	return EthTxWrapper{
		Tx: tx,
	}
}

func EtherBlock(block types.Block) Block {
	return ethBlockWrapper{
		block:      block,
		coinbase:   block.Coinbase().String(),
		timestamp:  time.Unix(int64(block.Time().Int64()), 0),
		parentHash: EtherHash(block.ParentHash()),
	}
}

type ethBlockWrapper struct {
	block        types.Block
	coinbase     string
	transactions []Tx
	timestamp    time.Time
	parentHash   Hash
}

func (b ethBlockWrapper) Coinbase() string {
	return b.coinbase
}
func (b ethBlockWrapper) Timestamp() time.Time {
	return time.Unix(int64(b.block.Time().Int64()), 0).UTC()
}
func (b ethBlockWrapper) Hash() Hash {
	// little hackity hack to set version 1 on all blocks
	if b.block.Version() != 0 {
		return b.block.Hash()
	}
	return b.block.SetVersion(1)
}
func (b ethBlockWrapper) Transactions() []Hash {
	btx := b.block.Transactions()
	hashes := make([]Hash, len(btx))
	for i, v := range btx {
		hashes[i] = EtherHash(v.Hash())
	}
	return hashes
}

type EthTxWrapper struct {
	Tx types.Transaction
}

func (e EthTxWrapper) Hash() Hash {
	return e.Tx.Hash()
}
