package vis

import (
	"fmt"
	"path/filepath"
)

// dotJoin is like filepath.Join, but it preserves the dot ('.'), which is
// needed to invoke correctly "go build".
// If we were to use directly filepath.Join("./cmd", foo), it would strip
// the dot.
func FilepathJoinDot(elem ...string) string {
	return fmt.Sprintf(".%s%s", string(filepath.Separator),
		filepath.Join(elem...))
}
