# GoLang 学习笔记

## 常用命令
| 命令                                                        | 描述              |
|-----------------------------------------------------------|-----------------|
| go run .                                                  | 运行              |
| go mod tidy                                               | 下载依赖            |
| go mod init 404coder.com/greetings                        | 创建go.mod管理依赖··· |
| go mod edit -replace 404coder.com/greetings=../greetings  | 修改module指向      |


## 测试文件
- 文件结尾为_test.go， go test command认为此文件包含测试函数
- Test function names have the form `TestName`, where Name says something about the specific test.

## discover Go install path
`go list -f '{{.Target}}'`

- 设定go install安装go程序的目录。实际安装到/Users/zephyr/go/bin
  - echo 'export GOPATH=$HOME/go' >> ~/.zshrc
  - echo 'export PATH=$GOPATH/bin:$PATH' >> ~/.zshrc
  - 可以通过go env -w GOBIN=/path 来修改
  - go install 安装的文件名字，取决于go.mod中的module名字


# Tour of Go
引入包，使用最小的sub genre，可以引入parent包。大写开头的函数、值会被export。
Go可以声明类型，位置在对象之后，如： y int。如果一次声明多个，类型相同可以shorten，like：x, y int。
声明变量用var，var c, python, java = true, false, "no!"可以声明并赋值，等价于 c, python, java := true, false, "no!"
常量使用const声明，Constants cannot be declared using the := syntax。
```go
package main

import (
	"fmt"
    "math"
    "runtime"
	"time"
    "strings"
)
import "golang.org/x/tour/pic"
import "golang.org/x/tour/wc
func split(sum int) (x, y int) { 
    //声明加赋值
    a := 1
    x = sum * 4 / 9 + a
    y = sum - x
    return
}
//具名返回值，尽量啥少用，减弱代码可读性。函数可以返回任意数量返回值。

//for循环，不要条件括号，但是表达式必须用花括号括起来。
//The init and post statements are optional. condition expression is required.
func forFn() {
    sum := 0
    for i := 0; i < 10; i++ {
      sum += i
    }
	for sum < 1000 {
		sum += sum
    }
	// infinite loop, 其他语言while(1){} 等于 Go的for{}
    for {
    }
    fmt.Println(sum)
}

// https://go.dev/tour/flowcontrol/8 平方根获取
func Sqrt(x float64) float64 {
    z := x
    for i := 0; i < 10; i++ {
      z -= (z*z - x) / (2*z)
      fmt.Println(z)
    }
    return z
}

// if 语句 条件括号非必要，同样花括号braces{} are required
func ifFn() {
	condition := true
	if condition {
		fmt.Println("condition is true")
    }
}

// if 语句可以有一个类似 init的short statement，作用范围只在if以及else的braces中，出了if else就无效
func pow(x, n, lim float64) float64 {
    if v := math.Pow(x, n); v < lim {
      return v
    } else if v > lim {
      fmt.Printf("%g >= %g\n", v, lim)
    } else {
      fmt.Printf("%g >= %g\n", v, lim)
    }
    return lim
}

// switch 语句类似其他语言，但是Go语言case内置了break，不需要手动return or break
//TODO 尚未理解 Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.
// 意思是 case或者switch是可以是一个函数的返回值？
func switchFn() {
    fmt.Print("Go runs on ")
    switch os := runtime.GOOS; os {
    case "darwin":
      fmt.Println("OS X.")
    case "linux":
      fmt.Println("Linux.")
    default:
      // freebsd, openbsd,
      // plan9, windows...
      fmt.Printf("%s.\n", os)
    }
}

// switch with no condition 可以当作一个写复杂if-else的简洁方式. = switch true
func switchNoConditionFn() {
    t := time.Now()
    switch {
    case t.Hour() < 12:
      fmt.Println("Good morning!")
    case t.Hour() < 17:
      fmt.Println("Good afternoon.")
    default:
      fmt.Println("Good evening.")
    }
}

// defer 用于调整函数执行时机，defer处理的语句，会待其他函数执行结束返回后执行，但参数的计算是即时的
// defer机制为stack处理，first-in-last-out 多处调用defer时，会先调用后入栈的函数
func deferFn() {
    fmt.Println("counting")
  
    for i := 0; i < 10; i++ {
      defer fmt.Println(i)
    }
  
    fmt.Println("done")
}

// # More types: structs, slices, and maps.

// ## Pointer
// This is known as "dereferencing" or "indirecting".
// Go语言中有指针类型，C的指针和Go pointer什么区别？https://www.cnblogs.com/cheyunhua/p/15302200.html
// & 运算符取出变量所在的内存地址 、 * 运算符取出指针变量所指向的内存地址里面的值，也叫 “解引用”
// Go指针has no pointer arithmetic. 不能直接操作指针，即指向内存地址，指向整个数组的值，C中指向的是数组第一个元素的内存地址
func pointerFn() {
    var p *int
  
    i := 42
    p = &i
  
    fmt.Println(*p) // read i through the pointer p
    *p = 21         // set i through the pointer p
}

// ## Structs: A struct is a collection of fields.
type Vertex struct {
    X ,Y int
}
// pointer + struct 时，讲道理应该用 p = &v/ *p.X，但为了方便，直接用p.X
// 可以具名赋值， prefix & return a pointer to the struct value
func structFn() {
    fmt.Println(Vertex{1, 2})
  
    v := Vertex{X:1, Y:2} 
    v.X = 4
    fmt.Println(v.X)
}

// # Array
// ## The type [n]T is an array of n values of type T.
func ArrayFn() {
    var a [10]int
	fmt.Println(a)

    primes := [6]int{2, 3, 5, 7, 11, 13}
	
	//Slices
    slicePrimes := []int{2, 3, 5, 7, 11, 13}
    var s []int = primes[1:4]
    fmt.Println(s)
    fmt.Println(slicePrimes)
}

//# Slices A slice is formed by specifying two indices, a low and high bound, separated by a colon:
//## a[low : high]  This selects a half-open range which includes the first element, but excludes the last one.
func SliceFn() {
  // 可以认为Slice是对数组的引用，创建Slice的时候，先创建一个数组，之后Slice指向数组
  // This is an array literal:
  arr := [3]struct {
    i int
    b bool
  }{{2, true}, {3, false}, {5, true}}

  //And this creates the same array as above, then builds a slice that references it:
  slices := []struct {
    i int
    b bool
  }{{2, true}, {3, false}, {5, true}}

  fmt.Println(arr, slices)

  //	The zero value of a slice is nil.
  var s []int
  fmt.Println(s, len(s), cap(s))
  if s == nil {
    fmt.Println("nil!")
  }
  // log: [] 0 0  nil!

  // ## make function: built-in方法，this is how you create dynamically-sized arrays.
  // 第三个参数可以指定capacity
  b := make([]int, 0, 5) // len(b)=0, cap(b)=5

  b = b[:cap(b)] // len(b)=5, cap(b)=5
  b = b[1:]      // len(b)=4, cap(b)=4

  // ## append function: built-in方法，可以向slice添加元素 link:https://golang.design/go-questions/slice/grow/
  // If the backing array of s is too small to fit all the given values， a bigger array will be allocated. 
  // The returned slice will point to the newly allocated array.

  // ## range for循环中使用range处理slice或者map的时候，返回i,v 分别为index和value，
  // like JS.map,但是v是深拷贝，修改不影响源值。 v可以省略，只用index，
  // When ranging over a slice, two values are returned for each iteration.
  // The first is the index, and the second is a copy of the element at that index.
  pow1 := make([]int, 10)
  for i := range pow1 {
    pow1[i] = 1 << uint(i) // == 2**i
  }

  fmt.Println(pow1)
}

// slice exercise
func Pic(dx, dy int) [][]uint8 {
  matrix := make([][]uint8, dy)

  for i := 0; i < dy; i++ {
    matrix[i] = make([]uint8, dx)

    for j := 0; j < dy; j++ {
      //matrix[i][j] = uint8((i + j) / 2)
      //matrix[i][j] = uint8(i * j)
      matrix[i][j] = uint8(i ^ j)

    }
  }

  return matrix
}

func printPic() {
  pic.Show(Pic)
}

//# Map
// 定义map，空值为nil，make function可以return a map of the given type，初始化并ready for use
type VertexMap struct {
  Lat, Long float64
}

var m map[string]VertexMap

// 初始化语句
var mLiteral = map[string]VertexMap{
  "Bell Labs": VertexMap{
    40.68433, -74.39967,
  },
  "Google": VertexMap{
    37.42202, -122.08408,
  },
}

// If the top-level type is just a type name, you can omit it from the elements of the literal.
var mLiteral = map[string]VertexMap{
  "Bell Labs": {
    40.68433, -74.39967,
  },
  "Google": {
    37.42202, -122.08408,
  },
}

func mapFn() {
  fmt.Println(m)
  if m == nil {
    fmt.Println("m is nil")
  }
  
  m = make(map[string]VertexMap)
  fmt.Println(m, len(m))

  m["Bell Labs"] = VertexMap{
    40.68433, -74.39967,
  }
  m["Bell Labs2"] = VertexMap{
    40.68433, -74.39967,
  }
  if m == nil {
    fmt.Println("m is nil")
  }
  if m["1"] == (VertexMap{0, 0}) {
    fmt.Println("nil!")
  }

  fmt.Println(m["1"])
  fmt.Println(m)

  fmt.Println(m["Bell Labs"])
  
  //  删除某值使用delete
  delete(m, "Bell Labs2")
  // Test that a key is present 如果存在那么ok为true，v为其值
  // 若不存在那么ok为false, v的值取决于类型？
  v, ok := m["Answer"]
  fmt.Println("The value:", v, "Present?", ok)
}

// map exercise 
func WordCount(s string) map[string]int {
  words := strings.Fields(s)
  wordMap := make(map[string]int)

  for _, v := range words {
    wordMap[v]++
  }
  return wordMap
}

func mapExercise() {
  wc.Test(WordCount)
}



```

point 2023/4/23
https://go.dev/doc/
https://go.dev/tour/moretypes/24
https://golang.design/go-questions/slice/grow/