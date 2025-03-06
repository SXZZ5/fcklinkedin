### LLM actually kaam aayega kya ?
- required experience, required skill, preferences and dealbreakers ki matching ke liye to hai useful. 
- bas itna hi chaiye i think baaki to fir kya hi kar sakte ho, deterministic problem to hai hi nhi waise bhi. System Prompt ko improve kar sakte ho thoda aur careful strcuturing se.
- basically llm ne mere liye thoda data extraction bhi kardiya aur thoda "intelligent" matching bhi kardi. 

### LLM_Reply_JSONSchema
```go
type LLM_replySchema struct {
	ReqdSkills []string `json:"reqdSkills`
	ReqdExperience int `json:"reqdExperience`
	MatchRating int `json:"matchRating`
	HaveSkills []string `json:"haveSkills`
	ShouldApply string `json:"shouldApply`
	HasDealbreakers string `json:"dealbreakers`

}
```
Abhi apan mast ye poora data har job description ke liye dalenge badhiya sqlite me. Aur frontend bana hi lunga do kaudi ka.

sqlite are you sure ? i mean postgres aur mysql ke bhari server/service se to simple hona chahiye mamla.

