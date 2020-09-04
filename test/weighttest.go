package main

import (
	"fmt"
	"github.com/skyhackvip/dispatcher/loadbalance"
)

func main() {
	weightLb := loadbalance.LoadBalanceFactory(loadbalance.Weight)
	weightLb.Add("a", "1")
	weightLb.Add("b", "2")
	weightLb.Add("c", "5")

	var count = make(map[string]int)
	for i := 0; i < 200000; i++ {
		weightRs, _ := weightLb.Get()
		count[weightRs]++
	}
	fmt.Println(count)

}
