package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   "eca",
		Short: "tool to audit containerd containers",
	}
)

func init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("scra")

	flags := rootCmd.PersistentFlags()
	flags.String("containerd-address", "/run/containerd/containerd.sock", "location of containerd endpoint")
	flags.String("root-prefix", "", "root prefix when accessing /proc et al from a hostPath mount")

	if err := viper.BindPFlags(flags); err != nil {
		panic(err)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
