package printer

import (
	"github.com/olivere/bodleiandiff/diff"
)

// Printer prints a diff using a specific output format, e.g. JSON.
type Printer interface {
	Print(diff.Diff) error
}
