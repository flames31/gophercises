package main

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestJsonConverter(t *testing.T) {
	jsonData := `
	{
		"intro": {
		"title": "The Little Blue Gopher",
		"story": [
			"Once upon a time, long long ago, there was a little blue gopher. Our little blue friend wanted to go on an adventure, but he wasn't sure where to go. Will you go on an adventure with him?",
			"One of his friends once recommended going to New York to make friends at this mysterious thing called \"GothamGo\". It is supposed to be a big event with free swag and if there is one thing gophers love it is free trinkets. Unfortunately, the gopher once heard a campfire story about some bad fellas named the Sticky Bandits who also live in New York. In the stories these guys would rob toy stores and terrorize young boys, and it sounded pretty scary.",
			"On the other hand, he has always heard great things about Denver. Great ski slopes, a bad hockey team with cheap tickets, and he even heard they have a conference exclusively for gophers like himself. Maybe Denver would be a safer place to visit."
		],
		"options": [
			{
			"text": "That story about the Sticky Bandits isn't real, it is from Home Alone 2! Let's head to New York.",
			"arc": "new-york"
			},
			{
			"text": "Gee, those bandits sound pretty real to me. Let's play it safe and try our luck in Denver.",
			"arc": "denver"
			}
		]
		}
	}`
	b := bytes.Buffer{}
	b.Write([]byte(jsonData))
	got, err := JsonConverter(&b)
	want := Adventure(make(map[string]Chapter))
	want["intro"] = Chapter{
		Title: "The Little Blue Gopher",
		Story: []string{
			"Once upon a time, long long ago, there was a little blue gopher. Our little blue friend wanted to go on an adventure, but he wasn't sure where to go. Will you go on an adventure with him?",
			"One of his friends once recommended going to New York to make friends at this mysterious thing called \"GothamGo\". It is supposed to be a big event with free swag and if there is one thing gophers love it is free trinkets. Unfortunately, the gopher once heard a campfire story about some bad fellas named the Sticky Bandits who also live in New York. In the stories these guys would rob toy stores and terrorize young boys, and it sounded pretty scary.",
			"On the other hand, he has always heard great things about Denver. Great ski slopes, a bad hockey team with cheap tickets, and he even heard they have a conference exclusively for gophers like himself. Maybe Denver would be a safer place to visit.",
		},
		Options: []Option{
			Option{
				Text: "That story about the Sticky Bandits isn't real, it is from Home Alone 2! Let's head to New York.",
				Arc:  "new-york",
			},
			Option{
				Text: "Gee, those bandits sound pretty real to me. Let's play it safe and try our luck in Denver.",
				Arc:  "denver",
			},
		},
	}
	if err != nil {
		fmt.Printf("ERR : %v\n", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("TEST FAILED!\ngot - %v\nwant - %v\n", got, want)
	} else {
		fmt.Println("TEST PASSED!")
	}
}
