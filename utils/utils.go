package utils

import (
	"fmt"
	"strings"
	"time"

	"github.com/oklog/ulid/v2"
)

const millisecondsPerSecond = 1000

func NewID(prefix string) string {
	return fmt.Sprintf("%s_%s", prefix, ulid.Make().String())
}

func GenerateHumanID(prefix string) string {
	if prefix == "" {
		prefix = "LGM" // LGM for log management
	}

	now := time.Now()
	ns := now.UnixMilli() % millisecondsPerSecond

	return fmt.Sprintf("%s%d-%03d-%02d%02d%02d%03d", strings.ToUpper(prefix), now.Year()-2000, now.YearDay(), now.Hour(), now.Minute(), now.Second(), ns)
}