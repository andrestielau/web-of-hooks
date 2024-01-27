package cmd

import (
	"github.com/andrestielau/web-of-hooks/package/app/flag"
	"github.com/andrestielau/web-of-hooks/package/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func New(use string, opts ...func(*cobra.Command)) *cobra.Command {
	cmd := &cobra.Command{Use: use}
	utils.Apply(cmd, opts)
	return cmd
}
func Use(u string) func(*cobra.Command) { return func(cmd *cobra.Command) { cmd.Use = u } }
func Add(cmds ...*cobra.Command) func(*cobra.Command) {
	return func(cmd *cobra.Command) { cmd.AddCommand(cmds...) }
}
func Alias(a ...string) func(*cobra.Command) {
	return func(cmd *cobra.Command) { cmd.Aliases = append(cmd.Aliases, a...) }
}

func Run(fn func(*cobra.Command, []string)) func(*cobra.Command) {
	return func(cmd *cobra.Command) { cmd.Run = fn }
}
func PreRun(fn func(*cobra.Command, []string)) func(*cobra.Command) {
	return func(cmd *cobra.Command) { cmd.PreRun = fn }
}
func PostRun(fn func(*cobra.Command, []string)) func(*cobra.Command) {
	return func(cmd *cobra.Command) { cmd.PostRun = fn }
}
func PPreRun(fn func(*cobra.Command, []string)) func(*cobra.Command) {
	return func(cmd *cobra.Command) { cmd.PersistentPreRun = fn }
}
func PPostRun(fn func(*cobra.Command, []string)) func(*cobra.Command) {
	return func(cmd *cobra.Command) { cmd.PersistentPostRun = fn }
}
func Flags(flags ...func(*pflag.FlagSet)) func(*cobra.Command) {
	return func(cmd *cobra.Command) { utils.Apply(cmd.Flags(), flags) }
}
func PFlags(flags ...func(*pflag.FlagSet)) func(*cobra.Command) {
	return func(cmd *cobra.Command) { utils.Apply(cmd.PersistentFlags(), flags) }
}
func BindFlags(v *viper.Viper) func(*cobra.Command) {
	return func(cmd *cobra.Command) { flag.BindSet(cmd.Flags(), v) }
}
