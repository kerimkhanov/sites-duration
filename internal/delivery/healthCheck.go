package delivery

import "github.com/gin-gonic/gin"

func (h *SiteHandler) HealthCheck(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}
	c.JSON(200, gin.H{
		"data": res,
	})
}

func (h *SiteHandler) HealthReady(c *gin.Context) {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}
	c.JSON(200, gin.H{
		"data": res,
	})
}
