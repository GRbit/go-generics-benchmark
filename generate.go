package main

import (
	"bufio"
	"os"
	"strconv"
)

const (
	rootFolder = "./generated/"
	nn         = 10002
)

func main() {
	err := os.MkdirAll(rootFolder, 0777)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(rootFolder + "dummy_main.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("package main\n\nfunc main() {\n}\n")
	if err != nil {
		panic(err)
	}
	defer writer.Flush()

	generateStatic()
	generateGenerics()
	generateFunctionCalls()
}

func generateStatic() {
	// generate huge number of static functions
	generate(rootFolder+"static.go", "package main\n\n", "",
		func(i int) string {

			return `
func MaxOfType` + strconv.Itoa(i) + `(a, b int) int {
	if a > b {
		return a
	}

	return b
}

`
		})
}

func generateGenerics() {
	s := "package main\n\n"
	s += `
type number interface {
	~int | ~int32 | ~int64 | ~float32 | ~float64 | ~uint
}
`
	generate(rootFolder+"generics.go", s, "",
		func(i int) string {

			return `
func MaxOfType` + strconv.Itoa(i) + `[T number](a, b T) T {
	if a > b {
		return a
	}

	return b
}

`
		})
}

func generateFunctionCalls() {
	// generate calls for every function
	generate(rootFolder+"calls.go", "package main\n\nfunc main() {\n", "}\n",
		func(i int) string {
			s := strconv.Itoa(i)
			return "\tMaxOfType" + s + "(" + s + ", " + s + ")\n"
		})

}

func generate(fileName, begining, ending string, f func(int) string) {
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(begining)
	if err != nil {
		panic(err)
	}
	defer writer.Flush()

	for i := 0; i < nn; i++ {
		_, err = writer.WriteString(f(i))
		if err != nil {
			panic(err)
		}
	}

	_, err = writer.WriteString(ending)
	if err != nil {
		panic(err)
	}
}

type xhf string

func (s xhf) String() string {
	return ""
}

type Stringer interface {
	String() string
	xhf
}

func Tos[T Stringer](s []T) []string {
	var ret []string
	for _, v := range s {
		ret = append(ret, v.String())
	}
	return ret
}
