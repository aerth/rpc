package rpc

import (
	"testing"
	"time"
)

func runAllTests(t *testing.T, testname string, client RPCClient) {
	t.Logf("%s", testname)
	if client == nil {
		t.Errorf("test=%s client is nil", testname)
	}
	// fetch block 1
	blocknum := int64(1)
	block, err := client.BlockByNumber(blocknum)
	if err != nil {
		t.Error(err)
		return
	}
	if block == nil {
		if testname == "DUMMY" {
			return
		}
		t.Error("block is nil")
		return
	}
	// hash it
	hash := block.Hash()

	// fetch it by hash
	block2, err := client.BlockByHash(hash.String())
	if err != nil {
		t.Error(err)
		return
	}

	// compare coinbases
	coinbase := block.Coinbase()
	if block2.Coinbase() != coinbase {
		t.Error("expected both blocks to have same coinbase (same blocks)")
	}

	// check balance
	bal, err := client.Balance(coinbase)
	if err != nil {
		t.Error(err)
		return
	}

	// print details
	t.Logf("%s block #%v", testname, blocknum)
	t.Logf("hash : %x", hash)
	t.Logf("minedby: %s", coinbase)
	t.Logf("time:  %s", block.Timestamp().Format(time.ANSIC))
	t.Logf("bal: %s", bal)

}
