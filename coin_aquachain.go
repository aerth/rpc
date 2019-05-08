package rpc

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"gitlab.com/aquachain/aquachain/common"
	"gitlab.com/aquachain/aquachain/core/types"
	aquaclient "gitlab.com/aquachain/aquachain/opt/aquaclient"
	"gitlab.com/aquachain/aquachain/params"
	rpcclient "gitlab.com/aquachain/aquachain/rpc/rpcclient"
)

// Aquachain implements the RPCClient interface
var _ RPCClient = &Aquachain{}

// AquaHeaderVersion is a type alias to aquachain/types.HeaderVersion
type AquaHeaderVersion = types.HeaderVersion

// AquaHash is a type alias to common.Hash
type AquaCommonHash = common.Hash

// MainnetAquachain is the ChainConfig for Aquachain's mainnet network
var MainnetAquachain = params.MainnetChainConfig

// Aquachain implements the RPCClient interface. It is similar to an Ethereum RPCClient but must retain a chainconfig for properly using the block.Hash() method (AQUA has header versions which provide the hash function)
type Aquachain struct {
	rpcclient *rpcclient.Client
	chaincfg  *params.ChainConfig
}

// AquaBlock is a wrapper around aquaBlockWrapper that holds HeaderVersion
type AquaBlock struct {
	aquaBlockWrapper
	HeaderVersion types.HeaderVersion
}

// NewAquachainF force create new aquachain
func NewAquachainF(rpchost string, params *params.ChainConfig) *Aquachain {
	r, err := NewAquachain(rpchost, params)
	if err != nil {
		fmt.Println("err+forced=panic")
		panic(err)
	}
	return r
}

// NewAquachain returns new RPCClient for connecting to Aquachain node with network/chain config
func NewAquachain(rpchost string, params *params.ChainConfig) (*Aquachain, error) {
	tghc, err := tgun.HTTPClient()
	if err != nil {
		return nil, err
	}

	rc, err := rpcclient.DialHTTPWithClient(rpchost, tghc)
	return &Aquachain{
		rpcclient: rc,
		chaincfg:  params,
	}, err
}

// Balance returns total coins (eg: 1.23456789) as a string
func (a *Aquachain) Balance(addr string) (string, error) {
	var (
		ctx, _ = context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
	)
	var address = common.HexToAddress(addr)
	var bignum, err = aquaclient.NewClient(a.rpcclient).Balance(ctx, address)
	if err != nil {
		log.Printf("aqua: %v", err)
		return "ERR", err
	}
	dots := new(big.Float).Quo(new(big.Float).SetInt(bignum), big.NewFloat(params.Aqua))
	return fmt.Sprintf("%.8f", dots), nil
}

// BlockByNumber returns Block
func (a *Aquachain) BlockByNumber(num int64) (Block, error) {
	var (
		ctx, _   = context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
		bignum   = big.NewInt(num)
		blk, err = aquaclient.NewClient(a.rpcclient).BlockByNumber(ctx, bignum)
	)
	if err != nil {
		return nil, err
	}
	return AquaBlock{Aqua2Block(*blk).(aquaBlockWrapper), a.headerversion(bignum)}, nil
}

// headerversion returns header version according to chainconfig
func (a *Aquachain) headerversion(num *big.Int) types.HeaderVersion {
	return a.chaincfg.GetBlockVersion(num)
}

// BlockByHash returns Block
func (a *Aquachain) BlockByHash(hashstr string) (Block, error) {
	var (
		hash     = common.HexToHash(hashstr)
		ctx, _   = context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
		blk, err = aquaclient.NewClient(a.rpcclient).BlockByHash(ctx, hash)
	)
	if err != nil {
		return nil, err
	}
	return AquaBlock{Aqua2Block(*blk).(aquaBlockWrapper), a.headerversion(blk.Number())}, nil
}

// Tx returns Tx if exists
func (a *Aquachain) Tx(hashstr string) (Tx, error) {
	var (
		hash               = common.HexToHash(hashstr)
		ctx, _             = context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
		tx, isPending, err = aquaclient.NewClient(a.rpcclient).TransactionByHash(ctx, hash)
	)
	if err != nil {
		return nil, err
	}
	if isPending {

	}
	return Aqua2Tx(*tx), errors.New("is pending")
}

// Call custom method
func (a *Aquachain) Call(resultptr interface{}, method string, args ...interface{}) error {
	return a.rpcclient.Call(resultptr, method, args...)
}

func Aqua2Hash(h common.Hash) Hash {
	return h
}

func Aqua2Tx(tx types.Transaction) Tx {
	return AquaTxWrapper{
		Tx: tx,
	}
}

func Aqua2Block(block types.Block) Block {
	return aquaBlockWrapper{
		block:      block,
		coinbase:   block.Coinbase().String(),
		timestamp:  time.Unix(block.Time().Int64(), 0),
		parentHash: Aqua2Hash(block.ParentHash()),
	}
}

type aquaBlockWrapper struct {
	block        types.Block
	coinbase     string
	transactions []Tx
	timestamp    time.Time
	parentHash   Hash
}

func (b aquaBlockWrapper) Coinbase() string {
	return b.coinbase
}
func (b AquaBlock) Hash() Hash {
	// aqua blocks received through the wire need set header version
	return b.block.SetVersion(b.HeaderVersion)
}
func (b aquaBlockWrapper) Timestamp() time.Time {
	return time.Unix(b.block.Time().Int64(), 0).UTC()
}
func (b aquaBlockWrapper) Hash() Hash {
	return b.block.Hash()
}
func (b aquaBlockWrapper) Transactions() []Hash {
	btx := b.block.Transactions()
	hashes := make([]Hash, len(btx))
	for i, v := range btx {
		hashes[i] = Aqua2Hash(v.Hash())
	}
	return hashes
}

type AquaTxWrapper struct {
	Tx types.Transaction
}

func (e AquaTxWrapper) Hash() Hash {
	return e.Tx.Hash()
}
