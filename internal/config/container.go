package config

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

func Start(lc fx.Lifecycle, sd fx.Shutdowner, e *echo.Echo, p *Properties) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			slog.Info(fmt.Sprintf("Process ID: %d on %s", os.Getpid(), p.Env.App.Environment))

			errChan := make(chan error, 1)

			go func() {
				if err := e.Start(fmt.Sprintf(":%d", p.Env.App.ListenPort)); err != nil && err != http.ErrServerClosed {
					errChan <- err
				}
			}()

			select {
			case <-ctx.Done():
				return ctx.Err()
			case err := <-errChan:
				return err
			case <-time.After(100 * time.Millisecond):
				go func() {
					if err := <-errChan; err != nil {
						slog.Error(err.Error())
						sd.Shutdown()
					}
				}()
				return nil
			}
		},
		OnStop: func(ctx context.Context) error {
			slog.Info("Shutting down HTTP server...")
			return e.Shutdown(ctx)
		},
	})
}
