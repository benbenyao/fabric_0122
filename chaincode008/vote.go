package main

import (
	"fmt"
	"bytes"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type VoteChaincode struct {

}

type Vote struct {
	Username string `json:"username"`
	VoteNum int `json:"votenum"`
}

func (t * VoteChaincode) Init (stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (t * VoteChaincode) Invoke (stub shim.ChaincodeStubInterface) peer.Response {

	fn , args := stub.GetFunctionAndParameters()

	if fn == "voteUser" {
		return t.voteUser(stub, args)
	} else if fn == "getUserVote" {
		return t.getUserVote(stub)
	}

	return shim.Error("调用方法不存在！")
}

func (t *VoteChaincode) voteUser (stub shim.ChaincodeStubInterface, args []string) peer.Response {

	fmt.Println("start voteUser")
	username := args[0]
	// 判断当前用户是否存在
	userAsBytes , err := stub.GetState(username)

	if err != nil {
		return shim.Error(err.Error())
	}

	vote := Vote{}

	if userAsBytes != nil {
		err = json.Unmarshal(userAsBytes , &vote)

		if err != nil {
			return shim.Error(err.Error())
		}

		vote.VoteNum += 1
	} else {
		vote = Vote{Username: username, VoteNum: 1}
	}

	voteJsonAsBytes , err := json.Marshal(vote)

	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(username, voteJsonAsBytes)

	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("vote user: "+username)
	fmt.Println("end voteUser")

	return shim.Success(nil)
}

func (t *VoteChaincode) getUserVote( stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("start getUserVote")
	resultIterator, err := stub.GetStateByRange("","")

	if err != nil {
		return shim.Error(err.Error())
	}

	defer resultIterator.Close()

	var buffer bytes.Buffer

	buffer.WriteString("[")

	isWrite := false
	for resultIterator.HasNext() {
		queryResponse , err := resultIterator.Next()

		if err != nil {
			return shim.Error(err.Error())
		}

		if isWrite == true {
			buffer.WriteString(",")
		}

		buffer.WriteString(string(queryResponse.Value))
		isWrite = true
	}

	buffer.WriteString("]")

	fmt.Println(buffer.String())
	fmt.Println("end getUserVote")
	return shim.Success(buffer.Bytes())
}

func main() {
	err := shim.Start(new(VoteChaincode))

	if err != nil {
		fmt.Println("chaincode start error")
	}
}