package cmd

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/lilpipidron/Host-DNS-Updater/proto/service/dns"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func addDNS(newDNS string) error {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("did not connect: %v", err)
	}
	defer conn.Close()

	dnsClient := dns.NewDNSServiceClient(conn)
	_, err = dnsClient.AddDNS(context.Background(), &dns.AddDNSRequest{Dns: newDNS})

	return err
}
