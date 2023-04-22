package main

import (
  "fmt"
  "time"
  // "bufio"
  "io/ioutil"
  "os"
  "os/exec"
  "github.com/pterm/pterm"
)


func check(e error) {
  if e != nil {
    panic(e)
  }
}


func check_same_file(pathA string, pathB string) (bool) {
  contA, errA := ioutil.ReadFile(pathA)
  contB, errB := ioutil.ReadFile(pathB)

  if errA != nil || errB != nil {
    return false;
  }

  if string(contA) == string(contB) {
    return true;
  } else { 
    return false;
  }
}


func main(){
  area, _ := pterm.DefaultArea.Start()
  fmt.Println("start")

  // Countdown
  for i := 3; i >= 0; i-- {
    base_str := fmt.Sprintf("Foo %d", i)

    str, _ := pterm.DefaultBigText.WithLetters(pterm.NewLettersFromString(base_str)).Srender()
    str = pterm.DefaultCenter.Sprint(str)
    area.Update(str)

    time.Sleep(time.Second)
  }

  startTime := time.Now()

  srcPath := "i.txt"
  dstPath := "t.txt"
  filePath := "target.txt"

  cmd := exec.Command("cp", srcPath, filePath)
  _, err := cmd.Output()
  check(err)

  f, err := os.Open(filePath)
  check(err)
  defer f.Close()

  for {
    isSame := check_same_file(filePath, dstPath)
    if isSame {
      break
    } else {
      time.Sleep(100 * time.Millisecond)
    }
  }

  endTime := time.Now()
  elapsedTime := endTime.Sub(startTime)

  pterm.DefaultBasicText.Println("DONE!!!")
  showStr := fmt.Sprintf("\n Time: %s", elapsedTime)
  
  pterm.DefaultBasicText.Println(showStr)
  time.Sleep(1 * time.Millisecond)


  area.Stop()

}
