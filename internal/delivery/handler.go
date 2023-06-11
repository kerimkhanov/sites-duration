package delivery

import "github.com/gin-gonic/gin"

func (h *SiteHandler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.GET("/health/live", h.HealthCheck)
	router.GET("/health/ready", h.HealthReady)

	router.GET("/time", h.GetSiteAccessTime)
	router.GET("/min", h.GetSiteWithMinAccessTime)
	router.GET("/max", h.GetSiteWithMaxAccessTime)
	router.GET("/admin/stats", h.GetAdminStats)

	return router
}
