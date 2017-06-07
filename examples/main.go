package main

import (
	"context"
	"fmt"
	"time"

	gpsd "github.com/dotpy3/go-gpsd"
)

func main() {
	var gps *gpsd.Session
	var err error

	if gps, err = gpsd.Dial(gpsd.DefaultAddress); err != nil {
		panic(fmt.Sprintf("Failed to connect to GPSD: %s", err))
	}

	gps.AddFilter("TPV", func(r interface{}) {
		tpv := r.(*gpsd.TPVReport)
		fmt.Println("TPV", tpv.Mode, tpv.Time)
	})

	skyfilter := func(r interface{}) {
		sky := r.(*gpsd.SKYReport)

		fmt.Println("SKY", len(sky.Satellites), "satellites")
	}

	gps.AddFilter("SKY", skyfilter)

	gps.OnError(func(err error) {
		fmt.Println("gpsd error:", err)
	})

	ctx, cancel := context.WithCancel(context.Background())
	gps.Watch(ctx)
	<-time.After(10 * time.Minute)
	fmt.Println("Shutting gpsd down...")
	cancel()
}
