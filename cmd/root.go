package cmd

import (
	"github.com/andrestielau/web-of-hooks/cmd/serve"
	"github.com/andrestielau/web-of-hooks/package/app/cmd"
	"github.com/samber/lo"
)

var root = cmd.New("woh", cmd.Add(serve.Root))

func Execute() { lo.Must0(root.Execute()) }
