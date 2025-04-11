/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/iamcaye/aitor-cli/client"
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
		projectPath, projectType, err := project.RunDetector()
		if err != nil {
			fmt.Println("Error detecting project:", err)
			return
		}

		if projectType == project.UNKNOWN_PRJ {
			fmt.Println("Unknown project type")
			return
		}

		fmt.Println("Project Path:", projectPath)
		fmt.Println("Project Type:", projectType)
		packageContent, err := utils.ReadFile(projectPath + "/package.json")
		if err != nil {
			panic(err)
		}

		lockContent, err := utils.ReadFile(
			fmt.Sprintf(
				"%s/%s",
				projectPath,
				project.ProjectLockFile[projectType],
			),
		)
		if err != nil {
			panic(err)
		}

		fmt.Println("Package Content (KB):", len(packageContent)/1024)
		fmt.Println("Lock Content (KB):", len(lockContent)/1024)

		dataString := fmt.Sprintf("[\n%s,\n\n%s\n]", packageContent, lockContent)
		client.SendAuditRequest([]byte(dataString))
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
