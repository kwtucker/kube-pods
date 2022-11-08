package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/kwtucker/kube-pods/config"
	"github.com/kwtucker/reef/kube"
	"github.com/spf13/cobra"
)

var (
	DryRun  bool
	Verbose bool
	Info    bool
)

func init() {
	RootCmd.PersistentFlags().BoolVarP(&DryRun, "dry-run", "n", false, "Dry run to inspect the result")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "Verbose. To list full objects and data.")
	RootCmd.PersistentFlags().BoolVarP(&Info, "info", "", false, "Info. information about the pod.`")
}

var RootCmd = &cobra.Command{
	Use:   "pods",
	Short: "Interacts with kubernetes pods.",
	Run: func(cmd *cobra.Command, args []string) {

		cfg := config.LoadConfig(config.Flags{
			DryRun:  DryRun,
			Verbose: Verbose,
			Info:    Info,
		})

		client, err := kube.NewClient(context.Background(), "")
		if err != nil {
			fmt.Println(err)
			return
		}

		podInfos := []kube.PodInfo{}

		if len(args) < 1 {
			fmt.Fprint(os.Stderr, cmd.Help())
			return
		}

		for _, name := range args {
			p, err := client.Pod(name)
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				return
			}

			podInfos = append(podInfos, p.Info())
		}

		if cfg.Info {
			byt, err := json.Marshal(podInfos)
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				return
			}

			fmt.Fprint(os.Stdout, "", string(byt))
		}
	},
}
