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
  // "github.com/pterm/pterm"
	"github.com/rivo/tview"
)

const start_str = `
╭━━━┳━━━━┳━━━┳━━━┳━━━━╮
┃╭━╮┃╭╮╭╮┃╭━╮┃╭━╮┃╭╮╭╮┃
┃╰━━╋╯┃┃╰┫┃╱┃┃╰━╯┣╯┃┃╰╯
╰━━╮┃╱┃┃╱┃╰━╯┃╭╮╭╯╱┃┃
┃╰━╯┃╱┃┃╱┃╭━╮┃┃┃╰╮╱┃┃
╰━━━╯╱╰╯╱╰╯╱╰┻╯╰━╯╱╰╯
`

var num = []string {
  `
  ╭━━━╮
  ┃╭━╮┃
  ┃┃┃┃┃
  ┃┃┃┃┃
  ┃╰━╯┃
  ╰━━━╯
  `,
  `
  ╱╭╮
  ╭╯┃
  ╰╮┃
  ╱┃┃
  ╭╯╰╮
  ╰━━╯
  `,
  `
  ╭━━━╮
  ┃╭━╮┃
  ╰╯╭╯┃
  ╭━╯╭╯
  ┃┃╰━╮
  ╰━━━╯
  `,
  `
  ╭━━━╮
  ┃╭━╮┃
  ╰╯╭╯┃
  ╭╮╰╮┃
  ┃╰━╯┃
  ╰━━━╯
  `,
}

var (
  app *tview.Application
  textView *tview.TextView
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

func run() {
  srcPath := "i.txt"
  dstPath := "t.txt"
  filePath := "target.txt"

  // fetcher
  chFetch := make(chan bool)
  // timer
  ticker := time.NewTicker(time.Second)
  fmt.Println("hoge")


  go fetch_challenge(0, srcPath, dstPath, chFetch)

  isFinish := false
  count := 3
  for {
    select {
      case _ = <- ticker.C:
        count --
        app.QueueUpdateDraw(func() {
          textView.Clear()
          fmt.Fprint(textView, num[count])
        })
      case isFinish = <- chFetch:
        // pterm.DefaultBasicText.Println("FETCH DONE!!!")
      default:
    }
    if isFinish && count <= 0 {
      app.QueueUpdateDraw(func() {
        textView.Clear()
        fmt.Fprint(textView, start_str)
      })
      break
    }
  }

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
  fmt.Println(elapsedTime)
}


func main(){
  app = tview.NewApplication()
  textView = tview.NewTextView().SetRegions(true)

  go run()
	// レイアウトを作成する
	flex := tview.NewFlex().
		AddItem(textView, 0, 1, true)

	// アプリケーションを開始する
	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}


}
