package fsrepo

import (
	"os"
	"runtime"
	"testing"

	config "github.com/ipfs/go-ipfs-config"
)

func TestConfig(t *testing.T) {
	const filename = ".ipfsconfig"
	cfgWritten := new(config.Config)
	cfgWritten.Identity.PeerID = "faketest"
	cfgWritten.P2P.Listen = map[string]string{
		"/x/one": "/ip4/127.0.0.1/tcp/1234",
	}
	cfgWritten.P2P.Forward = map[string]config.ForwardConfig{
		"/ip4/127.0.0.1/tcp/4411": {
			Peer:     "QmId",
			Protocol: "/x/another",
		},
	}

	err := WriteConfigFile(filename, cfgWritten)
	if err != nil {
		t.Fatal(err)
	}
	cfgRead, err := Load(filename)
	if err != nil {
		t.Fatal(err)
	}
	if cfgWritten.Identity.PeerID != cfgRead.Identity.PeerID {
		t.Fatal()
	}

	for protocol, addr := range cfgWritten.P2P.Listen {
		if cfgRead.P2P.Listen[protocol] != addr {
			t.Fatal()
		}
	}

	for addr, forward := range cfgWritten.P2P.Forward {
		if cfgRead.P2P.Forward[addr].Peer != forward.Peer {
			t.Fatal()
		}
	}

	st, err := os.Stat(filename)
	if err != nil {
		t.Fatalf("cannot stat config file: %v", err)
	}

	if runtime.GOOS != "windows" { // see https://golang.org/src/os/types_windows.go
		if g := st.Mode().Perm(); g&0117 != 0 {
			t.Fatalf("config file should not be executable or accessible to world: %v", g)
		}
	}
}
