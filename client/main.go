package main

import (
	"github.com/charmbracelet/log"
	"github.com/lilpipidron/Host-DNS-Updater/client/cmd"
	"github.com/lilpipidron/Host-DNS-Updater/proto/service/dns"
	"github.com/lilpipidron/Host-DNS-Updater/proto/service/hostname"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var dnsClient *dns.DNSServiceClient
var hostnameClient *hostname.HostnameServiceClient

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("did not connect: %v", err)
	}
	defer conn.Close()
	dnsClient := dns.NewDNSServiceClient(conn)
	hostnameClient := hostname.NewHostnameServiceClient(conn)

	cmd.Execute()
}
