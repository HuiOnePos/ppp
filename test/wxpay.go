package main

import (
	"github.com/HuiOnePos/ppp"
)

var wxpay *ppp.WXPay
//
// func main3() {
//
// 	config := ppp.LoadConfig("/Users/HuiOnePos/work/go/src/ppp/config.yml")
// 	ppp.NewLogger(config.Sys.LogLevel)
// 	ppp.NewDBPool(&config.DB)
//
// 	if !config.WXPay.Use {
// 		return
// 	}
// 	wxpay = ppp.NewWXPay(config.WXPay)
// 	// barPay()
// 	authSigner1()
// 	// tradeInfo()
// 	// cancel()
// 	// refund()
// 	// payParams(ppp.APPPAY)
// 	// payParams(ppp.WEBPAY)
//
// }
//
// func authSigner1() {
// 	fmt.Println(wxpay.AuthSigned(&ppp.Auth{MchID: "1490825832", Status: 1, Account: "闪收网络"}))
// }
