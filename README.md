# Habits

Make new terminal habits stick.


A simplified, terminal based, [spaced-repetition](http://en.wikipedia.org/wiki/Spaced_repetition) program for those who spend a lot of time in the terminal.

![](media/habits.gif)

Simplified in the sense that every time you've shown one item, it
is considered a successful recall. This reduces the model to just "show
me items, and repeat each N times".

## Usage

You can wire up `habits` to anywhere you think would be useful to you:

* Bash profile
* zshrc
* Cron + email

Here's a `.zshrc` based example, which will pop 2 items each time you
open a new terminal window:

```
~/bin/habits -repeat 5 -show 2
```

If you're looking for more, just `habits -h`:

```
$ habits -h
Usage of habits:
  -add="": Add a new habit
  -kill=-1: Remove a habit by index (use list)
  -list=false: List all
  -repeat=5: Repeat factor for each item
  -reporter="plain": A habit reporter. Use 'plain' or 'markdown'
  -show=1: How many items to show each time
```

**Note**: markdown support is WIP - don't use for now.

## Goals

Some goals were considered while building this

### Fast

When rigged into a `.zshrc` or any other shell profile - never hold back
the shell prompty.

For this, Go, Ruby, and Node.js sketches were
developed and benchmarked.

Go was obviously the better choice `by a factor
of 10`.

```
$ time habits
-> Use gh for github issues!
habits  0.00s user 0.00s system 75% cpu 0.006 total
```

### Simple

There are no extra features. It is minimal and will stay minimal.


## Next Up

Some ideas (pull req's welcome)

* Network based habit store
* Sharing habits
* Detect when you've used a CLI habit for a smarter repetition model



# Contributing

Fork, implement, add tests, pull request, get my everlasting thanks and a respectable place here :).


# Copyright

Copyright (c) 2014 [Dotan Nahum](http://gplus.to/dotan) [@jondot](http://twitter.com/jondot). See MIT-LICENSE for further details.

