package main

import(
  "encoding/json"
  "io/ioutil"
  "os"
  "fmt"
  "sort"
  "time"
)

type Reporter interface{
  Report([]*Habit, bool)
}


type Habits struct {
  HabitList []*Habit `json:"habits"`
  RepeatFactor int
}


type Habit struct{
  Content string `json:"content"`
  LastShow time.Time `json:"last_show"`
  Count int `json:"count"`
}

func NewHabitsDemo() *Habits{
  return &Habits{
    HabitList: []*Habit{
      &Habit{ Content: "To add a habit, use 'habits -add your-new-habit'", LastShow: time.Now(), Count: 0 },
      &Habit{ Content: "-> Reminders on zsh - wire habits to your ~/.zshrc", LastShow: time.Now(), Count: 0 },
      &Habit{ Content: "Take a break!", LastShow: time.Now(), Count: 0 },
    },
  }
}

func NewHabitsFromFile(file string) *Habits{
  text, err := ioutil.ReadFile(file)
  if err != nil {
    fmt.Errorf("Cannot read file: %s", file)
    os.Exit(1)
  }
  data := &Habits{}
  err = json.Unmarshal(text, &data)
  if err != nil {
    fmt.Errorf("cannot unmarshal json", err)
    os.Exit(1)
  }
  return data
}

type ByDate []*Habit

func (a ByDate) Len() int           { return len(a) }
func (a ByDate) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByDate) Less(i, j int) bool { return a[i].LastShow.Before(a[j].LastShow)}
func (self *Habits) SelectNext(num int) []*Habit{
  if self.RepeatFactor <= 0 {
    self.RepeatFactor = 1
  }

  sort.Sort(ByDate(self.HabitList))
  selected := []*Habit{}
  if len(self.HabitList) < num{
    selected = self.HabitList
  } else {
    selected = self.HabitList[:num]
  }
  for _,habit := range selected {
    habit.Count = habit.Count + 1
    if habit.Count >= self.RepeatFactor {
      habit.LastShow = time.Now()
      habit.Count = 0
    }
  }
  return selected
}

func (self *Habits) StoreToFile(path string) {
  b, err := json.Marshal(self)
  if err != nil {
    fmt.Errorf("Cannot serialize data: %v", err)
    os.Exit(1)
  }

  err = ioutil.WriteFile(path, b, 0666)
  if err != nil {
    fmt.Errorf("Cannot write file: %s\nReason: %v", path, err)
    os.Exit(1)
  }
}

func (self *Habits) Remove(idx int){
  if idx > len(self.HabitList)-1{
    return
  }

  self.HabitList = append(self.HabitList[:idx], self.HabitList[idx+1:]...)
}

func (self *Habits) Add(content string){
  h := &Habit{ Content: content, LastShow: time.Now(), Count: 0 }
  self.HabitList = append(self.HabitList, h)
}


