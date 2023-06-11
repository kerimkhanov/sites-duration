package service

import (
	"github.com/go-redis/redis"
	"github.com/kerimkhanov/sites-duration/internal/models"
	"strconv"
)

type SiteService struct {
	redisClient      *redis.Client
	endpointRequests map[string]int
}

func NewSiteService(redisClient *redis.Client) *SiteService {
	return &SiteService{
		redisClient:      redisClient,
		endpointRequests: make(map[string]int),
	}
}

func (s *SiteService) GetSiteAccessTime(site *models.SiteAccessTime) error {
	var (
		err             error
		accessTimeFloat float64
		accessTime      string
	)
	accessTime, err = s.redisClient.Get(site.Site).Result()
	if err != nil {
		if err == redis.Nil {
			return models.NotFound
		}
		return err
	}
	accessTimeFloat, err = strconv.ParseFloat(accessTime, 64)
	site.SetAccessTime(accessTimeFloat)
	s.incrementEndpointRequests("/time")
	return nil
}

func (s *SiteService) GetSiteMinDuration(site *models.SiteAccessTime) error {
	var (
		err          error
		minName      string
		minTime      string
		minTimeFloat float64
	)
	minName, err = s.redisClient.Get(models.MinName).Result()
	if err != nil {
		if err == redis.Nil {
			return models.NotFound
		}
		return err
	}
	minTime, err = s.redisClient.Get(models.MinTime).Result()
	if err != nil {
		if err == redis.Nil {
			return models.NotFound
		}
		return err
	}
	minTimeFloat, err = strconv.ParseFloat(minTime, 64)
	site.SetSite(minName)
	site.SetAccessTime(minTimeFloat)
	s.incrementEndpointRequests("/min")
	return nil
}

func (s *SiteService) GetSiteMaxDuration(site *models.SiteAccessTime) error {
	var (
		maxTimeFloat float64
		err          error
		maxName      string
		maxTime      string
	)
	maxName, err = s.redisClient.Get(models.MaxName).Result()
	if err != nil {
		if err == redis.Nil {
			return models.NotFound
		}
		return err
	}
	maxTime, err = s.redisClient.Get(models.MaxTime).Result()
	if err != nil {
		if err == redis.Nil {
			return models.NotFound
		}
		return err
	}
	maxTimeFloat, err = strconv.ParseFloat(maxTime, 64)
	site.SetSite(maxName)
	site.SetAccessTime(maxTimeFloat)
	s.incrementEndpointRequests("/max")
	return nil
}

func (s *SiteService) GetEndpointRequests(endpoint string) int {
	if count, ok := s.endpointRequests[endpoint]; ok {
		return count
	}
	return 0
}

func (s *SiteService) incrementEndpointRequests(endpoint string) {
	if s.endpointRequests == nil {
		s.endpointRequests = make(map[string]int)
	}
	s.endpointRequests[endpoint]++
}
