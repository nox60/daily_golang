package main

import "fmt"

type FlyWeight struct {
	data string
}

var F = make(map[string]*FlyWeight)

func GetFLyWeight(fileName string) *FlyWeight {
	image := F[fileName]
	if image == nil {
		image = &FlyWeight{data: fileName}
		F[fileName] = image
	}

	return image
}

func main() {
	fmt.Println(">>>>>>")
	a := *GetFLyWeight("a11")
	fmt.Println("a11.data: ", a.data, &a)

	a12 := *GetFLyWeight("a12")
	fmt.Println("a12.data: ", a12.data, &a12)

	a13 := *GetFLyWeight("a13")
	fmt.Println("a13.data: ", a13.data, &a13)

	a14 := *GetFLyWeight("a11")
	fmt.Println("a14.data: ", a14.data, &a14)

}
