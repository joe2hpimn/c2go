package main

type ConstAttr struct {
	Address  string
	Position string
	Tags     string
	Children []interface{}
}

func parseConstAttr(line string) *ConstAttr {
	groups := groupsFromRegex(
		"<(?P<position>.*)>(?P<tags>.*)",
		line,
	)

	return &ConstAttr{
		Address:  groups["address"],
		Position: groups["position"],
		Tags:     groups["tags"],
		Children: []interface{}{},
	}
}
