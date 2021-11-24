package cmd

import (
	"github.com/spf13/cobra"
)

var (
	dryRun bool
	url    string
)

func init() {
	rootCmd.Flags().BoolVar(&dryRun, "dry-run", false, "")
	rootCmd.Flags().StringVar(&url, "url", "", "")
}

var rootCmd = &cobra.Command{
	Use: "decorators",
	Run: func(cmd *cobra.Command, args []string) {
		var runner runner
		r := defaultRunner{url}

		if dryRun {
			runner = &dryRunner{&r}
		} else {
			runner = &fileRunner{&r, "output.txt"}
		}

		run(runner)
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func run(r runner) {
	err := r.GetData("key", "key", "key2", "key2", "key3")
	cobra.CheckErr(err)
}
