package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ydhnwb/golang_heroku/config"
	v1 "github.com/ydhnwb/golang_heroku/handler/v1"
	"github.com/ydhnwb/golang_heroku/middleware"
	"github.com/ydhnwb/golang_heroku/repo"
	"github.com/ydhnwb/golang_heroku/service"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB               = config.SetupDatabaseConnection()
	userRepo       repo.UserRepository    = repo.NewUserRepo(db)
	productRepo    repo.ProductRepository = repo.NewProductRepo(db)
	photoRepo      repo.PhotoRepository   = repo.NewPhotoRepo(db)
	authService    service.AuthService    = service.NewAuthService(userRepo)
	jwtService     service.JWTService     = service.NewJWTService()
	userService    service.UserService    = service.NewUserService(userRepo)
	productService service.ProductService = service.NewProductService(productRepo)
	photoService   service.PhotoService   = service.NewPhotoService(photoRepo)
	authHandler    v1.AuthHandler         = v1.NewAuthHandler(authService, jwtService, userService)
	userHandler    v1.UserHandler         = v1.NewUserHandler(userService, jwtService)
	productHandler v1.ProductHandler      = v1.NewProductHandler(productService, jwtService)
	photoHandler   v1.PhotoHandler        = v1.NewPhotoHandler(photoService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	authRoutes := server.Group("api/auth")
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
	}

	userRoutes := server.Group("api/user", middleware.AuthorizeJWT(jwtService))
	{
		userRoutes.GET("/profile", userHandler.Profile)
		userRoutes.PUT("/profile", userHandler.Update)
	}

	productRoutes := server.Group("api/product", middleware.AuthorizeJWT(jwtService))
	{
		productRoutes.GET("/", productHandler.All)
		productRoutes.POST("/", productHandler.CreateProduct)
		productRoutes.GET("/:id", productHandler.FindOneProductByID)
		productRoutes.PUT("/:id", productHandler.UpdateProduct)
		productRoutes.DELETE("/:id", productHandler.DeleteProduct)
	}

	photoRoutes := server.Group("api/photo", middleware.AuthorizeJWT(jwtService))
	{
		photoRoutes.GET("/", photoHandler.All)
		photoRoutes.POST("/", photoHandler.CreatePhoto)
		photoRoutes.GET("/:id", photoHandler.FindOnePhotoByID)
		photoRoutes.PUT("/:id", photoHandler.UpdatePhoto)
		photoRoutes.DELETE("/:id", photoHandler.DeletePhoto)
	}

	checkRoutes := server.Group("api/check")
	{
		checkRoutes.GET("health", v1.Health)
	}

	server.Run()
}
