# Go Quiz

A simple Go program to run a quiz from a JSON file.

## Usage

```sh
go run main.go -file=Notation2Solfège.json
```

```sh
go run main.go -file=Notation2Solfège.json -num=16
```

Replace `Notation2Solfège.json` with your quiz file.

## Requirements

- Go installed on your system

## Format of the quiz

```
[
  {
    "text": "What is the Solfège of Notion 1?",
    "answer": "do"
  },
  ...
]
```
