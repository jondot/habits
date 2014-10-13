package main

import(
  "flag"
  "github.com/mitchellh/go-homedir"
  "os"
  "fmt"
  "path/filepath"
)

var repeatFlag = flag.Int("repeat", 5, "Repeat factor for each item")
var showFlag = flag.Int("show", 1, "How many items to show each time")
var listFlag = flag.Bool("list", false, "List all")
var addFlag = flag.String("add", "", "Add a new habit")
var killFlag = flag.Int("kill", -1, "Remove a habit by index (use list)")
var reporterFlag = flag.String("reporter", "plain", "A habit reporter. Use 'plain' or 'markdown'")


func bootstrapIfNeeded(path string){
  _, err := os.Stat(path)
  if os.IsNotExist(err) {
    println("Bootstrapping to:")
    println(path)
    os.MkdirAll(filepath.Dir(path), 0755)
    habits := NewHabitsDemo()
    habits.StoreToFile(path)
  }
}


func main(){
  flag.Parse()

  dir, err := homedir.Dir()
  if err != nil{
    fmt.Errorf("Cannot open home directory: %s.\nReason: %v", dir, err)
    os.Exit(1)
  }

  var reporter Reporter = nil
  switch(*reporterFlag){
    case "markdown":
      reporter = &MarkdownReporter{}
    default:
      reporter = &PlainReporter{}
  }

  datafile := fmt.Sprintf("%s/.habits/local",dir)
  bootstrapIfNeeded(datafile)
  habits := NewHabitsFromFile(datafile)

  if *listFlag {
    reporter.Report(habits.HabitList, true)
    return
  }

  if *addFlag != ""{
    habits.Add(*addFlag)
    habits.StoreToFile(datafile)
    println(fmt.Sprintf("Added. You have a total of %d habits.", len(habits.HabitList)))
    return
  }

  if *killFlag != -1{
    habits.Remove(*killFlag)
    habits.StoreToFile(datafile)
    return
  }

  habits.RepeatFactor = *repeatFlag
  selected := habits.SelectNext(*showFlag)
  habits.StoreToFile(datafile)
  reporter.Report(selected,false)
}

