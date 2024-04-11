package ping

import (
	"net"
	"time"

	"golang.org/x/net/ipv4"
	"golang.org/x/net/ipv6"
)

func ListenPacket(network, address, sourceInterface string) (*PacketConn, error) {
	return nil, nil
}

var _ net.PacketConn = &PacketConn{}

// A PacketConn represents a packet network endpoint that uses either
// ICMPv4 or ICMPv6.
type PacketConn struct {
	c  net.PacketConn
	p4 *ipv4.PacketConn
	p6 *ipv6.PacketConn
}

func (c *PacketConn) ok() bool { return c != nil && c.c != nil }

// IPv4PacketConn returns the ipv4.PacketConn of c.
func (c *PacketConn) IPv4PacketConn() *ipv4.PacketConn {
	return nil
}

// IPv6PacketConn returns the ipv6.PacketConn of c.
func (c *PacketConn) IPv6PacketConn() *ipv6.PacketConn {
	return nil
}

// ReadFrom reads an ICMP message from the connection.
func (c *PacketConn) ReadFrom(b []byte) (int, net.Addr, error) {
	return 0, nil, nil
}

// WriteTo writes the ICMP message b to dst.
func (c *PacketConn) WriteTo(b []byte, dst net.Addr) (int, error) {
	return 0, nil
}

// Close closes the endpoint.
func (c *PacketConn) Close() error {
	return nil
}

// LocalAddr returns the local network address.
func (c *PacketConn) LocalAddr() net.Addr {
	return nil
}

// SetDeadline sets the read and write deadlines associated with the
// endpoint.
func (c *PacketConn) SetDeadline(t time.Time) error {
	return nil
}

// SetReadDeadline sets the read deadline associated with the
// endpoint.
func (c *PacketConn) SetReadDeadline(t time.Time) error {
	return nil
}

// SetWriteDeadline sets the write deadline associated with the
// endpoint.
func (c *PacketConn) SetWriteDeadline(t time.Time) error {
	return nil
}
