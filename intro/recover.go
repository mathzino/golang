package main

import (
  "fmt"
)

func endApp(){
  fmt.Println("End App")
  message := recover()
  if message != nil{
	  fmt.Println("Terjadi Error", message)
  }
}

func runApp(error bool){
  defer endApp()
  if error{
    panic("ERROR")
  }
}

func main(){
  runApp(true)
}