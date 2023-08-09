package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/jonathanblade/yagpt-shell/internal/chat/domain"
	"github.com/jonathanblade/yagpt-shell/internal/config"
	"github.com/jonathanblade/yagpt-shell/internal/logger"
)

type Client struct {
	config    *config.Config
	logger    *logger.Logger
	inner     *http.Client
	responses chan TextGenerationMsg
}

func NewClient(config *config.Config, logger *logger.Logger) *Client {
	c := http.Client{Timeout: time.Minute}
	return &Client{
		config:    config,
		logger:    logger,
		inner:     &c,
		responses: make(chan string),
	}
}

func (c *Client) textGenerationFromInstructionRequest(requestText string) *http.Request {
	url := "https://llm.api.cloud.yandex.net/llm/v1alpha/instruct"
	bodyBytes, err := json.Marshal(textGenerationFromInstructionRequest{
		Model: "general",
		GenerationOptions: textGenerationOptions{
			PartialResults: false,
			Temperature:    c.config.Temperature,
			MaxTokens:      2000,
		},
		InstructionText: "",
		RequestText:     requestText,
	})
	if err != nil {
		logger.FatalErr(err)
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		logger.FatalErr(err)
	}
	c.logger.Debugf("API request: %s", string(bodyBytes))
	authHeader := fmt.Sprintf("Api-Key %s", c.config.ApiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authHeader)
	req.Header.Add("x-folder-id", c.config.FolderID)
	return req
}

func (c *Client) textGenerationFromChatRequest(chatHistory []domain.Message, requestText string) *http.Request {
	url := "https://llm.api.cloud.yandex.net/llm/v1alpha/chat"
	textGenerationMessages := make([]textGenerationMessage, 0)
	for _, m := range chatHistory {
		textGenerationMessages = append(textGenerationMessages, textGenerationMessage{Role: m.Role.Name, Text: m.Text})
	}
	bodyBytes, err := json.Marshal(textGenerationFromChatRequest{
		Model: "general",
		GenerationOptions: textGenerationOptions{
			PartialResults: false,
			Temperature:    c.config.Temperature,
			MaxTokens:      2000,
		},
		InstructionText: "",
		Messages:        textGenerationMessages,
	})
	if err != nil {
		logger.FatalErr(err)
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		logger.FatalErr(err)
	}
	c.logger.Debugf("API request: %s", string(bodyBytes))
	authHeader := fmt.Sprintf("Api-Key %s", c.config.ApiKey)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", authHeader)
	req.Header.Add("x-folder-id", c.config.FolderID)
	return req
}

func (c *Client) generateTextFromInstruction(requestText string) string {
	apiReq := c.textGenerationFromInstructionRequest(requestText)
	apiRes, err := c.inner.Do(apiReq)
	if err != nil {
		return err.Error()
	}
	defer apiRes.Body.Close()
	bodyBytes, err := io.ReadAll(apiRes.Body)
	if err != nil {
		return err.Error()
	}
	c.logger.Debugf("API response: %s", string(bodyBytes))
	if apiRes.StatusCode != http.StatusOK {
		return string(bodyBytes)
	}
	var textGenerationResponse textGenerationFromInstructionResponse
	if err := json.Unmarshal(bodyBytes, &textGenerationResponse); err != nil {
		return err.Error()
	}
	return textGenerationResponse.Result.Alternatives[0].Text
}

func (c *Client) generateTextFromChat(chatHistory []domain.Message, requestText string) string {
	apiReq := c.textGenerationFromChatRequest(chatHistory, requestText)
	apiRes, err := c.inner.Do(apiReq)
	if err != nil {
		return err.Error()
	}
	defer apiRes.Body.Close()
	bodyBytes, err := io.ReadAll(apiRes.Body)
	if err != nil {
		return err.Error()
	}
	c.logger.Debugf("API response: %s", string(bodyBytes))
	if apiRes.StatusCode != http.StatusOK {
		return string(bodyBytes)
	}
	var textGenerationResponse textGenerationFromChatResponse
	if err := json.Unmarshal(bodyBytes, &textGenerationResponse); err != nil {
		return err.Error()
	}
	return textGenerationResponse.Result.Message.Text
}

func (c *Client) GenerateTextFromInstructionCmd(requestText string) tea.Cmd {
	return func() tea.Msg {
		text := c.generateTextFromInstruction(requestText)
		c.responses <- text
		return nil
	}
}

func (c *Client) GenerateTextFromChatCmd(chatHistory []domain.Message, requestText string) tea.Cmd {
	return func() tea.Msg {
		text := c.generateTextFromChat(chatHistory, requestText)
		c.responses <- text
		return nil
	}
}

func (c *Client) WaitResponseCmd() tea.Cmd {
	return func() tea.Msg {
		return <-c.responses
	}
}
