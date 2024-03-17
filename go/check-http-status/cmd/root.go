/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/supermarine1377/monitoring-scripts/go/check-http-status/internal/http_status"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "check-http-status <URL>",
	Short: "Monitors the HTTP status code of a specified website at regular intervals.",
	Long:  `Monitors the HTTP status code of a specified website at regular intervals.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Fprintln(cmd.OutOrStderr(), "no arguments provided")
			fmt.Fprintf(cmd.OutOrStderr(), "usage: %s\n", cmd.UseLine())
			os.Exit(1)
		}
		targetURL := args[0]
		intervalSeconds, err := cmd.Flags().GetInt(INTERVAL_SECONDS)
		if err != nil {
			fmt.Fprintln(cmd.OutOrStderr(), err)
			os.Exit(1)
		}
		m := http_status.NewMonitor(targetURL, intervalSeconds)
		ctx := context.Background()

		m.Writedown(ctx, os.Stdout)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

const INTERVAL_SECONDS = "interval-seconds"
const INTERVAL_SECONDS_SHORTHAND = "i"
const DEFAULT_INTERVAL_SECONDS = 60

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.check-http-status.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().IntP(
		"interval-seconds",
		"i",
		DEFAULT_INTERVAL_SECONDS,
		"interval_seconds are interval time between monitoring HTTP requests.",
	)
}
