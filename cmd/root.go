package cmd

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/go-ping/ping"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "goping [host]",
	Short: "helpful ping tool written in Go",
	Long:  `goping is a helpful ping tool written in Go.`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		host := args[0]
		pinger, err := ping.NewPinger(host)
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			os.Exit(1)
		}
		count, _ := cmd.Flags().GetInt("count")
		pinger.Count = count

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for range c {
				pinger.Stop()
			}
		}()

		pinger.OnRecv = func(pkt *ping.Packet) {
			fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
				pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
		}

		pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
			fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n",
				pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
		}

		pinger.OnFinish = func(stats *ping.Statistics) {
			fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
			fmt.Printf("%d packets transmitted, %d packets received, %v%% packet loss\n",
				stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
			fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n",
				stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
		}

		fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
		err = pinger.Run()
		if err != nil {
			fmt.Printf("ERROR: %s\n", err.Error())
			os.Exit(1)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntP("count", "c", 4, "Number of pings to send")
}
