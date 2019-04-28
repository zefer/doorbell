/*

Since the doorbell is on a mains circuit, a button press triggers more than 1
rising/falling edge on the GPIO pin.

Electrical interference from nearby appliances powering on/off can trigger
rising and falling edges on the GPIO pin. Generally, these are very brief and
therefore the logic used to filter them out is to ignore events that were not
followed by at least `debounceThreshold` events within `debounce` period.

*/
package main

import(
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"time"
	"github.com/stianeikeland/go-rpio"
)

var (
	pin = rpio.Pin(24)
	debounce = time.Millisecond * 300
	debounceThreshold = 3
)

func buttonPress(bounces int, elapsed time.Duration) {
	if bounces < debounceThreshold {
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
