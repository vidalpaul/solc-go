package data

import (
	"fmt"
	"strings"
)

func (o *output) Err() error {
	var fmtMsgs []string
	for _, err := range o.Errors {
		if strings.EqualFold(err.Severity, "error") {
			fmtMsgs = append(fmtMsgs, err.FormattedMessage)
		}
	}

	if len(fmtMsgs) == 0 {
		return nil
	}
	return fmt.Errorf("solc: compilation failed\n%s", strings.Join(fmtMsgs, "\n"))
}

type error_ struct {
	SourceLocation   sourceLocation `json:"sourceLocation"`
	Type             string         `json:"type"`
	Component        string         `json:"component"`
	Severity         string         `json:"severity"`
	Message          string         `json:"message"`
	FormattedMessage string         `json:"formattedMessage"`
}
