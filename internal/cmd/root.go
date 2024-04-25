package cmd

import (
	"context"
	"io/fs"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/Brix101/nestfile/frontend"
	nthttp "github.com/Brix101/nestfile/internal/http"
)

func APICmd(ctx context.Context) *cobra.Command {
	var port int

	cmd := &cobra.Command{
		Use:   "api",
		Args:  cobra.ExactArgs(0),
		Short: "Runs the RESTful API.",
		RunE: func(_ *cobra.Command, args []string) error {
			port = 5000
			if os.Getenv("PORT") != "" {
				port, _ = strconv.Atoi(os.Getenv("PORT"))
			}
			_, err := fs.Sub(frontend.Assets(), "dist")
			if err != nil {
				panic(err)
			}
			api := nthttp.NewHTTPHandler()
			srv := api.Server(port)

			go func() { _ = srv.ListenAndServe() }()

			<-ctx.Done()

			_ = srv.Shutdown(ctx)

			return nil
		},
	}

	return cmd
}
