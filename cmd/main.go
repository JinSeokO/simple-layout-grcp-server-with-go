package main

import (
	"log"
	"youtube/internal/adapters/app/api"
	"youtube/internal/adapters/core/arithmetic"
	"youtube/internal/adapters/framework/right/db"
	"youtube/internal/ports"

	gRPC "youtube/internal/adapters/framework/left/grpc"
)

func main() {
	var err error

	// "localhost:3307", "root", "adfaie83ma", "golang-test"
	var dbaseAdapter ports.DBPort
	var core ports.ArithmeticPort
	var appAdapter ports.APIPort
	var gRPCAdapter ports.GRPCPort
	//dbaseDriver := os.Getenv("DB_DRIVER")
	//dbSourceName := os.Getenv("DB_Name")
	dbaseDriver := "mysql"
	dbSourceName := "root:adfaie83ma@tcp(localhost:3307)/golang-test?parseTime=true"
	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dbSourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()
	appAdapter = api.NewAdapter(dbaseAdapter, core)

	gRPCAdapter = gRPC.NewAdapter(appAdapter)

	gRPCAdapter.RUN()
}
