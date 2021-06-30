package main

import (
	"log"

	"github.com/LunaYin/StockUpdate/stockupdate"
	"github.com/cloudstateio/go-support/cloudstate"
	"github.com/cloudstateio/go-support/cloudstate/crdt"
	"github.com/cloudstateio/go-support/cloudstate/protocol"
)

func main() {
	server, err := cloudstate.New(protocol.Config{
		ServiceName:    "stockupdate.StockUpdateService",
		ServiceVersion: "0.1.0",
	})
	if err != nil {
		log.Fatalf("cloudstate.New failed: %v", err)
	}
	err = server.RegisterCRDT(&crdt.Entity{
		ServiceName: "stockupdate.StockUpdateService",
		EntityFunc:  stockupdate.NewStock,
	}, protocol.DescriptorConfig{
		Service: "stockupdate.proto",
	})
	if err != nil {
		log.Fatalf("Cloudstate failed to register entity: %v", err)
	}
	err = server.Run()
	if err != nil {
		log.Fatalf("Cloudstate failed to run: %v", err)
	}
}
func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
}
