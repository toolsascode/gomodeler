package main

import (
	"errors"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"

	mango "github.com/muesli/mango-cobra"
	"github.com/muesli/roff"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"

	gomodeler "github.com/toolsascode/gomodeler/internal/template"
)

var rootCmd = &cobra.Command{
	Use:   "gomodeler",
	Short: "Go Modeler is a small CLI that brings the powerful features of the golang template into a simplified form.",
	Long: `GoModeler brings with it all the features that gotemplate provides in a less complex and easy to implement way, including extra features.
	Complete documentation is available at https://github.com/toolsascode/gomodeler`,
	Run: func(_ *cobra.Command, _ []string) {

		gomodeler.Run()

	},
}

var versionCmd = &cobra.Command{
	Use:                "version",
	Short:              "Print the version number of Go Modeler",
	Long:               `All software has versions. This is Go Modeler's`,
	DisableFlagParsing: true,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("Version: ", version)
		fmt.Println("Date: ", date)
		fmt.Println("Commit: ", commit)
		fmt.Println("Built by: ", builtBy)
	},
}

var docsCmd = &cobra.Command{
	Use:                   "docs",
	Short:                 "Generating Go Modeler CLI markdown documentation.",
	Long:                  `Allow generating documentation in markdown format for Go Modeler CLI internal commands`,
	Hidden:                true,
	DisableFlagParsing:    true,
	SilenceUsage:          true,
	DisableFlagsInUseLine: true,
	Args:                  cobra.NoArgs,
	ValidArgsFunction:     cobra.NoFileCompletions,
	Run: func(_ *cobra.Command, _ []string) {

		var path = "./docs"

		if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir(path, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		}

		log.Info("Generating markdown documentation")
		err := doc.GenMarkdownTree(rootCmd, path)
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("Documentation successfully generated in %s", path)
	},
}

var manCmd = &cobra.Command{
	Use:                   "man",
	Short:                 "Generates manpages",
	DisableFlagParsing:    true,
	SilenceUsage:          true,
	DisableFlagsInUseLine: true,
	Hidden:                true,
	Args:                  cobra.NoArgs,
	RunE: func(_ *cobra.Command, _ []string) error {
		manPage, err := mango.NewManPage(1, rootCmd.Root())
		if err != nil {
			return err
		}

		_, err = fmt.Fprint(os.Stdout, manPage.Build(roff.NewDocument()))
		return err
	},
}
