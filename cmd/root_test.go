package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRootCommand(t *testing.T) {
	tests := []struct {
		name    string
		command string
		want    string
	}{
		{
			name:    "normal question",
			command: "lamy test",
			want: `(類) test は何の一種か？

(種差) test は、同じグループの中で他と何が違うのか？

(部分) test を構成する部分を列挙すると？

(定義) test とは何か？

(語源) test の語源は？

(相反) test の反対は？

(原因・由来) test を生じさせる（た）ものは？

(結果・派生) test から生じる（た）ものは？

`,
		},
		{
			name:    "technical question",
			command: "lamy -t test",
			want: `(類) test は何の一種か？

(種差) test は、同じグループの中で他と何が違うのか？

(部分) test を構成する部分を列挙すると？

(連携) test と連携するものは？

(定義) test とは何か？

(ユースケース) test の具体的なユースケースは？

(利点) test を使うと何が嬉しいのか？

(欠点) test の欠点は？

`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			cmd := NewLamyCommand()
			cmd.SetOut(buf)
			cmdArgs := strings.Split(tt.command, " ")
			fmt.Printf("cmdArgs %+v\n", cmdArgs)
			cmd.SetArgs(cmdArgs[1:])
			cmd.Execute()

			get := buf.String()
			if tt.want != get {
				t.Errorf("unexpected response: want %v, got %v", tt.want, get)
			}
		})
	}
}
