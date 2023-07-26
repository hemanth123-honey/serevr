/*
package main

import "fmt"

func add(a, b int) int {
	sum := 0

	sum = a * b

	return sum
}

func main() {
	a := 5
	b := 6
	fmt.Println(add(a, b))
	fmt.Println("hi")
}
*/

/* package main

import "fmt"

func add(a, b int) {
	var c = a + b
	fmt.Println(c)

}

func main() {
	x := 9
	y := 10
	add(x, y)
}
*/

/* package main

import "fmt"

func rect(a, b int) (perimeter, area int) {

	len := a*b + 2
	fmt.Println(len)
	br := a + b
	fmt.Println(br)
	area = len + br
	perimeter = len * br

	return
}

func main() {
	a, b := rect(5, 6)
	fmt.Println(a, b)
}
*/
/*
package main

import "fmt"

func name(age *int, name *string) (int, string) {
	age1 := *age + 5
	name1 := *name + "kumar"
	return age1, name1
}

func main() {
	a := 20
	b := "hemanth"
	fmt.Println(a, b)
	c, d := name(&a, &b)
	fmt.Println(c, d)

}
*/
/*
package main

import "fmt"

func Sum(x ...int) int {
	fmt.Println(x)
	fmt.Println("%T", x)

	sum := 0

	for _, v := range x {
		sum = sum + v
	}
	return sum
}

func main() {

	x := []int{1, 2, 5}

	sum := Sum(x...)
	fmt.Println(sum)
}
*/

//anonymous function : write a function in main() without name
/*
package main

import "fmt"

func foo() {
	fmt.Println("how")

}

func main() {

	foo()

	func() {
		fmt.Println("hi")
	}()

	func(x int) {
		fmt.Println(x)
	}(10)

}
*/
// func expression  : write a function by declaring a variable
/*
package main

import "fmt"

func foo() {
	fmt.Println("how")

}

func main() {
	foo()

	f := func() {
		fmt.Println("hi")

	}
	f()

	f1 := func(x int) {
		fmt.Println("hello", x)
	}
	f1(5)

}
*/

//returning a func  :  return a function from a function
/*
package main

import "fmt"

func main() {
	f := func(x, y int) int {
		return x + y

	}(5, 6)
	fmt.Println(f)

	s := bar()
	fmt.Printf("%T\n", s)

	s1 := s()
	fmt.Println(s1)

}

func bar() func() int {
	return func() int {
		return 500
	}
}
*/

/*
package main

import "fmt"

func foo() func() (string, int) {
	return func() (string, int) {
		return "hemanth", 26
	}
}

func main() {

	fmt.Println(foo()())
	fmt.Println("=================")
	s := foo()
	fmt.Printf("%T\n", s)
	fmt.Println("=================")

	fmt.Println(s())
	fmt.Println("=================")

	x, y := s()
	fmt.Println(x, y)
}
*/

//Callback : passing a func as an argument
/*
package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8}
	tot := sum(x...)
	//tot := sum(1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println("all values sum: ", tot)

	xy := even(sum, x...)
	fmt.Println("even values sum: ", xy)

	xz := odd(sum, x...)
	fmt.Println("odd values sum: ", xz)



}

func sum(x ...int) int {
	tot := 0
	for _, v := range x {
		tot = tot + v
	}
	return tot
}

func even(f func(x ...int) int, v ...int) int {

	var y []int
	for _, vi := range v {

		if vi%2 == 0 {
			y = append(y, vi)
		}
	}
	return f(y...)

}

func odd(f func(x ...int) int, v ...int) int {

	var y []int
	for _, vi := range v {

		if vi%2 != 0 {
			y = append(y, vi)
		}
	}
	return f(y...)

}
*/

//closure : increment
/*
package main

import "fmt"

func main() {

	a := increntor()
	b := increntor()

	fmt.Println(a()) //1
	fmt.Println(a()) //2
	fmt.Println(a()) //3
	fmt.Println(b()) //1
	fmt.Println(b()) //2

}

func increntor() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
*/

// Recursion : func calls it self
/*
package main

import "fmt"

func main() {
	x := factorial(3)
	fmt.Println(x)

	y := factloop(4)
	fmt.Println(y)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
	// n * factorial(3-1) * factorial(2-1) * 1
}

func factloop(n int) int {
	f := 1

	for ; n > 0; n-- {
		f *= n
	}
	return f
}
*/