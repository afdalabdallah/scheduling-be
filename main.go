package main

import (
	"github.com/afdalabdallah/backend-web/controllers"
	"github.com/afdalabdallah/backend-web/initializers"
	"github.com/afdalabdallah/backend-web/repository/dosen_repository/dosen_pg"
	"github.com/afdalabdallah/backend-web/repository/matkul_repository/matkul_pg"
	"github.com/afdalabdallah/backend-web/repository/rumpun_repository/rumpun_pg"
	"github.com/afdalabdallah/backend-web/services"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabese()
}

func main() {

	rumpunRepo := rumpun_pg.NewPGRumpunRepository(initializers.DB)
	rumpunService := services.NewRumpunService(rumpunRepo)
	rumpunController := controllers.NewRumpunController(rumpunService)

	matkulRepo := matkul_pg.NewPGMatkulRepository(initializers.DB)
	matkulService := services.NewMatkulService(matkulRepo, rumpunRepo)
	matkulController := controllers.NewMatkulController(matkulService)

	dosenRepo := dosen_pg.NewPGDosenRepository(initializers.DB)
	dosenService := services.NewDosenService(dosenRepo, rumpunRepo)
	dosenController := controllers.NewDosenController(dosenService)

	route := gin.Default()
	// c := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:3000"}, // Replace with your frontend domain or address
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Add any other allowed headers
	// 	AllowCredentials: true,                                      // Set to true if you need to include credentials (cookies, HTTP authentication)
	// })
	// handler := c.Handler(route)

	route.Use(CORSMiddleware())
	rumpunRoute := route.Group("/rumpun")
	{
		rumpunRoute.POST("/", rumpunController.CreateRumpun)
		rumpunRoute.GET("/", rumpunController.GetAllRumpun)
		rumpunRoute.GET("/:rumpunID", rumpunController.GetRumpunById)
		rumpunRoute.PUT("/:rumpunID", rumpunController.UpdateRumpun)
		rumpunRoute.DELETE("/:rumpunID", rumpunController.DeleteRumpun)
	}

	matkulRoute := route.Group("/matkul")
	{
		matkulRoute.POST("/", matkulController.CreateMatkul)
		matkulRoute.GET("/", matkulController.GetAllMatkul)
		matkulRoute.GET("/:matkulID", matkulController.GetMatkulById)
		matkulRoute.PUT("/:matkulID", matkulController.UpdateMatkul)
		matkulRoute.DELETE("/:matkulID", matkulController.DeleteMatkul)
	}

	dosenRoute := route.Group("/dosen")
	{
		dosenRoute.POST("/", dosenController.CreateDosen)
		dosenRoute.GET("/", dosenController.GetAllDosen)
		dosenRoute.GET("/:dosenID", dosenController.GetDosenById)
		dosenRoute.PUT("/:dosenID", dosenController.UpdateDosen)
		dosenRoute.DELETE("/:dosenID", dosenController.DeleteDosen)
	}

	// Enable CORS with permissive options (allowing all origins, methods, and headers)

	// Wrap your mux with the CORS middleware

	// http.ListenAndServe(":5000", handler)
	// http.ListenAndServe(":5000", handler)
	// if err != nil {
	// 	panic(err)
	// }
	route.Run()
}

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		/*
		   c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		   c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		   c.Writer.Header().Set("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
		   c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, HEAD, POST, PUT, DELETE, OPTIONS, PATCH")
		*/

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
