package main 

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"

)

const (
	question_file = "questions/questions.json"
)

type question struct{
	Color			string `json:"color"`
	Backgroundcolor string `json:"backgroundcolor"`     //background-Backgroundcolor of the page
	Topic			string `json:"topic"`
	Question		string `json:"question"`
	Left			string `json:"left"`
	Middle			string `json:"middle"`
	Right			string `json:"right"`
}

type assesment struct {
	Name        string `json:"name"`
	Subject 	string `json:"subject"`
	Questions		[]*question `json:"questions"`
}

type  questionair struct {
	Min    		int `json:"min"`
	Max    		int `json:"max"`
	Step   		int  `json:"step"`
	Assesments		[]*assesment `json:"assesments"`
}


var questions = [...]string {
	"PURPLE",
	"RED",
	"BLUE",
	"ORANGE",
	"YELLOW",
	"GREEN",
}

func CreateDefaultQuestionair()(*questionair, bool){
	q1 := question{
		Color: "white",
		Backgroundcolor: "purple",    //"fuchsia",
		Topic: "Binding:",
		Question: "Hoe is onze onderlinge binding?",
		Left: "ieder voor zich",
		Middle: "echt onderling betrokken",
		Right: "een hechte familie",		
	}
	
	q2 := question{
		Color: "white",
		Backgroundcolor: "red",
		Topic: "Conflict:",
		Question: "Hoe gaan we om met wrevel en conflict?",
		Left: "conflictmijdend,besluiteloos",
		Middle: "stevig en besluitvaardig",
		Right: "Confronterend, conflictueus",		
	}
	
	q3 := question{
		Color: "white",
		Backgroundcolor: "blue",
		Topic: "Structurering:",
		Question: "Hoe gaan we om met regels en afspraken?",
		Left: "ongestructureerd en chaotisch",
		Middle: "Geordend en transparant",
		Right: "star en risico mijdend",		
	}
	
	q4 := question{
		Color: "white",
		Backgroundcolor: "orange",
		Topic: "Prestatie:",
		Question: "hoe prestatiegericht zijn we?",
		Left: "op de winkel Passend, defensief",
		Middle: "Effectief, gezond ambitieus",
		Right: "onderling Concurrerend, ratrace",		
	}
	
	q5 := question{
		Color: "grey",
		Backgroundcolor: "yellow",
		Topic: "Strategie en innovatie:",
		Question: "hoe gaan we om met ideeën en vernieuwing?",
		Left: "te weinig conceptueel denkend",
		Middle: "strategisch en innovatief",
		Right: "heel visionair maar weinig concreet",		
	}
	
	q6 := question{
		Color: "white",
		Backgroundcolor: "green",
		Topic: "Sociale verhouding:",
		Question: "hoe is de sociale verhouding tussen teamleden",
		Left: "formeel en onpersoonlijk",
		Middle: "sociaal en open",
		Right: "overmatig consensusgericht",		
	}
	
	questions := []*question{&q1, &q2, &q3, &q4, &q5, &q6}
	
	asses1 := assesment{
		Name: "Assesment1:",
		Subject: "Zelfbeeld: wij als team",
		Questions: questions,
	}
	
	asses2 := assesment{
		Name: "Assesment2:",
		Subject: "Hoe kijken de collega’s naar dit team?",
		Questions: questions,
	}
	
	asses3 := assesment{
		Name: "Assesment3:",
		Subject: "Hoe kijken de omgeving naar ons ?",
		Questions: questions,
	}
	
	assesments := []*assesment{ &asses1, &asses2, &asses3}
	
	questionair := questionair{
		Min: -5,
		Max: 5,
		Step: 1,
		Assesments: assesments,
	}
	
	
	return &questionair , true
	
}

func main() {
//	for _, Backgroundcolor := range questions {
//		fmt.Println(Backgroundcolor)
//	}

	questionair, _ := CreateDefaultQuestionair()
	
	fmt.Printf("%v", questionair.Assesments[0].Questions[0].Backgroundcolor)
	
	
	
	//spew.Dump(questionair)
	
	
}

