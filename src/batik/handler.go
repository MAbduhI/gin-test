package batik

import (
	"log"
	"net/http"

	"github.com/MAbduhI/gin-test/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Svc struct {
	Conf Conf
}
type Conf struct {
	DB  *gorm.DB
	Log log.Logger
}

func InitServices(c *config.Config) Svc {
	s := Svc{
		Conf: Conf{
			DB: c.DB,
		},
	}
	return s
}

func (s Svc) GetItem(ctx *gin.Context) {
	item, err := itemGetService(s.Conf.DB)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseItemDetail{
			Success: false,
		})
	}

	ctx.JSON(http.StatusOK, item)
}

func (s Svc) SaveItem(ctx *gin.Context) {
	var data Item

	// data := ctx.BindJSON(itemReq)
	ctx.BindJSON(&data)
	item, err := itemPostService(s.Conf.DB, data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseItemDetail{
			Success: false,
		})
	}

	ctx.JSON(http.StatusOK, item)
}

func (s Svc) UpdateItem(ctx *gin.Context) {
	var data Item

	// data := ctx.BindJSON(itemReq)
	ctx.BindJSON(&data)
	item, err := itemPatchService(s.Conf.DB, data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseItemDetail{
			Success: false,
		})
	}

	ctx.JSON(http.StatusOK, item)
}

func (s Svc) DeleteItem(ctx *gin.Context) {
	var data Item

	// data := ctx.BindJSON(itemReq)
	ctx.BindJSON(&data)
	item, err := itemDeleteService(s.Conf.DB, data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, ResponseItemDetail{
			Success: false,
		})
	}

	ctx.JSON(http.StatusOK, item)
}
