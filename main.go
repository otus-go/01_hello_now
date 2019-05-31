package main

import (
	"io"
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func PrintNow(out io.Writer) {
	res, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return
	}
	fmt.Fprintln(out, res.Format("2006-01-02 15:04:05"))
}

func main() {
	PrintNow(os.Stdout)
}
