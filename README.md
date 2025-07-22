# Go Quiz

A simple command line quiz application written in Go that reads a CSV file containing questions and answers.

## Installation

Make sure you have Go installed in your system.

```bash
git clone https://github.com/bkandh30/goquiz.git
cd goquiz
go build -o quiz
```

## Usage

### Basic Usage

```bash
./quiz
```

Note: This will use the default `problems.csv` file.

### Custom CSV File

```bash
./quiz -csv=filename.csv
```

Note: Check `problems.csv` for the CSV file format.
