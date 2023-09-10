package app

import (
	"context"
	"fmt"
	"github.com/Amore14rn/888Starz_test/internal/config"
	pb "github.com/Amore14rn/888Starz_test/internal/controllers/http/v1/product"
	ub "github.com/Amore14rn/888Starz_test/internal/controllers/http/v1/user"
	policy_product "github.com/Amore14rn/888Starz_test/internal/domain/policy/products"
	policy_user "github.com/Amore14rn/888Starz_test/internal/domain/policy/user"
	ppd "github.com/Amore14rn/888Starz_test/internal/domain/products/dao"
	spd "github.com/Amore14rn/888Starz_test/internal/domain/products/service"
	"github.com/Amore14rn/888Starz_test/internal/domain/user/dao"
	"github.com/Amore14rn/888Starz_test/internal/domain/user/service"
	"github.com/Amore14rn/888Starz_test/pkg/common/core/clock"
	"github.com/Amore14rn/888Starz_test/pkg/common/core/closer"
	"github.com/Amore14rn/888Starz_test/pkg/common/core/identity"
	"github.com/Amore14rn/888Starz_test/pkg/common/logging"
	"github.com/Amore14rn/888Starz_test/pkg/errors"
	"github.com/Amore14rn/888Starz_test/pkg/graceful"
	"github.com/Amore14rn/888Starz_test/pkg/metric"
	psql "github.com/Amore14rn/888Starz_test/pkg/postgresql"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
	"net"
	"net/http"
	"time"
)

type App struct {
	cfg        *config.Config
	pgClient   *pgxpool.Pool
	router     *gin.Engine
	httpServer *http.Server
}

func NewApp(ctx context.Context, cfg *config.Config) (App, error) {
	logging.L(ctx).Info("router initializing")

	router := gin.Default()

	logging.WithFields(ctx,
		logging.StringField("username", cfg.Postgres.User),
		logging.StringField("password", "<REMOVED>"),
		logging.StringField("host", cfg.Postgres.Host),
		logging.StringField("port", cfg.Postgres.Port),
		logging.StringField("database", cfg.Postgres.Database),
	).Info("PostgreSQL initializing")

	pgDsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.Database,
	)

	pgClient, err := psql.NewClient(ctx, 5, 3*time.Second, pgDsn, false)
	if err != nil {
		return App{}, errors.Wrap(err, "psql.NewClient")
	}

	closer.AddN(pgClient)

	logging.L(ctx).Info("heartbeat metric initializing")

	metricHandler := metric.Handler{}
	metricHandler.Register(router)

	cl := clock.New()
	generator := identity.NewGenerator()

	//User service
	userStorage := dao.NewUserStorage(pgClient)
	userService := service.NewUserService(userStorage)
	userPolicy := policy_user.NewUserPolicy(userService, generator, cl)
	userController := ub.NewUserHandler(userPolicy)

	logging.L(ctx).Info("handlers initializing")
	userGroup := router.Group("/user")
	{
		userGroup.POST("/create", userController.CreateUser)
		userGroup.GET("/all", userController.All)
		userGroup.GET("/get/:id", userController.GetUser)
		userGroup.POST("/get/:name", userController.GetUserByName)
		userGroup.PATCH("/update", userController.UpdateUser)
		userGroup.DELETE("/delete/:id", userController.DeleteUser)
		userGroup.POST("/create-order", userController.CreateOrder)
	}

	//Product service
	productStorage := ppd.NewProductDAO(pgClient)
	productService := spd.NewProductService(productStorage)
	productPolicy := policy_product.NewProductPolicy(productService, generator, cl)
	productController := pb.NewProductHandler(productPolicy)

	productGroup := router.Group("/product")
	{
		productGroup.POST("/create", productController.CreateProduct)
		productGroup.GET("/all", productController.All)
		productGroup.GET("/get/:id", productController.GetProduct)
		productGroup.PATCH("/update", productController.UpdateProduct)
		productGroup.DELETE("/delete/:id", productController.DeleteProduct)
	}

	return App{
		cfg:    cfg,
		router: router,
	}, nil

}

func (a *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.startHTTP(ctx)
	})
	return grp.Wait()
}

func (a *App) startHTTP(ctx context.Context) error {
	logger := logging.WithFields(ctx,
		logging.StringField("IP", a.cfg.Server.HOST),
		logging.StringField("Port", a.cfg.Server.PORT),
		logging.DurationField("WriteTimeout", a.cfg.Server.WriteTimeout),
		logging.DurationField("ReadTimeout", a.cfg.Server.ReadTimeout),
		logging.IntField("MaxHeaderBytes", a.cfg.Server.MaxHeaderBytes),
	)
	logger.Info("HTTP Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", a.cfg.Server.HOST, a.cfg.Server.PORT))
	if err != nil {
		logger.With(logging.ErrorField(err)).Fatal("failed to create listener")
	}

	handler := a.router

	a.httpServer = &http.Server{
		Handler:        handler,
		WriteTimeout:   a.cfg.Server.WriteTimeout,
		ReadTimeout:    a.cfg.Server.ReadTimeout,
		MaxHeaderBytes: a.cfg.Server.MaxHeaderBytes,
	}
	if err = a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.With(logging.ErrorField(err)).Fatal("failed to start server")
		}
	}

	httpErrChan := make(chan error, 1)
	httpShutdownChan := make(chan struct{})

	graceful.PerformGracefulShutdown(a.httpServer, httpErrChan, httpShutdownChan)

	return err
}
