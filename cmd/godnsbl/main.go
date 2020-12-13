package main

import (
	"encoding/json"
	"fmt"

	"sync"

	"github.com/alexanderkarlis/godnsbl"
)

func main() {

	wg := &sync.WaitGroup{}
	results := []godnsbl.Result{}
	for i := 1; i <= 255; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			rbl := godnsbl.Lookup("zen.spamhaus.org", fmt.Sprintf("127.0.0.%d", i))
			if len(rbl.Results) == 0 {
				results = append(results, godnsbl.Result{})
			} else {
				if rbl.Results[0].Code != "" {
					results = append(results, rbl.Results[0])
				}
			}
		}(i)
	}

	wg.Wait()

	pp, err := json.MarshalIndent(&results, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(pp))
}
