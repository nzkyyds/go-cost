package cost

import (
	"fmt"
	"log"
	"time"
)

var (
	s     time.Time
	token string
	last  time.Time
)

func Begin(tk string) {
	s = time.Now()
	last = time.Now()
	token = tk
	Printf("cost Begin")
}

func Printf(format string, v ...interface{}) {
	log.Printf("%s duration=%-15v, cost=%-15v, %s", token, time.Since(s), time.Since(last), fmt.Sprintf(format, v...))
	last = time.Now()
}

func End() {
	Printf("cost End")
}
