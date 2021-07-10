// Package cmd is the main part of lamy command.
package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "lamy",
		Short: "A CLI tool that asks a series of questions to clarify what you want to know.",
		Long: `A CLI tool that asks a series of questions to clarify what you want to know.
It contains questions about genre, difference, part, definition, etymology, opposite, cause, and effect.`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			lamyRun(args[0])
		},
	}
	isList = false
)

// Execute executes the lamy command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&isList, "list", "l", false, "list questions")
}

type Prompt struct {
	Label    string
	Question string
}

func formatQuestion(prompt Prompt, target string) string {
	return fmt.Sprintf(prompt.Question, target)
}

func generateQuestions(prompts []Prompt, target string) []*survey.Question {
	qs := []*survey.Question{}
	for _, p := range prompts {
		formattedQuestion := formatQuestion(p, target)
		qs = append(qs, &survey.Question{
			Name:   p.Label,
			Prompt: &survey.Input{Message: formattedQuestion},
		})
	}
	return qs
}

func lamyRun(target string) {
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

	if isList {
		for _, p := range prompts {
			fmt.Println(formatQuestion(p, target))
		}
		return
	}

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
