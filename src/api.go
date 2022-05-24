package src

import (
	"log"
	"net/http"

	"github.com/MAbduhI/gin-test/config"
	"github.com/MAbduhI/gin-test/src/batik"
	"github.com/gin-gonic/gin"
)

type check struct {
	Config *config.Config
}

type healthCheckResponse struct {
	DBStatus status `json:"db_status"`
}

type status struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

func (c check) healthCheck(ctx *gin.Context) {
	res := healthCheckResponse{
		DBStatus: status{
			Status: true,
		},
	}
	db, err := c.Config.DB.DB()
	if err != nil {
		log.Printf("[healthCheck] Error get DB:%+v", err)
		res.DBStatus = status{
			Message: err.Error(),
			Status:  false,
		}
	}

	err = db.Ping()
	if err != nil {
		log.Printf("[healthCheck] Error ping DB:%+v", err)
		res.DBStatus = status{
			Message: err.Error(),
			Status:  false,
		}
	}

	ctx.IndentedJSON(http.StatusOK, res)
}

func InitApi(conf *config.Config) {

	c := check{
		Config: conf,
	}

	batikSvc := batik.InitServices(conf)
	router := gin.Default()
	router.GET("/health", c.healthCheck)
	router.GET("/data", batikSvc.GetItem)
	router.POST("/datapost", batikSvc.SaveItem)
	router.PATCH("/datapost", batikSvc.UpdateItem)
	router.DELETE("/datapost", batikSvc.DeleteItem)
	router.Run("localhost:8001")

}
