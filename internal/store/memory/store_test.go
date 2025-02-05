package memory

import (
	"testing"

	"github.com/kellegous/golinks/internal/store/internal"
)

func TestS(t *testing.T) {
	var s Store
	internal.TestStore(t, &s)
}
