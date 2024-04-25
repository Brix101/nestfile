package cmd

import (
	"context"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/spf13/cobra"
)

func Execute(ctx context.Context) int {

	profile := false

	rootCmd := &cobra.Command{
		Use:   "nestfile",
		Short: "A stylish local web-based file browser",
		Long:  `Nest File CLI lets you create the database to use with File Browser`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if !profile {
				return nil
			}

			f, perr := os.Create("cpu.pprof")
			if perr != nil {
				return perr
			}

			_ = pprof.StartCPUProfile(f)
			return nil
		},
		PersistentPostRunE: func(cmd *cobra.Command, args []string) error {
			if !profile {
				return nil
			}

			pprof.StopCPUProfile()

			f, perr := os.Create("mem.pprof")
			if perr != nil {
				return perr
			}
			defer f.Close()

			runtime.GC()
			err := pprof.WriteHeapProfile(f)
			return err
		},
	}
	rootCmd.SetVersionTemplate("File Browser version {{printf \"%s\" .Version}}\n")

	persistent := rootCmd.PersistentFlags()

	persistent.BoolVarP(&profile, "profile", "p", false, "record CPU pprof")

	rootCmd.AddCommand(APICmd(ctx))

	if err := rootCmd.Execute(); err != nil {
		return 1
	}

	return 0
}
