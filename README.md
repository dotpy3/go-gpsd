# go-gpsd

*GPSD client for Go.*

## Installation

This version of `go-gpsd` requires Go 1.7.

```bash
go get github.com/dotpy3/go-gpsd
```

## Usage

go-gpsd is a streaming client for GPSD's JSON service and as such can be used only in async manner unlike clients for other languages which support both async and sync modes.

This fork of the original `go-gpsd` adds support for the `context` package, and programmatic error handling.

```go
import "github.com/dotpy3/go-gpsd"

func main() {
	gps := gpsd.Dial("localhost:2947")
}
```

After `Dial`ing the server, you should install stream filters. Stream filters allow you to capture only certain types of GPSD reports.

```go
gps.AddFilter("TPV", tpvFilter)
```

Filter functions have a type of `gps.Filter` and should receive one argument of type `interface{}`.

```go
tpvFilter := func(r interface{}) {
	report := r.(*gpsd.TPVReport)
	fmt.Println("Location updated", report.Lat, report.Lon)
}
```

Due to the nature of GPSD reports your filter will manually have to cast the type of the argument it received to a proper `*gpsd.Report` struct pointer.

After installing all needed filters, call the `Watch` method to start observing reports. Please note that at this time installed filters can't be removed.

```go
ctx, cancel := context.WithCancel()
gps.Watch(ctx)
<-time.After(30 * time.Second)
cancel()
```

`Watch()` will span a new goroutine in which all data processing will happen. In this case, this goroutine will be stopped after 30 seconds.

To handle errors, call the `OnError` method, with a function taking an `error` as an argument.

```go
gps.OnError(func (err error) {
	fmt.Println("Error:", err)
})
```

### Currently supported GPSD report types

* `VERSION` (`gpsd.VERSIONReport`)
* `TPV` (`gpsd.TPVReport`)
* `SKY` (`gpsd.SKYReport`)
* `ATT` (`gpsd.ATTReport`)
* `GST` (`gpsd.GSTReport`)
* `PPS` (`gpsd.PPSReport`)
* `Devices` (`gpsd.DEVICESReport`)
* `DEVICE` (`gpsd.DEVICEReport`)
* `ERROR` (`gpsd.ERRORReport`)

## Documentation

For complete library docs, visit [GoDoc.org](http://godoc.org/github.com/dotpy3/go-gpsd) or take a look at the `gpsd.go` file in this repository.

GPSD's documentation on their JSON protocol can be found at [http://catb.org/gpsd/gpsd_json.html](http://catb.org/gpsd/gpsd_json.html)

