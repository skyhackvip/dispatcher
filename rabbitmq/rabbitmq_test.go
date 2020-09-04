package rabbitmq

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	mq := New()
	defer mq.Close()

	msgBody := "test"
	mq.Send("heart-beat-exchange", "heart-beat-queue", msgBody)

	receviceMsg := <-mq.Receive("heart-beat-queue")
	var receiveMsgBody interface{}
	err := json.Unmarshal(receviceMsg.Body, &receiveMsgBody)
	if err != nil {
		panic(err)
	}
	if receiveMsgBody != msgBody {
		t.Errorf("receive %s,send %s", receiveMsgBody, msgBody)
	}
	fmt.Printf("receive %s,send %s\n", receiveMsgBody, msgBody)
}
