package education

import (
	"github.com/Reacti0n/SDJU-SDK/education/config"
	"github.com/Reacti0n/SDJU-SDK/education/context"
)

// Education 教务系统相关 API
type Education struct {
	ctx *context.Context
}

// NewEducation 实例化教务连接
func NewEducation(cfg *config.Config)*Education  {
	ctx := &context.Context{Config: cfg}
	return &Education{ctx: ctx}
}

// GetContext get Context
func (e *Education) GetContext()*context.Context  {
	return e.ctx
}