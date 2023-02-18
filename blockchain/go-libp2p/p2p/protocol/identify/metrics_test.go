//go:build nocover

package identify

import (
	"math/rand"
	"testing"

	"github.com/libp2p/go-libp2p/core/event"
	"github.com/libp2p/go-libp2p/core/network"
)

func TestMetricsNoAllocNoCover(t *testing.T) {
	events := []any{
		event.EvtLocalAddressesUpdated{},
		event.EvtLocalProtocolsUpdated{},
		event.EvtNATDeviceTypeChanged{},
	}
	dirs := []network.Direction{network.DirInbound, network.DirOutbound, network.DirUnknown}
	tr := NewMetricsTracer()
	tests := map[string]func(){
		"TriggeredPushes": func() { tr.TriggeredPushes(events[rand.Intn(len(events))]) },
		"Identify":        func() { tr.Identify(dirs[rand.Intn(len(dirs))]) },
		"IdentifyPush":    func() { tr.IdentifyPush(dirs[rand.Intn(len(dirs))]) },
	}
	for method, f := range tests {
		allocs := testing.AllocsPerRun(1000, f)
		if allocs > 0 {
			t.Fatalf("Alloc Test: %s, got: %0.2f, expected: 0 allocs", method, allocs)
		}
	}
}
