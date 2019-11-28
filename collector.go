package main

import (
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/prometheus/client_golang/prometheus"
)


var facebookStatus = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "facebook_status", Help: "Status code returned from Facebook status API"})

type FacebookStatus struct {
  Current FacebookCurrentStatus 
}

type FacebookCurrentStatus struct {
  Health int
  Subject string
}

func checkFacebook() {
        go func() {
                for {
                		res, err := http.Get("https://www.facebook.com/platform/api-status/")
						if err != nil {
						    panic(err.Error())
						}

						body, err := ioutil.ReadAll(res.Body)
						if err != nil {
						    panic(err.Error())
						}

						var status FacebookStatus
						json.Unmarshal([]byte(body), &status)
                        facebookStatus.Set(float64(status.Current.Health))

                        time.Sleep(60 * time.Second)
                }
        }()
}

func init() {
	//Register metrics with prometheus
	prometheus.MustRegister(facebookStatus)

	checkFacebook()
}