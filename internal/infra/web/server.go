package web

import (
	"context"
	"fmt"
	"go-htmx/internal/infra/config"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-contrib/multitemplate"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type (
	Middleware interface {
		Handler() gin.HandlerFunc
	}

	Handler interface {
		Group() string
		Method() string
		Path() string
		Handle(*gin.Context)
	}
)

const startTimeout = 2 * time.Second

func NewGinServer(
	middlewares []Middleware,
	handlers []Handler,
	lc fx.Lifecycle,
	renderer multitemplate.Renderer,
	log *zap.Logger,
	cfg *config.Config,
) *gin.Engine {
	r := gin.New()
	r.HTMLRender = renderer

	for _, m := range middlewares {
		r.Use(m.Handler())
	}

	for _, c := range handlers {
		r.Group(c.Group()).Handle(c.Method(), c.Path(), c.Handle)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r.Handler(),
	}

	lc.Append(fx.Hook{
		OnStart: onGinServerStart(server, log),
		OnStop: func(ctx context.Context) error {
			return server.Shutdown(ctx)
		},
	})
	return r
}

func NewRenderer(cfg *config.Config) multitemplate.Renderer {
	renderer := multitemplate.NewRenderer()

	// Load all template types
	layouts, _ := filepath.Glob(fmt.Sprintf("%s/layouts/*.html", cfg.TemplateFolder))
	partials, _ := filepath.Glob(fmt.Sprintf("%s/partials/*.html", cfg.TemplateFolder))
	components, _ := filepath.Glob(fmt.Sprintf("%s/components/**/*.html", cfg.TemplateFolder))
	pages, _ := filepath.Glob(fmt.Sprintf("%s/pages/**/*.html", cfg.TemplateFolder))

	// Generate template map
	for _, page := range pages {
		files := []string{page}
		files = append(files, layouts...)
		files = append(files, partials...)
		files = append(files, components...)

		name := strings.TrimPrefix(page, fmt.Sprintf("%s/pages/", cfg.TemplateFolder))
		name = strings.TrimSuffix(name, ".html")

		renderer.AddFromFiles(name, files...)
	}

	return renderer
}

func onGinServerStart(server *http.Server, log *zap.Logger) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		ch := make(chan error)

		ctx, cancel := context.WithTimeout(ctx, startTimeout)
		defer cancel()

		go func() {
			defer close(ch)

			if err := server.ListenAndServe(); err != http.ErrServerClosed {
				ch <- err
				log.Error("Failed to start server", zap.Error(err))
			}
		}()

		select {
		case <-ctx.Done():
			log.Info("web server started")
		case err := <-ch:
			return err
		}
		return nil
	}
}

type GinLoggerMiddleware struct {
	log *zap.Logger
}

func NewGinLoggerMiddleware(log *zap.Logger) *GinLoggerMiddleware {
	return &GinLoggerMiddleware{log: log}
}

func (m *GinLoggerMiddleware) Handler() gin.HandlerFunc {
	return ginzap.RecoveryWithZap(m.log, true)
}
