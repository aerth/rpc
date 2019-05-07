package rpc

import (
	"context"
	"errors"
	"math/big"
	"time"

	"gitlab.com/aquachain/aquachain/common"
	"gitlab.com/aquachain/aquachain/core/types"
	aquaclient "gitlab.com/aquachain/aquachain/opt/aquaclient"
	rpcclient "gitlab.com/aquachain/aquachain/rpc/rpcclient"
)

var _ RPCClient = &Aquachain{}

type AquaHeaderVersion = types.HeaderVersion
type AquaHash = common.Hash

type Aquachain struct {
	rpcclient *rpcclient.Client
}

func NewAquachainF(rpchost string) *Aquachain {
	r, err := NewAquachain(rpchost)
	if err != nil {
		panic(err)
	}
	return r
}
func NewAquachain(rpchost string) (*Aquachain, error) {
	tghc, err := tgun.HTTPClient()
	if err != nil {
		return nil, err
	}

	rc, err := rpcclient.DialHTTPWithClient(rpchost, tghc)
	return &Aquachain{
		rpcclient: rc,
	}, err
}

func (a *Aquachain) Balance(addr string) (string, error) {
	var (
		ctx, _      = context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
		address     = common.StringToAddress(addr)
		bignum, err = aquaclient.NewClient(a.rpcclient).Balance(ctx, address)
	)
	if err != nil {
		log.Printf("aqua: %v", err)
		return "", err
	}
	return bignum.String(), nil
}
func (a *Aquachain) BlockByNumber(num int64) (Block, error) {
	var (
		ctx, _   = context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
		bignum   = big.NewInt(num)
		blk, err = aquaclient.NewClient(a.rpcclient).BlockByNumber(ctx, bignum)
	)
	if err != nil {
		return nil, err
	}
	return blk, nil
}
func (a *Aquachain) BlockByHash(hashstr string) (Block, error) {
	var (
		hash     = common.HexToHash(hashstr)
		ctx, _   = context.WithDeadline(context.Background(), time.Now().Add(time.Second*30))
		blk, err = aquaclient.NewClient(a.rpcclient).BlockByHash(ctx, hash)
	)
	if err != nil {
		return nil, err
	}
	return blk, nil
}
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
	return tx, errors.New("is pending")
}

func (a *Aquachain) Call(resultptr interface{}, method string, args ...interface{}) error {
	return a.rpcclient.Call(resultptr, method, args...)
}
