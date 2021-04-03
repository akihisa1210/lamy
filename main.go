package main

import (
	"fmt"
	"os"
)

func main() {
	questions := []string{
		"[類] %s は何の一種か？\n",
		"[種差] %s は、同じグループの中で他と何が違うのか？\n",
		"[部分] %s を構成する部分を列挙すると？\n",
		"[定義] %s とは何か？\n",
		"[語源] %s の語源は？\n",
		"[相反] %s の反対は？\n",
		"[原因・由来] %s を生じさせる（た）ものは？\n",
	}

	if len(os.Args) != 2 {
		return
	}

	target := os.Args[1]
	var answer string
	for _, question := range questions {
		fmt.Printf(question, target)
		fmt.Scan(&answer)
	}
}
