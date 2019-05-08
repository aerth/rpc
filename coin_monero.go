package rpc

type Monero struct{}

func (b *Monero) Balance(addr string) (string, error)    { return "0.00000000", nil }
func (b *Monero) BlockByNumber(num int64) (Block, error) { return nil, nil }
func (b *Monero) BlockByHash(hash string) (Block, error) { return nil, nil }
func (b *Monero) Tx(hash string) (Tx, error)             { return nil, nil }

func (b *Monero) Call(resultptr interface{}, method string, args ...interface{}) error {
	return nil
}
