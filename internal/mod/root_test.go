package mod

import (
	"strings"
	"testing"

	"github.com/vidalpaul/solc-go/internal/mod"
)

func TestModRoot(t *testing.T) {
	if !strings.HasSuffix(mod.Root, "solc") {
		t.Fatalf("Unexpected module root: %q", mod.Root)
	}
}
