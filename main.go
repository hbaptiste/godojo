package main

import (
	"fmt"
	"mk-lang/lexer"
	"mk-lang/scraper"
)


func testLexer() {

	code := `
		const name = 'Harris'
		let isValid = false
		if (name.length > 10) {
			isValid = true
		}
		func test() {
			return name
		}`

	lexer := lexer.New(code)
	lexer.OnChar(func(char string) {
		fmt.Println("receive char:", char)
	})
	result := lexer.ScanTokens()

	for _, token := range result {
		fmt.Println(token)
	}
}

// function scraper
func testSCrapper() {
	scraper := scraper.CreateScraper()
	TEMPLATE := `
			<div>
				<p> this is it <a>You better know</a></p>
			</div>
				`
	scraper.Visit(TEMPLATE)

}	

func main() {
	testSCrapper()
}
