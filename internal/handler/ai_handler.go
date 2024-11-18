package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"log"
	"os"
)

// getLLMResponse sends the message to the Python API and gets the LLM response
func (h *ChatHandler) getLLMResponse(c echo.Context, userMessage string) (string, error) {
	log.Println("Getting LLM response...")
	ctx := c.Request().Context()
	apiKey := os.Getenv("LLM_API_KEY")
	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	aiResponse, err := llms.GenerateFromSinglePrompt(ctx, llm, userMessage)
	if err != nil {
		log.Fatal(err)
	}

	return aiResponse, nil
}

//
//func (h *ChatHandler) getAgentResponse(c echo.Context, userMessage string) (string, error) {
//	log.Println("Getting Agent response...")
//	ctx := c.Request().Context()
//	apiKey := os.Getenv("LLM_API_KEY")
//	llm, err := googleai.New(ctx, googleai.WithAPIKey(apiKey))
//	if err != nil {
//		log.Fatal(err)
//	}
//	agentTools := []tools.Tool{tools.Calculator{}
//	}
//	agent := agents.NewConversationalAgent()
//	aiResponse, err := llms.GenerateFromSinglePrompt(ctx, llm, userMessage)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	return aiResponse, nil
//}
