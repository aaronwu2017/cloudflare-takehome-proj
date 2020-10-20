package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"sort"
	"time"
)

var urlFlag = flag.String("url", "", "any web sites url")
var profileFlag = flag.Int("profile", 0, "number of requests(a positive integer)")

func main() {
	//for _, arg := range os.Args {
	//	fmt.Println(arg)
	//}
	flag.Parse()
	//fmt.Println(*urlFlag, *profileFlag)
	if len(os.Args) >= 2 && "--help" == os.Args[1] {
		usage()
	} else if *urlFlag == "" {
		fmt.Println("need url parameter")
		usage()
	} else if *profileFlag > 0 {
		profileFunc()
	} else if *profileFlag == 0 {
		urlFunc()
	} else if *profileFlag < 0 {
		fmt.Println("profile number must be a positive integer")
		usage()
	} else {
		usage()
	}

	//fmt.Println(*urlFlag, *helpFlag, *profileFlag)
}

func urlFunc() {
	//fmt.Println("urlFunc")
	err, profile := httpGet(*urlFlag)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fmt.Sprintf("Url : %s", profile.Url))
	fmt.Println(fmt.Sprintf("StatusCode : %d", profile.StatusCode))
	fmt.Println(fmt.Sprintf("Duratime : %dns", profile.Duratime))
	fmt.Println(fmt.Sprintf("ResponseSize : %d", profile.RespSize))
	fmt.Println(fmt.Sprintf("Response : %s", string(profile.Resp)))
}

//The number of requests
//The fastest time
//The slowest time
//The mean & median times
//The percentage of requests that succeeded
//Any error codes returned that weren't a success
//The size in bytes of the smallest response
//The size in bytes of the largest response

type TotalProfile struct {
	RequestNum     int
	FastestTime    int64
	SlowestTime    int64
	MeanTime       float64
	MedianTime     float64
	PercentSuccess float64
	SuccessRequest int
	ErrorRequest   int
	SmallestRes    int
	LargestRes     int
}

func profileFunc() {
	//fmt.Println("profileFunc")
	tp := TotalProfile{}
	start := time.Now()
	errCode := make([]int, 0)
	errUrl := make([]string, 0)
	var totalTime int64 = 0
	durationTimeArray := make([]float64, 0)
	for i := 0; i < *profileFlag; i++ {
		//fmt.Println("profileFunc", i+1)
		err, profile := httpGet(*urlFlag)
		tp.RequestNum++
		if err != nil {
			tp.ErrorRequest++
		} else if profile.StatusCode != http.StatusOK {
			tp.ErrorRequest++
			errCode = append(errCode, profile.StatusCode)
			errUrl = append(errUrl, *urlFlag)
		} else {
			tp.SuccessRequest++

			totalTime = totalTime + profile.Duratime
			durationTimeArray = append(durationTimeArray, float64(profile.Duratime))
			if profile.Duratime > tp.SlowestTime {
				tp.SlowestTime = profile.Duratime
			}

			if tp.FastestTime == 0 || profile.Duratime < tp.FastestTime {
				tp.FastestTime = profile.Duratime
			}

			if profile.RespSize > tp.LargestRes {
				tp.LargestRes = profile.RespSize
			}

			if tp.SmallestRes == 0 || profile.RespSize < tp.SmallestRes {
				tp.SmallestRes = profile.RespSize
			}
		}

	}
	tp.PercentSuccess = math.Round(float64(tp.SuccessRequest) * 100 / float64(tp.RequestNum))
	//tp.PercentSuccess = math.Round(float64(1)*100/float64(9))
	fmt.Println(tp.PercentSuccess)
	tp.MeanTime = math.Round(float64(totalTime) / float64(tp.SuccessRequest))
	//errcode
	//errCode := ""
	//for _, errCode := range errCodeProfile {
	//
	//}
	sort.Float64s(durationTimeArray) // sort the numbers

	mNumber := len(durationTimeArray) / 2
	if len(durationTimeArray)%2 == 0 {
		tp.MedianTime = (durationTimeArray[mNumber-1] + durationTimeArray[mNumber]) / 2
	} else {
		tp.MedianTime = durationTimeArray[mNumber]
	}

	duration := time.Since(start).Nanoseconds()
	fmt.Println(fmt.Sprintf("Url : %s", *urlFlag))
	fmt.Println(("================"))
	fmt.Println(fmt.Sprintf("TotalRequestNum : %d", tp.RequestNum))
	fmt.Println(fmt.Sprintf("FastestTime : %dns", tp.FastestTime))
	fmt.Println(fmt.Sprintf("SlowestTime : %dns", tp.SlowestTime))
	fmt.Println(fmt.Sprintf("MeanTime : %fns", tp.MeanTime))
	fmt.Println(fmt.Sprintf("MedianTime : %fns", tp.MedianTime))
	fmt.Println(fmt.Sprintf("SuccessRequest : %d", tp.SuccessRequest))
	fmt.Println(fmt.Sprintf("ErrorRequest : %d", tp.ErrorRequest))
	//fmt.Println(fmt.Sprintf("Percentage of requests : %f%%", tp.PercentSuccess))
	fmt.Println(fmt.Sprintf("Percentage of requests : %.0f%%", tp.PercentSuccess))
	fmt.Println(fmt.Sprintf("ErrorCode : %v", errCode))
	fmt.Println(fmt.Sprintf("ErrorUrl : %v", errUrl))
	fmt.Println(fmt.Sprintf("SmallestRes(bytes) : %d", tp.SmallestRes))
	fmt.Println(fmt.Sprintf("LargestRes(bytes) : %d", tp.LargestRes))
	fmt.Println(("================"))
	fmt.Println(fmt.Sprintf("Execute Duration : %dns", duration))
}

func usage() {
	//fmt.Println("usage")
	flag.Usage()
}

type Profile struct {
	Url        string
	StatusCode int
	Duratime   int64
	RespSize   int
	Resp       []byte
}

func httpGet(url string) (error, Profile) {
	start := time.Now()
	client := &http.Client{}
	client.Timeout = time.Second
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		// handle error
		return err, Profile{}
	}

	//req.Header.Set("Content-Type", "application/text")

	resp, err := client.Do(req)
	if err != nil {
		// handle error
		return err, Profile{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		return err, Profile{}
	}
	return nil, Profile{url, resp.StatusCode, time.Since(start).Nanoseconds(), len(body), body}
}
