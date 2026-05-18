package main

import "fmt"

func main() {
	websites := map[string]string{
		"google":              "google.com",
		"Amazon Web Services": "aws.com",
	}

	fmt.Println(websites)
	fmt.Println(websites["google"])

	websites["facebook"] = "facebook.com"
	fmt.Println(websites)
	delete(websites, "google")
	fmt.Println(websites)
}
