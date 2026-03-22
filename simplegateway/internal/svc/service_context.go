package svc

import (
	"github.com/polpo666/pzero/core/configcenter"

	"simplegateway/internal/config"
)

type ServiceContext struct {
	ConfigCenter configcenter.ConfigCenter[config.Config]
}

func NewServiceContext(cc configcenter.ConfigCenter[config.Config]) *ServiceContext {
	svcCtx := &ServiceContext{
		ConfigCenter: cc,
	}

	svcCtx.SetConfigListener()
	return svcCtx
}
