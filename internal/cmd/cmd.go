package cmd

import (
	"context"
	"net/http"

	"github.com/spf13/cobra"
)

func Execute(ctx context.Context) int {
	rootCmd := &cobra.Command{
		Use:   "nestfile",
		Short: "A stylish local web-based file browser",
		Long:  `Nest File CLI lets you create the database to use with File Browser`,
	}
	rootCmd.SetVersionTemplate("File Browser version {{printf \"%s\" .Version}}\n")

	flags := rootCmd.Flags()
	persistent := rootCmd.PersistentFlags()

	persistent.StringP("database", "d", "./nestfile.db", "database path")
	flags.StringP("port", "p", "8080", "port to listen on")
	flags.String("cache-dir", "", "file cache directory (disabled if empty)")
	flags.String("username", "admin", "username for the first user when using quick config")
	flags.String("password", "", "hashed password for the first user when using quick config (default \"admin\")")
	rootCmd.AddCommand(APICmd(ctx))

	go func() {
		_ = http.ListenAndServe("localhost:6060", nil)
	}()

	if err := rootCmd.Execute(); err != nil {
		return 1
	}

	return 0
}
