package services

import (
	"fmt"
	"net/http"
	"pentagi/pkg/hub"

	"github.com/gin-gonic/gin"
)

type HubService struct {
}

func NewHubService() *HubService {
	return &HubService{}
}

func (s *HubService) GetRSS(c *gin.Context) {
	targetURL := c.Query("url")
	if targetURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "url is required"})
		return
	}

	allowed, _ := hub.IsURLAllowed(targetURL)
	if !allowed {
		c.JSON(http.StatusForbidden, gin.H{"error": "URL not allowed"})
		return
	}

	rss, err := hub.FetchRSS(targetURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rss)
}

func (s *HubService) GetFlights(c *gin.Context) {
	flights, err := hub.GetLiveFlights()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, flights)
}

func (s *HubService) GetVessels(c *gin.Context) {
	vessels, err := hub.GetLiveVessels()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, vessels)
}

func (s *HubService) ListAPIs(c *gin.Context) {
	c.JSON(http.StatusOK, hub.GetRegistry())
}

func (s *HubService) GetNearbySignals(c *gin.Context) {
	latStr := c.Query("lat")
	lonStr := c.Query("lon")
	mode := c.DefaultQuery("mode", "wifi")

	var lat, lon float64
	fmt.Sscanf(latStr, "%f", &lat)
	fmt.Sscanf(lonStr, "%f", &lon)

	devices, err := hub.GetNearbyDevices(lat, lon, mode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"devices": devices})
}

func (s *HubService) ProxyJSON(c *gin.Context) {
	targetURL := c.Query("url")
	if targetURL == "" {
		// If no explicit URL, try to resolve from path (for /api/v1/... style calls)
		// This is a placeholder. In a real scenario, we'd map these to internal logic or upstream APIs.
		c.JSON(http.StatusNotImplemented, gin.H{
			"error": "Generic API proxying by path not fully implemented. Use ?url= or implement specific handlers.",
			"path":  c.Param("path"),
		})
		return
	}

	allowed, _ := hub.IsURLAllowed(targetURL)
	if !allowed {
		c.JSON(http.StatusForbidden, gin.H{"error": "URL not allowed"})
		return
	}

	resp, err := http.Get(targetURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	c.DataFromReader(resp.StatusCode, resp.ContentLength, "application/json", resp.Body, nil)
}

func (s *HubService) ListMilitaryFlights(c *gin.Context) {
	// Mock response based on WorldMonitor expected format
	flights, _ := hub.GetLiveFlights()
	c.JSON(http.StatusOK, gin.H{
		"flights": flights,
		"clusters": []interface{}{},
	})
}
