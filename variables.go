package main

import (
	gRPCServerApi "gRPC-test-server/simpleGrpcServer/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type gRPCServerObjectStruct struct {
	logger *logrus.Logger
}

// Variable holding everything together
var gRPCServerObject *gRPCServerObjectStruct

// gRPC variables
var (
	registerGrpcServerGrpcServer *grpc.Server
	lis                          net.Listener
)

// gRPC Server used for register clients Name, Ip and Por and Clients Test Enviroments and Clients Test Commandst
type gRPCServerGrpcServicesServer struct {
	gRPCServerApi.UnimplementedSimpleGrpcServerServicesServer
}

// executionLocationTypeType
// Definitions for where gRPC-server is running
type executionLocationTypeType int

const (
	LocalhostNoDocker executionLocationTypeType = iota
	LocalhostDocker
	GCP
)

// Environment variables
var (
	executionLocationForGrpcServer executionLocationTypeType
	gRPCServerPort                 string
)

// Unique ID for the application
var applicationRuntimeUuid string
