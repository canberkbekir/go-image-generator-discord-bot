package chatgpt

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/sashabaranov/go-openai"
)

var KeyOpenAI string

type OpenAIRunner interface {
	CreateImage(prompt string) (*[]byte, error)
}

type OpenAI struct {
}

func (o OpenAI) CreateImage(prompt string) (*[]byte, error) {
	c := openai.NewClient(KeyOpenAI)
	ctx := context.Background()

	reqBase64 := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize256x256,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	respBase64, err := c.CreateImage(ctx, reqBase64)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return nil, err
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
		return nil, err
	}

	return &imgBytes, nil
}

func NewOpenAI() OpenAIRunner {
	return &OpenAI{}
}
