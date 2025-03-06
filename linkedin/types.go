package linkedin 
import (
	"github.com/go-rod/rod"
)

type Navigator struct {
	Browser *rod.Browser
	Page *rod.Page
}

type OneJob struct {
	CompanyName string
	Link string
	Title string
	JobDescription string
	Freshness int
	FreshnessUnit string
}
//maybe later I will change this Freshness information into an actual timestamp 
//so that the true freshness at a later time can be calculated inplace without 
//having to view the link again or smth