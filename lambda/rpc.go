// Copyright 2020 Amazon.com, Inc. or its affiliates. All Rights Reserved

//go:build !lambda.norpc
// +build !lambda.norpc

package lambda

import (
	"errors"
	"log"
	"net"
	"net/rpc"
)

func init() {
	// Register `startFunctionRPC` to be run if the _LAMBDA_SERVER_PORT environment variable is set.
	// This happens when the runtime for the function is configured as `go1.x`.
	// The value of the environment variable will be passed as the first argument to `startFunctionRPC`.
	rpcStartFunction.f = startFunctionRPC
}

func startFunctionRPC(port string, handler Handler) error {
	lis, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	err = rpc.Register(NewFunction(handler))
	if err != nil {
		log.Fatal("failed to register handler function")
	}
	rpc.Accept(lis)
	return errors.New("accept should not have returned")
}
