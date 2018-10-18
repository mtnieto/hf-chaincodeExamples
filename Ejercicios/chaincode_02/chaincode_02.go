/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package example02

import (
	"fmt"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
	
}

type Person struct {
	Name string  `json:"Name"`
	Age string `json:"Age"`
	Address string `json:"Address"`
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Init")
	
	}

	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {


	return shim.Error("Invalid invoke function name")
}

// Transaction makes payment of X units from A to B
func (t *SimpleChaincode) modifyAddress(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	

	return shim.Success(nil)
}


// query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	

	return shim.Success(storedValue)
}