package pipeline

import (
	"testing"

	"gopcp.v2/chapter6/webcrawler/errors"
)

func TestErrorGenError(t *testing.T) {
	simpleErrMsg := "testing error"
	expectedErrType := errors.ERROR_TYPE_PIPELINE
	err := genError(simpleErrMsg)
	ce, ok := err.(errors.CrawlerError)
	if !ok {
		t.Fatalf("Inconsistent error type: expected: %T, actual: %T",
			errors.NewCrawlerError("", ""), err)
	}
	if ce.Type() != expectedErrType {
		t.Fatalf("Inconsistent error type string: expected: %q, actual: %q",
			expectedErrType, ce.Type())
	}
	expectedErrMsg := "crawler error: pipeline error: " + simpleErrMsg
	if ce.Error() != expectedErrMsg {
		t.Fatalf("Inconsistent error message: expected: %q, actual: %q",
			expectedErrMsg, ce.Error())
	}
}
