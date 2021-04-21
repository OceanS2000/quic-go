// +build linux

package quic

import "golang.org/x/sys/unix"

const msgTypeIPTOS = unix.IP_TOS

const (
	ipv4RECVPKTINFO = unix.IP_PKTINFO
	ipv6RECVPKTINFO = unix.IPV6_RECVPKTINFO
)

const (
	msgTypeIPv4PKTINFO = unix.IP_PKTINFO
	msgTypeIPv6PKTINFO = unix.IPV6_PKTINFO
)

const dontFragment = unix.IP_DONTFRAG
