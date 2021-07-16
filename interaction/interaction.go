// Package interaction contains interface for user interaction.
package interaction

import "github.com/akihisa1210/lamy/question"

type Interaction interface {
	// ask asks user the questions about target and store the answers.
	ask(string, []question.Question) []question.Answers
}
