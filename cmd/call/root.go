package call

import (
	"log"

	"woh/package/app/cmd"
	"github.com/spf13/cobra"
)

var Root = cmd.New("call",
	cmd.Add(
		cmd.New("endpoints", cmd.Alias("e"), cmd.Add(
			cmd.New("list", cmd.Alias("l"), cmd.Run(runListEndpoints)),
		)),
	),
)

func runListEndpoints(cmd *cobra.Command, _ []string) {
	log.Println("Hellos")
}
