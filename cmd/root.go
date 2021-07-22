// Package cmd is the main part of lamy command.
package cmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/akihisa1210/lamy/question"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "lamy",
		Short: "A CLI tool that asks a series of questions to clarify what you want to know.",
		Long: `A CLI tool that asks a series of questions to clarify what you want to know.
By default, the CLI asks you questions about genre, difference, part, definition, etymology, opposite, cause, and effect.`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := lamyRun(args[0])
			if err != nil {
				return err
			}
			return nil
		},
	}
	isList = false
	isTech = false
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
	rootCmd.Flags().BoolVarP(&isTech, "tech", "t", false, "use questions for technical topic")
}

func generateSurveyQuestions(qs []question.Question) []*survey.Question {
	sqs := []*survey.Question{}
	for _, q := range qs {
		sqs = append(sqs, &survey.Question{
			Name:   q.Name,
			Prompt: &survey.Input{Message: q.Content + "\n"},
		})
	}
	return sqs
}

func generateSurveyAnswer(qs []question.Question) map[string]interface{} {
	sas := map[string]interface{}{}
	for _, q := range qs {
		sas[q.Name] = ""
	}
	return sas
}

type TerminalInteraction struct{}

func (ti *TerminalInteraction) ask(qs []question.Question) (question.Answers, error) {
	surveyQuestions := generateSurveyQuestions(qs)
	surveyAnswers := generateSurveyAnswer(qs)

	err := survey.Ask(surveyQuestions, &surveyAnswers)
	if err != nil {
		return nil, err
	}

	ans := map[string]string{}
	for key, surveyAnswer := range surveyAnswers {
		ans[key] = surveyAnswer.(string)
	}

	return ans, nil
}

func lamyRun(target string) error {
	var qt []question.QuestionTemplate

	if isTech {
		qt = question.TechnicalQuestions
	} else {
		qt = question.DefaultQuestions
	}

	qs := question.NewQuestions(qt, target)

	if isList {
		for _, q := range qs {
			fmt.Println(q.Content + "\n")
		}
		return nil
	}

	var ti TerminalInteraction
	_, err := ti.ask(qs)
	if err != nil {
		return err
	}

	return nil
}
