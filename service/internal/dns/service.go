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

	log.Info("Successfully get dns list")

	return &dnsp.DNSListReply{DnsList: list}, nil
}

func (s *Service) AddDNS(ctx context.Context, req *dnsp.AddDNSRequest) (*emptypb.Empty, error) {
	file, err := os.OpenFile("/etc/resolv.conf", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Error("Failed open file", "error", err)
		return &emptypb.Empty{}, err
	}
	defer file.Close()

	_, err = file.WriteString("nameserver " + req.Dns + "\n")
	if err != nil {
		log.Error("Failed write file", "error", err)
		return &emptypb.Empty{}, err
	}

	log.Info("Successfully added DNS", "dns", req.Dns)

	return &emptypb.Empty{}, nil
}

func (s *Service) DeleteDNS(ctx context.Context, req *dnsp.DeleteDNSRequest) (*emptypb.Empty, error) {
	file, err := os.Open("/etc/resolv.conf")
	if err != nil {
		log.Error("Failed open file", "error", err)
		return &emptypb.Empty{}, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if !strings.Contains(line, req.Dns) {
			lines = append(lines, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error("Failed read file", "error", err)
		return &emptypb.Empty{}, err
	}

	outputFile, err := os.OpenFile("/etc/resolv.conf", os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		log.Error("Failed open file", "error", err)
		return &emptypb.Empty{}, err
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			log.Error("Failed write file", "error", err)
			return &emptypb.Empty{}, err
		}
	}

	if err := writer.Flush(); err != nil {
		log.Error("Failed flush", "error", err)
		return &emptypb.Empty{}, err
	}

	log.Info("Successfully deleted DNS", "dns", req.Dns)
	return &emptypb.Empty{}, nil
}
