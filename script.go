package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
		"net/http"
		"bytes"
		"io/ioutil"
		"strconv"
)

const username = "***"
const password = "***"
const url = "http://localhost"

type Data struct {
    Data  []DataInfo 
}

type DataInfo struct {
	user string
	password string
	email string
}

func createShoppingList(data [][]string) []DataInfo {
    var dataList []DataInfo
    for i, line := range data {
        if i > 0 {
            var rec DataInfo	
            for j, field := range line {
                if j == 0 {
                    rec.user = field
                } else if j == 1 {
                    rec.password = field
                } else {
										rec.email = field
								}
            }
            dataList = append(dataList, rec)
        }
    }
    return dataList
}

func main() {
		showMenu()
    f, err := os.Open("data.csv")
    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    requestList := createShoppingList(data)

		for i:= 0; i < len(requestList); i++ {

			user := requestList[i].user
			password := requestList[i].password
			email, _ := strconv.Atoi(requestList[i].email)

			stringTest := fmt.Sprintf(`{"data": {
				"user": "%s",
				"password": "%s",
				"email": %d
			}}`, user, password, email)
				
			req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(stringTest)))
			req.Header.Set("Content-Type", "application/json; charset=UTF-8")
			req.SetBasicAuth(username, password)
	
			client := &http.Client{}
			res, err := client.Do(req)
	
			if err != nil {
				fmt.Println(err)
			}
	
			defer res.Body.Close()
	
			fmt.Println("response Status:", res.Status)
			body, _ := ioutil.ReadAll(res.Body)
			fmt.Println("response Body:", string(body))
		}
}


func showMenu() {
	fmt.Println("Script request test")
	fmt.Println("===================")
	fmt.Println("Author: Gabriel Piassa")
	fmt.Println("Version: 0.0.1")
}