/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Accountinfo 账户信息
type Accountinfo struct {
	Name     string
	Time     int64
	Ballance int
}

// Init callback representing the invocation of a chaincode
// This chaincode will manage two accounts A and B and will transfer X units from A to B upon invoke
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	var err error
	_, args := stub.GetFunctionAndParameters()
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	// Initialize the chaincode
	A := args[0]
	Aval, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	B := args[2]
	Bval, err := strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("Expecting integer value for asset holding")
	}
	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

	aAccount := &Accountinfo{}
	aAccount.Name = A
	aAccount.Time = time.Now().Unix()
	aAccount.Ballance = Aval

	bAccount := &Accountinfo{}
	bAccount.Name = B
	bAccount.Time = time.Now().Unix()
	bAccount.Ballance = Bval

	// Write the state to the ledger
	aAccountByte, _ := json.Marshal(aAccount)
	err = stub.PutState(A, aAccountByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("PutState error:%s", err))
	}

	bAccountByte, _ := json.Marshal(bAccount)
	err = stub.PutState(B, bAccountByte)
	if err != nil {
		return shim.Error(fmt.Sprintf("PutState error:%s", err))
	}

	return shim.Success(nil)
}

// arg[0]=转出账户，arg[1]=转入账户，arg[2]=金额
func (t *SimpleChaincode) transferAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 3 {
		return shim.Error("must have 3 args")
	}

	A := args[0]
	B := args[1]
	val := args[2]

	Abytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error(fmt.Sprintf("Error Getstate of %s", A))
	}

	Bbytes, err1 := stub.GetState(B)
	if err1 != nil {
		return shim.Error(fmt.Sprintf("Error Getstate of %s", B))
	}

	var aAccount Accountinfo
	var bAccount Accountinfo

	err = json.Unmarshal(Abytes, &aAccount)
	if err != nil {
		return shim.Error(fmt.Sprintf("Error Unmarshal :", err.Error))
	}
	err = json.Unmarshal(Bbytes, &bAccount)
	if err != nil {
		return shim.Error(fmt.Sprintf("Error Unmarshal :", err.Error))
	}

	iVal, _ := strconv.Atoi(val)
	aAccount.Ballance -= iVal
	bAccount.Ballance += iVal

	aBytes, _ := json.Marshal(aAccount)
	bBytes, _ := json.Marshal(bAccount)
	stub.PutState(A, aBytes)
	stub.PutState(B, bBytes)

	return shim.Success(nil)
}

// arg[0]=账户名称
func (t *SimpleChaincode) queryAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		fmt.Printf("must have one arg")
		return shim.Error("must have one arg")
	}
	val, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(fmt.Sprintf("GetState error:%s", err))
	}
	fmt.Printf("key=%s, val=%s", args[0], val)
	return shim.Success([]byte(val))
}

//Invoke function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "transferAccount" {
		return t.transferAccount(stub, args)
	}
	if function == "queryAccount" {
		return t.queryAccount(stub, args)
	}
	return shim.Error("Invalid invoke function name.")
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
