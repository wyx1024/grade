package znet

import "go-growth/ziface"

type BaseRouter struct {}

func(r *BaseRouter) PreHeadler(req ziface.IRequeset) {}

func(r *BaseRouter)Headler(ziface.IRequeset){}

func(r *BaseRouter)PostHeadler(ziface.IRequeset){}