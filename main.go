package main

import (
	// "fcklinkedin/linkedin"
	// "fcklinkedin/utils"
	"fmt"
	// "time"
	// "github.com/go-rod/rod"
	"fcklinkedin/groqtalk"
)

var optum_prompt = `Optum
Software Engineer
About the job

Optum is a global organization that delivers care, aided by technology to help millions of people live healthier lives. The work you do with our team will directly improve health outcomes by connecting people with the care, pharmacy benefits, data and resources they need to feel their best. Here, you will find a culture guided by diversity and inclusion, talented peers, comprehensive benefits and career development opportunities. Come make an impact on the communities we serve as you help us advance health equity on a global scale. Join us to start Caring. Connecting. Growing together.



Required Qualifications

Degree in either Computer Science, Information Technology or a related discipline
2+ years of experience in Application support, Monitoring and Maintenance
Experience in incident management
Experience with monitoring and management tools
Solid knowledge of ITSM and processes
Proven solid analytical skills
Proven problem-solving and analytical skills
Proven excellent communication and teamwork abilities
Proven ability to understand technical terminology and ability to quickly learn and understand technical information

`

func main() {
	// navigator := linkedin.Navigator{}
	// navigator.Browser = rod.New().NoDefaultDevice().MustConnect()
	// navigator.Login()
	// navigator.NavigateToJobs()

	// utils.WriteToFile(string("hello"), string("cutie"), string("pie"))

	groqtalk.MakeChatRequest(optum_prompt)

	ch := make(chan int)
	w := <-ch
	fmt.Println(w)
}
