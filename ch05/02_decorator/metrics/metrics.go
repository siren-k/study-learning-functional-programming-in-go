package metrics

import (
	"github.com/admobi/easy-metrics"
	. "github.com/siren-k/study-learning-functional-programming-in-go/ch05/02_decorator/decorator"
	"math/rand"
	"net/http"
	"time"
)

type timing struct {
	count int64
	sum   time.Duration
}

func (t *timing) Observe(d time.Duration) {
	t.count++
	t.sum += d
	avgResponseTime.Set(t.sum.Seconds() / float64(t.count))
}

var (
	avgResponseTime = metrics.NewGauge("avgResponseTime")
	requests        = metrics.NewCounter("requests")
	responseTime    = &timing{}
)

func work() {
	randInt := rand.Intn(5000)
	Debug.Printf("- randInt: %v", randInt)

	workTime := time.Duration(randInt) * time.Millisecond
	time.Sleep(workTime)
}

func handler(w http.ResponseWriter, r *http.Request) {
	begin := time.Now()

	work()

	responseTime.Observe(time.Since(begin))
	requests.Inc()
}

func Serve(addr string) error {
	r, err := metrics.NewTrackRegistry("Stats", 100, time.Second, false)
	if err != nil {
		Error.Println(err)
	}

	err = r.AddMetrics(requests, avgResponseTime)
	if err != nil {
		Error.Println(err)
	}

	http.HandleFunc("/", handler)
	return http.ListenAndServe(addr, nil)
}

func DisplayResults(addr string) error {
	Info.Printf("Go to http://%s/easy-metrics?show=Stats", addr)
	return nil
}
