# go-ping-vrf

*The popular ICMP Ping library for Go, modified to work when using vrf (Virtual Routing and Forwarding) in Linux*

Original: [https://github.com/go-ping/ping]()

Only Linux, only ICMP (meaning no UDP). Should work for IPv4 and IPv6.

```shell
# build
go build cmd/ping/ping.go

# ping 10.0.0.2 using the interface vrf-priv1 with the ip 10.0.0.2 continuously
ping -src 10.0.0.3 -if vrf-priv1 --privileged 10.0.0.2
```
