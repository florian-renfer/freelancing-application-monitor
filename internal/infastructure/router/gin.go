package router

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/api/action"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/logger"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/presenter"
	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/repository"
	"github.com/florian-renfer/freelancing-application-monitor/internal/usecase"

	"github.com/gin-gonic/gin"
)

type ginEngine struct {
	router     *gin.Engine
	log        logger.Logger
	db         repository.SQL
	port       Port
	ctxTimeout time.Duration
}

func newGinServer(
	port Port,
	log logger.Logger,
	db repository.SQL,
	t time.Duration,
) *ginEngine {
	return &ginEngine{
		router:     gin.New(),
		log:        log,
		db:         db,
		port:       port,
		ctxTimeout: t,
	}
}

func (g ginEngine) Listen() {
	gin.SetMode(gin.ReleaseMode)
	gin.Recovery()

	g.setAppHandlers(g.router)

	server := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", g.port),
		Handler:      g.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		g.log.WithFields(logger.Fields{"port": g.port}).Infof("Starting HTTP Server")
		if err := server.ListenAndServe(); err != nil {
			g.log.WithError(err).Errorf("%s", "Error starting HTTP server")
			os.Exit(1)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		g.log.WithError(err).Errorf("%s", "Server Shutdown Failed")
		os.Exit(1)
	}

	g.log.Infof("Service down")
}

func (g ginEngine) setAppHandlers(router *gin.Engine) {
	// Authentication
	router.POST("/v1/auth/register", g.authRegister())
	router.POST("/v1/auth/login", g.authLogin())
	router.POST("/v1/auth/logout", g.authLogout())

	// Users
	router.GET("/v1/users", g.usersFindAll())
	router.GET("/v1/users/:id", g.usersFind())
	router.DELETE("/v1/users/:id", g.usersDelete())

	// Projects
	router.GET("/v1/projects", g.projectsFindAll())
	router.POST("/v1/projects", g.projectsCreate())
	router.PUT("/v1/projects/:id", g.projectsUpdate())
	router.GET("/v1/projects/:id", g.projectsFind())
	router.DELETE("/v1/projects/:id", g.projectsDelete())

	// Applications
	router.GET("/v1/applications", g.applicationsFindAll())
	router.POST("/v1/applications", g.applicationsCreate())
	router.PUT("/v1/applications/:id", g.applicationsUpdate())
	router.GET("/v1/applications/:id", g.applicationsFind())
	router.DELETE("/v1/applications/:id", g.applicationsDelete())

	// Healthcheck
	router.GET("/v1/health", g.healthcheck())
}

func (g ginEngine) healthcheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) authRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			uc = usecase.NewCreateUserInteractor(
				repository.NewUserSQL(g.db),
				presenter.NewCreateUserPresenter(),
				g.ctxTimeout,
			)

			act = action.NewCreateUserAction(uc, g.log)
		)

		act.Execute(c.Writer, c.Request)
	}
}

func (g ginEngine) authLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) authLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) usersFindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) usersFind() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) usersDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) projectsFindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) projectsCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) projectsUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) projectsFind() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) projectsDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) applicationsFindAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) applicationsCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) applicationsUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) applicationsFind() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}

func (g ginEngine) applicationsDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		action.HealthCheck(c.Writer, c.Request)
	}
}
