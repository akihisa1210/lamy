package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	target := os.Args[1]
	var answer string
	fmt.Printf("[類] %s は何の一種か？\n", target)
	fmt.Scan(&answer)
	fmt.Printf("[種差] %s は、同じグループの中で他と何が違うのか？\n", target)
	fmt.Scan(&answer)
	fmt.Printf("[部分] %s を構成する部分を列挙すると？\n", target)
	fmt.Scan(&answer)
	fmt.Printf("[定義] %s とは何か？\n", target)
	fmt.Scan(&answer)
	fmt.Printf("[語源] %s の語源は？\n", target)
	fmt.Scan(&answer)
	fmt.Printf("[相反] %s の反対は？\n", target)
	fmt.Scan(&answer)
	fmt.Printf("[原因・由来] %s を生じさせる（た）ものは？\n", target)
	fmt.Scan(&answer)
}
