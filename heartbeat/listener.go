package heartbeat

import (
	"fmt"
	"github.com/skyhackvip/dispatcher/rabbitmq"
	"strconv"
	"sync"
	"time"
)

//global variable
var addrList = make(map[string]time.Time)
var mutex sync.Mutex

//listen and revice heart beat
func ListenHeartbeat() {
	fmt.Println("start")
	queue := rabbitmq.New()
	defer queue.Close()
	//queue.BindExchange("heart-beat-exchange")

	//go clearExpiredAddr()

	fmt.Println("ea")
	msgList := queue.Receive("heart-beat-queue")
	fmt.Printf("%T\n", msgList)
	for msg := range msgList {
		fmt.Println("for")
		addr, err := strconv.Unquote(string(msg.Body))
		fmt.Println(addr, err)
		if err != nil {
			panic(err)
		}
		mutex.Lock()
		addrList[addr] = time.Now()
		mutex.Unlock()
	}
}

//clear expired addr,10 seconds no reply
func clearExpiredAddr() {
	for {
		time.Sleep(1 * time.Second)
		mutex.Lock()
		for addr, lastTime := range addrList {
			if lastTime.Add(10 * time.Second).Before(time.Now()) { //10 shoud be config
				delete(addrList, addr)
			}
		}
		mutex.Unlock()
	}
}

//get all live addr list
func GetAddrList() []string {
	mutex.Lock()
	defer mutex.Unlock()
	rs := make([]string, 0)
	for addr, _ := range addrList {
		rs = append(rs, addr)
	}
	return rs
}
