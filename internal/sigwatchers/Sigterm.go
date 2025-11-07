package sigwatchers

import (
	"os"
	"os/signal"

	"github.com/Moritisimor/EpsilonFetch/pkg/color"
)

func StartSigTermWatcher()  {
	sigTermWatcher := make(chan os.Signal, 1)
	signal.Notify(sigTermWatcher, os.Interrupt) 
	go func() {
		for range(sigTermWatcher) {
			color.PrintRedln("Please don't use Ctrl + C for quitting! Use the 'q' or 'q!' command instead.")
		} 
	}()
}