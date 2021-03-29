package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"time"

	ping "github.com/geberl/go-ping-vrf"
)

var usage = `
Usage:

    ping [-c count] [-i interval] [-t timeout] [--privileged] [-src ip] [-if interface] host

Examples:

    # ping 10.0.0.2 using the interface vrf-priv1 with the ip 10.0.0.2 continuously
    ping -src 10.0.0.3 -if vrf-priv1 --privileged 10.0.0.2

	# ping google continuously
    ping www.google.com

    # ping google 5 times
    ping -c 5 www.google.com

    # ping google 5 times at 500ms intervals
    ping -c 5 -i 500ms www.google.com

    # ping google for 10 seconds
    ping -t 10s www.google.com

    # Send a privileged raw ICMP ping
    sudo ping --privileged www.google.com

    # Send ICMP messages with a 100-byte payload
    ping -s 100 1.1.1.1
`

func main() {
	timeout := flag.Duration("t", time.Second*100000, "")
	interval := flag.Duration("i", time.Second, "")
	count := flag.Int("c", -1, "")
	size := flag.Int("s", 16, "")
	privileged := flag.Bool("privileged", false, "")
	source := flag.String("src", "", "")
	sourceInterface := flag.String("if", "", "")
	flag.Usage = func() {
		fmt.Print(usage)
	}
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		return
	}

	host := flag.Arg(0)
	pinger, err := ping.NewPinger(host)
	pinger.Source = *source
	pinger.SourceInterface = *sourceInterface
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}

	// listen for ctrl-C signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			pinger.Stop()
		}
	}()

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}
	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
			pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %d duplicates, %v%% packet loss\n",
			stats.PacketsSent, stats.PacketsRecv, stats.PacketsRecvDuplicates, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
			stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	pinger.Count = *count
	pinger.Size = *size
	pinger.Interval = *interval
	pinger.Timeout = *timeout
	pinger.SetPrivileged(*privileged)

	fmt.Printf("PING %s (%s) from %s (%s):\n", pinger.Addr(), pinger.IPAddr(), pinger.Source, pinger.SourceInterface)
	err = pinger.Run()
	if err != nil {
		fmt.Printf("Failed to ping target host: %s", err)
	}
}
