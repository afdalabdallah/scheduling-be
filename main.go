package main

import (
	"github.com/afdalabdallah/backend-web/controllers"
	"github.com/afdalabdallah/backend-web/initializers"
	"github.com/afdalabdallah/backend-web/repository/dosen_repository/dosen_pg"
	"github.com/afdalabdallah/backend-web/repository/jadwal_repository/jadwal_pg"
	"github.com/afdalabdallah/backend-web/repository/matkul_repository/matkul_pg"
	"github.com/afdalabdallah/backend-web/repository/perkuliahan_repository/perkuliahan_pg"
	"github.com/afdalabdallah/backend-web/repository/ruangan_repository/ruangan_pg"
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

	ruanganRepo := ruangan_pg.NewPGRuanganRepository(initializers.DB)
	ruanganService := services.NewRuanganService(ruanganRepo)
	ruanganController := controllers.NewRuanganController(ruanganService)

	matkulRepo := matkul_pg.NewPGMatkulRepository(initializers.DB)
	matkulService := services.NewMatkulService(matkulRepo, rumpunRepo)
	matkulController := controllers.NewMatkulController(matkulService)

	dosenRepo := dosen_pg.NewPGDosenRepository(initializers.DB)
	dosenService := services.NewDosenService(dosenRepo, rumpunRepo)
	dosenController := controllers.NewDosenController(dosenService)

	perkuliahanRepo := perkuliahan_pg.NewPGPerkuliahanRepository(initializers.DB)
	perkuliahanService := services.NewPerkuliahanService(rumpunRepo, matkulRepo, dosenRepo, perkuliahanRepo)
	perkuliahanController := controllers.NewPerkuliahanController(perkuliahanService)

	jadwalRepo := jadwal_pg.NewPGJadwalRepository(initializers.DB)
	jadwalService := services.NewJadwalService(jadwalRepo)
	jadwalController := controllers.NewJadwalController(jadwalService)

	route := gin.Default()

	route.Use(CORSMiddleware())

	jadwalRoute := route.Group("/jadwal")
	{
		jadwalRoute.POST("/", jadwalController.CreateJadwal)
		jadwalRoute.GET("/", jadwalController.GetAllJadwal)
		jadwalRoute.GET("/:jadwalID", jadwalController.GetJadwalById)
		jadwalRoute.PUT("/:jadwalID", jadwalController.UpdateJadwal)
		jadwalRoute.DELETE("/:jadwalID", jadwalController.DeleteJadwal)
	}

	ruanganRoute := route.Group("/ruangan")
	{
		ruanganRoute.POST("/", ruanganController.CreateRuangan)
		ruanganRoute.GET("/", ruanganController.GetAllRuangan)
		ruanganRoute.GET("/:ruanganID", ruanganController.GetRuanganById)
		ruanganRoute.PUT("/:ruanganID", ruanganController.UpdateRuangan)
		ruanganRoute.DELETE("/:ruanganID", ruanganController.DeleteRuangan)
	}

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

	perkuliahanRoute := route.Group("/perkuliahan")
	{
		perkuliahanRoute.POST("/", perkuliahanController.CreatePerkuliahan)
		perkuliahanRoute.GET("/", perkuliahanController.GetAllPerkuliahan)
		perkuliahanRoute.GET("/jadwal", perkuliahanController.GetPerkuliahanFormat)
		perkuliahanRoute.GET("/:perkuliahanID", perkuliahanController.GetPerkuliahanById)
		perkuliahanRoute.PUT("/:perkuliahanID", perkuliahanController.UpdatePerkuliahan)
		perkuliahanRoute.DELETE("/:perkuliahanID", perkuliahanController.DeletePerkuliahan)
	}

	route.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			print("masuk ke opstion")
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
