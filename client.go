package main

import (
    "fmt"
    "log"
    "net"
    "os"
    "strings"
    "net/rpc/jsonrpc"
    "strconv"
)

//response struct
type Request struct {
    NameandPercentage []string
    Budget            float64
}
//response struct
type Response struct {
    TradeId  int64
    Stocks  string
    Unvested float64
}
//portfolio response struct
type Check struct {
    Stocks  string
    Currentmarketvalue float64
    Unvested float64
    }

func main() {

    client, err := net.Dial("tcp", "127.0.0.1:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

    // Synchronous call
if((len(os.Args)>2)){
    fmt.Println("StockSymbol:Percentage",os.Args[1])
    fmt.Println("budget amount:",os.Args[2])
     w := strings.FieldsFunc(os.Args[1], func(r rune) bool {
        switch r {
        case ',', ':', '%':
            return true
        }
        return false
    })
    names := []string{}
    num := []string{}
    for i := 0; i < len(w); i=i+2 {
        names = append(names, w[i])
	
    }
     for i := 1; i < len(w); i = i + 2 {
        num = append(num, w[i])
    }
 
    str2 := append(names,num...)
    b,_ :=  strconv.ParseFloat(os.Args[2],64)
    request := Request{
        NameandPercentage: str2,
        Budget:            b,
    }
    var reply Response
    c := jsonrpc.NewClient(client)
    err = c.Call("Responses.BuyStock", request, &reply)
    if err != nil {
        log.Fatal("arith error:", err)
    } else {
	fmt.Println("TradeID:",reply.TradeId)
	 fmt.Println("stock details:",reply.Stocks)
	  fmt.Println("unvested amount:",reply.Unvested)
    }}else{

       var reply1 Check
    c1 := jsonrpc.NewClient(client)
    c,_ :=  strconv.ParseInt(os.Args[1],10,64)
    err = c1.Call("Checks.Checkprice", c, &reply1)
    if err != nil {
        log.Fatal("arith error:", err)
    } else {
        fmt.Println("stock details:",reply1.Stocks)
	 fmt.Println("current market value:",reply1.Currentmarketvalue)
	  fmt.Println("unvested amount:",reply1.Unvested)
    }
}
}