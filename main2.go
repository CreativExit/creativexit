package main

import (
	"fmt"
	"strings"
)

func main() {
	//var b bool
	//fmt.Println(b)
	//
	//c := true
	//fmt.Println(c)
	//
	//dog := "Rocoo"
	//
	//fmt.Printf("Tipo: %T, Valor %v\n", dog, dog)
	//
	//animales := [3]string{"perro", "gato", "delfin"}
	//
	//fmt.Println(animales)
	//
	//for i := 1; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//
	//comoSeLlama("Christian", "Gonzalez")
	//
	//total := sum(5, 10)
	//fmt.Println(total)
	//
	//texto := "AlejoDr"
	//
	//minusc, mayusc := convert(texto)
	//
	//fmt.Println(minusc, mayusc)
	//
	//content, err := os.ReadFile("./hello.txt")
	//if err != nil {
	//	fmt.Println("Ocurrio un error: %v", err)
	//	return
	//}
	//fmt.Println(string(content))

	//result, err := division(10, 0)
	//
	//if err != nil {
	//	fmt.Println("Ocurrio un error: %v", err)
	//	return
	//}
	//fmt.Println(result)

	// Funcuion que devuelve funcion
	//nums := []int{1, 4, 70, 5, 67, 90, 2}
	//result := filter(nums, func(num int) bool {
	//	if num >= 10 {
	//		return true
	//	}
	//	return false
	//})
	//
	//fmt.Println(result)

	//x := hello("Alejandro")("Gionzalesz")
	//fmt.Println(x)
	//fmt.Println(sum2(1, 4, 6, 7, 78, 8))

	//func(text string) {
	//	fmt.Println("Hello Christian" + text)
	//}("Gonzalez")
	//x := func() {
	//	fmt.Println("Hello Christian")
	//}
	//
	//x()

	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)

	//a := 5
	//defer fmt.Println("Defer", a)
	//
	//a = 10
	//fmt.Println(a)

	// defer
	//file, err := os.Create("hello.txt")
	//if err != nil {
	//	fmt.Printf("Ocurrio un error al crear el archivo %v", err)
	//	return
	//}
	//defer file.Close()
	//
	//_, err = file.Write([]byte("Hello Christian Gonzalez"))
	//if err != nil {
	//	fmt.Printf("Ocurrio un error al escribir el archivo %v", err)
	//	return
	//}

	// Panic
	//division(10, 2)
	//division(40, 5)
	//division(70, 0)
	//division(10, 2)

}

func division(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recuperandome del panico", r)
		}
	}()
	validarDivisor(divisor)
	fmt.Println(dividendo / divisor)

}

func validarDivisor(divisor int) {
	if divisor == 0 {
		panic("Panic")
	}
}

func sum2(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func hello(nombre string) func(string) string {
	return func(text string) string {
		return "Hello " + nombre + " " + text
	}
}

func filter(nums []int, callback func(int) bool) []int {
	result := []int{}
	for _, v := range nums {
		if callback(v) {
			result = append(result, v)
		}
	}
	return result

}

//func division(dividendo, divisor int) (int, error) {
//
//	if divisor == 0 {
//		return 0, errors.New("No puedes dividir por 0")
//	}
//
//	return dividendo / divisor, nil
//
//}

func convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)
	return min, may
}

func sum(sum1 int, sum2 int) int {
	return sum1 + sum2
}

func hola() {
	fmt.Println("Hola")
}

func comoSeLlama(firstName string, lastname string) {
	fmt.Println("Hello %s %s", firstName, lastname)
}
