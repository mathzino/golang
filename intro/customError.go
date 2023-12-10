package main

import (
  "fmt"
  "errors"
)

func pembagian(nilai uint, pembagi uint)(float64, error){
  if pembagi == 0 {
    return 0.0, errors.New("Maaf pembagi tidak boleh NOL")
  }else{
    return float64(nilai/pembagi), nil
  }
}

func main(){
  hasil, err := pembagian(8,0)
  if err == nil{
    fmt.Println("Hasil", hasil)
  }else{
    fmt.Println("Error", err.Error())
  }
}