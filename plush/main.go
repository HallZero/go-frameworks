package main

import (
	"fmt"
	"html/template"
	"log"
	"os"

	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/tags"
)

func main(){
	ctx := plush.NewContext();
	ctx.Set("name", "Mark")
	ctx.Set("names", []string{"John", "Paul", "George", "Ringo"})
	ctx.Set("js", func(s string) template.HTML {
		return template.HTML(`<alert>you've been pwned</alert>;`)
	})
	ctx.Set("div", func(opts map[string]interface{}, help plush.HelperContext) (template.HTML, error) {
		div := tags.New("div", opts)

		if help.Value("name") != nil && help.Value("name").(string) == "Mark" {
			
			s, err := help.Block()
	
			if err !=  nil {
				return template.HTML(""), err
			}
	
			div.Append(s)
		}


		return div.HTML(), nil

	})
	s, err := plush.Render(html(), ctx);
	
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(s)

}

func html() string {
	b, err := os.ReadFile("./index.html")

	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}