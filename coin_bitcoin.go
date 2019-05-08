package rpc

type Bitcoin struct{}

func (b *Bitcoin) Balance(addr string) (string, error)    { return "0.00000000", nil }
func (b *Bitcoin) BlockByNumber(num int64) (Block, error) { return nil, nil }
func (b *Bitcoin) BlockByHash(hash string) (Block, error) { return nil, nil }
func (b *Bitcoin) Tx(hash string) (Tx, error)             { return nil, nil }

func (b *Bitcoin) Call(resultptr interface{}, method string, args ...interface{}) error {
	return nil
}
