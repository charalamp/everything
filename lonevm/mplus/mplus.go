/* Copyright IBM Corp. All Rights Reserved. SPDX-License-Identifier: Apache-2.0 */


package main

import (

	"encoding/json"
	"fmt"
	"strings"

//	"github.com/hyperledger/fabric/protos/msp"
//	"github.com/hyperledger/fabric/core/chaincode/lib/cid"	
        "github.com/hyperledger/fabric/common/util"
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

type marblePrivateDetails struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"name"`    //the fieldtags are needed to keep case from bouncing around
	Price      string `json:"price"`
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
	case "initMarble":
		//create a new marble
		return t.initMarble(stub, args)
	case "readMarblePrivateDetails":
                return t.readMarblePrivateDetails(stub, args)
	default:
		//error
		fmt.Println("invoke did not find func: " + function)
		return shim.Error("Received unknown function invocation")


	}
}

// ============================================================
// initMarble - create a new marble, store into chaincode state
// ============================================================
func (t *SimpleChaincode) initMarble(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	var price string
	//   0       1        2
	// "asdf", "blue", "bob"

	

/*	// GET ID AND MSPID OF THE INVOKER
	MSPID, err := cid.GetMSPID(stub)
	if err != nil {
		fmt.Printf("Error getting MSP identity: %s\n", err.Error())
	}

	/*InvokerID, err := cid.GetID(stub)
        if err != nil {
                fmt.Printf("Error getting invoker identity: %s\n", err.Error())
        }


	fmt.Println(InvokerID)	



	if  MSPID == "TwoMSP"  {
        price = 2
	} else if MSPID == "OneMSP"  {
        price = 1
	} else {
        return shim.Error("Nor OneMSP or TwoMSP") 
	}

	erro := cid.AssertAttributeValue(stub, "hf.EnrollmentID", "one-admin")
        fmt.Printf("Error invoker: %s\n", erro.Error())

*/

	//if (cid.AssertAttributeValue(stub, "hf.EnrollmentID", "one-admin")) {
        //price = 22
 	//}
	/*if  InvokerID == "two-peer1"  {
        price = 22
        } else if  InvokerID == "one-peer1"  {
        price = 11
        } else {
        return shim.Error("Failed Invoker") 
        } */


	// INVOKE OTHER CHAINCODE
        chaincodeName := "pricecc"
	// Invoke price method without arguments
        chaincodeArgs := util.ToChaincodeArgs("putprice")
        chaincodeChannel := "pricechannel"


        response := stub.InvokeChaincode(chaincodeName, chaincodeArgs, chaincodeChannel)
        if response.Status != shim.OK {
                return shim.Error(fmt.Sprintf("Failed to invoke chaincode:"))
        }
	
	price = string(response.Payload) 
	// GET MARBLE ARGUMENTS
	marbleName := args[0]
	color := strings.ToLower(args[1])
	owner := strings.ToLower(args[2])


	// ==== Create marble object and marshal to JSON ====
	objectType := "marble"
	marble := &marble{objectType, marbleName, color, owner}
	marbleJSONasBytes, err := json.Marshal(marble)
	if err != nil {
		return shim.Error(err.Error())
	}

	//Alternatively, build the marble json string manually if you don't want to use struct marshalling
	//marbleJSONasString := `{"docType":"Marble",  "name": "` + marbleName + `", "color": "` + color + `", "size": ` + strconv.Itoa(size) + `, "owner": "` + owner + `"}`
	//marbleJSONasBytes := []byte(str)

	// === Save marble to state ===
	err = stub.PutState(marbleName, marbleJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

		
	// ==== Create marble private details object with price, marshal to JSON, and save to state ====
	marblePrivateDetails :=  &marblePrivateDetails{objectType, marbleName, price}	
	marblePrivateDetailsBytes, err := json.Marshal(marblePrivateDetails)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	err = stub.PutPrivateData("collectionMarblePrivateDetails", marbleName, marblePrivateDetailsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	
	fmt.Println("- end init marble")
	return shim.Success(nil)
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



func (t *SimpleChaincode) readMarblePrivateDetails(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var name string

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting name of the marble to query")
	}

	name = args[0]
	valAsbytes, err := stub.GetPrivateData("collectionMarblePrivateDetails", name) //get the marble private details from chaincode state
	if err != nil {
                jsonResp := "{\"Error\":\"Failed to get state for " + name + "\"}"
                return shim.Error(jsonResp)
        } else if valAsbytes == nil {
                jsonResp := "{\"Error\":\"Marble does not exist: " + name + "\"}"
                return shim.Error(jsonResp)
        }


	return shim.Success(valAsbytes)
}

