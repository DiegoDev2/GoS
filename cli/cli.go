package cli

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/briandowns/spinner"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gs",
	Short: "Go project generator",
	Long:  `A fast and flexible project generator built with Go.`,
	Run: func(cmd *cobra.Command, args []string) {
		color.Cyan.Println("GoS CLI - Use 'gs init <template>' to generate a project.")
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
		s := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		s.Suffix = fmt.Sprintf("  Generating %s template for project: %s", template, projectName)
		s.Start()

		time.Sleep(3 * time.Second)

		switch template {
		case "api":
			color.Green.Println("Generating API template...")
			// Aquí clonamos un repositorio o generamos archivos
			command := exec.Command("git", "clone", "https://github.com/DiegoDev2/gos-templates")
			command.Run()
		case "webapp":
			color.Green.Println("Generating WebApp template...")
			// Generación del template webapp
		case "microservice":
			color.Green.Println("Generating Microservice template...")
			// Generación de microservicio
		default:
			color.Red.Println("Unsupported template. Use 'api', 'webapp', or 'microservice'.")
		}

		s.Stop()
		color.Green.Println("Project generated successfully!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().StringP("name", "n", "", "Project name")
	if err := initCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println("Error marking 'name' flag as required:", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
