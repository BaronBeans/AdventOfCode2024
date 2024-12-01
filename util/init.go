package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"text/template"
)

func createDayFromTemplate(day int) {
	dayFilePath := fmt.Sprintf("days/day%d.go", day)
	file, err := os.Create(dayFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()
	tmplString := template.Must(template.ParseFiles("day.template"))
	err = tmplString.Execute(file, day)
	if err != nil {
		fmt.Println("Error writing file:", err)
	}
}

func updateMain(day int) {
	mainFilePath := "main.go"
	mainContents, errRead := os.ReadFile(mainFilePath)

	if errRead != nil {
		fmt.Println("Error reading file:", errRead)
		return
	}

	newMainContents := strings.Replace(string(mainContents), "//insert here", fmt.Sprintf("\"day%d\": days.Day%d,\n\t//insert here", day, day), 1)

	errWrite := os.WriteFile(mainFilePath, []byte(newMainContents), 0644)

	if errWrite != nil {
		fmt.Println("Error writing file:", errWrite)
		return
	}
}

func getInputFile(day int) {
	dayFilePath := fmt.Sprintf("days/day%d.input", day)
	file, err := os.Create(dayFilePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()

	client := &http.Client{}
	url := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request", err)
	}
	req.Header.Set("Content-Type", "application/json")
	cookieValue := os.Getenv("AOC_TOKEN")
	cookie := &http.Cookie{
		Name:  "session",
		Value: cookieValue,
	}
	req.AddCookie(cookie)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response", err)
		return
	}

	errWrite := os.WriteFile(dayFilePath, []byte(body), 0644)
	if errWrite != nil {
		fmt.Println("Error writing file:", errWrite)
		return
	}
}

func Init(day int) {
	fmt.Println(day)
	createDayFromTemplate(day)
	updateMain(day)
	getInputFile(day)
}
