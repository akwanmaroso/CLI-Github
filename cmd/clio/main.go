package main

import (
	"fmt"
	"os"

	clio "github.com/akwanmaroso/devops-go/clio"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd *cobra.Command

func main() {
	rootCmd.Execute()
}

func init() {
	rootCmd = &cobra.Command{
		Use:   "clio",
		Short: "DevOps Tool for Github",
	}
	rootCmd.AddCommand(clio.SearchCmd)
	rootCmd.AddCommand(clio.DocsCmd)
	rootCmd.AddCommand(clio.CloneCmd)
	rootCmd.AddCommand(clio.ForkCmd)
	rootCmd.AddCommand(clio.PullRequestCmd)
	rootCmd.AddCommand(clio.IssueCmd)
	rootCmd.AddCommand(clio.RepositoryCmd)
	viper.SetDefault("location", os.Getenv("HOME"))
	viper.SetConfigName("clio")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("No configuration file found")
	}
	viper.SetDefault("location", os.Getenv("HOME"))
}
