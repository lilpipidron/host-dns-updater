package cmd

import (
	"context"
	"github.com/lilpipidron/Host-DNS-Updater/proto/service/hostname"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func changeHostname(newHostname string) error {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer conn.Close()
	hostnameClient := hostname.NewHostnameServiceClient(conn)
	_, err = hostnameClient.SetHostName(context.Background(), &hostname.SetHostnameRequest{Hostname: newHostname})
	return err
}
