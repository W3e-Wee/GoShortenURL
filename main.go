package main

import "fmt"

/*
	URL SHORTENER
		Shortens a long URl to a shorter version
		e.g.
			https://chatgpt.com/c/xaosudnljp-asaksjd -> "goshort.lt/short-url"

		Features:
			[ ] Shortens inputed URL to a back-half of NO MORE THAN 16 CHARACTERS.
				e.g. "<domain name>/short-url"

			[ ] Option for custom urls
				- Must be no more than 16 characters

		Components
			1. Database/ Storage
				- Stores the map/ identifier and the long url

			2. Hashing
				- Hashes the identifier for each long url

			3. Unique ID generator
				- Generates a random ID for non-custom urls

*/

func main() {
	fmt.Println("Hello, World!")
}
