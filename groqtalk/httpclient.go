package groqtalk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var Myhttpclient = http.Client{
	Timeout: time.Second * 3,
} //single instance of http client hi use karunga throughout the program to aise hi export kare de rha hu.

var (
	groq_endpoint = `https://api.groq.com/openai/v1/chat/completions`
	groq_apikey   = `lollypop_logey?`
)

var systemPromptSet = false

/*=============== Groq_* => GROQ API REQUEST/RESPONSE se related types ================= */
type Groq_message struct {
	Role    string `json:"role"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type Groq_responseformat struct {
	Type string `json:"type"`
}

type Groq_jsonreq struct {
	Messages            []Groq_message      `json:"messages"`
	Model               string              `json:"model"`
	MaxCompletionTokens int                 `json:"max_completion_tokens"`
	ResponseFormat      Groq_responseformat `json:"response_format"`
}

type Groq_choice struct {
	Index   int          `json:"index"`
	Message Groq_message `json:"message"`
}

type Groq_response struct {
	ResponseObjType string        `json:"object"`
	Model           string        `json:"model"`
	Choices         []Groq_choice `json:"choices"`
}

/*================ LLM_* => LLM ka Structured Reply Message related types ============== */



func prepareRequestBody(msg string) Groq_jsonreq {
	msg += `
	Remember to respond in the json schema that I told you about at the start.`
	reqbody := Groq_jsonreq{
		Messages: []Groq_message{
			{
				Role:    "system",
				Name:    "system",
				Content: sysprompt,
			},
			{
				Role:    "user",
				Name:    "sushant",
				Content: msg,
			},
		},
		Model:               "gemma2-9b-it",
		MaxCompletionTokens: 1000,
		ResponseFormat:      Groq_responseformat{Type: "json_object"},
	}
	// fmt.Println("json_reqbody:", reqbody)
	return reqbody
}

func MakeChatRequest(msg string) {
	json_reqbody, err := json.Marshal(prepareRequestBody(msg))
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(json_reqbody))

	request, err := http.NewRequest("POST", groq_endpoint, bytes.NewReader(json_reqbody))
	if err != nil {
		fmt.Println("err in making request object")
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", "Bearer "+groq_apikey)
	resp, err := Myhttpclient.Do(request)
	if err != nil {
		fmt.Println("Erred Responese", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error reading error response:", err)
			return
		}
		fmt.Printf("Bad StatusCode %s, Error: %s\n", resp.Status, string(bodyBytes))
		return
	}

	var response_struct Groq_response
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&response_struct)
	// fmt.Println("Response received: \n", response_struct)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Bad StatusCode", resp.Status)
		return
	}

	var json_llm_reply LLM_replySchema
	llm_response := response_struct.Choices[0].Message.Content
	if err := json.Unmarshal([]byte(llm_response), &json_llm_reply); err != nil {
		fmt.Println("llm_reply_err:", err.Error())
	}
	fmt.Println("\n chat response: ", llm_response)
	fmt.Println("===============================\n")
	fmt.Println(json_llm_reply)
}

func TestRequest(msg string) {
	msg = "Hello How are you doing LLM boi ? "

	reqbody := Groq_jsonreq{
		Messages: []Groq_message{
			{
				Role:    "user",
				Name:    "sushant",
				Content: msg,
			},
		},
		Model:               "gemma2-9b-it",
		MaxCompletionTokens: 100,
	}

	json_reqbody, err := json.Marshal(reqbody)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	request, err := http.NewRequest("POST", groq_endpoint, bytes.NewReader(json_reqbody))
	if err != nil {
		fmt.Println("err in making request object")
	}
	request.Header.Add("Authorization", "Bearer "+groq_apikey)
	request.Header.Add("Content-Type", "application/json")
	resp, err := Myhttpclient.Do(request)
	if err != nil {
		fmt.Println("Erred Responese", err.Error())
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Bad StatusCode", resp.Status)
		return
	}

	var response_struct Groq_response
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&response_struct)
	fmt.Println("Response received: \n", response_struct)
	fmt.Println("\n chat response: ", response_struct.Choices[0].Message.Content)
}
