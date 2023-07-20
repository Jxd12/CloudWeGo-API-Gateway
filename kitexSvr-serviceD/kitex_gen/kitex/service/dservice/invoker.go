// Code generated by Kitex v0.6.1. DO NOT EDIT.

package dservice

import (
	server "github.com/cloudwego/kitex/server"
	service "kitexSvr-serviceD/kitex_gen/kitex/service"
)

// NewInvoker creates a server.Invoker with the given handler and options.
func NewInvoker(handler service.DService, opts ...server.Option) server.Invoker {
	var options []server.Option

	options = append(options, opts...)

	s := server.NewInvoker(options...)
	if err := s.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	if err := s.Init(); err != nil {
		panic(err)
	}
	return s
}
