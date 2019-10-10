/*
Copyright IBM Corp. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/


package main

import (

	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}


type marble struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"name"`    //the fieldtags are needed to keep case from bouncing around
	Color      string `json:"color"`
	Owner      string `json:"owner"`
}


// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}


// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	switch function {
	case "readMarble":
		//read a marble
		return t.readMarble(stub, args)
	default:
		//error
		fmt.Println("invoke did not find func: " + function)
		return shim.Error("Received unknown function invocation")


	}
}




func (t *SimpleChaincode) readMarble(stub shim.ChaincodeStubInterface, args []string) pb.Response {
        var name string

        if len(args) != 1 {
                return shim.Error("Incorrect number of arguments. Expecting name of the marble to query")
        }

        name = args[0]
        valAsbytes, err := stub.GetState(name) //get the marble from chaincode state
        if err != nil {
                jsonResp := "{\"Error\":\"Failed to get state for " + name + "\"}"
                return shim.Error(jsonResp)
        } else if valAsbytes == nil {
                jsonResp := "{\"Error\":\"Marble does not exist: " + name + "\"}"
                return shim.Error(jsonResp)
        }

        return shim.Success(valAsbytes)
}

