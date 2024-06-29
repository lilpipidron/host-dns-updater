package hostname

import (
	"context"
	"github.com/charmbracelet/log"
	hnp "github.com/lilpipidron/Host-DNS-Updater/proto/service/hostname"
	"os"
	"os/exec"
)

type Service struct {
	hnp.HostnameServiceServer
}

func (s *Service) SetHostName(ctx context.Context, req *hnp.SetHostnameRequest) (*hnp.SetHostnameReply, error) {
	log.Info("Called SetHostName", "hostname", req.Hostname)

	cmd := exec.Command("hostnamectl", "set-hostname", req.Hostname)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Error("Failed change hostname", "error", err)
		return &hnp.SetHostnameReply{ErrorMesage: err.Error()}, err
	}

	log.Info("Successfully change hostname", "hostname", req.Hostname)

	return &hnp.SetHostnameReply{ErrorMesage: ""}, nil
}
