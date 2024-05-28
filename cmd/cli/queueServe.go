/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cli

import (
	"github.com/spf13/cobra"
	"go-fiber-clean-architecture/application/queue/server/app"
)

// queueServeCmd represents the queueServe command
var queueServeCmd = &cobra.Command{
	Use:   "queue",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		app.AppRun()
	},
}

func init() {
	rootCmd.AddCommand(queueServeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// queueServeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// queueServeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
