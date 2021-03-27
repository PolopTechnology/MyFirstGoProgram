package main

import "strconv"
import "fmt"

func add(x int64, y int64) int64 {
        return x + y
}

func substract(x int64, y int64) int64 {
              return x - y
}

func multiply(x int64, y int64) int64 {
             return x * y
}

func divide(x int64, y int64) int64 {
           return x / y
}

func main() {

         fmt.Println("Insert first parameter.")
         var f1 string
         fmt.Scanln(&f1)
         fmt.Println("Insert operator.")
         var op string
         fmt.Scanln(&op)    
         fmt.Println("Insert second parameter.")
         var f2 string
         fmt.Scanln(&f2)       

         val1, err := strconv.ParseInt(f1, 0, 64)
         val2, err := strconv.ParseInt(f2, 0, 64)
         if err != nil {
             panic(err)
         }
         if op == "+" {
                  fmt.Println(add(val1, val2))
         } else if op == "-" {
                  fmt.Println(substract(val1, val2))
         } else if op == "*" {
                  fmt.Println(multiply(val1, val2))
         } else if op == "/" {
                  fmt.Println(divide(val1, val2))
         } else {
                  fmt.Println("Not an operator, dum dum")
                  fmt.Println("operators look like this: + * / -")
         }
}


