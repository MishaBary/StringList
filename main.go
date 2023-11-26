package main

import (
	"fmt"
	"myModules/str_list_hw/stringList"
)

func main() {
	strList := stringList.New("Misha Baryshev")
	fmt.Println("Initial string:", strList.String())

	strList.Append(' ')
	strList.Append('I')
	strList.Append('U')
	strList.Append('1')
	strList.Append('0')
	fmt.Println("strList.Append:", strList.String())

	strList.Prepend('_')
	fmt.Println("strList.Prepend:", strList.String())

	fmt.Println("strList.Len:", strList.Len())

	index := 1
	result, err := strList.At(index)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Symbol at index %d: %c\n", index, result)
	}

	strList.Insert('-', 20)
	fmt.Println("strList.Insert:", strList.String())

	strList.Remove(0)
	fmt.Println("strList.Remove:", strList.String())

	substring, _ := strList.Substring(6, 13)
	fmt.Println("strList.Substring:", substring.String())

	strList.Replace('M', 'm', 20)
	strList.Replace('B', 'b', 20)
	fmt.Println("strList.Replace:", strList.String())

	strList_2 := stringList.New("56")
	concatenated := strList.Concat(strList_2)
	fmt.Println("strList.Concat:", concatenated.String())
	fmt.Println("strList.Indexof:", strList.IndexOf('b'))

	strList_3 := stringList.New("65")
	fmt.Println("strList_2.Equals:", strList_2.Equals(strList_3))
	fmt.Println("strList_2.Equals:", strList_2.Equals(strList_2))
}
