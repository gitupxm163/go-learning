package main

import (
	"fmt"
	"sort"
)

func main() {

	var map1 map[string]int
	fmt.Println(map1, map1 == nil)

	map2 := make(map[string]int)
	fmt.Println(map2, map2 == nil)

	map2["k1"] = 1
	fmt.Println(map2)

	v, ok := map2["k2"]

	if ok {
		fmt.Println("get v ok v = ",v)
	}else {
		fmt.Println("get v not ok")
	}

	 map3 :=  map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
	}

	var keys []string
	for key, val := range map3 {
		keys = append(keys, key)
		fmt.Println(key,val)
	}

	fmt.Println(keys)

	sort.Strings(keys)

	fmt.Println(keys)

	fmt.Println("=======")
	for _, key := range keys {
		fmt.Println(key, map3[key])
	}

	delete(map3,"d")

	fmt.Println(map3)

	map3["d"] = 4

	fmt.Println(map3)
}
