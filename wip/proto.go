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

func fetch_challenge(idx int, srcPath string, dstPath string, ch chan bool) {
  url := fmt.Sprintf("https://5yy10qutb4.execute-api.us-east-1.amazonaws.com/typing_challenge?idx=%d", idx)
  res, err := http.Get(url)
  check(err)
  defer res.Body.Close()
  

  var cont Post
  body, _ := ioutil.ReadAll(res.Body)
  err = json.Unmarshal(body, &cont)
  check(err)

  file, err := os.Create(srcPath)
  check(err)
  _, err = file.WriteString(cont.Src)
  check(err)

  file, err = os.Create(dstPath)
  check(err)
  _, err = file.WriteString(cont.Dst)
  check(err)
  ch <- true
}



func main(){
  fmt.Println("start")

  srcPath := "i.txt"
  dstPath := "t.txt"
  filePath := "target.txt"

  // fetcher
  chFetch := make(chan bool)
  // timer
  ticker := time.NewTicker(time.Second)

  go fetch_challenge(0, srcPath, dstPath, chFetch)

  isFinish := false
  count := 4

  for {
    select {
      case _ = <- ticker.C:
        count --
        base_str := fmt.Sprintf("Foo %d", count)

        pterm.DefaultBasicText.Println(base_str)
      case isFinish = <- chFetch:
        pterm.DefaultBasicText.Println("FETCH DONE!!!")
      default:
    }
    if isFinish && count <= 0 {
      break
    }
  }

  pterm.DefaultBasicText.Println("START!")

  startTime := time.Now()

  // start
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


}
