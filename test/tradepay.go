package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/panjjo/ppp"
)

func main() {
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	tradePayResult := ppp.PayResult{}
	err = client.Call("AliPay.BarCodePay", ppp.BarCodePayRequest{
		OutTradeId: "20178752312435",
		TradeName:  "Craig",
		ItemDes:    "快速收银,临时商品",
		AuthCode:   "286048035225904833",
		Amount:     100,
		UserId:     "2088102169330843"}, &tradePayResult)
	fmt.Println(tradePayResult)
	if err != nil {
		fmt.Println(err)
	}
}