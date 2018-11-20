package main

import(
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct{

}

type Person struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Sex string `json:"sex"`
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response{
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke( stub shim.ChaincodeStubInterface) peer.Response{

	// GetFunctionAndParameters
	//{"Args":["set","tom","100"]}
	fn, args := stub.GetFunctionAndParameters()
	fmt.Println(fn, args)

	args = stub.GetStringArgs()
	fmt.Println(args)


	// PutState 
	err := stub.PutState("str",[]byte("hello"))
	if err != nil {
		fmt.Println("str PutState error: "+err.Error())
	}else{
		fmt.Println("str PutState success!")
	}

	// PutState 
	err = stub.PutState("str1",[]byte("hello1"))
	if err != nil {
		fmt.Println("str1 PutState error: "+err.Error())
	}else{
		fmt.Println("str1 PutState success!")
	}

	// PutState 
	err = stub.PutState("str2",[]byte("hello2"))
	if err != nil {
		fmt.Println("str2 PutState error: "+err.Error())
	}else{
		fmt.Println("str2 PutState success!")
	}

	// GetState 

	strValue , err := stub.GetState("str")
	if err != nil {
		fmt.Println("str GetState error: "+err.Error())
	}else {
		fmt.Printf("str value: %s \n",string(strValue))
	}

	// DelState
	err = stub.DelState("str1")

	if err != nil {
		fmt.Println("Delete str1 err")
	}else {
		fmt.Println("Delete str1 success!")
	}

	// GetStateByRange

	resultIterator , err := stub.GetStateByRange("str" , "str2")

	defer resultIterator.Close()
	fmt.Println("-----start resultIterator-----")
	for resultIterator.HasNext() {
		item, _ := resultIterator.Next()
		fmt.Println(string(item.Value))
	}
	fmt.Println("-----end resultIterator-----")

	// GetHistoryForKey
	historyIterator,err := stub.GetHistoryForKey("str")
	defer historyIterator.Close()
	fmt.Println("-----start historyIterator-----")
	for resultIterator.HasNext() {
		item, _ := historyIterator.Next()
		fmt.Println(string(item.TxId))
		fmt.Println(string(item.Value))
	}
	fmt.Println("-----end historyIterator-----")

	// CreateCompositeKey
	indexName := "sex~name"
	indexKey , err := stub.CreateCompositeKey(indexName,[]string{"boy","xiaowang"})

	value := []byte{0x00}
	stub.PutState(indexKey,value)
	fmt.Println(indexKey)
	indexKey , err = stub.CreateCompositeKey(indexName,[]string{"boy","xiaoli"})
	stub.PutState(indexKey,value)
	fmt.Println(indexKey)
	indexKey , err = stub.CreateCompositeKey(indexName,[]string{"girl","xiaofang"})
	fmt.Println(indexKey)
	stub.PutState(indexKey,value)

	// GetStateByPartialCompositeKey
	resultIterator,err = stub.GetStateByPartialCompositeKey(indexName, []string{"boy"})
	defer resultIterator.Close()
	fmt.Println("-----start resultIterator-----")
	for resultIterator.HasNext() {
		item, _ := resultIterator.Next()

		objectType, compositeKeyParts, err := stub.SplitCompositeKey(item.Key)
		if err != nil {
			return shim.Error(err.Error())
		}
		fmt.Println("objectType: "+objectType)
		fmt.Println("sex : "+compositeKeyParts[0])
		fmt.Println("name : "+compositeKeyParts[1])
		
	}
	fmt.Println("-----end resultIterator-----")

	// GetQueryResult
	resultIterator , err = stub.GetQueryResult("{\"selector\": {\"sex\": \"boy\"}}" )

	defer resultIterator.Close()
	fmt.Println("-----start resultIterator-----")
	for resultIterator.HasNext() {
		item, _ := resultIterator.Next()
		fmt.Println(string(item.Value))
	}
	fmt.Println("-----end resultIterator-----")

	// InvokeChaincode
	trans:=[][]byte{[]byte("invoke"),[]byte("a"),[]byte("b"),[]byte("11")}
	stub.InvokeChaincode("mycc",trans,"mychannel")

	return shim.Success(nil)
}

func main(){
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Println("start err")
	}
}