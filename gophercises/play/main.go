package main

import (
	"fmt"
	"net/url"
)

func main() {
	myUrl := "https://www.google.com:3000/jdahbsnfsv"
	u, _:= url.Parse(myUrl)
	fmt.Println(u.Scheme, u.Hostname(), u.Fragment)


}
