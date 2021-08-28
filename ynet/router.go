package ynet

import "yinx/yiface"

type BaseRouter struct{}

func (router *BaseRouter) PreHandle(request yiface.IRequest) {}

func (router *BaseRouter) Handle(request yiface.IRequest) {}

func (router *BaseRouter) PostHandle(request yiface.IRequest) {}
