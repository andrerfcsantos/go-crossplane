package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/aluttik/go-crossplane"
	"github.com/spf13/cobra"
)

var (
	tabsFlag      *bool
	noHeadersFlag *bool
	identFlag     *int
)

func init() {
	tabsFlag = buildCmd.Flags().BoolP("tabs", "t", false, "indent with tabs instead of spaces")
	noHeadersFlag = buildCmd.Flags().Bool("no-headers", false, "do not write header to configs")
	identFlag = buildCmd.Flags().IntP("ident", "i", 4, "number of spaces to indent output")
	rootCmd.AddCommand(buildCmd)
}

var (
	buildFilePath = ""
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "The build command builds an nginx configuration from a json payload.",
	Long:  `All software has versions. This is Hugo's`,
	Args:  buildFileArg,
	RunE:  runBuildCmd,
}

func buildFileArg(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("parse requires the path to a nginx config file to be passed as argument")
	}
	buildFilePath = args[0]
	return nil
}

func runBuildCmd(cmd *cobra.Command, args []string) error {

	file, err := os.Open(buildFilePath)
	if err != nil {
		log.Fatal(err)
	}

	input, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var payload crossplane.Payload
	if err = json.Unmarshal(input, &payload); err != nil {
		log.Fatal(err)
	}

	combined, err := payload.Combined()
	if err != nil {
		log.Fatal(err)
	}
	config := combined.Config[0]
	options := &crossplane.BuildOptions{
		Indent: *identFlag,
		Tabs:   !*tabsFlag,
		Header: *noHeadersFlag,
	}

	var output bytes.Buffer
	if err = crossplane.Build(&output, config, options); err != nil {
		log.Fatal(err)
	}

	fmt.Println(output.String())
	return nil
}
