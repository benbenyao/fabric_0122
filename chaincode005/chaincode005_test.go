package main

import (
    "fmt"
    "testing"
    "github.com/hyperledger/fabric/core/chaincode/shim"
)

func checkInvoke(t *testing.T, stub *shim.MockStub, args [][]byte) {
    res := stub.MockInvoke("1", args)
    if res.Status != shim.OK {
        fmt.Println("Invoke", args, "failed", string(res.Message))
        t.FailNow()
    }
}

func Test_Chaincode005(t *testing.T) {

    hello := new(SimpleChaincode)
    stub := shim.NewMockStub("hello", hello)

    checkInvoke(t, stub, [][]byte{[]byte("set"),[]byte("tom"),[]byte("100")})
}
