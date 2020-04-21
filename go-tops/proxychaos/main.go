package proxychaos

import (
	"flag"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/elazarl/goproxy"
)

func main() {
	StartProxyForChaosCtl()
}

func StartProxyForChaosCtl() {
	proxyAddr := flag.String("proxychaos.addrss", ":8080", "proxychaos address")
	lr := flag.Float64("latency.rate", .3, "latency rate")
	lc := flag.Int64("latency.sec", 3, "latency in seconds")
	br := flag.Float64("block.rate", 0.3, "block rate")
	flag.Parse()
	latencyRate := *lr
	latencySec := *lc
	blockRate := *br

	if latencyRate+blockRate > 1.0 || latencySec < 0 {
		flag.Usage()
		os.Exit(1)
	}
	normalRate := 1.0 - latencyRate - blockRate
	handler := NewChaosHandler(normalRate, blockRate, latencyRate, latencySec)

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = true
	proxy.OnRequest().Do(handler)

	log.Printf("proxychaos listens %s", *proxyAddr)
	log.Fatal(http.ListenAndServe(*proxyAddr, proxy))
}

type chaosHandler struct {
	normalRate  float64
	blockRate   float64
	latencyRate float64
	latencySec  int64
}

func (h *chaosHandler) Handle(r *http.Request, c *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	normalChan := make(chan bool, 1)
	blockChan := make(chan bool, 1)

	rv := rand.Float64()

	if rv < h.normalRate {
		normalChan <- true
	} else if rv < h.normalRate+h.blockRate {
		blockChan <- true
	}

	select {
	case <-normalChan:
		return r, nil
	case <-blockChan:
		log.Printf("block request %s", r.RequestURI)
		return nil, nil
	case <-time.After(time.Duration(h.latencySec) * time.Second):
		log.Printf("delayed request %s", r.RequestURI)
		return r, nil
	}
}

// NewChaosHandler creates a chaos handler that forwards, block or add latency to http request
func NewChaosHandler(normalRate, blockRate, latencyRate float64, latencySec int64) goproxy.ReqHandler {
	return &chaosHandler{
		normalRate:  normalRate,
		blockRate:   blockRate,
		latencyRate: latencyRate,
		latencySec:  latencySec,
	}
}
