package main

import (
	"fmt"

	"github.com/nothyphen/idpays/idpay"
)

func main() {
	idpay := idpay.NewIdPays()
	id, link, _ := idpay.Payment("1", "d50a77ba-124a-40ef-9e7b-b326a721921b", 101, 10000, "porya", "09111111111", "pourya@ikernel.ir", "hichi", "http://ikernel.ir/callback")
	fmt.Println(id, link)
}