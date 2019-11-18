package main

type FlyWeight struct {
	data string
}

var F map[string]*FlyWeight

func GetFLyWeight(fileName string) *FlyWeight {
	image := F[fileName]
	if image == nil {

	}
	/*
		image := F[filename]
		if image == nil {
			image = NewImageFlyweight(filename)
			f.maps[filename] = image
		}

	*/
}
