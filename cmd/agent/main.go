package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
	"strconv"
	"time"
)

var pollCount int
var randomValue int
var pollInterval int = 2
var reportInterval int = 10

type metrics struct {
	Alloc         uint64
	BuckHashSys   uint64
	Frees         uint64
	GCCPUFraction float64
	GCSys         uint64
	HeapAlloc     uint64
	HeapIdle      uint64
	HeapInuse     uint64
	HeapObjects   uint64
	HeapReleased  uint64
	HeapSys       uint64
	LastGC        uint64
	Lookups       uint64
	MCacheInuse   uint64
	MCacheSys     uint64
	MSpanInuse    uint64
	MSpanSys      uint64
	Mallocs       uint64
	NextGC        uint64
	NumForcedGC   uint32
	NumGC         uint32
	OtherSys      uint64
	PauseTotalNs  uint64
	StackInuse    uint64
	StackSys      uint64
	Sys           uint64
	TotalAlloc    uint64
	/* PollCount     int
	RandomValue   int */
}

func SendPost(metricType string, metric string, value string) {
	url := "http://localhost:8080/update/" + metricType + "/" + metric + "/" + value
	response, err := http.Post(url, "text/plain; charset=UTF-8", nil)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf(url+" Status Code: %d\r\n", response.StatusCode)
}

func memInfo(metricName string) {
	stats := runtime.MemStats{}
	runtime.ReadMemStats(&stats)
	//fmt.Printf("%+v\n", stats)

	//time.Sleep(time.Duration(reportInterval) * time.Second)

	/* fmt.Printf("Alloc: %+v\n", stats.Alloc)
	fmt.Printf("BuckHashSys: %+v\n", stats.BuckHashSys)
	fmt.Printf("Frees: %+v\n", stats.Frees)
	fmt.Printf("GCCPUFraction: %+v\n", stats.GCCPUFraction)
	fmt.Printf("GCSys: %+v\n", stats.GCSys)
	fmt.Printf("HeapAlloc: %+v\n", stats.HeapAlloc)
	fmt.Printf("HeapIdle: %+v\n", stats.HeapIdle)
	fmt.Printf("HeapInuse: %+v\n", stats.HeapInuse)
	fmt.Printf("HeapObjects: %+v\n", stats.HeapObjects)
	fmt.Printf("HeapReleased: %+v\n", stats.HeapReleased)
	fmt.Printf("HeapSys: %+v\n", stats.HeapSys)
	fmt.Printf("LastGC: %+v\n", stats.LastGC)
	fmt.Printf("Lookups: %+v\n", stats.Lookups)
	fmt.Printf("MCacheInuse: %+v\n", stats.MCacheInuse)
	fmt.Printf("MCacheSys: %+v\n", stats.MCacheSys)
	fmt.Printf("MSpanInuse: %+v\n", stats.MSpanInuse)
	fmt.Printf("MSpanSys: %+v\n", stats.MSpanSys)
	fmt.Printf("Mallocs: %+v\n", stats.Mallocs)
	fmt.Printf("NextGC: %+v\n", stats.NextGC)
	fmt.Printf("NumForcedGC: %+v\n", stats.NumForcedGC)
	fmt.Printf("NumGC: %+v\n", stats.NumGC)
	fmt.Printf("OtherSys: %+v\n", stats.OtherSys)
	fmt.Printf("PauseTotalNs: %+v\n", stats.PauseTotalNs)
	fmt.Printf("StackInuse: %+v\n", stats.StackInuse)
	fmt.Printf("StackSys: %+v\n", stats.StackSys)
	fmt.Printf("Sys: %+v\n", stats.Sys)
	fmt.Printf("TotalAlloc: %+v\n", stats.TotalAlloc)

	fmt.Printf("PollCount: %+v\n", pollCount)
	fmt.Printf("RandomValue: %+v\n", randomValue) */

	var metr metrics = metrics{
		Alloc:         stats.Alloc,
		BuckHashSys:   stats.BuckHashSys,
		Frees:         stats.Frees,
		GCCPUFraction: stats.GCCPUFraction,
		GCSys:         stats.GCSys,
		HeapAlloc:     stats.HeapAlloc,
		HeapIdle:      stats.HeapIdle,
		HeapInuse:     stats.HeapInuse,
		HeapObjects:   stats.HeapObjects,
		HeapReleased:  stats.HeapReleased,
		HeapSys:       stats.HeapSys,
		LastGC:        stats.LastGC,
		Lookups:       stats.Lookups,
		MCacheInuse:   stats.MCacheInuse,
		MCacheSys:     stats.MCacheSys,
		MSpanInuse:    stats.MSpanInuse,
		MSpanSys:      stats.MSpanSys,
		Mallocs:       stats.Mallocs,
		NextGC:        stats.NextGC,
		NumForcedGC:   stats.NumForcedGC,
		NumGC:         stats.NumGC,
		OtherSys:      stats.OtherSys,
		PauseTotalNs:  stats.PauseTotalNs,
		StackInuse:    stats.StackInuse,
		StackSys:      stats.StackSys,
		Sys:           stats.Sys,
		TotalAlloc:    stats.TotalAlloc,
	}

	/* metr.PollCount += 1
	metr.RandomValue = rand.Intn(10000) */

	values := reflect.ValueOf(metr)
	types := values.Type()

	for i := 0; i < values.NumField(); i++ {
		// types.Field(i).Index[0] 	- порядковый номер
		// types.Field(i).Name 		- наименование поля
		// values.Field(i) 			- значение поля
		//fmt.Println(types.Field(i).Index[0], types.Field(i).Name, values.Field(i))
		if types.Field(i).Name == "GCCPUFraction" {
			SendPost("gauge", types.Field(i).Name, strconv.FormatFloat(values.Field(i).Interface().(float64), 'f', 6, 64))
		} else if types.Field(i).Name == "NumForcedGC" || types.Field(i).Name == "NumGC" {
			SendPost("gauge", types.Field(i).Name, strconv.FormatUint(uint64(values.Field(i).Interface().(uint32)), 10))
		} else {
			SendPost("gauge", types.Field(i).Name, strconv.FormatUint(values.Field(i).Interface().(uint64), 10))
		}

	}

}

func main() {
	//response, err := http.Get("https://practicum.yandex.ru")

	/* if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Status Code: %d\r\n", response.StatusCode) */

	for {
		memInfo("Alloc")
		memInfo("BuckHashSys")
		time.Sleep(time.Duration(reportInterval) * time.Second)
	}

}
