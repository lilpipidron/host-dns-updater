package cmd

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/lilpipidron/Host-DNS-Updater/proto/service/dns"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func getDNSList() ([]string, error) {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("did not connect: %v", err)
		return nil, err
	}
	defer conn.Close()
	dnsClient := dns.NewDNSServiceClient(conn)
	list, err := dnsClient.GetListDNS(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, err
	}
	return list.DnsList, nil
}
