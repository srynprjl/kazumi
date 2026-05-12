package cli

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/srynprjl/kazumi/lib/creation"
)

var rootCmd = &cobra.Command{
	Use:     "kazumi",
	Args:    cobra.MinimumNArgs(1),
	Example: "kazumi -s=1.25 -p=1.33 -i=<url>  <url>`",
	Short:   "Video to Nightcore ",
	Run: func(cmd *cobra.Command, args []string) {

		var opt creation.Options = creation.Options{}
		cmd.Flags().VisitAll(func(f *pflag.Flag) {
			if f.Changed {
				if f.Name == "speed" {
					opt.Speed.Enabled = true
					opt.Speed.Value, _ = cmd.Flags().GetFloat64(f.Name)
				}
				if f.Name == "pitch" {
					opt.Pitch.Enabled = true
					opt.Pitch.Value, _ = cmd.Flags().GetFloat64(f.Name)
				}
				if f.Name == "reverb" {
					opt.Reverb.Enabled = true
					opt.Reverb.InGain = 1.0
					opt.Reverb.OutGain = 0.8
					opt.Reverb.Delay = 40
					opt.Reverb.Decay = 0.3
				}
			}
		})
		// fmt.Println(opt)
		audio := args[0]
		img, _ := cmd.Flags().GetString("image")
		creation.FullProcedure(audio, img, opt)

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().Float64P("speed", "s", 1.00, "Adjust speed for the video")
	rootCmd.Flags().Lookup("speed").NoOptDefVal = "1.25"

	rootCmd.Flags().Float64P("pitch", "p", 1.00, "Adjust pitch for the video")
	rootCmd.Flags().Lookup("pitch").NoOptDefVal = "1.33"

	rootCmd.Flags().BoolP("reverb", "r", false, "Add reverb for to the video")

	rootCmd.Flags().StringP("image", "i", "", "Image url for video")

	rootCmd.AddCommand(jsonCmd)
	rootCmd.AddCommand(cacheCmd)

	cacheCmd.Flags().BoolP("cache", "c", false, "Clean the cache directory")
	cacheCmd.Flags().BoolP("logs", "l", false, "Clean the logs directory")
	cacheCmd.MarkFlagsMutuallyExclusive("cache", "logs")
	cacheCmd.Flags().BoolP("all", "a", false, "Clean all caches")
}
