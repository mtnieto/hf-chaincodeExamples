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
	return shim.Success(nil)
}

func (t *HelloWorldCC) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

}

// Transaction makes payment of X units from A to B
func (t *HelloWorldCC) greet(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	
	return shim.Success()
}

// Deletes an entity from state
func (t *HelloWorldCC) setGreet(stub shim.ChaincodeStubInterface, args []string) pb.Response {


	return shim.Success(nil)
}

