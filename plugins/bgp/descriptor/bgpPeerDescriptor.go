package descriptor

import (
	"context"

	"github.com/contiv/bgp-vpp/plugins/bgp/descriptor/adapter"
	"github.com/contiv/bgp-vpp/plugins/bgp/model"
	"github.com/ligato/cn-infra/logging"
	kvs "github.com/ligato/vpp-agent/plugins/kvscheduler/api"
	bgpapi "github.com/osrg/gobgp/api"
	gobgp "github.com/osrg/gobgp/pkg/server"
)

const (
	peerDescriptorName = "peer-conf"
)

//our descriptor
type PeerDescriptor struct {
	log    logging.Logger
	server *gobgp.BgpServer
}

// NewGlobalConfDescriptor creates a new instance of the descriptor.
func NewPeerConfDescriptor(log logging.PluginLogger, server *gobgp.BgpServer) *kvs.KVDescriptor {
	d := &PeerDescriptor{log: log, server: server}

	// Set plugin descriptor init values
	gcd := &adapter.PeerConfDescriptor{
		Name:          peerDescriptorName,
		NBKeyPrefix:   model.ModelBgpPeer.KeyPrefix(),
		ValueTypeName: model.ModelBgpPeer.ProtoName(),
		KeySelector:   model.ModelBgpPeer.IsKeyValid,
		KeyLabel:      model.ModelBgpPeer.StripKeyPrefix,
		Create:        d.Create,
		Delete:        d.Delete,
		UpdateWithRecreate: func(key string, oldValue, newValue *model.PeerConf, metadata interface{}) bool {
			// Modify always performed via re-creation
			return true
		},
	}
	return adapter.NewPeerConfDescriptor(gcd)
}

// Create creates new value.
func (d *PeerDescriptor) Create(key string, value *model.PeerConf) (metadata interface{}, err error) {
	d.log.Infof("Creating Peer %s,  neighbor_address = %s, peer_as = %d",
		value.Name, value.NeighborAddress, value.PeerAs)
	n := &bgpapi.Peer{
		Conf: &bgpapi.PeerConf{
			NeighborAddress: value.NeighborAddress,
			PeerAs:          value.PeerAs,
		},
	}
	err = d.server.AddPeer(context.Background(), &bgpapi.AddPeerRequest{
		Peer: n,
	})
	if err != nil {
		d.log.Errorf("Error creating PeerConf = %s", err)
		return nil, err
	}

	return nil, nil
}

// Delete removes an existing value.
func (d *PeerDescriptor) Delete(key string, value *model.PeerConf, metadata interface{}) error {
	d.log.Infof("Deleting Peer %s", value.Name)
	err := d.server.DeletePeer(context.Background(), &bgpapi.DeletePeerRequest{})
	if err != nil {
		return err
	}
	return nil
}

// UpdateWithRecreate returns true if value update requires full re-creation.
/*func (d *GlobalDescriptor) UpdateWithRecreate(key string, old, new *model.GlobalConf, metadata interface{}) bool {
	return true
}*/

// Retrieve retrieves values from SB.
/*func (d *GlobalDescriptor) Retrieve(correlate []adapter.PluginKVWithMetadata) (retrieved []adapter.PluginKVWithMetadata, err error) {
	return retrieved, nil
}*/

// Dependencies lists dependencies of the given value.
func (d *PeerDescriptor) Dependencies(key string, value *model.PeerConf) (deps []kvs.Dependency) {
	return deps
}
