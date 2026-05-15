package cli

import (
	"github.com/spf13/cobra"
	"github.com/srynprjl/kazumi/ui/gui"
)

var guiCmd = &cobra.Command{
	Use: "gui",
	Run: func(cmd *cobra.Command, args []string) {
		gui.Gui()
	},
}
