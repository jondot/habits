package main

import(
  "fmt"
  "github.com/russross/blackfriday"
)

var renderer blackfriday.Renderer = NewTerminalRenderer(0)

func markdown(text string) string{
  return string(blackfriday.Markdown([]byte(text), renderer, 0))
}

type MarkdownReporter struct{
}

func (self *MarkdownReporter) Report(selected []*Habit, withIndex bool) {
  for idx,habit := range selected {
    if withIndex {
      fmt.Printf("%d. ", idx)
    }
    fmt.Printf("%s\n", markdown(habit.Content))
  }
}
