package main

import (
	"b2b-reputation-system/blockchain"
	"fmt"
	"os"
)

func main() {
	// Definition of the Fabric SDK properties
	fSetup := blockchain.FabricSetup{
		// Network parameters
		OrdererID: "orderer.bom2buy-chain.com",

		// Channel parameters
		ChannelID:     "b2bchannel",
		ChannelConfig: os.Getenv("GOPATH") + "/src/b2b-reputation-system/fixtures/artifacts/b2b.channel.tx",

		// Chaincode parameters
		ChainCodeID:     "b2b-reputation-system-test",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath:   "b2b-reputation-system/chaincode/",
		OrgAdmin:        "Admin",
		OrgName:         "Bom2buy",
		ConfigFile:      "config.yaml",

		// User parameters
		UserName: "User1",
	}

	// Initialization of the Fabric SDK from the previously set properties
	err := fSetup.Initialize()
	if err != nil {
		fmt.Printf("Unable to initialize the Fabric SDK: %v\n", err)
		return
	}

	// Install and instantiate the chaincode
	err = fSetup.InstallAndInstantiateCC()
	if err != nil {
		fmt.Printf("Unable to install and instantiate the chaincode: %v\n", err)
		return
	}

	// Query the chaincode
	response, err := fSetup.QueryHello()
	if err != nil {
		fmt.Printf("Unable to query hello on the chaincode: %v\n", err)
	} else {
		fmt.Printf("Response from the query hello: %s\n", response)
	}

	// Invoke the chaincode
	txId, err := fSetup.InvokeHello("b2b-reputation-system-test")
	if err != nil {
		fmt.Printf("Unable to invoke hello on the chaincode: %v\n", err)
	} else {
		fmt.Printf("Successfully invoke hello, transaction ID: %s\n", txId)
	}

	// Query again the chaincode
	response, err = fSetup.QueryHello()
	if err != nil {
		fmt.Printf("Unable to query hello on the chaincode: %v\n", err)
	} else {
		fmt.Printf("Response from the query hello: %s\n", response)
	}
	// Close SDK
	defer fSetup.CloseSDK()
}
