package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "math"
    "net"
    "net/http"
    "net/rpc"
    "net/rpc/jsonrpc"
    "strconv"
   "strings"
    "time"
)

var UserId int64
var remaining float64
//request struct
type Request struct {
    NameandPercentage []string
    Budget            float64
}
type Stockprice struct {
    stockprice  []float64
    stockamount []float64
    noofstocks  []float64
}
type Responses struct{}
//response struct
type Response struct {
    TradeId  int64
    Stocks  string
    Unvested float64
}

//json stock struct
type Stock struct {
    List struct{
	Meta struct{
		Count int   `json:"count"`
		Start int   `json:"start"`
		Type string `json:"type"`
	} `json:"meta"`
	Resources []struct{
		Resource struct{
			Classname string `json:"classname"`
			Fields struct{
				Name string    `json:"name"`
				Price string   `json:"price"`
				Symbol string  `json:"symbol"`
				Ts string      `json:"ts"`
				Type string    `json:"type"`
				UTCtime string `json:"utctime"`
				Volume string  `json:"volume"`
			}`json:"fields"`
		}`json:"resource"`
	}`json:"resources"`
    }`json:"list"`
}
//function to create price string
func Pricestr(s Stock) (prices string) {
	
	var priceStr  = ""
	count := s.List.Meta.Count
	for i:=0;i<count;i++ {
		priceStr = priceStr + s.List.Resources[i].Resource.Fields.Price + ","
		
		
	}

	return priceStr
}
//function to create price 
func Pricestr1(s1 Stock) (prices []float64) {
	
	var priceStr  []float64
	count := s1.List.Meta.Count
	for i:=0;i<count;i++ {
		if n, err := strconv.ParseFloat((s1.List.Resources[i].Resource.Fields.Price), 64); err == nil {
		priceStr = append(priceStr,n)
		//priceStr = priceStr + s.List.Resources[i].Resource.Fields.Price + ","
		//symbolStr= symbolStr + stock.List.Resources[i].Resource.Fields.Symbol + ","
		
	}}
	return priceStr
	}
//function to caluculate no of stocks
func noofstock(x float64, y float64, balance float64)(tt float64){
return math.Floor(x*balance/y)
}
//function to caluculate sum of stock prices
func total(x string, y []float64)(t []float64){

var numbers []float64
rq := strings.Split(x,",")
	for _, arg := range rq {
        if n, err := strconv.ParseFloat(arg, 64); err == nil {
            numbers = append(numbers, n)

        }
    }
var nu []float64
for i:=0;i<len(y);i++{
n:=numbers[i]*y[i]
nu = append(nu,n)

}
return nu
}
//function to create intended output string
func Strfunc(m []string, n []float64, p string)(t2 string){
var numbers []float64
var strn []string
var strval string
rq := strings.Split(p,",")
	for _, arg := range rq {
        if n, err := strconv.ParseFloat(arg, 64); err == nil {
            numbers = append(numbers, n)
        }
    }
   for i:=0;i<len(numbers);i++{
    q:=fmt.Sprintf("%v:%v:%v",m[i],n[i],numbers[i])
    strn = append(strn,q)
    }
  strval = strings.Join(strn, ", ")
    return strval
}

func NoofStocks(priceStr string, percntStr string, balance float64)(tt []float64){
var numbers []float64
 bud:=balance
var nos []float64
	rq := strings.Split(priceStr,",")
	for _, arg := range rq {
        if n, err := strconv.ParseFloat(arg, 64); err == nil {
            numbers = append(numbers, n)

        }
    }
var numbers1 []float64
	
	rq1 := strings.Split(percntStr,",")
	for _, arg1 := range rq1 {
        if n1, err := strconv.ParseFloat(arg1, 64); err == nil {
            numbers1 = append(numbers1, n1)

        }
}

var numbers2 []float64
for _, arg2 := range numbers1 {
         	n2:= arg2/100.00
            numbers2 = append(numbers2, n2)
        }
for _, arg3 := range numbers2  {
for _,arg4 := range numbers{
n3 :=noofstock(arg3,arg4,bud)
nos = append(nos, n3)
}
}
return nos	
}


const (
    timeout = time.Duration(time.Second * 100)
)
//method for geting stock details
func (t *Responses) BuyStock(req *Request, reply *Response) error {

    client := http.Client{Timeout: timeout}
    var names []string
    for i := 0; i < ((len(req.NameandPercentage))/2); i++{
        names = append(names, req.NameandPercentage[i])
    }
    s := names
    str := strings.Join(s, ",")
    var numb []string
    n := len(req.NameandPercentage)/2
   
    for i := n; i < (n*2); i++{
        numb = append(numb, req.NameandPercentage[i])
    }
    url := fmt.Sprintf("http://finance.yahoo.com/webservice/v1/symbols/%s/quote?format=json",str )
     
    res, err := client.Get(url)
    if err != nil {
        fmt.Errorf("Stocks cannot access yahoo finance API: %v", err)
    }
    defer res.Body.Close()
    content, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Errorf("Stocks cannot read json body: %v", err)
    }
    var stock Stock
    err = json.Unmarshal(content, &stock)
    if err != nil {
        fmt.Errorf("Stocks cannot parse json data: %v", err)
    }
   var nos []float64
  s1 := numb
str1 := strings.Join(s1, ",")
   pstr := Pricestr(stock)
  tstr := NoofStocks(pstr,str1,req.Budget) 
  n1:=len(tstr)
for i:=0;i<n1;i++{
if(i%4==0){
n3:=tstr[i]
nos=append(nos,n3)
}
}
total1:=0.0
rem:=total(pstr,nos)
for _,ar:=range rem{
total1 += ar
}
    remaining = (req.Budget) - total1
    strn1:=Strfunc(names,nos,pstr)
    UserId += 1
   id := UserId
   
cprice[id]=strn1
    result := Response{
        TradeId: id,
        Stocks:strn1,
        Unvested: remaining,
    }
    *reply = result
    return nil
}
// method caluclating marketcurrent value
func cmarket(m []float64,n []float64)(z float64){
var m1 []float64
 total := 0.
for i:=0;i<len(m);i++{
n:=m[i]*n[i]
m1=append(m1,n)
}
for j:=0;j<len(m1);j++{
total = total+m1[j]
}
return total
}

//portfolio method
func compareprice(x string)(y Check){
var names,per,num,stre []string
var oprice,nos []float64
 w := strings.FieldsFunc(x, func(r rune) bool {
        switch r {
        case ',', ':', ' ':
            return true
        }
        return false
    })
     for i := 0; i < len(w); i=i+3 {
	n1:=w[i]
        names = append(names, n1)
    }
     for i := 1; i < len(w); i = i + 3 {
	n2:=w[i]
        per = append(per, n2)
    }
     for i := 2; i < len(w); i = i + 3 {
	n3:=w[i]
        num = append(num, n3)
    }
    strjoin:=strings.Join(names,",")
    client1 := http.Client{Timeout: timeout}
    url := fmt.Sprintf("http://finance.yahoo.com/webservice/v1/symbols/%s/quote?format=json",strjoin )
    fmt.Printf(url)
    res, err := client1.Get(url)
    if err != nil {
        fmt.Errorf("Stocks cannot access yahoo finance API: %v", err)
    }
    defer res.Body.Close()
    content, err := ioutil.ReadAll(res.Body)
    if err != nil {
        fmt.Errorf("Stocks cannot read json body: %v", err)
    }
    var stock Stock
    err = json.Unmarshal(content, &stock)
    if err != nil {
        fmt.Errorf("Stocks cannot parse json data: %v", err)
    }
    neprice := Pricestr1(stock)
    nprice:=Pricestr(stock)
    for q:=0;q<len(per);q++ {
        if n2, err := strconv.ParseFloat(per[q], 64); err == nil {
            nos = append(nos, n2)
        }
    for p:=0;p<len(num);p++ {
        if n1, err := strconv.ParseFloat(num[p], 64); err == nil {
            oprice = append(oprice, n1)
        }
	
    }
    for i:=0;i<len(neprice);i++{
	if(neprice[i]<oprice[i]){
	s:=fmt.Sprintf("-%v",neprice[i])
	stre = append(stre,s)
	} else if(neprice[i]>oprice[i]){
	s1:=fmt.Sprintf("+%v",neprice[i])
	stre = append(stre,s1)
	} else{
	s2:=fmt.Sprintf("%v",neprice[i])
	stre = append(stre,s2)
	}

	}
	
	}
	str := Strfunc(names,nos,nprice)
	cmarketval:=cmarket(neprice,nos)
	result:=Check{
	Stocks:str,
	Currentmarketvalue:cmarketval,
	Unvested:remaining,
	}
	return result

}

type Checks struct{}
//portfolio response struct
type Check struct {
    Stocks  string
    Currentmarketvalue float64
    Unvested float64
    }
var cprice = map[int64]string{}
//function calling portfolio method
func (t1 *Checks) Checkprice(u int64, reply1 *Check) error {
var ex string
if val, ok := cprice[u]; ok{
ex=val
}
var resstr Check
resstr=compareprice(ex)
*reply1= resstr
return nil
}

func main() {
    st := new(Responses)
    st1 := new(Checks)
    server := rpc.NewServer()
    server.Register(st)
     server.Register(st1)
    server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
    listener, e := net.Listen("tcp", ":1234")
    if e != nil {
        log.Fatal("listen error:", e)
    }
    for {
        if conn, err := listener.Accept(); err != nil {
            log.Fatal("accept error: " + err.Error())
        } else {
            log.Printf("new connection established\n")
            go server.ServeCodec(jsonrpc.NewServerCodec(conn))

        }
    }
}