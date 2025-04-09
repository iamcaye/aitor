/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/iamcaye/aitor-cli/pkg/project"
	"github.com/iamcaye/aitor-cli/utils"
	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Checks current workspace to make a security audition",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check called")
		// Check if the current directory is a project folder
		currentDir, err := os.Getwd()
		if err != nil {
			fmt.Println("Error getting current directory:", err)
			return
		}

		projectPath, err := project.SearchProjectFolder(currentDir)
		if err != nil {
			fmt.Println("Error searching project folder:", err)
			return
		}
		if projectPath == "" {
			fmt.Println("No project folder found")
			return
		}
		fmt.Println("Project folder found:", projectPath)

		// get package.json file content
		packageJsonPath := projectPath + "/package.json"
		packageJsonContent, err := utils.ReadFile(packageJsonPath)
		if err != nil {
			fmt.Println("Error reading package.json file:", err)
			return
		}
		fmt.Println("package.json content: ", len(packageJsonContent), "bytes")

		// compress the package.json file
		compressedPackageJson, err := utils.Compress([]byte(packageJsonContent))
		if err != nil {
			fmt.Println("Error compressing package.json file:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
