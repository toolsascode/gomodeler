package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile     string
	logLevel    string
	mergeFiles  bool
	tmplPath    string
	outputPath  string
	envs        []string
	envFiles    []string
	tmplFiles   []string
	outputFiles []string
)

func initFlag() {

	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file (default is $HOME/.gomodeler.yaml or ./.configs/gomodeler.yaml or .gomodeler.yaml)")
	rootCmd.PersistentFlags().StringVarP(&logLevel, "log-level", "l", "info", "Log level [debug, info, warn, error, fatal, panic]")
	rootCmd.PersistentFlags().BoolVarP(&mergeFiles, "merge", "m", false, "Merge the output files into a single file. (Default: false)")
	rootCmd.PersistentFlags().StringSliceVarP(&envs, "env", "e", []string{}, "List of environment variables in `key=value` format separated by commas.")
	rootCmd.PersistentFlags().StringSliceVarP(&envFiles, "env-file", "f", []string{}, "`List` of environment variable files. Supported extensions JSON or YAML.")
	rootCmd.PersistentFlags().StringSliceVarP(&tmplFiles, "template-file", "t", []string{}, "`List` of template files for rendering separated by commas. Optional: filename.ext.exttmpl")
	rootCmd.PersistentFlags().StringSliceVarP(&outputFiles, "output-file", "o", []string{}, "`List` of files that will be saved to disk after rendering. The order specified in the file list entry should be followed. templates.")
	rootCmd.PersistentFlags().StringVar(&tmplPath, "template-path", "", "`Directory` of template files to be rendered. Optional: filename.ext.exttmpl")
	rootCmd.PersistentFlags().StringVar(&outputPath, "output-path", "", "`Location` where rendered files are saved. To use this functionality, use the following format in the template {file name}.{extension file}.{extension template file}.")

	// At least one is required.
	rootCmd.MarkFlagsOneRequired("output-path", "output-file")
	rootCmd.MarkFlagsOneRequired("template-path", "template-file")

	_ = viper.BindPFlag("env", rootCmd.PersistentFlags().Lookup("env"))
	_ = viper.BindPFlag("envFile", rootCmd.PersistentFlags().Lookup("env-file"))
	_ = viper.BindPFlag("template.files", rootCmd.PersistentFlags().Lookup("template-file"))
	_ = viper.BindPFlag("template.path", rootCmd.PersistentFlags().Lookup("template-path"))
	_ = viper.BindPFlag("output.files", rootCmd.PersistentFlags().Lookup("output-file"))
	_ = viper.BindPFlag("output.path", rootCmd.PersistentFlags().Lookup("output-path"))
	_ = viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))

	viper.SetDefault("author", "Carlos Freitas <carlosrfjunior@gmail.com>")
	viper.SetDefault("license", "Apache License 2.0")

	viper.AutomaticEnv()

	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(docsCmd)
	rootCmd.AddCommand(manCmd)

}
