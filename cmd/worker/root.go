package worker

import (
	"syscall"

	"github.com/andrestielau/web-of-hooks/package/actor"
	"github.com/andrestielau/web-of-hooks/package/app/cmd"
	"github.com/andrestielau/web-of-hooks/package/utils"
	"github.com/andrestielau/web-of-hooks/webhooks/cron"
	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var Root = cmd.New("work",
	cmd.Alias("w"),
	cmd.Run(runWork),
)

func runWork(cmd *cobra.Command, args []string) {
	sys := actor.NewSystem(cmd.Context())          // Create Actor System
	defer sys.Stop()                               // Wait for Adapters to close
	lo.Must0(sys.Start(Adapters()))                // Start Adapters
	utils.WaitSig(syscall.SIGINT, syscall.SIGTERM) // Wait for signals
}

func ChooseAdapters(c *cron.Adapter) actor.Actors {
	return actor.Actors{
		"cron": c,
	}
}
