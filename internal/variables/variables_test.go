package variables_test

import (
	"fmt"
	"github.com/iamkaran/tb-override/internal/variables"
	"log"
	"testing"
)

func TestJSONFetch(t *testing.T) {
	variables, err := variables.LoadMap("variables.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("data: %v", variables.Data)
}
