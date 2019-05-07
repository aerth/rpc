// rpc package - chains in unison
package rpc

import (
	logpkg "log"
	"os"

	tgunpkg "github.com/aerth/tgun"
)

/* dis gon be tuf */

type Block interface{}
type Tx interface{}

// RPCClient easy to implement ..?
type RPCClient interface {
	Balance(addr string) (string, error)
	BlockByNumber(num int64) (Block, error)
	BlockByHash(hash string) (Block, error)
	Tx(txhash string) (Tx, error)
	Call(result interface{}, method string, args ...interface{}) error
}

// log can be used by all files in this package, instead of loading new loggers per file.
var log = logpkg.New(os.Stderr, "", 0)

// tgun may be used for easy proxied requests.
// the http.go file contains safe aliases for use instead of importing net/http directly
var tgun = &tgunpkg.Client{
	Proxy: "socks5://127.0.0.1:1080",
}
