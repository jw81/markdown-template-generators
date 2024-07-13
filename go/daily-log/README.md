# Daily Log Markdown Generator
- Utility for generating a standardized daily log markdown template.

## Usage
`daily-log "Stand-Up Meeting" "1:1 with Bubba" "Backlog Grooming"`

Will generate output that looks like this and you can copy/paste it into a markdown editor, like Obsidian ::
```
Day# 752


# Stand-Up Meeting
---
- 


# 1:1 with Bubba
---
- 


# Backlog Grooming
---
- 


# Other
---
- 


```

## Setup as global executable
```
# build an executable
go build -o daily-log go/daily-log/main.go

# move the executable into the user's local binaries
sudo mv daily-log /usr/local/bin/

# execute the binary file
daily-log "Stand-Up Meeting" "1:1 with Bubba" "Backlog Grooming"
```

## Features
- I like to keep a running count of how many daily logs I've created.  (basically, how many work days have I had here?).  To keep this running day count, this executable creates a `/usr/local/bin/counter.txt` file where that value is managed.  Each run of the executable will increment this value.  You can directly reset the value to whatever is needed by editing that file.
- An `Other` section is added at the end of each log to provide a section for tracking those unplanned things that pop up every day.