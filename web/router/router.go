package router

//import (
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//)
//
//func RegisterRouter(r *gin.Engine) {
//	// 探活用
//	r.GET("/", func(c *gin.Context) {
//		c.JSON(http.StatusOK, gin.H{
//			"message": "Hello World!",
//		})
//	})
//
//	api := r.Group("/api/v1")
//	{
//		System := api.Group("/system")
//		{
//			RegisterSystemAPI(System)
//		}
//		SQL := api.Group("/sql")
//		{
//			RegisterSQLAPI(SQL)
//		}
//	}
//}
