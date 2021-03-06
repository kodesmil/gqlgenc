package main

import (
	"context"
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/kodesmil/gqlgenc/clientgen"
	"github.com/kodesmil/gqlgenc/config"
	"github.com/kodesmil/gqlgenc/generator"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err.Error())
		os.Exit(2)
	}

	clientPlugin := clientgen.New(cfg.Query, cfg.Client, cfg.Generate)
	if err := generator.Generate(ctx, cfg, api.AddPlugin(clientPlugin)); err != nil {
		fmt.Fprintf(os.Stderr, "%+v", err.Error())
		os.Exit(4)
	}
}
