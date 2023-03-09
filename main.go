package main

import (
	"fmt"
	uuidGenerator "github.com/google/uuid"
	"log"
	"os"
	"strconv"
)

func main() {

	// Set up logging
	initLogger()

	// Stop gRPC-server before exiting
	defer gRPCServerObject.StopGrpcServer()

	// Start gRPC-server
	gRPCServerObject.InitGrpcServer()

}

// mustGetEnv is a helper function for getting environment variables.
// Displays a warning if the environment variable is not set.
func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func init() {

	// Create Unique Uuid for run time instance
	applicationRuntimeUuid = uuidGenerator.New().String()
	fmt.Println("ApplicationRuntimeUuid: " + applicationRuntimeUuid)

	var err error

	// Get Environment variable to tell were gRPC-Server is running
	var environmentVariableExecutionLocationForGrpcServer = mustGetenv("ExecutionLocationForGrpcServer")

	switch environmentVariableExecutionLocationForGrpcServer {
	case "LOCALHOST_NODOCKER":
		executionLocationForGrpcServer = LocalhostNoDocker

	case "LOCALHOST_DOCKER":
		executionLocationForGrpcServer = LocalhostDocker

	case "GCP":
		executionLocationForGrpcServer = GCP

	default:
		fmt.Println("Unknown Execution location for gRPC-Server: " + environmentVariableExecutionLocationForGrpcServer + ". Expected one of the following: 'LOCALHOST_NODOCKER', 'LOCALHOST_DOCKER', 'GCP'")
		os.Exit(0)
	}

	// Port for gRPC-Server
	_, err = strconv.Atoi(mustGetenv("GrpcServerPort"))
	if err != nil {
		fmt.Println("Couldn't convert environment variable 'GrpcServerPort' to an integer, error: ", err)
		os.Exit(0)
	}

	// We now know that port is an integer
	gRPCServerPort = mustGetenv("GrpcServerPort")

}
