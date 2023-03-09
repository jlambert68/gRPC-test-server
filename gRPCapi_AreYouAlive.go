package main

import (
	gRPCServerApi "gRPC-test-server/simpleGrpcServer/go_grpc_api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	time "time"
)

// AreYouAlive - *********************************************************************
// Anyone can check if Fenix TestCase Builder server is alive with this service
func (s *gRPCServerGrpcServicesServer) AreYouAlive(ctx context.Context, emptyParameter *gRPCServerApi.EmptyParameter) (*gRPCServerApi.AckNackResponse, error) {

	logger.WithFields(logrus.Fields{
		"id":                                 "1ff67695-9a8b-4821-811d-0ab8d33c4d8b",
		"emptyParameter.MessageToGrpcServer": emptyParameter.MessageToGrpcServer,
	}).Debug("Incoming 'gRPC - AreYouAlive'")

	defer logger.WithFields(logrus.Fields{
		"id": "9c7f0c3d-7e9f-4c91-934e-8d7a22926d84",
	}).Debug("Outgoing 'gRPC - AreYouAlive'")

	time := time.Now().String()

	logger.WithFields(logrus.Fields{
		"id":   "a36a4003-3340-41dd-8b70-d9acf052f415",
		"time": time,
	}).Debug("Time sent back to gRPC-client")

	return &gRPCServerApi.AckNackResponse{AckNack: true, Comments: "I'm alive and my time is: " + time +
			". My 'ApplicationRuntimeUuid' is: '" + applicationRuntimeUuid +
			"'. The message you sent to me was: '" + emptyParameter.MessageToGrpcServer + "'"},
		nil
}
