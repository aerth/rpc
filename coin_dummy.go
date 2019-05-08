package rpc

type dummyCoin struct{}

func (d dummyCoin) Balance(addr string) (string, error)    { return "0.00000000", nil }
func (d dummyCoin) BlockByNumber(num int64) (Block, error) { return nil, nil }
func (d dummyCoin) BlockByHash(hash string) (Block, error) { return nil, nil }
func (d dummyCoin) Tx(hash string) (Tx, error)             { return nil, nil }

func (d dummyCoin) Call(resultptr interface{}, method string, args ...interface{}) error {
	return nil
}
