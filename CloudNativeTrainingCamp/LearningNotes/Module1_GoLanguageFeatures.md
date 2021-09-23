# GO

## 统一思想-12factor

[12factor](https://12factor.net/)

## Why GO

GO Principles:

`Less is exponentially more`  – Rob Pike, Go Designer

`Do Less, Enable More`  – Russ Cox, Go Tech Lead

其他编程语言的弊端:

- 硬件发展速度远远超过软件。
- C 语言等原生语言缺乏好的依赖管理 (依赖头文件）。
- Java 和 C++ 等语言过于笨重。
- 系统语言对垃圾回收和并行计算等基础功能缺乏支持。
- 对多核计算机缺乏支持。

Go 语言是一个可以编译高效，支持高并发的，面向垃圾回收的全新语言。

- 秒级完成大型程序的单节点编译。
- 依赖管理清晰。
- 不支持继承，程序员无需花费精力定义不同类型之间的关系。
- 支持垃圾回收，支持并发执行，支持多线程通讯。
- 对多核计算机支持友好。

Go 语言不支持的特性

- 不支持函数重载和操作符重载
- 为了避免在 C/C++ 开发中的一些 Bug 和混乱，不支持隐式转换
- 支持接口抽象，不支持继承
- 不支持动态加载代码
- 不支持动态链接库
- 通过 `recover` 和 `panic` 来替代异常机制
- 不支持断言
- 不支持静态变量

## Start GO

1. Download and Install GO
2. Environment Variables:
   1. `GOROOT` -> GO安装目录
   2. `GOPATH` ->
      1. `src` -> 存放源代码
      2. `pkg` -> 存放依赖包
      3. `bin` -> 存放可执行文件
   3. Others Variables：
      1. `GOOS`, `GOARCH`, `GOPROXY`

### SOME CMD

```text
bug           start a bug report
**build**     compile packages and dependencies
clean         remove object files and cached files
doc           show documentation for package or symbol
env           print Go environment information
fix           update packages to use new APIs
**fmt**       gofmt (reformat) package sources
generate      generate Go files by processing source
**get**       add dependencies to current module and install them
**install**   compile and install packages and dependencies
list list     packages or modules
**mod**       module maintenance
run           compile and run Go program
**test**      test packages
**tool**      run specified go tool
version       print Go version
vet           report likely mistakes in packages
```

#### Go build

- Go 语言不支持动态链接，因此编译时会将所有依赖编译进同一个二进制文件。
- 指定输出目录
  - `go build –o bin/mybinary .`
- 常用环境变量设置编译操作系统和 CPU 架构
  - `GOOS=linux GOARCH=amd64 go build`
- 全支持列表
  - `$GOROOT/src/go/build/syslist.go`

#### Go test

Go 语言原生自带测试

```go
import "testing"
func TestIncrease(t *testing.T) {
    t.Log("Start testing")
    increase(1, 2)
}
```

`go test ./… -v` 运行测试

`go test` 命令扫描所有 `*_test.go` 为结尾的文件，惯例是将测试代码与正式代码放在同目录，
如 `foo.go` 的测试代码一般写在 `foo_test.go`

#### Go vet

代码静态检查，发现可能的 bug 或者可疑的构造

- `Print-format` 错误，检查类型不匹配的`print`

    ```go
    str := "hello world!"
    fmt.Printf("%d\n", str)
    ```

- `Boolean` 错误，检查一直为 `true`、`false` 或者冗余的表达式

    ```go
    fmt.Println(i != 0 || i != 1)
    ```

- `Range` 循环，比如如下代码主协程会先退出，`go routine`无法被执行

    ```go
    words := []string{"foo", "bar", "baz"}
    for _, word := range words {
        go func() {
            fmt.Println(word).
        }()
    }
    ```

- `Unreachable` 的代码，如 `return` 之后的代码
- 其他错误，比如变量自赋值，error 检查滞后等

    ```go
    res, err := http.Get("https://www.spreadsheetdb.io/")
    defer res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }
    ```

## Control Structure

### if

基本形式

```go
if condition1 {
    // do something 
} else if condition2 {
    // do something else 
} else {
    // catch-all or default
}
```

`if` 的简短语句

- 同 `for` 一样， `if` 语句可以在条件表达式前执行一个简单的语句

```go
if v := x - 100; v < 0 {
    return v
}
```

### switch

```go
switch var1 {
case val1: //空分支
case val2:
    fallthrough //执行case3中的f()
case val3:
    f()
default: //默认分支
    ...
}
```

### for

Go 只有一种循环结构：for 循环

计入计数器的循环

- for 初始化语句; 条件语句; 修饰语句 {}

    ```go
    for i := 0; i < 10; i++ {
        sum += i
    }
    ```

- 初始化语句和后置语句是可选的，此场景与 `while` 等价（Go 语言不支持 `while`）

    ```go
    for ; sum < 1000; {
        sum += sum
    }
    ```

- 无限循环

    ```go
    for {
        if condition1 {
            break
        }
    }
    ```

### for-range

遍历数组，切片，字符串，Map 等

```go
for index, char := range myString {
    ...
}
for key, value := range MyMap {
    ...
}
for index, value := range MyArray {
    ...
}
```

Notes: 如果 `for range` 遍历指针数组，则 `value` 取出的指 针地址为原指针地址的拷贝

## Common Data Structures

### 常量与变量

常量

`const identifier type`

变量

`var identifier type`

### 变量定义

- 变量
  - `var` 语句用于声明一个变量列表，跟函数的参数列表一样，类型在最后。
  - `var c, python, java bool`
- 变量的初始化
  - 变量声明可以包含初始值，每个变量对应一个。
  - 如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。
  - `var i, j int = 1, 2`
- 短变量声明
  - 在函数中，简洁赋值语句 `:=` 可在类型明确的地方代替 `var` 声明。
  - 函数外的每个语句都必须以关键字开始（`var`, `func` 等等），因此 `:=` 结构不能在函数外使用。
  - `c, python, java := true, false, "no!"`

### 类型转换与推导

- 类型转换 - 表达式 `T(v)` 将值 `v` 转换为类型 `T`。
  - 一些关于数值的转换：
    - `var i int = 42`
    - `var f float64 = float64(i)`
    - `var u uint = uint(f)`
  - 或者，更加简单的形式：
    - `i := 42`
    - `f := float64(i)`
    - `u := uint(f)`
- 类型推导
  - 在声明一个变量而不指定其类型时（即使用不带类型的 `:=` 语法或 `var =` 表达式语法），变量的类型由右值推导得出。
    - `var i int`
    - `j := i // j 也是一个 int`

### 数组

- 相同类型且长度固定连续内存片段
- 以编号访问每个元素
- 定义方法
  - `var identifier [len]type`
- 示例
  - `myArray := [3]int{1,2,3}`

### 切片(slice)

- 切片是对数组一个连续片段的引用
- 数组定义中不指定长度即为切片
  - `var identifier []type`
- 切片在未初始化之前默认为 `nil`， 长度为0
- 常用方法

    ```go
    func main() {
        myArray := [5]int{1, 2, 3, 4, 5}
        mySlice := myArray[1:3]
        fmt.Printf("mySlice %+v\n", mySlice)
        fullSlice := myArray[:]
        remove3rdItem := deleteItem(fullSlice, 2)
        fmt.Printf("remove3rdItem %+v\n", remove3rdItem)
    }
    
    func deleteItem(slice []int, index int) []int {
        return append(slice[:index], slice[index+1:]...)
    }
    ```

### Make 和 New

- New 返回指针地址
- Make 返回第一个元素，可预设内存空间，避免未来的内存拷贝
- 示例

    ```go
    mySlice1 := new([]int)
    mySlice2 := make([]int, 0)
    mySlice3 := make([]int, 10)
    mySlice4 := make([]int, 10, 20)
    // res
    // > MySlice1: <*[]int>(0xc00000c030)
    // > MySLice2: <[]int>(length: 0, cap: 0)
    // > MySLice3: <[]int>(length: 10, cap: 10)
    // > MySLice4: <[]int>(length: 10, cap: 20)
    ```

### 关于切片一些问题

```go
func main() {
    mySlice := []int{10, 20, 30, 40, 50}
    for _, value := range mySlice {
        value *= 2
    }
    fmt.Printf("mySlice %+v\n", mySlice)
    for index, _ := range mySlice {
        mySlice[index] *= 2
    }
    fmt.Printf("mySlice %+v\n", mySlice)
}
// mySlice [10 20 30 40 50]
// mySlice [20 40 60 80 100]
```

### Map

- 声明方法
- var map1 map[keytype]valuetype
- 示例

    ```go
    myMap := make(map[string]string, 10)
    myMap["a"] = "b"
    myFuncMap := map[string]func() int{
        "funcA": func() int { return 1 },
    }
    fmt.Println(myFuncMap)
    f := myFuncMap["funcA"]
    fmt.Println(f())
    ```

### 访问 Map 元素

- 按 Key 取值

    ```go
    value, exists := myMap["a"]
    if exists {
        println(value)
    }
    ```

- 遍历 Map

    ```go
    for k, v := range myMap {
        println(k, v)
    }
    ```

### 结构体和指针

- 通过 type … struct 关键字自定义结构体
- Go 语言支持指针，但不支持指针运算
  - 指针变量的值为内存地址
  - 未赋值的指针为 nil

``` go
type MyType struct {
Name string
}
func printMyType(t *MyType){
println(t.Name)
}
func main(){
t := MyType{Name: "test"}
printMyType(&t)
}
```

### 结构体标签

- 结构体中的字段除了有名字和类型外，还可以有一个可选的标签（tag）
- 使用场景：`Kubernetes APIServer` 对所有资源的定义都用 `Json tag` 和 `protoBuff tag`

```go
type MyType struct {
    // NodeName string `json:"nodeName,omitempty" protobuf:"bytes,10,opt,name=nodeName"`
    Name string `json:"name"`
}

func main() {
    mt := MyType{Name: "test"}
    myType := reflect.TypeOf(mt)
    name := myType.Field(0)
    tag := name.Tag.Get("json")
    println(tag)
}
```

### 类型重命名

```go
// Service Type string describes ingress methods for a service
type ServiceType string
const (
    // ServiceTypeClusterIP means a service will only be accessible inside the
    // cluster, via the ClusterIP.
    ServiceTypeClusterIP ServiceType = "ClusterIP"
    // ServiceTypeNodePort means a service will be exposed on one port of
    // every node, in addition to 'ClusterIP' type.
    ServiceTypeNodePort ServiceType = "NodePort"
    // ServiceTypeLoadBalancer means a service will be exposed via an
    // external load balancer (if the cloud provider supports it), in addition
    // to 'NodePort' type.
    ServiceTypeLoadBalancer ServiceType = "LoadBalancer"
    // ServiceTypeExternalName means a service consists of only a reference to
    // an external name that kubedns or equivalent will return as a CNAME
    // record, with no exposing or proxying of any pods involved.
    ServiceTypeExternalName ServiceType = "ExternalName"
)

```
