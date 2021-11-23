# Multiple request test script

## This script running in GO if not have compiler execute this:

*** Open url: https://go.dev/dl/ download GO to your OS ***

### Linux
```bash
$ tar -xzf downloadfile.tar && mv go /usr/local/
# Open with your editor ~/.profile or your path variable ambience example:
$ vim ~/.profile
#add export PATH=$PATH:/usr/local/go/bin
export PATH=$PATH:/usr/local/go/bin
$ source ~/.profile

#execute test
$ go version
# output: go version go1.14.2 linux/amd64
```

## Now change file data.csv with you paylod

column 1 | column 2 | column 3

#### If your payload have more or less column, it will be necessary to change script

```go
//Change this conditions to switch case
for j, field := range line {
    if j == 0 {
        rec.column1 = field
    } else if j == 1 {
        rec.column2 = field
    } else {
        rec.column3 = field
    }
}

//update all data struct 
type Data struct {
    Data  []DataInfo 
}

type DataInfo struct {
	column1 string
	column2 string
	column3 string
}

//Change this data structure
column1 := requestList[i].column1
column2 := requestList[i].column2
column3, _ := strconv.Atoi(requestList[i].column3) //if your column3 is integer

stringTest := fmt.Sprintf(`{"data": {
  "column1": "%s",
  "column2": "%s",
  "column3": %d
}}`, column1, column2, port)
```

### Add URL, Username and Password
```go
const username = "username"
const password = "password"
const url = "http://localhost"
```


## After install and updated running script

```bash
$ go build test.go
$ go run test.go
```
