package message_handler

type FunctionCall struct {
	Name      string `json:"name,omitempty"`
	Arguments string `json:"arguments"`
}

type FunctionReturnMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
	Name    string `json:"name"`
}

type FunctionCallMessage struct {
	Role         string       `json:"role"`
	Content      string       `json:"content"`
	FunctionCall FunctionCall `json:"function_call"`
}

type PlainTextMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type MessageType struct {
	FunctionReturnMessage FunctionReturnMessage `json:"function_return_message,omitempty"`
	FunctionCallMessage   FunctionCallMessage   `json:"function_call_message,omitempty"`
	PlainTextMessage      PlainTextMessage      `json:"plain_text_message,omitempty"`
}

type Message struct{}

func (m *Message) newText(role string, content string) PlainTextMessage {
	return PlainTextMessage{
		Role:    role,
		Content: content,
	}
}

func (m *Message) System(content interface{}) PlainTextMessage {
	return m.newText("system", content.(string))
}

func (m *Message) User(content interface{}) PlainTextMessage {
	return m.newText("user", content.(string))
}

func (m *Message) Assistant(content interface{}) PlainTextMessage {
	return m.newText("assistant", content.(string))
}

func (m *Message) FunctionCall(call interface{}) FunctionCallMessage {
	callData := call.(map[string]interface{})
	return FunctionCallMessage{
		Role:         "assistant",
		FunctionCall: FunctionCall{Name: callData["name"].(string), Arguments: callData["arguments"].(string)},
		Content:      "",
	}
}

func (m *Message) FunctionReturn(name interface{}, content interface{}) FunctionReturnMessage {
	return FunctionReturnMessage{
		Role:    "function",
		Name:    name.(string),
		Content: content.(string),
	}
}
