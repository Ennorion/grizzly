package main

import (
	"log"

	"github.com/go-clix/cli"
	"github.com/grafana/grizzly/pkg/grafana"
	"github.com/grafana/grizzly/pkg/grizzly"
	"github.com/grafana/grizzly/pkg/prometheus"
)

// Version is the current version of the grr command.
// To be overwritten at build time
var Version = "dev"

func main() {
	log.SetFlags(0)

	rootCmd := &cli.Command{
		Use:     "grr",
		Short:   "Grizzly",
		Version: Version,
	}

	registry, err := GetProviderRegistry()
	if err != nil {
		log.Fatalln(err)
	}

	config := grizzly.Config{
		Registry: registry,
		Notifier: grizzly.Notifier{},
	}
	// workflow commands
	rootCmd.AddCommand(
		getCmd(config),
		listCmd(config),
		showCmd(config),
		diffCmd(config),
		applyCmd(config),
		watchCmd(config),
		listenCmd(config),
		exportCmd(config),
		previewCmd(config),
		providersCmd(config),
	)

	// Run!
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

// GetProviderRegistry registers all known providers
func GetProviderRegistry() (grizzly.Registry, error) {
	registry := grizzly.NewProviderRegistry()
	registry.RegisterProvider(&grafana.Provider{})
	registry.RegisterProvider(&prometheus.Provider{})
	return registry, nil
}
