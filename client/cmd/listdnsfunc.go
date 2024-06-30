package cmd

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lilpipidron/Host-DNS-Updater/proto/service/dns"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getDNSList() []string {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("did not connect: %v", err)
	}
	defer conn.Close()
	dnsClient := dns.NewDNSServiceClient(conn)
	list, err := dnsClient.GetListDNS(context.Background(), &empty.Empty{})
	if err != nil {
		log.Error("Failed to get DNS List", "error", err)
	}

	return list.DnsList
}
