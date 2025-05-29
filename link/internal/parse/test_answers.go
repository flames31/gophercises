package parse

func getAnswers() map[int][]Link {
	return map[int][]Link{
		1: []Link{{
			Href: "/other-page",
			Text: "A link to another page",
		}},

		2: []Link{{
			Href: "https://www.twitter.com/joncalhoun",
			Text: "Check me out on twitter",
		},
			{
				Href: "https://github.com/gophercises",
				Text: "Gophercises is on Github!",
			}},
	}
}
