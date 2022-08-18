package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type res1 struct {
	Page   int
	Fruits []string
	some   string
}

type res2 struct {
	page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	resD := &res1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"},
		some:   "thing",
	}
	resB, _ := json.Marshal(resD)
	// some is not in the output.
	fmt.Println(string(resB))

	res2D := &res2{
		page:   2,
		Fruits: []string{"apple", "peach", "pear"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var dat map[string]interface{}
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Println(num)

	j := `{"page":1, "Fruits": ["apple", "peach"]}`
	res := res2{}
	json.Unmarshal([]byte(j), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)
}
