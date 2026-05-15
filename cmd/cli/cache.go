package cli

import (
	"github.com/spf13/cobra"
	"github.com/srynprjl/kazumi/lib/misc"
)

var cacheCmd = &cobra.Command{
	Use: "cache",
	Run: func(cmd *cobra.Command, args []string) {
		cache, _ := cmd.Flags().GetBool("clean")
		if cache {
			misc.CleanVideoCaches()
			return
		}

	},
}
