package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filePath = "/usr/local/bin/counter.txt"

func main() {
  output := dayCountContent()
  args := os.Args[1:]

  if len(args) == 0 {
    fmt.Println("No arguments provided")
  } else {
    for _, arg := range args {
      output += addEvent(arg)
    }
  }
  output += addEvent("Other")

  fmt.Println(output)
}

func dayCountContent() string {
  count := getDayCount()
  count = updateDayCount(count)

  return fmt.Sprintf("Day# %d\n\n\n", count)
}

func getDayCount() int {
  var count int
  data, err := os.ReadFile(filePath)

  if err != nil {
    count = 0
  } else {
    valueFromFile := strings.TrimSpace((string(data)))
    count, err = strconv.Atoi(valueFromFile)

    if err != nil {
      fmt.Println("Error reading count:", err)
    }
  }

  return count
}

func updateDayCount(currentCount int) int {
  newCount := currentCount + 1

  err := os.WriteFile(filePath, []byte(strconv.Itoa(newCount)), 0644)
  if err != nil {
    fmt.Println("Error writing count:", err)
  }

  return newCount
}

func addEvent(eventName string) string {
  heading := fmt.Sprintf("# %s\n", eventName) 
  heading += "---\n"
  heading += "- \n\n\n"
  return heading
}