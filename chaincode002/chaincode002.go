package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type Helloworld struct {

}

func (t * Helloworld) Init(stub shim.ChaincodeStubInterface) peer.Response{

	args:= stub.GetStringArgs()

	err := stub.PutState(args[0],[]byte(args[1]))

	if err != nil {
		shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (t *Helloworld) Invoke (stub shim.ChaincodeStubInterface) peer.Response{

	fn, args := stub.GetFunctionAndParameters()

	if fn =="set" {
		return t.set(stub, args)
	}else if fn == "get"{
		return t.get(stub , args)
	}
	return shim.Error("Invoke fn error")
}

func (t *Helloworld) set(stub shim.ChaincodeStubInterface , args []string) peer.Response{
	err := stub.PutState(args[0],[]byte(args[1]))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *Helloworld) get (stub shim.ChaincodeStubInterface, args [] string) peer.Response{

	value, err := stub.GetState(args[0])

	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(value)
}

func main(){
	err := shim.Start(new(Helloworld))
	if err != nil {
		fmt.Println("start error")
	}
}