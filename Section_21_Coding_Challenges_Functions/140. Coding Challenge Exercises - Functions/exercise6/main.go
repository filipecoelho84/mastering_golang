package main

func searchItem(a []string, b string) (c bool) {

	for _, v := range a {
		if b == v {
			return true
		} else {
			return false
		}
	}
	return
}

func main() {

	animals := []string{"lion", "tiger", "bear"}
	result := searchItem(animals, "bear")

}
