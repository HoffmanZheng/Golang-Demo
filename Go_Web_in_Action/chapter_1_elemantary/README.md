# Getting started with Go

### Package

1. If you want initialize with a certain package without using it's function or variable, an **underscore** could be used during import process

   ```go
   import (
     _ "os"
     alias "fmt"
   )
   ```

2. A Go application could only have one main package, where the only one main function located at, without which an error occurs during compile process.

### Sematic

1. The left braces must located at the same line with function name, `if` and `for`, right braces must locate at the same line with `else if`.

2. Short variable declaration `name := expression` could only used inside the function, can't used as a global variable.
3. Local variable with the same name has a higher priority than global variable.
4. **Constants** could only be boolean, numeric and string. The expression declaring a constant must be able to be evaluated by the compiler.

### Control Statement

1. If lots of `else if` structures must be used, put the case with large possibility in the front.

2. Advanced control of loop is provided by Golang: `break <Tag>` and `goto`, both could interrupt a loop without any condition.

   ```go
   JumpLoop:
   for j := 0; j < 5; j++ {
     for i := 0; i < 5; i++ {
       if i > 2 {
       	break JumpLoop
       }
       fmt.Println(i)
     }
   }
   
   func main() {
     for x := 0; x < 20; x++ {
       for y := 0; y < 20; y++ {
         if y == 2 {
         	goto breakTag
         }
       }
     }
     return
     breakTag:
       fmt.Println("done")
   }
   ```

3. The `val` in the `for-range` is a copy and read only, thus the original value can not be modified during iteration.

4. Anonymous variable is a placeholder, which does not participate in the space allocation.

   ```go
   m := map[string]int{
     "shirdon": 100,
     "ronger": 98,
   }
   
   for _, value := range m {
     fmt.Println(value)
   }
   ```

5. Go improved the traditional `switch-case` statement. Code blocks between cases are independent, just like multiple `if-else`. Therefore it's no need to use break to jump out current case.

6. Multiple case value and expression could also be used in switch

   ```go
   var language = "golang"
   switch language {
     case "golang", "java":
       fmt.Println("popular languages")
   }
   
   var r int = 6
   switch {
     case r > 1 && r < 10:
       fmt.Println(r)
   }
   ```

### Data Type

1. Go also has several **schema-based** data types. Int and uint have 4-bytes on 32-bits OS, but have 8-bytes on 64-bits OS. 

2. String declared with double quote is interpretable, and supports escaping, whereas string declared with **backquote** is used to create native string literals and does not support escaping.

3. Byte in string could be exacted by using index: `str[index]`, but it's not reliable, because several characters may have multiple bytes. Thus a rune slice could be used:

   ```go
   str := "i love go web"
   chars := []rune(str)
   for _, char := range chars {
     fmt.Println(string(char))
   }
   ```

4. String is also immutable in Golang. Operation `+` and `+=` is not a efficient way, it's better to use StringBuilder:

   ```go
   var buffer bytes.Buffer
   for {
     if piece, ok := getNextString();ok {
       buffer.WriteString(piece)
     } else {
       break
     }
   }
   fmt.Println(buffer.String())
   ```

5. Pointer, See: [4.pointer.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go%20Web%20in%20Action/chapter_1_elemantary/4.pointer.go)

6. Array

   ```go
   var name[SIZE] type    // declaration
   var numbers = [5]float32{100.0, 8.0, 9.4, 6.8, 30.1}
   var numbers = [...]float32{100.0, 8.0, 9.4, 6.8, 30.1}
   ```

7. Slice is a reference to a continuous segment

   ```go
   var slicestr []string  // declaration 
   make([]Type, size, cap)
   ```

8. Map

   ```go
   var name map[key_type]value_type  // declaration
   createdMap := make(map[string]float32)  // could specify cap in the end
   literalMap = map[string]string{"first": "go", "second": "web"}
   ```

### Function

1. Function in Golang could return multiple value

   ```go
   func compare(x, y int) (int, int) {
     return x+y, x*y
   }
   ```

2. Variable parameter is also supported in Golang function

   ```go
   func myFunc(arg ...string) {
     ...   // arg is a string slice
   }
   ```

3. Function parameters are value transfered by default, which does not change the original variable. Reference transfer could be implemented by using pointer as parameters, which pass the address of variable and modify the variable during the call process.

   ```go
   func exchange(x *int, y *int) int {...}
   ```

4. Anonymous function (closure): could use variable in function without declaration. It could also be used for callback, passing the function as parameter.

   ```go
   func main() {
     x, y := 6, 8
     defer func(a int) {
       fmt.Println("defer x, y = ", a, y) //y is closure reference
       // defer x, y = 6 108
     }(x)   // passing parameter
     x += 10
     y += 100
     fmt.Println(x, y)   // 16 108
   }
   
   ////////
   
   func main() {
     //call after declaration
     f := func(data int) {
       fmt.Println("closure", data)
     }
     f(6)
   
     // call directly
     func(data int) {
       fmt.Println("closure, directly", data)
     }(8)
   }
   
   /////////
   
   func visitPrint(list []int, f func(int)) {  // callback function
     for _, value := range list {
       f(value)
     }
   }
   
   func main() {
     sli := []int{1, 6, 8}
     visitPrint(sli, func(value int) {
       fmt.Println(value)
     })
   }
   ```

5. `defer` delay function performs action **after** return or panic of function. defer statements would be pushed into a stack, and pop out when function finished, so it's in LIFO order.

   ```go
   var name string = "go"
   
   func myfunc() string {
     defer func() {
       name = "python"	
     }()
   
     fmt.Printf("name in myfunc: %s\n", name)
     return name	
   }
   
   func main() {
     myname := myfunc()
     fmt.Printf("name in main func: %s\n", name)
     fmt.Printf("myname in main func: %s\n", myname)
   }
   
   // name in myfunc: go
   // name in main func: python
   // myname in main func: go    // return before defer
   ```

### OOP

1. Encapsulation:

   1. Encapsulate attributes in struct, the capitalization of the first letter is used to control the accessibility, rather than public, private modifier.

   2. Method in Golang is a function working on a receiver.

      ```go
      func (recv recv_type) methodName(parameter_list) (return_value_list) {...}
      ```

2. Inherit

   1. There is no `extends` keyword in Golang, inherit could be implemented by embedding anonymous types in struct.

      ```go
      type Engine interface {
        Run()
        Stop()
      }
      
      type Bus struct {
        Engine
      }
      
      func (c *Bus) Working() {
        c.Run()
        c.Stop()
      }
      ```

3. Polymorphism

   1. Implements of interface with different receiver 

### Reflection

1. Reflection is that the computer observes and modifies it's behavior in it's runtime. This makes it possible for the variable type to be unknown at compile time. 
2. `TypeOf()` and `ValueOf()` could extract type and value info from a variable, which would be transfered to `interface{}` at the first place.
3. Reflected objectcould also be re-transfered to interface type, via `Value.Interface()`.
4. Reflected object is only modifiable if it's value is settable， which could be tested via `Value.CanSet()`. To enable the modifiability of reflected object, object address should be used to create reflected object, and `Elem()` to return the pointed data. See: [5.reflect.go](https://github.com/HoffmanZheng/Golang-Demo/blob/master/Go%20Web%20in%20Action/chapter_1_elemantary/5.reflect.go)





