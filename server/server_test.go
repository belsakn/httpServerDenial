package server

import (
	"strconv"
	"testing"
	"time"
)

var want200 int = 200
var want503 int = 503

func TestPrepMapOfClients(t *testing.T) {
	clientMap = make(Clients)
	for i := 1; i < 5; i++ {
		clientMap[strconv.Itoa(i)] = &Client{startTime: time.Now(), requestCount: 1}
	}
}

func TestFirstReq(t *testing.T) {
	got := responseCodeForClient("1", time.Now())
	if want200 != got {
		t.Errorf("want: %d, got: %d", want200, got)
	}
}

func TestToMuchReq(t *testing.T) {
	clientMap["2"].requestCount = 6
	got := responseCodeForClient("2", time.Now().Local().Add(time.Second*time.Duration(1)))
	if want503 != got {
		t.Errorf("want: %d, got: %d", want503, got)
	}
}

func TestTimeHasPased(t *testing.T) {
	clientMap["3"].requestCount = 6
	got := responseCodeForClient("3", time.Now().Local().Add(time.Second*time.Duration(9)))
	if want200 != got {
		t.Errorf("want: %d, got: %d", want200, got)
	}
}
