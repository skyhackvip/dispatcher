package healthcheck

import (
	"log"
	"net/http"
	"sync"
	"time"
)

var addrList = make([]string, 0)
var failCount = make(map[string]int)
var mutex sync.Mutex
var maxRetrytimes int = 3

//get request keep alive
func pingCheck(addr string) bool {
	resp, err := http.Get(addr)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return false
	}
	return true
}

//loop addrlist and check alive
func loopCheck() {
	for idx, addr := range addrList {
		if pingCheck(addr) {
			log.Println(addr, "ok")
			continue
		}
		log.Println(addr, "fail")
		mutex.Lock()
		failCount[addr]++
		if failCount[addr] > maxRetrytimes {
			addrList = append(addrList[:idx], addrList[idx+1:]...)
			log.Println(addr, "removed")
		}
		mutex.Unlock()
	}
}

//health check
func HealthCheck() {
	ticker := time.NewTicker(time.Second * 5) //every 5 second send a ping request
	for {
		select {
		case <-ticker.C:
			log.Println("start health check...")
			loopCheck()
			log.Println("health check end")
		}
	}
}

//add/init addrlist
func AddAddr(addr ...string) {
	mutex.Lock()
	defer mutex.Unlock()
	addrList = append(addrList, addr...)
}

//get alive addrlist
func GetAliveAddrList() []string {
	mutex.Lock()
	defer mutex.Unlock()
	return addrList
}
