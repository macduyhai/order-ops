package main

import (
	"order-ops/controllers"
	"order-ops/daos"
	"order-ops/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func CORSMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func InitGin(db *gorm.DB) *gin.Engine {
	orderDao := daos.NewOrderDao(db)
	orderService := services.NewOrderService(orderDao)

	branchSellDao := daos.NewBranchSellDao(db)
	branchSellService := services.NewBranchSellService(branchSellDao)

	typeProductDao := daos.NewTypeProductDao(db)
	typeProductService := services.NewTypeProductService(typeProductDao)

	sellerDao := daos.NewSellerDao(db)
	sellerService := services.NewSellerService(sellerDao)

	ctl := controllers.Controller{
		OrderService:       orderService,
		BranchSellService:  branchSellService,
		TypeProductService: typeProductService,
		SellerService:      sellerService,
	}

	engine := gin.Default()
	engine.Use(CORSMiddleWare())

	engine.GET("/health", ctl.HealthCheck)
	apiGroup := engine.Group("/api/v1")
	{
		orderGroup := apiGroup.Group("/orders")
		{
			orderGroup.POST("", ctl.AddOrder)
			orderGroup.DELETE("", ctl.Delete)
			orderGroup.PUT("", ctl.UpdateOrders)
			orderGroup.GET("/search", ctl.Search)
			// orderGroup.GET("/number-orders", ctl.NumberOrder)
			orderGroup.POST("/make-done", ctl.MakeDone)
			orderGroup.POST("/delay", ctl.MakeDelay)
			orderGroup.POST("/shipping-time", ctl.AddShippingTime)
		}
		branchSellGroup := apiGroup.Group("/branchsells")
		{
			branchSellGroup.POST("", ctl.AddBranchSell)
			branchSellGroup.PUT("", ctl.UpdateBranch)
			branchSellGroup.GET("/search-branch", ctl.SearchBranch)
			branchSellGroup.POST("/delete", ctl.DeleteBranchSell)

		}
		typeProductGroup := apiGroup.Group("/typeproducts")
		{
			typeProductGroup.POST("", ctl.AddTypeProduct)
			typeProductGroup.PUT("", ctl.UpdateTypeProduct)
			typeProductGroup.GET("/search-type", ctl.SearchType)
			typeProductGroup.POST("/delete", ctl.DeleteTypeProduct)

		}
		sellerGroup := apiGroup.Group("/sellers")
		{
			sellerGroup.POST("", ctl.AddSeller)
			sellerGroup.PUT("", ctl.UpdateSeller)
			sellerGroup.GET("/search-seller", ctl.SearchSeller)
			sellerGroup.POST("/delete", ctl.DeleteSeller)

		}

		labelGroup := apiGroup.Group("/labels")
		{
			labelGroup.POST("", ctl.AddLabelToOrder)
		}
	}
	return engine
}
