// Go Modeler is a small CLI that brings the powerful features of the golang template into a simplified form.
package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// Start application
func init() {

	initFlag()

	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	log.SetReportCaller(false)
	log.SetOutput(os.Stdout)

}

func Execute() {

	err := rootCmd.Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func main() {
	Execute()
}
