package lib

type Item struct {
	Name string `json:"Name"`
}

func GetFacts() []Item {
    var items = []Item {
        {
            Name: "Harry",
        },
        {
            Name: "Connor",
        },
        {
            Name: "Matt",
        },
    }

    return items
}
