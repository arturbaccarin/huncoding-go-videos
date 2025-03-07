// https://youtu.be/sTrEfZL7XyM
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "k8s-cli",
	Short: "k8s-cli is a CLI tool for Kubernetes",
	Long:  "k8s-cli is a CLI tool for Kubernetes",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
