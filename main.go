package main

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

type Prompt struct {
	Label    string
	Question string
}

func generateQuestions(prompts []Prompt, target string) []*survey.Question {
	qs := []*survey.Question{}
	for _, p := range prompts {
		formattedQuestion := fmt.Sprintf(p.Question, target)
		qs = append(qs, &survey.Question{
			Name:   p.Label,
			Prompt: &survey.Input{Message: formattedQuestion},
		})
	}
	return qs
}

func main() {
	prompts := []Prompt{
		{
			"Genre",
			"[類] %s は何の一種か？\n",
		},
		{
			"Difference",
			"[種差] %s は、同じグループの中で他と何が違うのか？\n",
		},
		{
			"Part",
			"[部分] %s を構成する部分を列挙すると？\n",
		},
		{
			"Definition",
			"[定義] %s とは何か？\n",
		},
		{
			"Etymology",
			"[語源] %s の語源は？\n",
		},
		{
			"Opposite",
			"[相反] %s の反対は？\n",
		},
		{
			"Cause",
			"[原因・由来] %s を生じさせる（た）ものは？\n",
		},
		{
			"Effect",
			"[結果・派生] %s から生じる（た）ものは？\n",
		},
	}

	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s [target]", os.Args[0])
		os.Exit(1)
	}

	target := os.Args[1]

	qs := generateQuestions(prompts, target)

	answers := struct {
		Genre      string
		Difference string
		Part       string
		Definition string
		Etymology  string
		Opposite   string
		Cause      string
		Effect     string
	}{}

	err := survey.Ask(qs, &answers)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
