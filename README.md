# gomail
A simple golang package to send email using gmail

## Usage

```
package main

import (
	"fmt"

	"github.com/souvikhaldar/gomail"
)

func main() {
	e, config := gomail.New("from@eg.com", []string{"to@eg.com"}, "Hello Gopher!", "Testing the fantastic gomail lib", "<password>")
	if e != nil {
		fmt.Print(fmt.Errorf("Error in creating config %v", e))
	}
	if e := config.SendMail(); e != nil {
		fmt.Print(fmt.Errorf("Error in sending mail %v", e))
	}
}
```
