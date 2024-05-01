package cmd

import (
	"context"
	"io/fs"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/Brix101/nestfile/frontend"
	"github.com/Brix101/nestfile/internal/api"
	"github.com/Brix101/nestfile/internal/util"
)

func APICmd(ctx context.Context) *cobra.Command {
	var port int

	apiCmd := &cobra.Command{
		Use:   "api",
		Args:  cobra.ExactArgs(0),
		Short: "Runs the RESTful API.",
		RunE: func(_ *cobra.Command, args []string) error {
			port = 5000

			logger := util.NewLogger("api")
			defer func() { _ = logger.Sync() }()

			db, err := util.NewDatabase(ctx)
			if err != nil {
				return err
			}
			db.Close()

			assetsFs, err := fs.Sub(frontend.Assets(), "dist")
			if err != nil {
				return err
			}

			api := api.NewHTTPHandler(ctx, logger, db, assetsFs)
			srv := api.Server(port)

			go func() { _ = srv.ListenAndServe() }()

			logger.Info("🚀🚀🚀 Server at port: ", zap.Int("port", port))

			<-ctx.Done()

			_ = srv.Shutdown(ctx)

			return nil
		},
	}

	return apiCmd
}
