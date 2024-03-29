package call

import (
	"log"

	"woh/package/app/cmd"

	webhooksv1 "woh/webhooks/adapt/grpc/v1"

	"github.com/samber/lo"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client webhooksv1.WebHookServiceClient
var Root = cmd.New("call",
	cmd.PPreRun(func(cmd *cobra.Command, _ []string) {
		client = webhooksv1.NewWebHookServiceClient(lo.Must(grpc.Dial("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()))))
	}),
	cmd.Add(
		cmd.New("applications", cmd.Alias("a"), cmd.Add(
			cmd.New("list", cmd.Alias("l"), cmd.Run(runListApplications)),
		)),
		cmd.New("endpoints", cmd.Alias("e"), cmd.Add(
			cmd.New("list", cmd.Alias("l"), cmd.Run(runListEndpoints)),
		)),
	),
)

func runListApplications(cmd *cobra.Command, _ []string) {
	log.Println(client.ListApps(cmd.Context(), &webhooksv1.ListAppsRequest{}))
}

func runListEndpoints(cmd *cobra.Command, _ []string) {
	log.Println(client.ListEndpoints(cmd.Context(), &webhooksv1.ListEndpointsRequest{}))
}
