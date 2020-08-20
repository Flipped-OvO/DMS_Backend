package router

import (
	"backend/pkg/fs/handle"
	"github.com/gin-gonic/gin"
)

func Initialization(r *gin.RouterGroup) {
	company := r.Group("/company")
	company.GET("", handle.SearchCompanyByNameOrCode)
	company.GET("/", handle.SearchCompanyByNameOrCode)

	company.GET("/uninit", handle.GetUnInitCompany)
	company.GET("/collection", handle.GetCollectionCompany)
	company.POST("/collection", handle.NewCollectionCompany)
	company.DELETE("/collection", handle.RemoveCollectionCompany)

	company.GET("/profit", handle.GetCompanyProfit)
}
