package main

import (
	"auth-and-gateway-microservice/auth"
	db "auth-and-gateway-microservice/db/sqlc"
	"auth-and-gateway-microservice/logging"
	"auth-and-gateway-microservice/middleware"
	"auth-and-gateway-microservice/utils"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"net/http"
)

var (
	store          *db.Store
	logger         logging.Logger
	ctx            context.Context
	authController *auth.Controller
	authRoutes     auth.Routes
)

func init() {
	ctx = context.TODO()
	logger = logging.GetLogger()
	config, err := utils.LoadConfig(".")
	if err != nil {
		logger.Fatalf("could not load config: %v", err)
	}
	conn, err := sql.Open(config.PostgresDriver, config.PostgresSource)
	if err != nil {
		logger.Fatalf("could not connect to postgres database: %v", err)
	}
	// Ping database
	err = conn.Ping()
	if err != nil {
		logger.Fatalf("could not ping to postgres database: %v", err)
	}
	store = db.NewStore(conn)
	authController = auth.NewController(store)
	authRoutes = auth.NewRoutes(authController, store)
}

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		logger.Fatalf("could not load config: %v", err)
	}
	// migrations
	err = runDBMigration(config.MigrationURL, config.PostgresSource)
	if err != nil {
		logger.Fatalf("migrations error: %s", err.Error())
	}
	// setup router, and run server
	server := setupRouter()
	logger.Fatal(server.Run(fmt.Sprintf("%s:%s", config.HttpServerAddress, config.Port)))
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.OpenCORSMiddleware())
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Route %s not found", ctx.Request.URL)})
	})
	v1 := router.Group("/v1")
	apiV1 := v1.Group("/api")
	authRoutes.AuthRoute(apiV1)
	return router
}

func runDBMigration(migrationURL string, dbSource string) error {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		return err
	}
	//
	if err = migration.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}
