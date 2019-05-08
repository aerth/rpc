# xchains/rpc

  * connect to all of your rpcs from one place

### usage

  * import "github.com/xchains/rpc"
  * initialize one or more RPCClient(s)
  * use Balance, BlockByNumber, BlockByHash, or lookup Tx
  * send signed transaction

### adding a new coin
  * add `coin_mycoin.go` and implement the RPCClient interface
  * make wrappers for Tx, Block, Hash data types
  * add to `allcoins_test`

**Note: this is a work in progress and the API should not be considered stable**
