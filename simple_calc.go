package main
import (
  "fmt"
  "strconv"
)

/*

You are building an educational website and want to create a simple calculator for students to use. The calculator will only allow addition and subtraction of non-negative integers.

Given an expression string using the "+" and "-" operators like "5+16-2", write a function to find the total.

Sample input/output:
calculate("6+9-12")  => 3
calculate("1+2-3+4-5+6-7") => -2
calculate("100+200+300") => 600
calculate("1-2-3-0") => -4
calculate("255") => 255


*/

func main() {
  const expression1 = "6+9-12"; // = 3
  const expression2 = "1+2-3+4-5+6-7"; // = -2
  const expression3 = "100+200+300"; // = 600
  const expression4 = "1-2-3-0"; // = -4
  const expression5 = "255"; // = 255
  
//   calculate(expression1)
//   calculate(expression2)
  calculate(expression5)
//   calculate(expression1)

}

// n: length of input string
// Time: O(n)
// Space O(1)

func calculate(s string) int {
  r := 0
  curr := ""
  plus := true
  lastExecuted := true
  
  for i, _ := range s {

    if s[i] == '+' {
      currVal, _ := strconv.Atoi(curr)
      fmt.Println("add = ",r, currVal)
//       r = r + currVal
       if plus {
        r = r + currVal
      }else{
        r = r - currVal
      }
      fmt.Println("added r = ", r)
      plus = true
      curr = ""
      lastExecuted = false
      
    }else if s[i] == '-' {
      currVal, _ := strconv.Atoi(curr)
      fmt.Println("sub = ",r, currVal)
//       r = r - currVal
       if plus {
        r = r + currVal
      }else{
        r = r - currVal
      }
      fmt.Println("substrct r = ", r)
      curr = ""
      plus = false
      lastExecuted = false
      
    }else{
      curr = curr + string(s[i])
      fmt.Println("curr = ", curr)
      lastExecuted = true
    }
  }
  
  fmt.Println("r == ", r)
  
  if lastExecuted {
    fmt.Println("curr in lastexec = ", curr, r)
    currVal, _ := strconv.Atoi(curr)
    if plus {
      r = r + currVal
    }else{
      r = r - currVal
    }
  }
  
  fmt.Println("final r = ", r)
  return r
}
