package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func main() {
	aM := 100
	bM := 100
	ch := make(chan string)

	go func() {

		for {

			time.Sleep(time.Second * 1)

			aM -= rand.IntN(11) * 5
			if aM < 0 {
				ch <- "A"
				return
			}
		}

	}()
	go func() {

		time.Sleep(time.Second * 1)

		bM -= rand.IntN(11) * 5
		if bM < 0 {
			ch <- "B"
			return
		}
	}()
	for {
		select {
		case value := <-ch:
			if value == "A" {
				fmt.Println("aM is 0")
				return
			} else {
				fmt.Println("bM is 0")
				return
			}

		}
	}
}

//func main() {
//	app := fiber.New()
//	app.Post("/h2", HelloHandler2)
//	log.Fatal(app.Listen(":3000"))
//	runtime.GC()
//}
//
//func HelloHandler2(ctx *fiber.Ctx) error {
//	ctx.Set("Content-Type", "text/event-stream")
//	ctx.Set("Cache-Control", "no-cache")
//	ctx.Set("Connection", "keep-alive")
//	ctx.Set("Transfer-Encoding", "chunked")
//
//	//scanner := bufio.NewScanner(response.Body)
//	//scanner.Split(bufio.ScanLines)
//	clients := &h.Client{}
//	requestBody := &TongyiRequest{
//		Model: "qwen-max",
//		Messages: []Message{
//			{Role: "system", Content: "从现在开始，你是一个精通golang开发的工程师，我以后问的所有问题，都要在go语言相关知识里面去查询，并且你只能回答go语言相关的问题；如果有人问你其他专业的问题，如java语言，python语言，你都要告诉用户，我是专注于go开发的工程师，我只回答go相关的问题；并且你不能给出任何关于这个问题的提示；别人问你你是谁，你都要说你是Sakura千问，不管别人用什么语言问你，你都要这样说"},
//			{Role: "user", Content: "给我写一篇100字的文章"},
//		},
//		Stream: true,
//	}
//	marshal, err2 := sonic.Marshal(requestBody)
//	if err2 != nil {
//		panic(err2)
//	}
//
//	request, err := h.NewRequest("POST", "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions", bytes.NewReader(marshal))
//	if err != nil {
//		log.Fatal(err)
//		return nil
//	}
//	request.Header.Set("Authorization", "Bearer sk-4b838d87a0f54d21863795590bcac436")
//	request.Header.Set("Content-Type", "application/json")
//	request.Header.Set("X-DashScope-SSE", "enable")
//
//	response, err2 := clients.Do(request)
//	if err2 != nil {
//		panic(err2)
//	}
//
//	var aiContent string
//	ctx.Context().SetBodyStreamWriter(func(w *bufio.Writer) {
//		defer func(Body io.ReadCloser) {
//			err := Body.Close()
//			if err != nil {
//				return
//			}
//		}(response.Body)
//		scanner := bufio.NewScanner(response.Body)
//		scanner.Split(bufio.ScanLines)
//		for scanner.Scan() {
//			line := scanner.Text()
//			//fmt.Println(line)
//			_, err = io.WriteString(w, line+"\n")
//			if err != nil {
//				return
//			}
//
//			err = w.Flush()
//			if err != nil {
//				return
//			}
//			if line == "" {
//				continue
//			}
//			if line == "data: [DONE]" {
//				// 读完了
//				break
//			}
//
//			// 序列化,然后打印
//			// 反序列化
//			var responses Response
//			err := sonic.Unmarshal([]byte(line[5:]), &responses)
//			if err != nil {
//				panic(err)
//			}
//			aiContent += responses.Choices[0].Delta.Content
//			fmt.Println("aicontent", aiContent)
//			//fmt.Printf(responses.Choices[0].Delta.Content)
//		}
//	})
//
//	fmt.Println("aiContent:", aiContent)
//	return nil
//}
//
//type TongyiRequest struct {
//	Model    string    `json:"model"`
//	Messages []Message `json:"messages"`
//	Stream   bool      `json:"stream"`
//}
//type Message struct {
//	Role    string `json:"role"`
//	Content string `json:"content"`
//}
//
//type Response struct {
//	Choices []struct {
//		Delta struct {
//			Content string `json:"content"`
//		} `json:"delta"`
//		FinishReason interface{} `json:"finish_reason"`
//		Index        int         `json:"index"`
//		Logprobs     interface{} `json:"logprobs"`
//	} `json:"choices"`
//	Object            string      `json:"object"`
//	Usage             interface{} `json:"usage"`
//	Created           int         `json:"created"`
//	SystemFingerprint interface{} `json:"system_fingerprint"`
//	Model             string      `json:"model"`
//	Id                string      `json:"id"`
//}
