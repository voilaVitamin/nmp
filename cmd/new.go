package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [post-name]",
	Short: "Create a new blog post",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		postName := args[0]
		date := time.Now().Format("2006-01-02")
		content := fmt.Sprintf(`---
draft: false
date: %s
categories:
  - Trivial
---

# %s
`, date, postName)

		// Ensure the directory exists
		dir := "docs/posts"
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}

		// Create the new file
		filePath := fmt.Sprintf("%s/%s.md", dir, postName)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			fmt.Println("Error creating file:", err)
			return
		}

		fmt.Printf("New post '%s.md' created in %s/\n", postName, dir)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
}
