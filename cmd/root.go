package cmd

import (
	"woh/cmd/serve"
	"woh/package/app/cmd"

	"woh/cmd/call"

	"github.com/samber/lo"
)

var root = cmd.New("woh", cmd.Add(serve.Root, call.Root))

func Execute() { lo.Must0(root.Execute()) }
