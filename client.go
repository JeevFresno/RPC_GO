package main

import (
	"net/rpc/jsonrpc"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Request1 struct {
	Symb string
	Budget float64
}

type Response1 struct {
	Remainder float64
	Tradeid int
	Rep string
}




func main() {
	//var length int
	service :="127.0.0.1:1400"
	//length = len(os.Args)
	client, err := jsonrpc.Dial("tcp", service)
	if err != nil {
		log.Fatal("dialing:", err)
	}

		inparg := os.Args[1]
		inpbud,err := strconv.ParseFloat(os.Args[2],64)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		req1 := Request1{inparg,inpbud}
		res1 := Response1{0.00,0,""}
		err = client.Call("Stock.Stkresponse", req1, &res1)
		if err != nil {
			log.Fatal("stock error:", err)
		}

		fmt.Println("TRADE ID::",res1.Tradeid)
		fmt.Println("STOCK::",res1.Rep)
		fmt.Println("UNINVESTED AMOUNT",res1.Remainder)

}

