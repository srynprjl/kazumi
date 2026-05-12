package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/srynprjl/kazumi/lib/misc"
)

var cacheCmd = &cobra.Command{
	Use: "cache",
	Run: func(cmd *cobra.Command, args []string) {
		all, _ := cmd.Flags().GetBool("all")
		cache, _ := cmd.Flags().GetBool("cache")
		logs, _ := cmd.Flags().GetBool("logs")
		if all {
			misc.CleanVideoCaches()
			misc.CleanLogCache()
			return
		}
		if cache {
			misc.CleanVideoCaches()
			return
		}
		if logs {
			misc.CleanLogCache()
			return
		}
		if len(args) == 0 {
			fmt.Println("Deleting all cache")
			misc.CleanVideoCaches()
			misc.CleanLogCache()
			return
		}
	},
}
