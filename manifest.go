package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/genuinetools/reg/registry"
)

const manifestHelp = `Get the json manifest for a repository.`

func (cmd *manifestCommand) Name() string      { return "manifest" }
func (cmd *manifestCommand) Args() string      { return "[OPTIONS] NAME[:TAG|@DIGEST]" }
func (cmd *manifestCommand) ShortHelp() string { return manifestHelp }
func (cmd *manifestCommand) LongHelp() string  { return manifestHelp }
func (cmd *manifestCommand) Hidden() bool      { return false }

func (cmd *manifestCommand) Register(fs *flag.FlagSet) {
}

type manifestCommand struct {
}

func (cmd *manifestCommand) Run(ctx context.Context, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("pass the name of the repository")
	}

	image, err := registry.ParseImage(args[0])
	if err != nil {
		return err
	}

	// Create the registry client.
	r, err := createRegistryClient(ctx, image.Domain)
	if err != nil {
		return err
	}

	var manifest interface{}
	// Get the v2 manifest.
	manifest, err = r.Manifest(ctx, image.Path, image.Reference())
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(manifest, " ", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
