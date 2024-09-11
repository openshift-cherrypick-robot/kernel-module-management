package main

import (
	"fmt"

	"github.com/rh-ecosystem-edge/kernel-module-management/internal/utils"
	"github.com/rh-ecosystem-edge/kernel-module-management/internal/worker"
	"github.com/spf13/cobra"
)

func rootFuncPreRunE(cmd *cobra.Command, args []string) error {
	logger.Info("Starting worker", "version", Version, "git commit", GitCommit)

	mr := worker.NewModprobeRunner(logger)
	fsh := utils.NewFSHelper(logger)
	w = worker.NewWorker(mr, fsh, logger)

	return nil
}

func kmodLoadFunc(cmd *cobra.Command, args []string) error {
	cfgPath := args[0]

	logger.Info("Reading config", "path", cfgPath)

	cfg, err := configHelper.ReadConfigFile(cfgPath)
	if err != nil {
		return fmt.Errorf("could not read config file %s: %v", cfgPath, err)
	}

	mountPathFlag := cmd.Flags().Lookup(worker.FlagFirmwarePath)
	if mountPathFlag.Changed {
		logger.V(1).Info(worker.FlagFirmwarePath + " set, setting firmware_class.path")

		if err := w.SetFirmwareClassPath(mountPathFlag.Value.String()); err != nil {
			return fmt.Errorf("could not set the firmware_class.path parameter: %v", err)
		}
	}

	return w.LoadKmod(cmd.Context(), cfg, mountPathFlag.Value.String())
}

func kmodUnloadFunc(cmd *cobra.Command, args []string) error {
	cfgPath := args[0]

	logger.Info("Reading config", "path", cfgPath)

	cfg, err := configHelper.ReadConfigFile(cfgPath)
	if err != nil {
		return fmt.Errorf("could not read config file %s: %v", cfgPath, err)
	}

	return w.UnloadKmod(cmd.Context(), cfg, cmd.Flags().Lookup(worker.FlagFirmwarePath).Value.String())
}

func setCommandsFlags() {
	kmodLoadCmd.Flags().String(
		worker.FlagFirmwarePath,
		"",
		"if set, this value will be written to "+worker.FirmwareClassPathLocation+" and it is also the value that firmware host path is mounted to")

	kmodLoadCmd.Flags().Bool(
		"tarball",
		false,
		"If true, extract the image from a tarball image instead of pulling from the registry",
	)

	kmodUnloadCmd.Flags().String(
		worker.FlagFirmwarePath,
		"",
		"if set, this the value that firmware host path is mounted to")

	kmodUnloadCmd.Flags().Bool(
		"tarball",
		false,
		"If true, extract the image from a tarball image instead of pulling from the registry",
	)
}
