package main

import (
	"fmt"
	"github.com/OrgDeBassac/releasy-ai/internal/git"
)

func main() {
	b, err := git.GetReleaseBranches()
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("branches: %#v\n", b)
}
