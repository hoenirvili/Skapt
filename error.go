package Skapt

import (
	"errors"
	"fmt"
	"os"
)

// Error flags
var (
	errNFlagAlias = errors.New("Inappropriate number of aliases")
	errTINT       = errors.New("Can't parse the INT from this flag")
	unkFLAG       = errors.New("Unknow type flag")
)

// C style exit flgs
const (
	EXIT_SUCCESS = iota
	EXIT_FAILURE
)

func errOnExit(err error) {
	fmt.Fprintf(os.Stderr, "%s\n", err)
	os.Exit(EXIT_FAILURE)
}
