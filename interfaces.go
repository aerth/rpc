// rpc package - chains in unison
package rpc

import (
	logpkg "log" // imported only in this file
	"os"         // imported only in this file
	"time"

	tgunpkg "github.com/aerth/tgun" // only here
)

/* dis gon be tuf */

// Block ....
type Block interface {
	Hash() Hash
	Transactions() []Hash
	Coinbase() string
	Timestamp() time.Time
}

// Tx ....
type Tx interface {
	Hash() Hash
}

// Hash ....
type Hash interface {
	String() string
} // multihash

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
// use instead of importing net/http directly
var Tgun = &tgunpkg.Client{
	Proxy: "socks5://127.0.0.1:1080",
}

var tgun = Tgun
