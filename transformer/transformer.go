package transformer

import (
	"encoding/json"
	"fmt"
)

type Metadata struct {
	Source       string        `json:"source"`
	Errors       []*SylvaError `json:"errors"`
	ErrorStrings []string      `json:"error-strings"`
}

type ProgramData struct {
	Successful bool     `json:"successful"`
	Metadata   Metadata `json:"meta"`
	AST        Node     `json:"ast"`
}

// returns json corresponding to a ProgramData object
func TransformJSON(code, source string, minify bool) (jsonData []byte, err error) {
	programData := Transform(code, source)

	if minify {
		jsonData, err = json.Marshal(programData)
	} else {
		jsonData, err = json.MarshalIndent(programData, "", "  ")
	}

	if err != nil {
		err = fmt.Errorf("error while marshaling json: %v", err)
	}

	return jsonData, err
}

func Transform(code, source string) *ProgramData {
	lexer := MakeLexer(code)
	lexer.Lex()

	errors := []*SylvaError{}
	errors = append(errors, lexer.Errors...)

	parser := MakeParser(lexer.Tokens)
	ast := parser.Parse()

	errors = append(errors, parser.Errors...)
	errorStrings := make([]string, 0, len(errors))

	if len(errors) != 0 {
		for _, err := range errors {
			errorStrings = append(errorStrings, err.Format(code, source, "\t"))
		}
	}

	ast.Marshal()

	programData := &ProgramData{
		AST:        ast,
		Successful: len(errors) == 0,
		Metadata: Metadata{
			Source:       source,
			Errors:       errors,
			ErrorStrings: errorStrings,
		},
	}

	return programData
}
