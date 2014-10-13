package main
import (
  . "gopkg.in/check.v1"
  "time"
)



type HabitsSuite struct{}
var _ = Suite(&HabitsSuite{})

func (s *HabitsSuite) TestHabitsSelectNext(c *C) {
  habits := NewHabitsDemo()

  selected := habits.SelectNext(1)
  c.Check(len(selected), Equals, 1)

  selected = habits.SelectNext(2)
  c.Check(len(selected), Equals, 2)
}

func (s *HabitsSuite) TestHabitsRepeatFactor(c *C) {
  habits := NewHabitsDemo()
  habits.RepeatFactor = 1
  one := habits.SelectNext(1)[0]
  time.Sleep(1*time.Millisecond)
  two := habits.SelectNext(1)[0]
  time.Sleep(1*time.Millisecond)
  c.Check(one, Not(Equals), two)

  habits.RepeatFactor = 2
  one = habits.SelectNext(1)[0]
  time.Sleep(1*time.Millisecond)
  two = habits.SelectNext(1)[0]
  time.Sleep(1*time.Millisecond)
  c.Check(one, Equals, two)
  one = habits.SelectNext(1)[0]
  time.Sleep(1*time.Millisecond)
  two = habits.SelectNext(1)[0]
  time.Sleep(1*time.Millisecond)
  c.Check(one, Equals, two)

  //break the duality
  _ = habits.SelectNext(1)[0]

  one = habits.SelectNext(1)[0]
  time.Sleep(1*time.Millisecond)
  two = habits.SelectNext(1)[0]
  time.Sleep(1*time.Millisecond)
  c.Check(one, Not(Equals), two)
}

