package dns

import (
	"context"
	dnsp "github.com/lilpipidron/Host-DNS-Updater/proto/service/dns"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Service struct {
	dnsp.DNSServiceServer
}

/*GetListDNS(context.Context, *emptypb.Empty) (*DNSListReply, error)
DeleteDNS(context.Context, *DeleteDNSRequest) (*DeleteDNSReply, error)
AddDNS(context.Context, *AddDNSRequest) (*AddDNSReply, error)*/

func (s *Service) GetListDNS(ctx context.Context, *emptypb.Empty) (*DNSListReply, error)
