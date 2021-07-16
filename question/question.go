// Package question contains questions and related types.
package question

import (
	"fmt"
)

// Question has question name and its template.
type Question struct {
	Name     string
	Template string
}

// Format returns question from question template and given target.
func (q Question) Format(target string) string {
	return fmt.Sprintf(q.Template, target)
}

// DefaultQuestions are the question templates used by default.
var DefaultQuestions = []Question{
	{
		"Genre",
		"(類) %s は何の一種か？",
	},
	{
		"Difference",
		"(種差) %s は、同じグループの中で他と何が違うのか？",
	},
	{
		"Part",
		"(部分) %s を構成する部分を列挙すると？",
	},
	{
		"Definition",
		"(定義) %s とは何か？",
	},
	{
		"Etymology",
		"(語源) %s の語源は？",
	},
	{
		"Opposite",
		"(相反) %s の反対は？",
	},
	{
		"Cause",
		"(原因・由来) %s を生じさせる（た）ものは？",
	},
	{
		"Effect",
		"(結果・派生) %s から生じる（た）ものは？",
	},
}

// TechnicalQuestions are the question templates for a technical topic.
var TechnicalQuestions = []Question{
	{
		"Genre",
		"(類) %s は何の一種か？",
	},
	{
		"Difference",
		"(種差) %s は、同じグループの中で他と何が違うのか？",
	},
	{
		"Part",
		"(部分) %s を構成する部分を列挙すると？",
	},
	{
		"Cooperation",
		"(連携) %s と連携するものは？",
	},
	{
		"Definition",
		"(定義) %s とは何か？",
	},
	{
		"UseCase",
		"(ユースケース) %s の具体的なユースケースは？",
	},
	{
		"Pros",
		"(利点) %s を使うと何が嬉しいのか？",
	},
	{
		"Cons",
		"(欠点) %s の欠点は？",
	},
}

// Answers have the answer for the corresponding question.
type Answers = map[string]string
