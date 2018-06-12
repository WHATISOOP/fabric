package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/whatisoop/fabric/core/chaincode/shim"
	pb "github.com/whatisoop/fabric/protos/peer"
)

type Token struct {
	Lock        bool    `json:"Lock"`
	TokenName   string  `json:"TokenName"`
	TokenSymbol string  `json:"TokenSymbol"`
	TotalSupply float64 `json:"TotalSupply"`
	Owner       string  `json:Owner`
}

type Account struct {
	Name      string             `json:"Name"`
	Frozen    bool               `json:"Frozen"`
	BalanceOf map[string]float64 `json:"BalanceOf"`
}

func (account *Account) balance(_currency string) map[string]float64 {
	bal := map[string]float64{_currency: account.BalanceOf[_currency]}
	return bal
}

func (account *Account) balanceAll() map[string]float64 {
	return account.BalanceOf
}

type SmartContract struct {
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := stub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "createAccount" {
		return s.createAccount(stub, args)
	} else if function == "createToken" {
		return s.createToken(stub, args)
	} else if function == "setLock" {
		return s.setLock(stub, args)
	} else if function == "transferToken" {
		return s.transferToken(stub, args)
	} else if function == "frozenAccount" {
		return s.frozenAccount(stub, args)
	} else if function == "balance" {
		return s.balance(stub, args)
	} else if function == "balanceAll" {
		return s.balanceAll(stub, args)
	} else if function == "showAccount" {
		return s.showAccount(stub, args)
	} else if function == "showToken" {
		return s.showToken(stub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) createAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1: accountName.")
	}

	key := args[0]
	name := args[0]
	existAsBytes, err := stub.GetState(key)
	fmt.Printf("GetState(%s) %s \n", key, string(existAsBytes))
	if string(existAsBytes) != "" {
		fmt.Println("账户名已经存在")
		return shim.Error("账户名已经存在")
	}

	account := Account{
		Name:      name,
		Frozen:    false,
		BalanceOf: map[string]float64{}}

	accountAsBytes, _ := json.Marshal(account)
	err = stub.PutState(key, accountAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("createAccount %s \n", string(accountAsBytes))

	return shim.Success([]byte("创建账户成功"))
}
func (s *SmartContract) showToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1: tokensymbol.")
	}
	_symbol := args[0]
	tokenAsBytes, err := stub.GetState(_symbol)
	if err != nil {
		return shim.Error(err.Error())
	} else {
		fmt.Printf("GetState(%s)) %s \n", _symbol, string(tokenAsBytes))
	}
	return shim.Success(tokenAsBytes)
}

func (s *SmartContract) createToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4: tokenname, tokensymbol, totalsupply, accountname.")
	}

	_name := args[0]
	_symbol := args[1]
	_supply, _ := strconv.ParseFloat(args[2], 64)
	_account := args[3]

	tokenBytes, err := stub.GetState(_symbol)
	if err != nil {
		return shim.Error(err.Error())
	}
	if tokenBytes != nil {
		return shim.Error("代币已经存在")
	}

	coinbaseAsBytes, err := stub.GetState(_account)
	if err != nil {
		return shim.Error(err.Error())
	}
	if coinbaseAsBytes == nil {
		return shim.Error("账户不存在")
	}
	fmt.Printf("Coinbase before %s \n", string(coinbaseAsBytes))

	coinbase := &Account{}

	json.Unmarshal(coinbaseAsBytes, &coinbase)

	token := Token{}
	token.Lock = false
	token.TokenName = _name
	token.TokenSymbol = _symbol
	token.TotalSupply = _supply
	token.Owner = _account

	tokenAsBytes, _ := json.Marshal(token)
	err = stub.PutState(_symbol, tokenAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	} else {
		fmt.Printf("Init Token %s \n", string(tokenAsBytes))
	}

	coinbase.BalanceOf[_symbol] = _supply
	coinbaseAsBytes, _ = json.Marshal(coinbase)
	err = stub.PutState(_account, coinbaseAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("Coinbase after %s \n", string(coinbaseAsBytes))

	return shim.Success([]byte("创建token成功"))
}

func (s *SmartContract) transferToken(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4： from_account_name, to_account_name, tokensymbol, amouont.")
	}
	_from := args[0]
	_to := args[1]
	_currency := args[2]
	_amount, _ := strconv.ParseFloat(args[3], 32)

	if _amount <= 0 {
		return shim.Error("Amount must > 0.")
	}

	fromAsBytes, err := stub.GetState(_from)
	if err != nil {
		return shim.Error(err.Error())
	}
	if fromAsBytes == nil {
		return shim.Error("转出人不存在!")
	}
	fmt.Printf("fromAccount %s \n", string(fromAsBytes))
	fromAccount := &Account{}
	json.Unmarshal(fromAsBytes, &fromAccount)

	toAsBytes, err := stub.GetState(_to)
	if err != nil {
		return shim.Error(err.Error())
	}
	if toAsBytes == nil {
		return shim.Error("接收人不存在!")
	}
	fmt.Printf("toAccount %s \n", string(toAsBytes))
	toAccount := &Account{}
	json.Unmarshal(toAsBytes, &toAccount)

	tokenAsBytes, err := stub.GetState(_currency)
	if err != nil {
		return shim.Error(err.Error())
	}
	if tokenAsBytes == nil {
		return shim.Error("token不存在")
	}
	fmt.Printf("Token %s \n", string(toAsBytes))
	token := &Token{}
	json.Unmarshal(tokenAsBytes, &token)

	if token.Lock {
		return shim.Error("锁仓状态，停止一切转账活动")
	}
	if fromAccount.Frozen {
		return shim.Error("From 账号冻结")
	}
	if toAccount.Frozen {
		return shim.Error("To 账号冻结")
	}
	if fromAccount.BalanceOf[_currency] >= _amount {
		fromAccount.BalanceOf[_currency] -= _amount
		toAccount.BalanceOf[_currency] += _amount
	} else {
		return shim.Error("余额不足")
	}

	fromAsBytes, err = json.Marshal(fromAccount)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(_from, fromAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	} else {
		fmt.Printf("fromAccount %s \n", string(fromAsBytes))
	}

	toAsBytes, err = json.Marshal(toAccount)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(_to, toAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	} else {
		fmt.Printf("toAccount %s \n", string(toAsBytes))
	}

	return shim.Success([]byte("转账成功"))
}

func (s *SmartContract) setLock(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2： tokensymbol, islock(true/false).")
	}

	tokenKey := args[0]

	lock := args[1]

	tokenAsBytes, err := stub.GetState(tokenKey)
	if err != nil {
		return shim.Error(err.Error())
	}
	if tokenAsBytes == nil {
		return shim.Error("token不存在")
	}

	token := Token{}

	json.Unmarshal(tokenAsBytes, &token)

	if lock == "true" {
		token.Lock = true
	} else if lock == "false" {
		token.Lock = false
	} else {
		return shim.Error("锁定状态只能是true/false.")
	}

	tokenAsBytes, err = json.Marshal(token)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(tokenKey, tokenAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("setLock - end %s \n", string(tokenAsBytes))

	return shim.Success(nil)
}
func (s *SmartContract) frozenAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	_account := args[0]
	_status := args[1]

	accountAsBytes, err := stub.GetState(_account)
	if err != nil {
		return shim.Error(err.Error())
	}

	account := Account{}

	json.Unmarshal(accountAsBytes, &account)

	var status bool
	if _status == "true" {
		status = true
	} else if _status == "false" {
		status = false
	} else {
		return shim.Error("账户冻结状态只能是true/false.")
	}

	account.Frozen = status

	accountAsBytes, err = json.Marshal(account)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(_account, accountAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	} else {
		fmt.Printf("frozenAccount - end %s \n", string(accountAsBytes))
	}

	return shim.Success([]byte("frozeAccount successfully."))
}

func (s *SmartContract) showAccount(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1： accountname.")
	}
	_account := args[0]

	accountAsBytes, err := stub.GetState(_account)
	if err != nil {
		return shim.Error(err.Error())
	} else {
		fmt.Printf("Account balance %s \n", string(accountAsBytes))
	}
	return shim.Success(accountAsBytes)
}

func (s *SmartContract) balance(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2: accountname, tokensymbol.")
	}
	_account := args[0]
	_currency := args[1]

	accountAsBytes, err := stub.GetState(_account)
	if err != nil {
		return shim.Error(err.Error())
	} else {
		fmt.Printf("Account balance %s \n", string(accountAsBytes))
	}

	account := Account{}
	json.Unmarshal(accountAsBytes, &account)
	result := account.balance(_currency)

	resultAsBytes, _ := json.Marshal(result)
	fmt.Printf("%s balance is %s \n", _account, string(resultAsBytes))

	return shim.Success(resultAsBytes)
}

func (s *SmartContract) balanceAll(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1: accountname.")
	}
	_account := args[0]

	accountAsBytes, err := stub.GetState(_account)
	if err != nil {
		return shim.Error(err.Error())
	} else {
		fmt.Printf("Account balance %s \n", string(accountAsBytes))
	}

	account := Account{}
	json.Unmarshal(accountAsBytes, &account)
	result := account.balanceAll()
	resultAsBytes, _ := json.Marshal(result)
	fmt.Printf("%s balance is %s \n", _account, string(resultAsBytes))

	return shim.Success(resultAsBytes)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
