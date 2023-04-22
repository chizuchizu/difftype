package main

import (
  "fmt"
	"time"

	"github.com/pterm/pterm"
)


func main(){
  area, _ := pterm.DefaultArea.Start()
	// area, _ := pterm.DefaultArea.WithCenter().Start()

  for i := 5; i >= 0; i-- {
    base_str := fmt.Sprintf("Foo %d", i)
    str, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString(base_str)).Srender()
    str = pterm.DefaultCenter.Sprint(str)
    area.Update(str)
    time.Sleep(time.Second)
  }

  area.Stop()

}
