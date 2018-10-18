/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package example00

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// HelloWorldCC example simple Chaincode implementation
type HelloWorldCC struct {
}

func (t *HelloWorldCC) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	_, args := stub.GetFunctionAndParameters()
	var A, greet string    // Entities
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	A = args[0]
	greet = args[1]
	
	err = stub.PutState(A, []byte(greet))
	if err != nil {
		return shim.Error(err.Error())
	}


	return shim.Success(nil)
}

func (t *HelloWorldCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "greet" {
		// Make payment of X units from A to B
		return t.greet(stub, args)
	} else if function == "setGreet" {
		// Deletes an entity from its state
		return t.setGreet(stub, args)
	}
	return shim.Error("Invalid invoke function name. Expecting \"greet\" \"setGreet\" ")
}

// Transaction makes payment of X units from A to B
func (t *HelloWorldCC) greet(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var A string    // Entities

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	A = args[0]
	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	greet, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	fmt.Println("Greet: ", string(greet))
	return shim.Success(greet)
}

// Deletes an entity from state
func (t *HelloWorldCC) setGreet(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]
	greet := args[1]

	oldGreet, err := stub.GetState(A)
	// Delete the key from the state in ledger
	if err != nil {
		return shim.Error("Failed to get state")
	}
	fmt.Println("The old greet was " + string(oldGreet))

	err = stub.PutState(A, []byte(greet))
	return shim.Success(nil)
}

