package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	username := string(os.Args[1])
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("request error:", err)
		return
	}
	defer resp.Body.Close()
	var Events []Event
	if err := json.NewDecoder(resp.Body).Decode(&Events); err != nil {
		fmt.Println("json decode error:", err)
		return
	}
	if len(Events) == 0 {
		fmt.Println("no such user:", username)
	}
	for _, event := range Events {
		fmt.Println(event)
	}
}
