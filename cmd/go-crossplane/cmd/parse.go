package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aluttik/go-crossplane"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(parseCmd)
}

var (
	parseFilePath = ""
)

// -h, --help            show this help message and exit
// -o OUT, --out OUT     write output to a file
// -i NUM, --indent NUM  number of spaces to indent output -> argument to marshal?
// --ignore DIRECTIVES   ignore directives (comma-separated) -> IgnoreDirectives []string
// --no-catch            only collect first error in file
// --tb-onerror          include tracebacks in config errors
// --combine             use includes to create one single file -> CombineConfigs bool
// --single-file         do not include other config files -> SingleFile bool
// --include-comments    include comments in json -> ParseComments bool
// --strict              raise errors for unknown directives -> ErrorOnUnknownDirectives bool

// If true, parsing will stop immediately if an error is found.
// StopParsingOnError bool

// If true, checks that directives are in valid contexts.
// SkipDirectiveContextCheck bool

// If true, checks that directives have a valid number of arguments.
// SkipDirectiveArgsCheck bool

// If an error is found while parsing, it will be passed to this callback
// function. The results of the callback function will be set in the
// PayloadError struct that's added to the Payload struct's Errors array.
// ErrorCallback func(error) interface{}

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parses an nginx configuration into a json payload",
	Long:  `The parse command parses an nginx configuration into a json payload`,
	Args:  parseFileArg,
	RunE:  runParseCmd,
}

func parseFileArg(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("parse requires the path to a nginx config file to be passed as argument, %d positional argument(s) passed", len(args))
	}
	parseFilePath = args[0]
	return nil
}

func runParseCmd(cmd *cobra.Command, args []string) error {
	payload, err := crossplane.Parse(parseFilePath, &crossplane.ParseOptions{})
	if err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

	return nil
}
