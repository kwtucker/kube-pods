package cmd

import (
	"context"
	"fmt"

	"github.com/kwtucker/kube-pods/config"
	"github.com/kwtucker/reef/kube"
	"github.com/spf13/cobra"
)

var (
	DryRun  bool
	Verbose bool
)

func init() {
	RootCmd.PersistentFlags().BoolVarP(&DryRun, "dry-run", "n", false, "Dry run to inspect the result")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose. To list full objects and data.")
}

var RootCmd = &cobra.Command{
	Use:   "pods",
	Short: "Interacts with kubernetes pods.",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := config.LoadConfig(config.Flags{
			DryRun:  DryRun,
			Verbose: Verbose,
		})

		client, err := kube.NewClient(context.Background(), "")
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(args) < 1 {
			for _, name := range args {
				p, err := client.Pod(name)
				if err != nil {
					fmt.Printf("%+v", err)
					return
				}

				fmt.Printf("%+v\n", p.Info())
			}
		}

		fmt.Println(cfg)

	},
}
