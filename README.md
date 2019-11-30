# gomail
A simple golang package to send email using gmail 

## Installation 

`go get github.com/souvikhaldar/gomail`  

## Usage

```
package main

import (
	"fmt"

	"github.com/souvikhaldar/gomail"
)

func main() {
	e, config := gomail.New("<from>", "<password>")
	if e != nil {
		fmt.Print(fmt.Errorf("Error in creating config %v", e))
	}
	if e := config.SendMail([]string{"<to>"},"<subject>","<body>"); e != nil {
		fmt.Print(fmt.Errorf("Error in sending mail %v", e))
	}
}
```

## Note
Gmail need to be allowed access to unsafe app.

https://serverfault.com/questions/635139/how-to-fix-send-mail-authorization-failed-534-5-7-14

#### Reference
https://gist.github.com/jpillora/cb46d183eca0710d909a

