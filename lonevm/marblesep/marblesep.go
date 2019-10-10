

package main

import (
	"fmt"
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
	Size       int    `json:"size"`
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

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "initMarble" { //create a new marble
		return t.initMarble(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initMarble - create a new marble, store into chaincode state
// ============================================================
func (t *SimpleChaincode) initMarble(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0       1       2     3
	// "asdf", "blue", "35", "bob"
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// ==== Input sanitation ====
	fmt.Println("- start init marble")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty string")
	}


	chaincodeName := "marbles-one"
        chaincodeArgs := util.ToChaincodeArgs("initMarble", args[0], args[1], args[2], args[3])
        chaincodeChannel := "onetwochannel"


        response := stub.InvokeChaincode(chaincodeName, chaincodeArgs, chaincodeChannel)
        if response.Status != shim.OK {
                return shim.Error(fmt.Sprintf("Failed to invoke chaincode:"))
        }
        return shim.Success(nil)        



}











//	"github.com/hyperledger/fabric/common/util"

// ============================================================
// initMarble - create a new marble, store into chaincode state
// ============================================================
/* func (t *SimpleChaincode) initMarble(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//   0       1       2     3
	// "asdf", "blue", "35", "bob"
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	chaincodeName := "marbles-one"
	chaincodeArgs := toChaincodeArgs("initMarble", args[0], args[1], args[2], args[3])
	chaincodeChannel := "onetwochannel"

	
        response := stub.InvokeChaincode(chaincodeName, chaincodeArgs, chaincodeChannel)
        if response.Status != shim.OK {
	        return shim.Error(fmt.Sprintf("Failed to invoke chaincode:"))
    	}
	return shim.Success(nil)	
}*/

