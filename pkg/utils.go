package utils

import (
	"github.com/go-redis/redis"
	"github.com/kerimkhanov/sites-duration/internal/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	protocol = "http://"
)

func CheckAndSave(client *redis.Client, config *viper.Viper, lg *logrus.Logger) {
	var (
		err        error
		site       string
		startTime  time.Time
		resp       *http.Response
		accessTime float64
		minSite    string
		minTime    float64
		maxSite    string
		maxTime    float64
		minTimeStr string
		maxTimeStr string
	)

	sites := config.GetString("APP.SITES.LISTS")
	sitesArr := strings.Split(strings.ReplaceAll(sites, " ", ""), ",")

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	siteAccessTimes := make([]models.SiteAccessTime, 0, len(sitesArr))

	for range ticker.C {
		siteAccessTimes = nil
		for _, site = range sitesArr {
			startTime = time.Now()

			resp, err = http.Get(protocol + site)
			if err != nil {
				lg.Println("Error:", err)
				continue
			}
			resp.Body.Close()

			accessTime = time.Since(startTime).Seconds()

			err = client.Set(site, accessTime, 0).Err()
			if err != nil {
				lg.Println("Error saving access time:", err)
			}
			lg.Println("Access time for", site, ":", accessTime)

			siteAccessTimes = append(siteAccessTimes, models.SiteAccessTime{
				Site:       site,
				AccessTime: accessTime,
			})
		}

		quickSort(siteAccessTimes)

		if len(siteAccessTimes) > 0 {
			minSite = siteAccessTimes[0].Site
			minTime = siteAccessTimes[0].AccessTime
			maxSite = siteAccessTimes[len(siteAccessTimes)-1].Site
			maxTime = siteAccessTimes[len(siteAccessTimes)-1].AccessTime
		}

		// Сохраняем минимальное и максимальное время доступа и соответствующие им сайты в Redis
		minTimeStr = strconv.FormatFloat(minTime, 'f', -1, 64)
		err = client.Set(models.MinName, minSite, 0).Err()
		if err != nil {
			lg.Println("Error saving min site:", err)
		}
		err = client.Set(models.MinTime, minTimeStr, 0).Err()
		if err != nil {
			lg.Println("Error saving min site:", err)
		}

		maxTimeStr = strconv.FormatFloat(maxTime, 'f', -1, 64)
		err = client.Set(models.MaxName, maxSite, 0).Err()
		if err != nil {
			lg.Println("Error saving max site:", err)
		}
		err = client.Set(models.MaxTime, maxTimeStr, 0).Err()
		if err != nil {
			lg.Println("Error saving max site:", err)
		}
	}
}

func quickSort(arr []models.SiteAccessTime) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[len(arr)-1].AccessTime
	left := 0
	right := len(arr) - 2

	for left <= right {
		for left <= right && arr[left].AccessTime < pivot {
			left++
		}

		for left <= right && arr[right].AccessTime > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	arr[left], arr[len(arr)-1] = arr[len(arr)-1], arr[left]

	quickSort(arr[:left])
	quickSort(arr[left+1:])
}
