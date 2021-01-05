/*

Since the doorbell is on a mains circuit, a button press triggers more than 1
rising/falling edge on the GPIO pin.

Electrical interference from nearby appliances powering on/off can trigger
rising and falling edges on the GPIO pin. Generally, these are very brief and
therefore the logic used to filter them out is to ignore events that were not
followed by at least `debounceThreshold` events within `debounce` period.

*/
package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/stianeikeland/go-rpio"
)

var (
	pin               = rpio.Pin(24)
	debounce          = time.Millisecond * 300
	debounceThreshold = 3

	influxDBAddr = flag.String("influxdbaddr", "127.0.0.1:8086", "InfluxDB address")
	influxDBAuth = flag.String("influxdbauth", "user:token", "InfluxDB auth string")
	influxDBName = flag.String("influxdbname", "mydb", "InfluxDB database name")
)

func buttonPress(bounces int, elapsed time.Duration) {
	rejected := bounces < debounceThreshold

	// Write the event to InfluxDB, for monitoring in Grafana.
	client := influxdb2.NewClient(*influxDBAddr, *influxDBAuth)
	defer client.Close()
	influx := client.WriteAPIBlocking("", *influxDBName)
	p := influxdb2.NewPoint(
		"doorbell1",
		map[string]string{
			"rejected": strconv.FormatBool(rejected),
		},
		map[string]interface{}{
			"duration": elapsed.Milliseconds(),
			"bounces":  bounces,
		},
		time.Now())
	err := influx.WritePoint(context.Background(), p)
	if err != nil {
		fmt.Printf("Error writing to influxdb: %s\n", err.Error())
	}

	if rejected {
		fmt.Printf("Ignoring %o bounces in %s\n", bounces, elapsed)
		return
	}

	fmt.Printf("Accepting %o bounces in %s\n", bounces, elapsed)

	doorbell()
}

func doorbell() {
	file, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	dir := filepath.Dir(file)
	script := path.Join(dir, "notify.sh")
	fmt.Printf("Running %s\n", script)

	cmd := exec.Command("/bin/sh", "-c", script)
	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	flag.Parse()

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer rpio.Close()

	pin.Input()
	pin.PullDown()
	pin.Detect(rpio.AnyEdge)

	bounces := 0
	var bounceStart time.Time

	for {
		if pin.EdgeDetected() && bounces == 0 {
			bounceStart = time.Now()
			continue
		}

		if pin.EdgeDetected() {
			bounces++
		}

		if bounces == 0 {
			continue
		}

		elapsed := time.Since(bounceStart)

		if elapsed >= debounce {
			buttonPress(bounces, elapsed)
			bounces = 0
		}

		// time.Sleep(time.Millisecond * 20)
	}
}
