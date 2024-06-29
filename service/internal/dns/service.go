package dns

import (
	"bufio"
	"context"
	"github.com/charmbracelet/log"
	dnsp "github.com/lilpipidron/Host-DNS-Updater/proto/service/dns"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
	"strings"
)

type Service struct {
	dnsp.DNSServiceServer
}

/*
DeleteDNS(context.Context, *DeleteDNSRequest) (*DeleteDNSReply, error)
*/

func (s *Service) GetListDNS(ctx context.Context, empty *emptypb.Empty) (*dnsp.DNSListReply, error) {
	file, err := os.Open("/etc/resolv.conf")
	if err != nil {
		log.Error("Failed open file", "error", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var list []string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "nameserver") {
			fields := strings.Fields(line)
			if len(fields) > 1 {
				list = append(list, fields[1])
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error("Failed read file", "error", err)
		return nil, err
	}

	return &dnsp.DNSListReply{DnsList: list}, nil
}

func (s *Service) AddDNS(ctx context.Context, req *dnsp.AddDNSRequest) (*dnsp.AddDNSReply, error) {
	file, err := os.OpenFile("/etc/resolv.conf", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Error("Failed open file", "error", err)
		return &dnsp.AddDNSReply{ErrorMessage: err.Error()}, err
	}
	defer file.Close()

	_, err = file.WriteString("nameserver " + req.Dns + "\n")
	if err != nil {
		log.Error("Failed write file", "error", err)
		return &dnsp.AddDNSReply{ErrorMessage: err.Error()}, err
	}

	return &dnsp.AddDNSReply{ErrorMessage: ""}, nil
}
