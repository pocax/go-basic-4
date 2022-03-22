package main

func main() {
	var randomValues interface{}
	_ = randomValues
	randomValues = "Jalan thamrin"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Jalan", "thamrin"}
}

func main() {
	var v interface{}
	v = 20
	v = v * 9
}

//type assertion
func main() {
	var v interface{}
	v = 20

	if value, ok := v.(int); ok == true {
		v = value * 9
	}
}

func main() {
	rs := []interface{}{20, "Jalan", true, 2, "thamrin", true}
	rm := map[string]interface{}
	{
		name: "Jalan",
		age:  20,
		status: true,
	}

	_, _ = rs, rm
}
