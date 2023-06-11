package delivery

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/kerimkhanov/sites-duration/internal/models"
	"github.com/kerimkhanov/sites-duration/internal/service"
	"net/http"
	"strings"
)

type SiteHandler struct {
	siteService *service.SiteService
}

func NewSiteHandler(siteService *service.SiteService) *SiteHandler {
	return &SiteHandler{
		siteService: siteService,
	}
}

func (h *SiteHandler) GetSiteAccessTime(c *gin.Context) {
	var (
		site models.SiteAccessTime
		err  error
	)
	siteName := strings.ToLower(c.Query("name"))
	site.SetSite(siteName)
	err = h.siteService.GetSiteAccessTime(&site)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrGettingSite)
		return
	}
	c.JSON(http.StatusOK, site)
}

func (h *SiteHandler) GetSiteWithMinAccessTime(c *gin.Context) {
	var site models.SiteAccessTime
	err := h.siteService.GetSiteMinDuration(&site)
	if err != nil {
		if errors.Is(err, models.NotFound) {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusInternalServerError, models.ErrGettingSite)
		return
	}

	c.JSON(http.StatusOK, site)
}

func (h *SiteHandler) GetSiteWithMaxAccessTime(c *gin.Context) {
	var site models.SiteAccessTime
	err := h.siteService.GetSiteMaxDuration(&site)
	if err != nil {
		if errors.Is(err, models.NotFound) {
			c.JSON(http.StatusBadRequest, err)
		}
		c.JSON(http.StatusInternalServerError, models.ErrGettingSite)
		return
	}

	c.JSON(http.StatusOK, site)
}

func (h *SiteHandler) GetAdminStats(c *gin.Context) {
	var (
		stats = make(map[string]int)
	)
	stats = make(map[string]int)

	timeRequests := h.siteService.GetEndpointRequests("/time")
	stats["time"] = timeRequests

	minRequests := h.siteService.GetEndpointRequests("/min")
	stats["min"] = minRequests

	maxRequests := h.siteService.GetEndpointRequests("/max")
	stats["max"] = maxRequests

	c.JSON(http.StatusOK, stats)
}
