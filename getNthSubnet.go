package codesnippets

import (
	"fmt"

	"github.com/dspinhirne/netaddr-go"
)

func GetNthSubnet(superNet *netaddr.IPv4Net, subPl int, subIndex int) (*netaddr.IPv4Net, error) {
	if superNet.Netmask().PrefixLen() == uint(subPl) && subIndex == 0 {
		return superNet, nil
	}
	subnet0 := superNet.NthSubnet(uint(subPl), uint32(subIndex))
	if subnet0 != nil {
		return subnet0, nil
	}
	return nil, fmt.Errorf("Supernet(%s) does not contain subnet prefixlen(%d) at index(%d)", superNet.String(), subPl, subIndex)
}
