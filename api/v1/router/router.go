package router

import (
	"Server/api/v1/handler"
	_ "Server/cmd/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Server() {
	engine := gin.Default()
	//SetupSwagger()
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/admin", handler.GetAdminHandle)
	engine.Run(":10101")
}

//// SetupSwagger
//func SetupSwagger() {
//	docs.SwaggerInfo.Title = "Swagger Example API"
//	docs.SwaggerInfo.Description = "This is a sample server for Swagger."
//	docs.SwaggerInfo.Version = "1.0"
//	docs.SwaggerInfo.Host = "petstore.swagger.io"
//	docs.SwaggerInfo.BasePath = "/v2"
//}

//func Server() {
//	mux := http.NewServeMux()
//	mux.HandleFunc("admin", handler.AdminHandler)
//	err := http.ListenAndServe(":10101", http.HandlerFunc(handler.AdminHandler))
//	if errors.Is(err, http.ErrServerClosed) {
//		fmt.Printf("Server closed\n")
//	} else if err != nil {
//		fmt.Printf("error starting server: %s\n", err)
//		os.Exit(1)
//	}
//}
