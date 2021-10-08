// +build amd64,!windows arm64,!windows

package machine

import (
	"net/url"

	"github.com/containers/common/pkg/config"
	"github.com/containers/podman/v3/cmd/podman/registry"
	"github.com/containers/podman/v3/pkg/machine"
	"github.com/containers/podman/v3/pkg/machine/qemu"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	sshCmd = &cobra.Command{
		Use:   "ssh [NAME] [COMMAND [ARG ...]]",
		Short: "SSH into an existing machine",
		Long:  "SSH into a managed virtual machine ",
		RunE:  ssh,
		Example: `podman machine ssh myvm
  podman machine ssh myvm echo hello`,
		ValidArgsFunction: autocompleteMachineSSH,
		PreRunE:           noAarch64,
	}
)

var (
	sshOpts machine.SSHOptions
)

func init() {
	sshCmd.Flags().SetInterspersed(false)
	registry.Commands = append(registry.Commands, registry.CliCommand{
		Command: sshCmd,
		Parent:  machineCmd,
	})
}

func ssh(cmd *cobra.Command, args []string) error {
	var (
		err     error
		validVM bool
		vm      machine.VM
		vmType  string
	)

	// Set the VM to default
	vmName := defaultMachineName

	// If we're not given a VM name, use the remote username from the connection config
	if len(args) == 0 {
		sshOpts.Username, err = remoteConnectionUsername()
		if err != nil {
			return err
		}
	}
	// If len is greater than 0, it means we may have been
	// provided the VM name.  If so, we check.  The VM name,
	// if provided, must be in args[0].
	if len(args) > 0 {
		switch vmType {
		default:
			validVM, err = qemu.IsValidVMName(args[0])
			if err != nil {
				return err
			}
			if validVM {
				vmName = args[0]
			} else {
				sshOpts.Username, err = remoteConnectionUsername()
				if err != nil {
					return err
				}
				sshOpts.Args = append(sshOpts.Args, args[0])
			}
		}
	}

	// If len is greater than 1, it means we might have been
	// given a vmname and args or just args
	if len(args) > 1 {
		if validVM {
			sshOpts.Args = args[1:]
		} else {
			sshOpts.Username, err = remoteConnectionUsername()
			if err != nil {
				return err
			}
			sshOpts.Args = args
		}
	}

	switch vmType {
	default:
		vm, err = qemu.LoadVMByName(vmName)
	}
	if err != nil {
		return errors.Wrapf(err, "vm %s not found", vmName)
	}
	return vm.SSH(vmName, sshOpts)
}

func remoteConnectionUsername() (string, error) {
	cfg, err := config.ReadCustomConfig()
	if err != nil {
		return "", err
	}
	dest, _, err := cfg.ActiveDestination()
	if err != nil {
		return "", err
	}
	uri, err := url.Parse(dest)
	if err != nil {
		return "", err
	}
	username := uri.User.String()
	return username, nil
}
