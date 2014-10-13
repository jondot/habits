package main

import(
  "fmt"
)


type PlainReporter struct {
}

func (self *PlainReporter) Report(selected []*Habit, withIndex bool){
  for idx,habit := range selected {
    if withIndex {
      fmt.Printf("%d. ", idx)
    }
    fmt.Printf("%s\n", habit.Content)
  }
}
