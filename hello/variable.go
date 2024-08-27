package main

import (
	"fmt"
	"math"
	"strconv"
	// "hex"
)

var f, g = 1, "durian"

// 方法 结构体
type Circle struct {
	radius float64
}

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {

	// vari()
	// operator()
	// ifSwitch()
	// fortest()
	// funtest()
	// arr()
	// structLearn()
	// sliceLearn()
	// rangeLearn()
	// MapLearn()

	// recursion()

	// strconvLearn()

}

func strconvLearn() {
	// 整型转浮点型
	a := 10
	b := float64(a) / 25
	fmt.Println(b)
	// 浮点型转整型
	c := 10.5
	d := int(c)
	fmt.Println(d)
	// 字符串转整型
	str := "123h"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(num)
	// 整型转字符串
	str = strconv.Itoa(a)
	fmt.Println(str)

	// 字符串转换为浮点数
	str = "3.14159x"
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)

	// 浮点数转换为字符串
	str = strconv.FormatFloat(f, 'f', 6, 64)

	fmt.Println(str)

	// 接口转换
	i := interface{}("3.14")
	fmt.Println(i)
	stri, ok := i.(string)
	if ok {
		fmt.Println(stri)
	} else {
		fmt.Println("not string")
	}

	// 字符串转布尔值
	str = "true"
	var boo bool
	boo, err = strconv.ParseBool(str)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(boo)
	}

	// 布尔值转字符串
	str = strconv.FormatBool(true)
	fmt.Println(str)

	// 字符串转rune
	str = "abc"
	runes := []rune(str)
	fmt.Println(runes)
	str2 := string(runes)
	fmt.Println(str2)

	// rune转字符串
	str = string(runes[0])
	fmt.Println(str)

	// r, size, err1 := strconv.UnquoteChar(str, 0)
	// if err1 != nil {
	// 	fmt.Println(err1)
	// } else {
	// 	fmt.Println(r, size)
	// }

	// rune转字符串
	str = strconv.QuoteRune('a')
	fmt.Println(str)

	// 字符串转字节切片
	// str = "hello world"
	// var slice []byte
	// slice, err = hex.DecodeString(str)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(slice)
	// }

	// 字节切片转字符串
	// b = []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64}
	// str = hex.EncodeToString(b)
	// fmt.Println(str)

	// 字符串转字节切片
	str = "hello world"
	b2 := []byte(str)
	fmt.Println(b2)

	// 字节切片转字符串
	str = string(b2)
	fmt.Println(str)

}

func recursion() {
	// fmt.Println(factorial(3))
	// fmt.Println(fibonacci(4))
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fibonacci(i))
	}
}

func factorial(n uint64) uint64 {
	if n > 1 {
		return n * factorial(n-1)
	}
	return n
}

/*
斐波那契：
n = 4，取第四位斐波那契数列
进入斐波那契函数，返回a：fib（3）+ b：fib（2）
a: fib（3）返回a1:fib（2）+ a2:fib（1）
a1:fib（2）返回a11: fib（1）+ a12: fib（0）= 1
a2:fib（1）= 1,返回1
a = a1 + a2 = 2
b = fib（1）+ fib（0）= 1
返回a + b = 3
*/
func fibonacci(n int) int {
	if n <= 1 {
		// fmt.Println("n <= 1 = ",n)
		return n
	}
	// fmt.Println(n)
	return fibonacci(n-1) + fibonacci(n-2)
}

func MapLearn() {
	// m := make(map[string]int)
	m1 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	k := m1["a"]
	fmt.Println(k)
	k1, v1 := m1["a"]
	fmt.Println(k1, v1)
	m1["a"] = 111
	// fmt.Println(m1["a"])
	fmt.Println(m1)
	fmt.Println("map 长度", len(m1))
	m1["dd"] = 77
	fmt.Println("map 长度", len(m1))
	fmt.Println(m1)
	// 删除元素
	delete(m1, "b")

	fmt.Println(m1)
	fmt.Println("map 长度", len(m1))

	if v, ok := m1["rr"]; ok {
		fmt.Println("value is", v)
	} else {
		fmt.Println("key is not exist", v, ok)
	}

}

func rangeLearn() {
	// map1 := make(map[string]int)
	map1 := map[string]int{}

	map1["1"] = 111
	map1["key2"] = 222
	// map1["name3"] = 333
	// map1["k4"] = 444

	for k, v := range map1 {
		fmt.Println(k, v)
	}
	for k := range map1 {
		fmt.Println(k)
	}
	for _, v := range map1 {
		fmt.Println(v)
	}

	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}

}

func sliceLearn() {
	/*
		声明一个切片
		var s []string = make([]string, len, cap)
		s := make([]string, len, cap)

		切片和数组的区别：数组是固定长度，切片是动态长度可改变
		声明数组：arr := [3]int{1,2,3}
		声明切片：s := []int{1,2,3}

		len和cap的区别：len表示切片的长度，也就是元素个数，cap表示切片的容量，容量是指切片的最大长度，当切片的长度超过了容量时，会重新分配内存，扩大容量

		切片扩容策略：
		当切片的当前长度超过容量时，Go 会分配一个新的底层数组。
		新数组的容量通常是当前容量的两倍（如果当前容量小于 1024）。
		如果当前容量大于或等于 1024，新数组的容量通常是当前容量的 1.25 倍。
	*/
	arr := [3]int{1, 2, 3}
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(arr), cap(arr), arr)
	// arr = append(arr, 4)   报错，不可改变长度，不能再次添加元素
	// fmt.Printf("len = %d, cap = %d, slice = %v\n", len(arr), cap(arr), arr)

	numbers := make([]int, 3, 5)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	s := []string{"a", "b"}
	fmt.Println(s[0:1])
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s), cap(s), s)
	s = append(s, "c")
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s), cap(s), s)
	s = append(s, "d", "e", "f")
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s), cap(s), s)

	// 获取子切片，初始化新切片
	s1 := s[1:3]
	fmt.Printf("s1: len = %d, cap = %d, slice = %v\n", len(s1), cap(s1), s1)
	s1 = append(s1, "g", "h")
	fmt.Printf("s1: len = %d, cap = %d, slice = %v\n", len(s1), cap(s1), s1)

	// 复制切片
	// var s2 []string
	s2 := make([]string, 3)
	copy(s2, s[3:4])
	fmt.Printf("s2: len = %d, cap = %d, slice = %v\n", len(s2), cap(s2), s2)
	copy(s2[1:], []string{"a"})
	fmt.Printf("s2: len = %d, cap = %d, slice = %v\n", len(s2), cap(s2), s2)

}

// 结构体
func structLearn() {

	// type Books struct {
	// 	title string
	// 	author string
	// 	subject string
	// 	book_id int
	// }

	// 同时创建并输出结构体
	fmt.Println(Books{"Java 语言", "you", "code", 2089})
	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Python 语言", author: "he", subject: "code", book_id: 10034})

	// 忽略字段
	fmt.Println(Books{title: "Python 语言", subject: "code"})

	// 创建并输出结构体
	book1 := Books{"Go 语言", "me", "code", 10045}
	fmt.Println(book1)
	fmt.Println(book1.title)

	// 结构体作为参数
	printBook(book1)
	fmt.Println(getSubject(book1))

	// 结构体指针
	book2 := &book1
	fmt.Println(*book2)
	updateBooks(book2)
	fmt.Println(*book2)

}
func updateBooks(book *Books) {
	book.title = "golang语言"
}

func printBook(book Books) {
	fmt.Println(book.title)
}

func getSubject(book Books) string {
	return book.subject
}

func arr() {
	var numbers [3]int
	numbers = [3]int{3, 6, 9}
	var numbers1 [3]int = [3]int{1, 2, 3}
	numbers2 := [3]int{5, 5, 5}
	numbers3 := [...]int{6, 6, 6}
	numbers4 := []int{}
	numbers4 = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println(numbers1)
	fmt.Println(numbers2)
	fmt.Println(numbers3)

	fmt.Println(getSum(numbers4))
	fmt.Println(numbers4)
	updateArr(&numbers4)
	fmt.Println(numbers4)

	for i := 0; i < len(numbers); i++ {
		numbers[i] += i
	}
	fmt.Println(numbers)

	// 多维数组
	var arr1 [2][1]int = [2][1]int{{1}, {2}}
	fmt.Println(arr1)

	arr2 := [][]string{}
	fmt.Println(arr2)

	// arr2[0][0] = "01"
	row1 := []string{"a", "b", "c"}
	// row2 := []string{"d", "e", "f"}
	arr2 = append(arr2, row1)
	// arr2 = append(arr2, row1)
	arr2 = append(arr2, []string{"g", "h"})
	arr2 = append(arr2, []string{"j", "k", "l", "m"})
	fmt.Println(arr2)

	for i := range arr2 {
		for j := range arr2[i] {
			fmt.Println(arr2[i][j])
		}
	}

	for i, value := range arr2 {
		for j, v := range value {
			fmt.Println(i, j, v)
		}
	}

}

// 函数接收数组的指针
func updateArr(arr *[]int) {
	sum := 0
	for i, v := range *arr {
		sum += v
		(*arr)[i] = sum
	}
	fmt.Println(*arr)
	// return sum
}

// 函数接收数组
func getSum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

// 函数
func funtest() {
	fmt.Printf("%d\n", max(10, 20))
	fmt.Println(swap1(3, "bf"))

	// 函数作为实参
	getSquare := func(x float64) float64 {
		return math.Sqrt(x)
	}
	fmt.Println(getSquare(9))

	// 引用传递
	x := 10
	y := 20
	fmt.Println(x, y)
	swap2(&x, &y)
	fmt.Println(x, y)

	// 闭包，匿名函数
	//nextNumber 为一个函数，函数 i 为 0
	nextNumber := getSequence()
	// 调用 nextNumber 函数，i 变量自增 1 并返回
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	// 创建新的函数 nextNumber1，并查看结果
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())

	// 匿名函数
	anonymous()

	// 方法
	method()

}

func method() {
	var c1 Circle
	c1.radius = 10.00
	fmt.Println("Area of circle is:", c1.getArea())

	var c2 Circle
	c2.radius = 111.00
	fmt.Println("Radius: ", c2.getRadius())

}

// 定义方法，该方法属于Circle类型对象中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

func (cc Circle) getRadius() float64 {
	return cc.radius
}

// 匿名函数，也就是闭包，匿名函数的优点是可以直接使用函数内的变量，不必申明
// 通常用于在函数内部定义函数，或者作为函数参数进行传递
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 多个匿名函数实例
func anonymous() {
	// 定义一个匿名函数并将其赋值给变量add
	add := func(a, b int) int {
		return a + b
	}

	// 调用匿名函数add
	result := add(3, 4)
	fmt.Println("3 + 4 = ", result)

	// 在函数内部使用匿名函数
	multiply := func(x, y int) int {
		return x * y
	}

	result = multiply(2, 3)

	// 将匿名函数作为参数传递给其他函数

	calculate := func(operation func(a, b int) int, x, y int) int {
		return operation(x, y)
	}

	sum := calculate(add, 3, 4)
	fmt.Println("3 + 4 = ", sum)

	// 也可以直接在函数调用中定义匿名函数
	different := calculate(func(a, b int) int {
		return a - b
	}, 10, 4)

	fmt.Println("10 - 4 = ", different)

}

// 引用传递，可改变实参的值
func swap2(x *int, y *int) {
	*x, *y = *y, *x
}

func swap1(x int, y string) (int, string) {
	return x, y
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func fortest() {
	for i := 10; i > 0; i-- {
		fmt.Printf("%d", i)
		if i == 1 {
			fmt.Println("")
		}
	}
	i := 3
	for i < 30 {
		i += i
		fmt.Printf("%d ", i)
	}
	fmt.Println("")
	i = 3
	for i < 30 {
		i += i
		fmt.Printf("%d ", i)
	}
	fmt.Println("")

	str := []string{"go", "python", "java"}
	for i, s := range str {
		fmt.Println(i, s)
	}

	numbers := []int{1, 2, 3}

	for i, num := range numbers {
		fmt.Printf("第%d个数为：%d\n", i, num)
	}

	//re:标记
	// break使用标记时，只会退出当前标记所在的for循环;
	// 如下例，标记在i循环上，使用break时，直接退出i循环，标记在j循环上，退出j循环，继续下一个i循环，如不加标记的break
	// continue使用标记时，会直接跳过当前标记所在的for循环，继续执行下一个循环。
	// 如下例，标记在i循环上，使用continue时，会跳过当前i循环，执行下一个i循环，功能如break；标记在j循环，相当于未使用标记，跳过当前j循环，执行下一个j循环，功能如不加标记的continue。
re:
	for i := 0; i < 2; i++ {
		// re:
		for j := 0; j < 5; j++ {
			// if j == 3 {
			// 	break re
			// }
			if j == 3 {
				continue re
			}
			// fmt.Printf("i=%d,j=%d\n",i,j)
		}
	}

	i = 10
	fmt.Println("i = ", i)
LOOP111:
	for i > 0 {
		i -= 3
		if i < 3 {
			break LOOP111 // 退出循环，同break
		}
		fmt.Println("循环里的i = ", i)
	}

}

func operator() {

	a := 1.2
	b := 2.0
	fmt.Println("a * b =", a*b)
	fmt.Println("a / b =", a/b)

	a = 5
	b = 2
	a--
	fmt.Println(a)

	if a != b {
		fmt.Println("a != b", a != b)
	} else if a <= b {
		fmt.Println("a <= b", a <= b)
	} else {
		fmt.Println("a > b", a > b)
	}

	c := true
	d := true
	if c && d {
		fmt.Println("c && d", c && d)
	} else if c || d {
		fmt.Println("c || d", c || d)
	} else if !(c && d) {
		fmt.Println("!(c & d)", !(c && d))
	}

	// 位运算符
	// & 按位与运算符，参与运算的两个值,如果两个相应位都为1,则该位的结果为1,否则为0.
	// | 按位或运算符，参与运算的两个值,如果两个相应位中有1,则该位的结果为1,否则为0.
	// ^ 按位异或运算符，参与运算的两个值,如果两个相应位相异,则该位的结果为1,否则为0.

	// 左移运算符 << ：将运算数的各二进位全部左移若干位，高位丢弃，低位补0。
	// 右移运算符 >> ：将运算数的各二进位全部右移若干位，低位丢弃，高位补0。
	// 无符号右移运算符 >>> ：将运算数的各二进位全部右移若干位，低位丢弃，高位补0。

	e := 12
	e1 := 1

	fmt.Println(e << 2)
	fmt.Println(e1 >> 1)
	fmt.Println(string(48))

	a += b
	fmt.Println(a)
	a1 := 10
	b1 := 3
	a1 %= b1
	a1 = 10
	fmt.Println(a1)
	a1 /= b1
	fmt.Println("a1 = ", a1)
	a1 <<= 2
	fmt.Println("a1 <<= 2: ", a1)
	a1 ^= b1
	fmt.Println("a1 ^= b1: ", a1)
	a1 |= b1
	fmt.Println("a1 |= b1: ", a1)

	fmt.Printf("%T\n", a)

	ptr := &a1
	fmt.Println(a1)
	fmt.Println(*ptr)

}

func ifSwitch() {
	grade := "B"
	marks := 90
	switch {
	case true:
		grade = "A+"
		// fmt.Println("grade = A")
	case marks >= 90:
		grade = "A"
	case marks >= 80, marks >= 70:
		grade = "B"
	case false, marks < 50:
		grade = "C"
	default:
		grade = "D"
	}
	fmt.Println("grade = ", grade)

	var i interface{}
	i = 10
	switch r := i.(type) {
	case int:
		fmt.Println("marks is int")
	case float64:
		fmt.Println("marks is float64")
	default:
		fmt.Println("marks is other type", r)
	}

	// fallthrough 会强制执行下一条case，无论case的表达式结果是否为true
	switch {
	case false:
		fmt.Println("true")
		fallthrough
	case false:
		fmt.Println("false1")
		fallthrough
	case false:
		fmt.Println("fasle2")
	case true:
		fmt.Println("true2")
	default:
		fmt.Println("default")
	}
}

func vari() {
	var a string = "durain"
	fmt.Println(a)

	var b, c = 1, 2
	fmt.Println(b, c)

	var d bool
	fmt.Println(d)

	var e []int
	fmt.Println(e)

	// var f = [1, 2]int
	// fmt.println(f)

	v_name := "durain"
	fmt.Println(v_name)

	var a1, a2 int
	a1, a2 = 3, 4
	fmt.Println(a1, a2)

	a3, a4 := 5, 6
	fmt.Println(a3, a4)

	var (
		b1 = 1
		d1 = true
	)
	fmt.Println(b1, d1)

	fmt.Println(f, g)

	var f1 = f

	fmt.Println(&f, &f1)

	//定义常量
	const LENGTH = 10
	fmt.Println(LENGTH)

	const (
		PI = 3.14159265358979323846
		PI1
	)
	fmt.Println(PI, PI1)

	// 变量的值不可延用上一个变量的值，如y必须要赋值，不能直接跟在x下面
	// var (
	// 	x = 1
	// 	y
	// )
	// fmt.Println(x, y)

	const (
		h1 = iota
		h2
		h3 = "hi"
		h4
		h5 = iota
	)
	fmt.Println(h1, h2, h3, h4, h5)

	// << 左移运算符，将运算数的各二进位全部左移若干位，由"<<"右边的数字指定移动的位数，高位丢弃，低位补0。
	// i1 = 1<<iota: 1左移0位，结果为1,
	// i2 = 1<<iota: 1左移1位，结果为2，1左移一位变为二进制10，即2
	// i3 = 1<<iota: 1左移2位，结果为4，1左移两位变为二进制100，即4
	// i4 = 2<<iota: 2左移3位，结果为16，2左移3位变为二进制10000，即16
	// i5 = 2<<iota: 2左移4位，结果为32，1左移三位变为二进制100000，即32
	// i6 = 2<<iota: 2左移5位，结果为64，1左移四位变为二进制1000000，即64
	const (
		i1 = 1 << iota
		i2
		i3
		i4 = 2 << iota
		i5
		i6
	)
	fmt.Println(i1, i2, i3, i4, i5, i6)
}
