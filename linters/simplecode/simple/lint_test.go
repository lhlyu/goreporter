package simple

import (
	"testing"

	"github.com/lhlyu/goreporter/linters/simplecode/lint/testutil"
)

func TestAll(t *testing.T) {
	testutil.TestAll(t, Funcs, "../../")
}
