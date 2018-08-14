package cmd

import (
	"bytes"
	"fmt"
	"github.com/golang/glog"
	"github.com/hasusuf/dnsotls/util"
	"github.com/spf13/cobra"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
)

type DoTClient struct {
	http.Client
	Endpoints []string
}

var (
	debugFlag bool
	baseCmd   *cobra.Command
	doTClient = &DoTClient{
		Endpoints: []string{
			"https://1.0.0.1/dns-query",
			"https://1.1.1.1/dns-query",
		},
	}
)

func NewDnsOtlsCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dnsotls",
		Short: "DNS-over-TLS proxy",
		Run: func(cmd *cobra.Command, args []string) {
			runDnsOtls(cmd)
		},
	}

	cmd.PersistentFlags().BoolVarP(
		&debugFlag,
		"debug",
		"",
		false,
		"Turn on debug logging.")

	addDnsOtlsFlags(cmd)

	cmd.AddCommand(NewCmdVersion())

	return cmd
}

func addDnsOtlsFlags(cmd *cobra.Command) {
	cmd.Flags().String(
		"bind",
		"127.0.0.1",
		"Binding IP address")
	cmd.Flags().Int(
		"port",
		53,
		"Listening port")
}

func (c *DoTClient) RunQuery(query []byte) []byte {
	endpoint := c.Endpoints[rand.Int()%len(c.Endpoints)]
	r, err := c.Client.Post(
		endpoint,
		"application/dns-udpwireformat",
		bytes.NewBuffer(query),
	)
	errorHandler(err)

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	errorHandler(err)

	return body
}

func runDnsOtls(cmd *cobra.Command) {
	bindIp := util.GetFlagString(cmd, "bind")
	listenPort := util.GetFlagInt(cmd, "port")
	source := net.UDPAddr{IP: net.ParseIP(bindIp), Port: listenPort}
	fmt.Printf("Lisning on %s port %d \n", source.IP, source.Port)

	conn, err := net.ListenUDP("udp", &source)
	errorHandler(err)
	defer conn.Close()

	for {
		query := make([]byte, 128)

		n, _, _, addr, err := conn.ReadMsgUDP(query, nil)
		errorHandler(err)

		query = query[:n]

		go func(query []byte, addr *net.UDPAddr) {
			var err error
			resp := doTClient.RunQuery(query)

			_, _, err = conn.WriteMsgUDP(resp, nil, addr)
			errorHandler(err)
		}(query, addr)
	}
}

func Execute() {
	baseCmd = NewDnsOtlsCommand()
	err := baseCmd.Execute()
	errorHandler(err)
}

func errorHandler(err error) {
	if err != nil && debugFlag {
		glog.Errorf("something went wrong: %v", err)
	}
}
