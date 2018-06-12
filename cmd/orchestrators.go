package cmd

import (
	"fmt"

	"github.com/Azure/dcos-engine/pkg/api"
	"github.com/Azure/dcos-engine/pkg/helpers"
	"github.com/spf13/cobra"
)

const (
	orchestratorsName             = "orchestrators"
	orchestratorsShortDescription = "Display info about supported orchestrators"
	orchestratorsLongDescription  = "Display supported versions and upgrade versions for each orchestrator"
)

type orchestratorsCmd struct {
	// user input
	orchestrator string
	version      string
}

func newOrchestratorsCmd() *cobra.Command {
	oc := orchestratorsCmd{}

	command := &cobra.Command{
		Use:   orchestratorsName,
		Short: orchestratorsShortDescription,
		Long:  orchestratorsLongDescription,
		RunE: func(cmd *cobra.Command, args []string) error {
			return oc.run(cmd, args)
		},
	}

	f := command.Flags()
	f.StringVar(&oc.orchestrator, "orchestrator", "", "orchestrator name (optional) ")
	f.StringVar(&oc.version, "version", "", "orchestrator version (optional)")

	return command
}

func (oc *orchestratorsCmd) run(cmd *cobra.Command, args []string) error {
	orchs, err := api.GetOrchestratorVersionProfileListVLabs(oc.orchestrator, oc.version)
	if err != nil {
		return err
	}

	data, err := helpers.JSONMarshalIndent(orchs, "", "  ", false)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	return nil
}