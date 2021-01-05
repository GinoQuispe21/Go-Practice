//In golang only exist the loop for, not exist while and do-while

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Number of the iteration: %d\n", i)
	}

	//We can use "for" as a "while" using the following format
	j := 0
	for {

		//if j == 2 {
		//	continue
		//}

		fmt.Println(j)
		j++
		if j > 10 {
			break
		}
	}
}
