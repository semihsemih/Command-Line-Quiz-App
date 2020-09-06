# Command-Line-Quiz-App

Quiz App ask a question from CSV file and then give result of how many questions they get right and how many they get incorrect.

### Usage

```sh
go run main.go
```

###### Read a CSV File (Default "questions.csv")
##

```sh
go run main.go -csv="questions.csv"
```

###### Set Time Limit (Default 30 Second)
##

```sh
go run main.go -limit=30
```

###### Shuffle Questions Order (Default "on")
##

```sh
go run main.go -shuffle="off"
```