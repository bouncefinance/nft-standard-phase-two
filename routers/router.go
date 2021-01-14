package routers

import (
	"github.com/gin-gonic/gin"
	"nft_standard/config"
	"nft_standard/handler"
)


func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(middleware.TlsHandler())

	gin.SetMode(config.RunMode)

	r.GET("/erc721",handler.GetERC721)
	r.GET("/erc1155",handler.GetERC1155)
	r.GET("/nft",handler.GetAllNFT)

	//assetsGroup:=r.Group("/assets")
	//{
	//	assetsGroup.GET("/erc721",handler.GetMetadata721)
	//	assetsGroup.GET("/erc1155",handler.GetMetadata1155)
	//}

	return r
}

