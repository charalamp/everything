/*
Copyright IBM Corp. All Rights Reserved.
SPDX-License-Identifier: Apache-2.0
*/


package main

import (

	"fmt"

//	"github.com/hyperledger/fabric/protos/msp"
	"github.com/hyperledger/fabric/core/chaincode/lib/cid"	
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
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
	case "putprice":
		//create a new marble
		return t.price(stub, args)
	default:
		//error
		fmt.Println("invoke did not find func: " + function)
		return shim.Error("Received unknown function invocation")


	}
}

// ============================================================
// initMarble - create a new marble, store into chaincode state
// ============================================================
func (t *SimpleChaincode) price(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var price string
	//   0       1        2
	// "asdf", "blue", "bob"

	

	// GET ID AND MSPID OF THE INVOKER
	MSPID, err := cid.GetMSPID(stub)
	if err != nil {
		fmt.Printf("Error getting MSP identity: %s\n", err.Error())
	}


	if  MSPID == "TwoMSP"  {
        price = "2"
	} else if MSPID == "OneMSP"  {
        price = "1"
	} else {
        return shim.Error("Nor OneMSP or TwoMSP") 
	}

	err1 := cid.AssertAttributeValue(stub, "hf.EnrollmentID", "one-admin")

	if err1 == nil {
        price = "10"
 	}

        err2 := cid.AssertAttributeValue(stub, "hf.EnrollmentID", "two-admin")
        if err2 == nil {
        price = "20"
        }

	priceAsBytes := []byte(price)
	fmt.Println("- end init marble")
	return shim.Success(priceAsBytes)

}





