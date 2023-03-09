package main

import (
	gRPCServerApi "gRPC-test-server/simpleGrpcServer/go_grpc_api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

// InitGrpcServer - Set up and start Backend gRPC-server
func (gRPCServerObject *gRPCServerObjectStruct) InitGrpcServer() {

	var err error

	// Find first non allocated port from defined start port
	logger.WithFields(logrus.Fields{
		"Id": "054bc0ef-93bb-4b75-8630-74e3823f71da",
	}).Debug("Backend Server tries to start")

	logger.WithFields(logrus.Fields{
		"Id":               "ca3593b1-466b-4536-be91-5e038de178f4",
		"gRPCServerPort: ": gRPCServerPort,
	}).Debug("Start listening on:")
	lis, err = net.Listen("tcp", ":"+gRPCServerPort)

	if err != nil {
		logger.WithFields(logrus.Fields{
			"Id":    "ad7815b3-63e8-4ab1-9d4a-987d9bd94c76",
			"err: ": err,
		}).Error("failed to listen:")
	} else {
		logger.WithFields(logrus.Fields{
			"Id":               "ba070b9b-5d57-4c0a-ab4c-a76247a50fd3",
			"gRPCServerPort: ": gRPCServerPort,
		}).Info("Success in listening on port:")

	}

	// Creates a new RegisterWorkerServer gRPC server
	//go func() {
	logger.WithFields(logrus.Fields{
		"Id": "b0ccffb5-4367-464c-a3bc-460cafed16cb",
	}).Debug("Starting Backend gRPC Server")

	registerGrpcServerGrpcServer = grpc.NewServer()
	gRPCServerApi.RegisterSimpleGrpcServerServicesServer(registerGrpcServerGrpcServer, &gRPCServerGrpcServicesServer{})

	// Register RouteGuide on the same server.
	reflection.Register(registerGrpcServerGrpcServer)

	logger.WithFields(logrus.Fields{
		"Id":               "e843ece9-b707-4c60-b1d8-14464305e68f",
		"gRPCServerPort: ": gRPCServerPort,
	}).Info("registerGrpcServerGrpcServer for TestExecution Server started")
	registerGrpcServerGrpcServer.Serve(lis)

}

// StopGrpcServer - Stop Backend gRPC-server
func (gRPCServerObject *gRPCServerObjectStruct) StopGrpcServer() {

	logger.WithFields(logrus.Fields{}).Info("Gracefully stop for: registerGrpcServerGrpcServer")
	registerGrpcServerGrpcServer.GracefulStop()

	logger.WithFields(logrus.Fields{
		"gRPCServerPort: ": gRPCServerPort,
	}).Info("Close net.Listing")
	_ = lis.Close()

}
