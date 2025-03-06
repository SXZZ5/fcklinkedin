package main

import (
	"fcklinkedin/linkedin"
	// "fcklinkedin/utils"
	"fmt"
	// "time"
	"github.com/go-rod/rod"
)


func main() {
	navigator := linkedin.Navigator{}
	navigator.Browser = rod.New().NoDefaultDevice().MustConnect()
	navigator.Login()
	navigator.NavigateToJobs()

	// utils.WriteToFile(string("hello"), string("cutie"), string("pie"))

	// page.MustWaitStable().MustScreenshot("a.png")

	ch := make(chan int)
	w := <-ch
	fmt.Println(w)
}
