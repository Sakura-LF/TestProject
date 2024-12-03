package AiTest

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3/client"
	"log"
	h "net/http"
)

type TongyiRequest struct {
	Model    string `json:"model"`
	Messages []struct {
		Role    string `json:"role"`
		Content string `json:"content"`
	} `json:"messages"`
	Stream bool `json:"stream"`
}

func NewRequest() {
	c := client.New()
	request := &TongyiRequest{
		Model: "qwen-max",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: "给我写一篇100字的文章"},
		},
		Stream: true,
	}
	header := map[string]string{
		"Authorization":   "Bearer sk-4b838d87a0f54d21863795590bcac436",
		"Content-Type":    "application/json",
		"X-DashScope-SSE": "enable",
	}
	response, err := c.Post("https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions", client.Config{
		Header: header,
		Body:   request,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	stream := response.RawResponse.BodyStream()
	defer response.RawResponse.CloseBodyStream()

	//reader := bytes.NewReader(response.Body())
	scanner := bufio.NewScanner(stream)
	scanner.Split(bufio.ScanLines)
	for {
		line, err := scanner.Text(), scanner.Err()
		if !scanner.Scan() {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		//去除前缀
		//line = strings.TrimPrefix(strconv.Itoa(line), "data: ")

		fmt.Println(line)
	}

	//Responses := make([]Response, 0, 20)
	//err = sonic.Unmarshal(response.Body(), &Responses)
	//if err != nil {
	//	panic(err)
	//}
	//bufio.NewReader()

	//bufio.NewReader(response.Body())
	//fmt.Println("Response:", response.JSON(&Responses))
}

type Response struct {
	Data struct {
		Choices []struct {
			Delta struct {
				Content string `json:"content"`
			} `json:"delta"`
			FinishReason interface{} `json:"finish_reason"`
			Index        int         `json:"index"`
			Logprobs     interface{} `json:"logprobs"`
		} `json:"choices"`
		Object            string      `json:"object"`
		Usage             interface{} `json:"usage"`
		Created           int         `json:"created"`
		SystemFingerprint interface{} `json:"system_fingerprint"`
		Model             string      `json:"model"`
		Id                string      `json:"id"`
	} `json:"data"`
}

func NewRequest2() {
	clients := &h.Client{}

	requestBody := &TongyiRequest{
		Model: "qwen-max",
		Messages: []struct {
			Role    string `json:"role"`
			Content string `json:"content"`
		}{
			{Role: "system", Content: "You are a helpful assistant."},
			{Role: "user", Content: "给我写一篇100字的文章"},
		},
		Stream: true,
	}
	marshal, err2 := sonic.Marshal(requestBody)
	if err2 != nil {
		panic(err2)
	}

	request, err := h.NewRequest("POST", "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions", bytes.NewReader(marshal))
	if err != nil {
		log.Fatal(err)
		return
	}
	request.Header.Set("Authorization", "Bearer sk-4b838d87a0f54d21863795590bcac436")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-DashScope-SSE", "enable")

	response, err2 := clients.Do(request)
	if err2 != nil {
		panic(err2)
	}
	defer response.Body.Close()

	scanner := bufio.NewScanner(response.Body)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		// 处理每一行数据
		fmt.Println(line)

		// 在这里可以添加更多的逻辑来解析事件
		// 例如检查特定的前缀如 "data:" 等等
	}
}
