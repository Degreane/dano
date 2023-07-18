package main

import (
	"fmt"
	"time"

	"github.com/degreane/dano/utils"
)

func main() {
	start := time.Now()
	// utils.RunFyne()
	utils.InitGui()
	// for i := 1; i <= 40000; i++ {
	// 	// fmt.Printf("%d=>%s\n========\n\n", i, utils.GetName(i, ""))
	// 	utils.GetName(i, "")
	// }
	fmt.Println(utils.GetName(18000, ""))

	duration := time.Since(start)
	fmt.Printf("Took %+v\n", duration)
	utils.RunGUI()
}
