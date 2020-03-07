package escapeshellchar_test

import (
	"strconv"
	"testing"

	"github.com/bgpat/escapeshellchar"
)

func Test(t *testing.T) {
	for i := 0; i < 0x100; i++ {
		t.Run(strconv.FormatInt(int64(i), 16), func(t *testing.T) {
			want := string([]rune{rune(i)})
			escaped := escapeshellchar.EscapeShellString(want)
			got := escapeshellchar.UnEscapeShellString(escaped)
			if got != want {
				t.Errorf("0x%x: want:%q, got:%q, escaped:%q", i, want, got, escaped)
			}
		})
	}
}
