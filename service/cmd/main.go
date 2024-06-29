package main

import (
	dnsp "github.com/lilpipidron/Host-DNS-Updater/proto/service/dns"
	hnp "github.com/lilpipidron/Host-DNS-Updater/proto/service/hostname"
	"github.com/lilpipidron/Host-DNS-Updater/service/internal/dns"
	"github.com/lilpipidron/Host-DNS-Updater/service/internal/hostname"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	hnp.RegisterHostnameServiceServer(s, &hostname.Service{})
	dnsp.RegisterDNSServiceServer(s, &dns.Service{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
