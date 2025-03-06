package linkedin

import (
	"fcklinkedin/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/proto"
)

/*
									FLOW
NavigateToJobs() called by main
NavigateToJobs() -> IterateListings() ---> VisitElements() ---> ParseJD()

*/

var (
	JobButtonTopBar_cssel = `#global-nav > div > nav > ul > li:nth-child(3)`

	JobSearchBox_xpath = `//input[@aria-label="Search by title, skill, or company"]`

	JobListItemsLi_xpath = `//li[@data-occludable-job-id]`

	JobHeading_xpath = `//main//h1` //dar lagta hai ek din ye dhokha dega

	JobDescription_cssel = `#job-details`

	JobFreshness_xpath = `//main//span[contains(text(), 'ago')]`

	NextPageButton_xpath = `//button[@aria-label="View next page"]`
	CompanyName_xpath    = `//div[contains(@class, "company-name")]`
)

func (z *Navigator) NavigateToJobs() {
	element := z.Page.MustElement(JobButtonTopBar_cssel)
	element.Click(proto.InputMouseButtonLeft, 1)
	element = z.Page.MustElementX(JobSearchBox_xpath)
	keyactions := element.MustKeyActions()
	element.MustInput("Golang")
	keyactions.Press(input.Enter).MustDo()
	z.IterateListings()
}

func (z *Navigator) IterateListings() {
	z.Page.MustWaitDOMStable()
	base := z.Page.MustElementX(JobListItemsLi_xpath)
	for {
		z.VisitElement(base)
		var err error
		base, err = base.Next()
		if err != nil {
			break
		}
	}
}

func (z *Navigator) VisitElement(element *rod.Element) {
	//scroll to the element and click on it to populate the JD region
	element.ScrollIntoView()
	element.MustWaitStable()
	outer := element.MustGetXPath(true)
	outer += `//a`
	fmt.Println("xpath to element's linktag", outer)
	linkelem := z.Page.MustElementX(outer)
	linkelem.MustClick()
	link, err := linkelem.Attribute("href")
	if err != nil {
		fmt.Println(err.Error())
	}

	job := OneJob{}
	job.Link = `www.linkedin.com` + *link
	z.ParseJD(&job)
	utils.WriteToFile(job.CompanyName, job.Title, job.JobDescription)
}

func (z *Navigator) ParseJD(job *OneJob) error {
	jobHeading, err := z.Page.MustElementX(JobHeading_xpath).Text()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	job.Title = jobHeading

	companyName, err := z.Page.MustElementX(CompanyName_xpath).Text()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	job.CompanyName = companyName

	jobDescription, err := z.Page.MustElement(JobDescription_cssel).Text()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	job.JobDescription = jobDescription

	freshness := strings.Split(z.Page.MustElementX(JobFreshness_xpath).MustText(), string(" "))
	numeric, err := strconv.Atoi(freshness[0])
	if err != nil {
		return err
	}
	job.Freshness = numeric
	if freshness[1] == "day" {
		job.FreshnessUnit = string("DAY")
	} else {
		job.FreshnessUnit = string("HOURS")
	}
	return nil
}
