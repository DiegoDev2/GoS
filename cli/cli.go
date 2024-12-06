package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gs",
	Short: "Go project generator",
	Long:  `A fast and flexible project generator built with Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'gs init <template>' to generate a project.")
	},
}

var initCmd = &cobra.Command{
	Use:   "init [template type]",
	Short: "Generate a new project template",
	Long:  `Generate a new Go project based on a template (e.g., API, WebApp, Microservice, etc.).`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		template := args[0]
		projectName, _ := cmd.Flags().GetString("name")

		switch template {
		case "api":
			fmt.Printf("Generating API template for project: %s\n", projectName)
		case "webapp":
			fmt.Printf("Generating WebApp template for project: %s\n", projectName)
		case "microservice":
			fmt.Printf("Generating Microservice template for project: %s\n", projectName)
		default:
			fmt.Println("Unsupported template. Use 'api', 'webapp', or 'microservice'.")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("name", "n", "", "Project name")
	initCmd.MarkFlagRequired("name")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
