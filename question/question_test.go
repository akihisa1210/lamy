package question

import "testing"

func TestNewQuestions(t *testing.T) {
	tests := []struct {
		name   string
		qts    []QuestionTemplate
		target string
		want   []Question
	}{
		{
			name: "One question",
			qts: []QuestionTemplate{
				{
					Name:     "Test name",
					Template: "Test template %s",
				},
			},
			target: "target",
			want: []Question{
				{
					Name:    "Test name",
					Content: "Test template target",
				},
			},
		},
		{
			name: "Two questions",
			qts: []QuestionTemplate{
				{
					Name:     "Test name 1",
					Template: "Test template 1 %s",
				},
				{
					Name:     "Test name 2",
					Template: "Test template 2 %s",
				},
			},
			target: "target",
			want: []Question{
				{
					Name:    "Test name 1",
					Content: "Test template 1 target",
				},
				{
					Name:    "Test name 2",
					Content: "Test template 2 target",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs := NewQuestions(tt.qts, tt.target)
			for i, q := range qs {
				if q.Name != tt.want[i].Name {
					t.Errorf("Invalid question name: %v", q.Name)
				}
				if q.Content != tt.want[i].Content {
					t.Errorf("Invalid question content: %v", q.Content)
				}
			}
		})
	}
}
