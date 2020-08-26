package pork

import (
	"fmt"
	"github.com/spf13/cobra"
)

var SearchCmd = &cobra.Command{
	Use: "search",
	Short: "search for github repository by keyword",
	Run: func(cmd *cobra.Command, args []string) {
		repositoryList := SearchByKeyword(args)
		for _, repository := range repositoryList {
			fmt.Println(repository)
		}
	},
}

func SearchByKeyword(keyword []string) []string {
	return []string{"myrepository"}
}

