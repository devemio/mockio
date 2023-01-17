package config

import (
	"flag"
	"fmt"
	"os"

	"github.com/devemio/mockio/internal/types"
)

const (
	Version = "0.7.0"

	help = `Usage: %s -i routes.json

Options:
`
)

type Config struct {
	Port            int
	Path            string
	ResponseLatency int
	IsOutputColored bool
	Verbose         types.Verbose
}

func New() *Config {
	port := flag.Int("p", 8080, "server port")
	path := flag.String("i", "", "path to mock file")
	responseLatency := flag.Int("l", 0, "response latency (ms)")
	isOutputColored := flag.Bool("c", true, "color output")
	v := flag.Bool("v", false, "verbose output")
	vv := flag.Bool("vv", false, "very verbose output (print response headers)")

	flag.Usage = usage
	flag.Parse()

	validate(*path)

	return &Config{
		Port:            *port,
		Path:            *path,
		ResponseLatency: *responseLatency,
		IsOutputColored: *isOutputColored,
		Verbose:         level(*v, *vv),
	}
}

func usage() {
	_, _ = fmt.Fprintf(flag.CommandLine.Output(), help, os.Args[0])

	flag.PrintDefaults()
}

func validate(path string) {
	if path != "" {
		return
	}

	usage()

	os.Exit(-1)
}

func level(v, vv bool) types.Verbose {
	if vv {
		return types.VerbosityVeryVerbose
	}

	if v {
		return types.VerbosityVerbose
	}

	return types.VerbosityNormal
}
