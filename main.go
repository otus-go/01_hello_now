package main

import (
	"fmt"
	"io"
	"os"

	"github.com/beevik/ntp"
)

// for tests
var ntpTime = ntp.Time

func PrintNow(out io.Writer) {
	res, err := ntpTime("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return
	}
	fmt.Fprintln(out, res.Format("2006-01-02 15:04:05"))
}

func main() {
	PrintNow(os.Stdout)
}
