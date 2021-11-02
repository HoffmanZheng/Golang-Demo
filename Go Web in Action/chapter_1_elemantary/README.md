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

1. Go also has several schema-based data types. Int and uint have 4-bytes on 32-bits OS, but have 8-bytes on 64-bits OS. 

