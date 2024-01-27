package serve

import (
	"syscall"

	"woh/package/actor"
	"woh/package/app/cmd"
	"woh/package/utils"
	"woh/webhooks/grpc"
	"woh/webhooks/http"
	"woh/webhooks/subs"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
)

var Root = cmd.New("serve",
	cmd.Alias("s"),
	cmd.Run(runServe),
)

func runServe(cmd *cobra.Command, args []string) {
	sys := actor.NewSystem(cmd.Context())          // Create Actor System
	defer sys.Stop()                               // Wait for Adapters to close
	lo.Must0(sys.Start(Adapters()))                // Start Adapters
	utils.WaitSig(syscall.SIGINT, syscall.SIGTERM) // Wait for signals
}

func ChooseAdapters(
	h *http.Adapter,
	_ *grpc.Adapter,
	s *subs.Adapter,
) actor.Actors {
	return actor.Actors{
		"http": h,
		//"grpc": g,
		"subs": s,
	}
}
