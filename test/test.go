package main

import (
	"fmt"

	"github.com/nothyphen/idpays/idpay"
)

func main() {
	idpay := idpay.NewIdPays("d50a77ba-124a-40ef-9e7b-b326a721921b")
	
	id, link, _ := idpay.Payment("1", 101, 10000, "porya", "09111111111", "pourya@ikernel.ir", "hichi", "http://ikernel.ir/callback")
	fmt.Println(id, link)

	_, err := idpay.Verify("1", "d2e353189823079e1e4181772cff5292", "101")
	if err != nil {
		panic(err)
	}
}