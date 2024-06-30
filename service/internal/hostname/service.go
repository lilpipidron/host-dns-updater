package hostname

import (
	"context"
	"github.com/charmbracelet/log"
	hnp "github.com/lilpipidron/Host-DNS-Updater/proto/service/hostname"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
	"os/exec"
)

type Service struct {
	hnp.HostnameServiceServer
}

func (s *Service) SetHostName(ctx context.Context, req *hnp.SetHostnameRequest) (*emptypb.Empty, error) {
	log.Info("Called SetHostName", "hostname", req.Hostname)

	cmd := exec.Command("hostnamectl", "set-hostname", req.Hostname)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Error("Failed change hostname", "error", err)
		return &emptypb.Empty{}, err
	}

	log.Info("Successfully change hostname", "hostname", req.Hostname)

	return &emptypb.Empty{}, nil
}
