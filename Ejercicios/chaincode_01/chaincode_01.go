/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package bankChaincode

import (
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type BankChaincode struct {
}

func (t *BankChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	
}

func (t *BankChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "" {
		// Make payment of X units from A to B
		return t.initAccount(stub, args)
	} else if function == "" {
		// Deletes an entity from its state
		return t.withdraw(stub, args)
	} else if function == "" {
		// the old "Query" is now implemtned in invoke
		return t.sendMoney(stub, args)
	} else if function == "" {
		// the old "Query" is now implemtned in invoke
		return t.getBalance(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
}

// Init account
func (t *BankChaincode) initAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
}

// withdraw money
func (t *BankChaincode) withdraw(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
}

// query callback representing the query of a chaincode
func (t *BankChaincode) getBalance(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var account string // Entities
	
}

func (t *BankChaincode) sendMoney(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
}