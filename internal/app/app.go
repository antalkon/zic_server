package app

import (
	"backend/internal/transport/rest/router"
	"backend/pkg/cache"
	"backend/pkg/config"
	"backend/pkg/db"
	"backend/pkg/logger"
	"backend/pkg/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

type App struct {
	Config *config.Config
	Logger *logger.Logger
	DB     *db.Database
	Cache  *cache.RedisClient
	S3     *storage.MinIOClient
}

func NewApp() (*App, error) {
	// Load config
	cfg := config.LoadConfig()

	// Init logger
	l := logger.NewLogger(cfg.ServerEnv)
	l.Info("ZAP Logger initialized")

	// Init db (Postgres)
	database, err := db.NewDatabase(cfg)
	if err != nil {
		l.Fatal("Database initialization failed", zap.Error(err))
		return nil, err
	}
	l.Info("Database (Postgres) initialized")

	// Init Redis
	redisClient, err := cache.NewRedisClient(cfg)
	if err != nil {
		l.Fatal("Redis initialization failed", zap.Error(err))
		return nil, err
	}
	l.Info("Redis initialized")

	// Init MinIO (S3)
	s3Client, err := storage.NewMinIOClient(cfg)
	if err != nil {
		l.Fatal("MinIO initialization failed", zap.String("error", err.Error()))
		return nil, err
	}
	l.Info("MinIO initialized")

	return &App{
		Config: cfg,
		Logger: l,
		DB:     database,
		Cache:  redisClient,
		S3:     s3Client,
	}, nil
}

// Run app
func (a *App) Run() {
	a.Logger.Info("Application is running...")
}

func (a *App) RunServer() error {
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(a.Logger.WithEchoMiddleware())

	// Routes
	router.SetupRouter(e, a.Config, a.Logger, a.DB, a.Cache, a.S3)

	return e.Start(":" + a.Config.ServerAddress)
}
