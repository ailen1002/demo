package controller

import (
	"context"
	"demo/api/backend"
	"demo/internal/service"
)

var (
	Portlist = cPortlist{}
)

type cPortlist struct{}

func (a *cPortlist) Postlist(ctx context.Context, req *backend.PortlistReq) (res *backend.PortlistRes, err error) {
	out := service.Portlist()

	if err != nil {
		return nil, err
	}

	return &backend.PortlistRes{Lists: out}, nil
}
