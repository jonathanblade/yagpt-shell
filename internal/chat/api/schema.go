package api

type TextGenerationMsg = string

type textGenerationFromInstructionRequest struct {
	Model             string                `json:"model"`
	GenerationOptions textGenerationOptions `json:"generationOptions"`
	InstructionText   string                `json:"instructionText"`
	RequestText       string                `json:"requestText"`
}

type textGenerationFromChatRequest struct {
	Model             string                  `json:"model"`
	GenerationOptions textGenerationOptions   `json:"generationOptions"`
	InstructionText   string                  `json:"instructionText"`
	Messages          []textGenerationMessage `json:"messages"`
}

type textGenerationFromInstructionResponse struct {
	Result textGenerationFromInstructionResult `json:"result"`
}

type textGenerationFromChatResponse struct {
	Result textGenerationFromChatResult `json:"result"`
}

type textGenerationFromChatResult struct {
	NumTokens int                   `json:"numTokens"`
	Message   textGenerationMessage `json:"message"`
}

type textGenerationFromInstructionResult struct {
	NumPromptTokens int                         `json:"numPromptTokens"`
	Alternatives    []textGenerationAlternative `json:"alternatives"`
}

type textGenerationAlternative struct {
	Text      string  `json:"text"`
	Score     float64 `json:"score"`
	NumTokens int     `json:"numTokens"`
}

type textGenerationOptions struct {
	PartialResults bool    `json:"partialResults"`
	Temperature    float64 `json:"temperature"`
	MaxTokens      int     `json:"maxTokens"`
}

type textGenerationMessage struct {
	Role string `json:"role"`
	Text string `json:"text"`
}
