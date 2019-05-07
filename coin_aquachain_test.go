package rpc

import (
	"os"
	"testing"
	"time"

	"gitlab.com/aquachain/aquachain/common"
)

func TestAquachain(t *testing.T) {
	t1 := time.Now()
	defer func() {
		t.Logf("AQUA test took %s", time.Since(t1))
	}()
	node := os.Getenv("AQUA_NODE")
	if node == "" {
		t.Skip("AQUA_NODE env not set, skipping.")
	}
	aqua, err := NewAquachain(node)
	if err != nil {
		t.Error(err)
	}

	blk, err := aqua.BlockByNumber(0)
	if err != nil {
		t.Error(err)
	}
	t.Logf("Connect took %s", time.Since(t1))

	// decode genesis block
	type VersionSetter interface {
		SetVersion(n AquaHeaderVersion) AquaHash
	}
	block, ok := blk.(VersionSetter)
	if !ok || block == nil {
		t.Errorf("Could not convert interface to block type")
	}

	// compare known genesis to result
	hash := block.SetVersion(1) // use hash version 1
	gen := common.HexToHash("0x381c8d2c3e3bc702533ee504d7621d510339cafd830028337a4b532ff27cd505")
	if hash != gen {
		t.Errorf("%x != %x", hash, gen)
	}
	t.Log("AQUA genesis hash:", hash.String())
}
