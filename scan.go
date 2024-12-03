package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func (app *App) GetOwnersFromScanV2(owners map[string]any, url, addr string) {
	var perPairAddrs int
	for i := 1; ; i++ {
		addrs := rawReqV2(url, addr, i)
		for _, addr := range addrs {
			if _, ok := owners[strings.ToLower(addr)]; !ok {
				perPairAddrs++
				owners[strings.ToLower(addr)] = nil
			}
		}
		if len(addrs) < 50 {
			fmt.Printf("%s: %d addresses for pool %s\n", url[8:], perPairAddrs, addr)
			break
		}
		time.Sleep(35 * time.Millisecond)
	}
}

// works only with V2 pools
func rawReqV2(url, addr string, page int) []string {
	get, err := http.Get(url + "/token/generic-tokenholders2?a=" + addr + "&p=" + strconv.Itoa(page))
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := io.ReadAll(get.Body)
	if err != nil {
		log.Fatalln(err)
	}
	// parsing from data-clipboard-text
	r := regexp.MustCompile(`data-clipboard-text="(0x[a-fA-F0-9]{40})"`)
	result := r.FindAllStringSubmatch(string(resp), -1)
	var addrs = make([]string, 0)
	for _, v := range result {
		addrs = append(addrs, v[1])
	}
	return addrs
}
