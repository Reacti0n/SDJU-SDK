package vpn

import (
	"github.com/Reacti0n/SDJU-SDK/vpn/config"
	"github.com/Reacti0n/SDJU-SDK/vpn/context"
)

// VPN 连接内网相关 API
type VPN struct {
	ctx *context.Context
}

// NewVPN 实例化 VPN 连接
func (v *VPN) NewVPN(cfg *config.Config)*VPN  {
	ctx := &context.Context{Config: cfg}
	return &VPN{ctx: ctx}
}

// GetContext get Context
func (v *VPN) GetContext()*context.Context  {
	return v.ctx
}