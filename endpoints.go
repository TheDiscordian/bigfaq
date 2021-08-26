package main

import (
	"fmt"
	"net/http"
	//"html"
	"encoding/json"
	//"net/url"
)

/*func dumpLinkInfo(w http.ResponseWriter, r *http.Request) {
	getVars := r.URL.Query()
	link := getVars.Get("link")
	fmt.Fprintf(w, html.EscapeString(link))
}*/

type ChallengeAnswerBundle struct {
	Challenge string
	Answer string
	URL string
	Author string
	Github string
	Resolved bool
}

func queryQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	getVars := r.URL.Query()
	question := getVars.Get("q")
	fmt.Println(question);
	challenge := searchChallenges(question)
	var out []byte
	if challenge.Resolved {
		answer := searchAnswersByChallengeID(challenge.Id)
		out, _ = json.Marshal(&ChallengeAnswerBundle{Challenge:challenge.Name, Answer:answer.Text, URL:answer.URL, Author:answer.Name, Github:challenge.Github, Resolved:challenge.Resolved})
	} else {
		out, _ = json.Marshal(&ChallengeAnswerBundle{Challenge:challenge.Name, Github:challenge.Github, Resolved:challenge.Resolved})
	}
	fmt.Fprintf(w, string(out))
}

func initHandlers() {
	//http.HandleFunc("/linkInfo", dumpLinkInfo)
	http.HandleFunc("/", queryQuestion)
}
