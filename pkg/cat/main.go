package main

import (
	"fmt"
	"github.com/cat-go/cat"
	"math/rand"
	"time"
)

func init() {
	cat.DebugOn()
	cat.Init(&cat.Options{
		AppId:      "demo-service",
		Port:       2280,
		HttpPort:   8180,
		ServerAddr: "192.168.1.1",
	})
}
const TestType = "foo"

func main() {
	for true {
		fmt.Println(time.Now())
		t := cat.NewTransaction(TestType, "test")
		if rand.Int31n(100) == 0 {
			t.SetStatus(cat.FAIL)
		}
		t.AddData("foo", "bar")
		t.NewEvent(TestType, "event-1")
		t.Complete()
		if rand.Int31n(100) == 0 {
			t.LogEvent(TestType, "event-2", cat.FAIL)
		} else {
			t.LogEvent(TestType, "event-2")
		}
		t.LogEvent(TestType, "event-3", cat.SUCCESS, "k=v")
		t.SetDurationStart(time.Now().Add(-5 * time.Second))
		t.SetTime(time.Now().Add(-5 * time.Second))
		t.SetDuration(time.Millisecond * 500)
		t.Complete()
	}

}
