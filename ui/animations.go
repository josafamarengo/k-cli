package ui

import (
    "fmt"
    "github.com/briandowns/spinner"
    "time"
)

func showSpinner(message string, task func()) {
    sp := spinner.New(spinner.CharSets[11], 100*time.Millisecond)
	sp.Color("bgBlack", "bold", "fgHiYellow")
	sp.Suffix = fmt.Sprintf(" %s", message)

  	sp.Start()

	task()
        
    sp.Stop()
}
