package cli

import (
	"github.com/spf13/cobra"
	"github.com/srynprjl/kazumi/lib/creation"
)

var jsonCmd = &cobra.Command{
	Use:     "json",
	Args:    cobra.MinimumNArgs(1),
	Example: "kazumi json <url>",
	Run: func(cmd *cobra.Command, args []string) {
		val := args[0]
		if val != "" {
			jsonData := creation.ParseJSON(val)
			creation.DownloadUsingJSON(jsonData)
		}
	},
}
