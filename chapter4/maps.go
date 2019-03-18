// p. 123
package main

import "fmt"

func main() {
	// ages := make(map[string]int)
	// ages["alice"] = 31
	// ages["charlie"] = 34
	
	ages := map[string]int{
		"alice": 31,
		"charlie": 34,
	}

	emptyMap := map[string]int{}

	fmt.Println(ages)
	delete(ages, "alice")
	fmt.Println(ages)
	fmt.Println(emptyMap)
	fmt.Println(emptyMap["Bob"] + 1)

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	if _, ok := ages["Bog"]; !ok {
		fmt.Println("No key")
	}
}
