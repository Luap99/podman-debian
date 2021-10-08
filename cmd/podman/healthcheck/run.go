package healthcheck

import (
	"context"
	"fmt"

	"github.com/containers/podman/v3/cmd/podman/common"
	"github.com/containers/podman/v3/cmd/podman/registry"
	"github.com/containers/podman/v3/libpod/define"
	"github.com/containers/podman/v3/pkg/domain/entities"
	"github.com/spf13/cobra"
)

var (
	runCmd = &cobra.Command{
		Use:               "run CONTAINER",
		Short:             "run the health check of a container",
		Long:              "run the health check of a container",
		Example:           `podman healthcheck run mywebapp`,
		RunE:              run,
		Args:              cobra.ExactArgs(1),
		ValidArgsFunction: common.AutocompleteContainersRunning,
	}
)

func init() {
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Command: runCmd,
		Parent:  healthCmd,
	})
}

func run(cmd *cobra.Command, args []string) error {
	response, err := registry.ContainerEngine().HealthCheckRun(context.Background(), args[0], entities.HealthCheckOptions{})
	if err != nil {
		return err
	}
	if response.Status == define.HealthCheckUnhealthy {
		registry.SetExitCode(1)
		fmt.Println(response.Status)
	}
	return err
}
