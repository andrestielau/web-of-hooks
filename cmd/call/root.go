package call

import (
	"log"

	"woh/package/app/cmd"

	webhooksv1 "woh/webhooks/adapt/grpc/v1"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var client webhooksv1.WebHookServiceClient
var Root = cmd.New("call",
	cmd.PPreRun(func(cmd *cobra.Command, _ []string) {
		client = webhooksv1.NewWebHookServiceClient(lo.Must(grpc.Dial("http://localhost:3001")))
	}),
	cmd.Add(
		cmd.New("endpoints", cmd.Alias("e"), cmd.Add(
			cmd.New("list", cmd.Alias("l"), cmd.Run(runListEndpoints)),
		)),
	),
)

func runListEndpoints(cmd *cobra.Command, _ []string) {
	log.Println(client.ListEndpoints(cmd.Context(), &webhooksv1.ListEndpointsRequest{}))
}
