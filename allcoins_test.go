package rpc

import (
	"os"
	"testing"
	"time"
)

func TestDummy(t *testing.T) {
	t1 := time.Now()
	defer func() {
		t.Logf("DUMMY test took %s", time.Since(t1))
	}()
	dummy := new(dummyCoin)
	runAllTests(t, "DUMMY", dummy)
}
func TestAquachain(t *testing.T) {
	t1 := time.Now()
	node := os.Getenv("AQUA_NODE")
	if node == "" {
		t.Skip("AQUA_NODE env not set, skipping.")
		return
	}
	defer func() {
		t.Logf("AQUA test took %s", time.Since(t1))
	}()
	aqua, err := NewAquachain(node, MainnetAquachain)
	if err != nil {
		t.Error(err)
		return
	}

	runAllTests(t, "AQUA", aqua)
}
func TestEthereum(t *testing.T) {
	t1 := time.Now()
	node := os.Getenv("ETH_NODE")
	if node == "" {
		t.Skip("ETH_NODE env not set, skipping.")
		return
	}
	defer func() {
		t.Logf("ETH test took %s", time.Since(t1))
	}()
	eth, err := NewEthereum(node)
	if err != nil {
		t.Error(err)
		return
	}

	runAllTests(t, "ETH", eth)
}
