package main

import (
  "fmt"
  "time"
  // "bufio"
  "io/ioutil"
  "net/http"
  "os"
  "os/exec"
  "encoding/json"
  "github.com/pterm/pterm"
)


func check(e error) {
  if e != nil {
    panic(e)
  }
}

type Post struct {
  Src string `json:"src"`
  Dst string `json:"dst"`
}

func (me *Post) String() string {
  return fmt.Sprintf("\n===Src===\n%s \n ===Dst=== \n%s", me.Src, me.Dst)
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

func fetch_challenge(idx int, srcPath string, dstPath string) {
  url := fmt.Sprintf("http://127.0.0.1:8000/typing_challenge?idx=%d", idx)
  res, err := http.Get(url)
  check(err)
  defer res.Body.Close()
  

  var cont Post
  body, _ := ioutil.ReadAll(res.Body)
  err = json.Unmarshal(body, &cont)
  check(err)
  pterm.DefaultBasicText.Println(cont.Src)
  pterm.DefaultBasicText.Println(cont.Dst)

  file, err := os.Create(srcPath)
  check(err)
  _, err = file.WriteString(cont.Src)
  check(err)

  file, err = os.Create(dstPath)
  check(err)
  _, err = file.WriteString(cont.Dst)
  check(err)
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

  chFetch := make(chan bool)
    select {
      case isFinish := <- chFetch:
          pterm.DefaultBasicText.Println("FETCH DONE!!!")
      default:
    }


  startTime := time.Now()

  srcPath := "i.txt"
  dstPath := "t.txt"
  filePath := "target.txt"

  cmd := exec.Command("cp", srcPath, filePath)
  _, err := cmd.Output()
  check(err)

  // hogehogehoge
  fetch_challenge(0, srcPath, dstPath)

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
