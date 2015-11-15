package netUtil

import "net"

// IPNet is a wrapper object for the IPNet of the golang network package
type IPNet struct {
	network  *net.IPNet
	ipLenght int // 4 or 16
}

func newNetwork(n *net.IPNet) *IPNet {
	net := new(IPNet)

	net.network = n
	net.ipLenght = len(n.IP)

	return net
}

// ParseCIDR parses s as a CIDR notation IP address and mask,
// like "192.168.100.1/24" or "2001:DB8::/48", as defined in
// RFC 4632 and RFC 4291.
//
// It returns the IP address and the network implied by the IP
// and mask.  For example, ParseCIDR("192.168.100.1/16") returns
// the IP address 192.168.100.1 and the network 192.168.0.0/16.
func ParseCIDR(s string) (net.IP, *IPNet, error) {
	ip, n, err := net.ParseCIDR(s)

	return ip, newNetwork(n), err
}

// Contains reports whether the network includes ip
func (n *IPNet) Contains(ip net.IP) bool {
	return n.network.Contains(ip)
}

// Network returns the address's network name, "ip+net"
func (n *IPNet) Network() string {
	return n.network.String()
}

// StartAddress returns the start address of the network
func (n *IPNet) StartAddress() net.IP {
	return n.IP()
}

// EndAddress returns the broadcast address of a network
func (n *IPNet) EndAddress() net.IP {
	out := make([]byte, len(n.network.Mask))

	for i, b := range n.network.Mask {
		out[i] = (0xff ^ b) | n.network.IP[i]
	}

	return out
}

// IP returns the network ip address and is the same es StartAddress
func (n *IPNet) IP() net.IP {
	return n.network.IP
}

// Mask returns the network mask
func (n *IPNet) Mask() net.IPMask {
	return n.network.Mask
}

// IsIPv4 returns true if the network is an IPv4 network otherwise false
func (n *IPNet) IsIPv4() bool {
	return (n.ipLenght == net.IPv4len)
}

// IsIPv6 returns true if the network is an IPv6 network otherwise false
func (n *IPNet) IsIPv6() bool {
	return (n.ipLenght == net.IPv6len)
}

// String returns the CIDR notation of n like "192.168.100.1/24"
// or "2001:DB8::/48" as defined in RFC 4632 and RFC 4291.
// If the mask is not in the canonical form, it returns the
// string which consists of an IP address, followed by a slash
// character and a mask expressed as hexadecimal form with no
// punctuation like "192.168.100.1/c000ff00".
func (n *IPNet) String() string {
	return n.network.String()
}
