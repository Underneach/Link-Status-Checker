package src

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func SetupNetClient() {
	client.SetTimeout(15 * time.Second)
	client.SetHeaders(map[string]string{
		"User-Agent":                "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:127.0) Gecko/20100101 Firefox/127.0",
		"Accept":                    "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
		"Accept-Language":           "en-US;q=0.5,en;q=0.3",
		"Accept-Encoding":           "gzip, deflate, br, zstd",
		"Referer":                   "https://www.google.com/",
		"DNT":                       "1",
		"Connection":                "keep-alive",
		"Upgrade-Insecure-Requests": "1",
		"Sec-Fetch-Dest":            "document",
		"Sec-Fetch-Mode":            "navigate",
		"Sec-Fetch-Site":            "cross-site",
		"Sec-Fetch-User":            "?1",
		"Priority":                  "u=1",
	})
}

func RunMultiThread() {
	startTime = time.Now()
	for _, link := range linksList {
		wg.Add(1)
		workPool.Invoke(link)
	}
	wg.Wait()
}

func Worker(link string) {
	resp, err = client.R().Get(link)
	switch {
	case err != nil:
		fmt.Printf("[-] [%d/%d] [%s] - %s\n", checkedLinks, totalLinks, link, err)
	case resp.IsSuccess():
		fmt.Printf("[+] [%d/%d] [%d] - %s\n", checkedLinks, totalLinks, resp.StatusCode(), link)
		resultList = append(resultList, link)
		validLinks++
	case resp.IsError():
		fmt.Printf("[-] [%d/%d] [%d] - %s\n", checkedLinks, totalLinks, resp.StatusCode(), link)
		invalidLinks++
	default:
		fmt.Printf("[*] [%d/%d] [%d] - %s\n", checkedLinks, totalLinks, resp.StatusCode(), link)
	}
	checkedLinks++
	wg.Done()
}

func Save() {
	fmt.Println("\n\nSaving results...")
	ctime := time.Now().String()
	resFile, fileErr := os.Open("result_" + ctime + ".txt")
	if fileErr != nil {
		fmt.Println(err)
		return
	}

	write := bufio.NewWriter(resFile)
	write.WriteString(strings.Join(resultList, "\n"))
	write.Flush()
	fmt.Println("Results saved to: result_" + ctime + ".txt")
	workTime := int(startTime.Sub(time.Now()) / time.Minute)
	fmt.Printf("\nTime taken: %s\n", time.Duration(workTime))
}
