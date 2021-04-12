package flags

import (
	"flag"
	"fmt"
	"os"
)

func FatalFlagValue(msg, name string, val interface{}) {
	_, _ = fmt.Fprintf(os.Stderr, "invalid value %#v for flag -%s: %s\n", val, name, msg)
	flag.Usage()
	os.Exit(1)
}
