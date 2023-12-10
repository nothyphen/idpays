# Go [IDPay](https://idpay.ir/) API Wrapper

[![Go Report Card](https://goreportcard.com/badge/github.com/nothyphen/idpays)](https://goreportcard.com/report/github.com/nothyphen/idpays)

### Installation
```
go get github.com/nothyphen/idpays
```


------

## Usage

```go
package main

import (
    "fmt"
    "github.com/nothyphen/idpays"
)

func main() {
    idp := idpays.NewIdPays()

    id, link, error := idp.Payment(orderID, callbackURL, amount, mail, phone, desc)
}
```