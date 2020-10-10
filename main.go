package main

import (
	"fmt"

	"github.com/someshkoli/httpreq/pkg/http"
	"github.com/spf13/pflag"
)

type reqPair struct {
	request  http.Request
	response *http.Response
}

func run(url string, count int64, full bool) {
	var fastest, slowest, timesum int64
	fastest = -9223372036854775808
	slowest = 9223372036854775807
	timesum = 0
	smallest := 4294967295
	biggest := 0
	var failed []reqPair
	var success []reqPair
	var total []reqPair
	if count == 1 {
		req, err := http.MakeRequest(url)
		if err != nil {
			fmt.Println(err)
		}
		res := req.Call()
		if res.Err != nil {
			fmt.Println(res.Err)
			fmt.Println("Use --help for more options")
			return
		}
		if full {
			fmt.Println(res.Data)
			return
		}
		fmt.Println(res.Body)
		return
	}
	for i := 0; i < int(count); i++ {
		req, err := http.MakeRequest(url)
		if err != nil {
			fmt.Println(err)
		}
		res := req.Call()
		if res.Err != nil {
			fmt.Println(res.Err)
			fmt.Println("Use --help for more options")
			return
		}
		pair := reqPair{request: req, response: res}
		total = append(total, pair)
		if res.Status != 200 {
			failed = append(failed, pair)
		} else {
			success = append(success, pair)
		}
		if res.Time.Milliseconds() < slowest {
			slowest = res.Time.Milliseconds()
		}
		if res.Time.Milliseconds() > fastest {
			fastest = res.Time.Milliseconds()
		}
		if len(res.Body) < smallest {
			smallest = len(pair.response.Body)
		}
		if len(res.Body) > biggest {
			biggest = len(pair.response.Body)
		}
		timesum += res.Time.Milliseconds()
	}
	mean := float64(timesum) / float64(count)
	percentage := float64(len(success)) / float64(count)
	median := total[(count+1)/2].response.Time
	fmt.Printf("Mean of all response time: %f ms\n", mean)
	fmt.Printf("Median of all response time: %d ms\n", median)
	fmt.Printf("Percentage of successfull response: %d%%\n", int(percentage*100))
	fmt.Printf("Request with fastest response: %d \n", fastest)
	fmt.Printf("Request with slowest response: %d \n", slowest)
	fmt.Printf("Request with biggest size: %d \n", biggest)
	fmt.Printf("Request with smallest size: %d \n", smallest)
	if len(failed) != 0 {
		fmt.Printf("Failed Responses: ")
		for _, f := range failed {
			fmt.Printf("%d ", f.response.Status)
		}
		fmt.Print("\n")
	}
}

func main() {
	var url string
	var full bool
	var profile int64
	pflag.StringVar(&url, "url", "https://red-tree-f56a.someshkoli.workers.dev/links", "URL to make HTTP request to.")
	pflag.Int64Var(&profile, "profile", 1, "Number of requests to make.")
	pflag.BoolVar(&full, "full", false, "If mentioned prints full response.")
	pflag.Parse()
	run(url, profile, full)
}
