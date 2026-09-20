package cmds

import "os"

type EditorFunc func(state *EditorState, args []string) error

type EditorState struct {
	FileName        string
	Buffer          []string
	Modified        bool
	CommandRegistry map[string]EditorFunc
}

func NewEditorState(buf []string, fileName string) *EditorState {
	return &EditorState{
		FileName: fileName,
		Buffer:   buf,
		Modified: false,
		CommandRegistry: map[string]EditorFunc{
			"h": PrintHelp,
			"w": Write,
			"r": Read,
			"e": Edit,
			"x": Execute,
			"c": Clear,
			"d": Delete,
			"a": Append,
			"i": Insert,
			"f": Find,
			"p": Replace,
			"q": func(state *EditorState, args []string) error {
				os.Exit(0)
				return nil
			},
		},
	}
}
