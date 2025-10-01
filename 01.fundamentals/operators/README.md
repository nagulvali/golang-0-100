# Go Operators
- Operators are the foundation of any programming language.
- Thus the functionality of the Go language is incomplete without the use of operators.
- Operators allow us to perform different kinds of operations on operands.
- In the Go language, operators Can be categorized based on their different functionality:
  - Arithmetic Operators
  - Relational Operators
  - Logical Operators
  - Bitwise Operators
  - Assignment Operators
  - Misc Operators


## Arithmetic Operators
- These are used to perform arithmetic/mathematical operations on operands in Go language: 
  - Addition: The ‘+’ operator adds two operands. For example, x+y.
  - Subtraction: The ‘-‘ operator subtracts two operands. For example, x-y.
  - Multiplication: The '\*' operator multiplies two operands. For example, x*y.
  - Division: The ‘/’ operator divides the first operand by the second. For example, x/y.
  - Modulus: The ‘%’ operator returns the remainder when the first operand is divided by the second. For example, x%y.

> -, +, !, &, *, <-, and ^ are also known as unary operators and the precedence of unary operators is higher.

> ++ and -- operators are from statements they are not expressions, so they are out from the operator hierarchy.


```go
// Go program to illustrate the
// use of arithmetic operators
package main
 
import "fmt"
 
func main() {
   p:= 34
   q:= 20
    
   // Addition
   result1:= p + q
   fmt.Printf("Result of p + q = %d", result1)
    
   // Subtraction
   result2:= p - q
   fmt.Printf("\nResult of p - q = %d", result2)
    
   // Multiplication
   result3:= p * q
   fmt.Printf("\nResult of p * q = %d", result3)
    
   // Division
   result4:= p / q
   fmt.Printf("\nResult of p / q = %d", result4)
    
   // Modulus
   result5:= p % q
   fmt.Printf("\nResult of p %% q = %d", result5)
}
```



## Relational Operators
- Relational operators are used for the comparison of two values. Let’s see them one by one:
  - ‘=='(Equal To) operator checks whether the two given operands are equal or not. If so, it returns true. Otherwise, it returns false. For example, 5==5 will return true.
  - ‘!='(Not Equal To) operator checks whether the two given operands are equal or not. If not, it returns true. Otherwise, it returns false. It is the exact boolean complement of the ‘==’ operator. For example, 5!=5 will return false.
  - ‘>'(Greater Than) operator checks whether the first operand is greater than the second operand. If so, it returns true. Otherwise, it returns false. For example, 6>5 will return true.
  - ‘<‘(Less Than) operator checks whether the first operand is lesser than the second operand. If so, it returns true. Otherwise, it returns false. For example, 6<5 will return false.
  - ‘>=' (Greater Than Equal To)operator checks whether the first operand is greater than or equal to the second operand. If so, it returns true. Otherwise, it returns false. For example, 5>=5 will return true.
  - ‘<=' (Less Than Equal To)operator checks whether the first operand is lesser than or equal to the second operand. If so, it returns true. Otherwise, it returns false. For example, 5<=5 will also return true.

```go
// Go program to illustrate the
// use of relational operators
package main
 
import "fmt"
 
func main() {
   p:= 34
   q:= 20
    
   // ‘=='(Equal To)
   result1:= p == q
   fmt.Println(result1)
    
   // ‘!='(Not Equal To)
   result2:= p != q
   fmt.Println(result2)
    
   // ‘<‘(Less Than)
   result3:= p < q
   fmt.Println(result3)
    
   // ‘>'(Greater Than)
   result4:= p > q
   fmt.Println(result4)
    
   // ‘>='(Greater Than Equal To)
   result5:= p >= q
   fmt.Println(result5)
    
   // ‘<='(Less Than Equal To)
   result6:= p <= q
   fmt.Println(result6)
    
    
}
```



## Logical Operators
- They are used to combine two or more conditions/constraints or to complement the evaluation of the original condition in consideration.  
  - Logical AND: The ‘&&’ operator returns true when both the conditions in consideration are satisfied. Otherwise it returns false. For example, a && b returns true when both a and b are true (i.e. non-zero).
  - Logical OR: The ‘||’ operator returns true when one (or both) of the conditions in consideration is satisfied. Otherwise it returns false. For example, a || b returns true if one of a or b is true (i.e. non-zero). Of course, it returns true when both a and b are true.
  - Logical NOT: The ‘!’ operator returns true the condition in consideration is not satisfied. Otherwise it returns false. For example, !a returns true if a is false, i.e. when a=0.

```go
// Go program to illustrate the
// use of logical operators
package main
import "fmt"
func main() {
    var p int = 23
    var q int = 60
      
    if(p!=q && p<=q){ 
        fmt.Println("True")
    }
      
    if(p!=q || p<=q){ 
        fmt.Println("True")
    }
      
    if(!(p==q)){ 
        fmt.Println("True")
    }
      
}
```


## Bitwise Operators
- In Go language, there are 6 bitwise operators which work at bit level or used to perform bit by bit operations. Following are the bitwise operators :
  - `&` (bitwise AND): Takes two numbers as operands and does AND on every bit of two numbers. The result of AND is 1 only if both bits are 1.
  - `|` (bitwise OR): Takes two numbers as operands and does OR on every bit of two numbers. The result of OR is 1 any of the two bits is 1.
  - `^` (bitwise XOR): Takes two numbers as operands and does XOR on every bit of two numbers. The result of XOR is 1 if the two bits are different.
  - `<<` (left shift): Takes two numbers, left shifts the bits of the first operand, the second operand decides the number of places to shift.
  - `>>` (right shift): Takes two numbers, right shifts the bits of the first operand, the second operand decides the number of places to shift.
  - `&^` (AND NOT): This is a bit clear operator.

```go
// Go program to illustrate the
// use of bitwise operators
package main
 
import "fmt"
 
func main() {
   p:= 34
   q:= 20
    
   // & (bitwise AND)
   result1:= p & q
   fmt.Printf("Result of p & q = %d", result1)
    
   // | (bitwise OR)
   result2:= p | q
   fmt.Printf("\nResult of p | q = %d", result2)
    
   // ^ (bitwise XOR)
   result3:= p ^ q
   fmt.Printf("\nResult of p ^ q = %d", result3)
    
   // << (left shift)
   result4:= p << 1
   fmt.Printf("\nResult of p << 1 = %d", result4)
    
   // >> (right shift)
   result5:= p >> 1
   fmt.Printf("\nResult of p >> 1 = %d", result5)
    
   // &^ (AND NOT)
   result6:= p &^ q
   fmt.Printf("\nResult of p &^ q = %d", result6)
    
    
}
```

## Assignment Operators
- Assignment operators are used to assigning a value to a variable.
- The left side operand of the assignment operator is a variable and right side operand of the assignment operator is a value.
- The value on the right side must be of the same data-type of the variable on the left side otherwise the compiler will raise an error.
- Different types of assignment operators are shown below:
  - `“=”` (Simple Assignment): This is the simplest assignment operator. This operator is used to assign the value on the right to the variable on the left.
  - `“+=”` (Add Assignment): This operator is a combination of ‘+’ and ‘=’ operators. This operator first adds the current value of the variable on left to the value on the right and then assigns the result to the variable on the left.
  - `“-=”` (Subtract Assignment): This operator is a combination of ‘-‘ and ‘=’ operators. This operator first subtracts the current value of the variable on left from the value on the right and then assigns the result to the variable on the left.
  - `“*=”` (Multiply Assignment): This operator is a combination of ‘*’ and ‘=’ operators. This operator first multiplies the current value of the variable on left to the value on the right and then assigns the result to the variable on the left.
  - `“/=”` (Division Assignment): This operator is a combination of ‘/’ and ‘=’ operators. This operator first divides the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
  - `“%=”` (Modulus Assignment): This operator is a combination of ‘%’ and ‘=’ operators. This operator first modulo the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
  - `“&=”` (Bitwise AND Assignment): This operator is a combination of ‘&’ and ‘=’ operators. This operator first “Bitwise AND” the current value of the variable on the left by the value on the right and then assigns the result to the variable on the left.
  - `“^=”` (Bitwise Exclusive OR): This operator is a combination of ‘^’ and ‘=’ operators. This operator first “Bitwise Exclusive OR” the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
  - `“|=”` (Bitwise Inclusive OR): This operator is a combination of ‘|’ and ‘=’ operators. This operator first “Bitwise Inclusive OR” the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
  - `"<<="` (Left shift AND assignment operator): This operator is a combination of ‘<<’ and ‘=’ operators. This operator first “Left shift AND” the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.
  - `">>="` (Right shift AND assignment operator): This operator is a combination of ‘>>’ and ‘=’ operators. This operator first “Right shift AND” the current value of the variable on left by the value on the right and then assigns the result to the variable on the left.

```go
// Go program to illustrate the
// use of assignment operators
package main
  
import "fmt"
  
func main() {
   var p int = 45
    var q int = 50
     
   // “=”(Simple Assignment)
   p = q
   fmt.Println(p)
     
   // “+=”(Add Assignment)
    p += q
   fmt.Println(p)
     
   //“-=”(Subtract Assignment)
   p-=q
   fmt.Println(p)
     
   // “*=”(Multiply Assignment)
   p*= q
   fmt.Println(p)
     
   // “/=”(Division Assignment)
    p /= q
   fmt.Println(p)
    
    // “%=”(Modulus Assignment)
    p %= q
   fmt.Println(p)
    
}

```

## Misc Operators
- `&:` This operator returns the address of the variable.
- `*:` This operator provides pointer to a variable.
- `<-:` The name of this operator is receive. It is used to receive a value from the channel.

```
// Go program to illustrate the
// use of Misc Operators
package main
  
import "fmt"
  
func main() {
  a := 4
   
  // Using address of operator(&) and 
  // pointer indirection(*) operator
  b := &a 
  fmt.Println(*b) 
  *b = 7 
  fmt.Println(a) 
}
```


