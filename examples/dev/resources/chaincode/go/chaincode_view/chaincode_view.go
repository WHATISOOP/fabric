package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/whatisoop/fabric/core/chaincode/shim"
	pb "github.com/whatisoop/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type TestInfo struct {
	TxId  string `json:"txId"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Date  string `json:"date"`
}

// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	switch function {
	case "add":
		return t.addInfo(stub, args)
	case "get":
		return t.getInfo(stub, args)
	case "searchByView":
		return t.searchByView(stub, args)
	default:
		return shim.Error("Unsupported operation")
	}
}

func (t *SimpleChaincode) getInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	key := args[0]
	fmt.Println("===getInfo===" + key)
	if len(args) < 1 {
		return shim.Error("getInfo operation must have 1 arg")
	}

	value, err := stub.GetState(key)
	if err != nil {
		return shim.Error("getInfo operation failed while getting the state : " + err.Error())
	}

	return shim.Success(value)
}

func (t *SimpleChaincode) addInfo(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("===addInfo===")
	if len(args) < 1 {
		return shim.Error("addInfo operation must have 1 arg")
	}
	// get the args
	bytetestinfo := []byte(args[0])
	//get some info
	testInfo := &TestInfo{}
	err := json.Unmarshal(bytetestinfo, &testInfo)
	if err != nil {
		fmt.Println(err)
		return shim.Error("Unmarshal failed")
	}

	testInfo.TxId = stub.GetTxID()
	dataInfo, err2 := json.Marshal(testInfo)
	if err2 != nil {
		fmt.Println(err2)
		return shim.Error("Marshal testInfo failed")
	}

	//save the json info
	err = stub.PutState(testInfo.TxId, dataInfo)
	if err != nil {
		return shim.Error("putting state err: " + err.Error())
	}
	fmt.Println("===addInfo successful===" + testInfo.TxId)
	return shim.Success(nil)
}

func (t *SimpleChaincode) searchByView(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) < 1 {
		return shim.Error("searchByView operation must have 1 ars")
	}
	opt := args[0]
	fmt.Println("===searchByView===" + opt)
	res, err := stub.QueryByView(opt)
	if err != nil {
		return shim.Error(err.Error())
	}

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("{\"list\":[")

	bArrayMemberAlreadyWritten := false
	for res.HasNext() {
		queryResponse, err := res.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]}")
	fmt.Printf("searchByView queryResult:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())

}
