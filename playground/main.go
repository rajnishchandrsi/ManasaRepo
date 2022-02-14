package main

import (
   "fmt"
   "regexp"
)

func main() {
   s := "Abraham Lincoln* &^ $#@"
   reg := regexp.MustCompile("[a-zA-Z0-9]+")
   res := reg.ReplaceAllString(s,"")
   fmt.Println(res) // Abraham Lincoln
   //fmt.Println(m.FindStringIndex("Hello, Welcome"))

}