package telegram

type apiMessageFrom struct {
	Id    int  `json:"id"`
	IsBot bool `json:"is_bot"`
}

type apiMessageChat struct {
	Id int `json:"id"`
}

type apiMessage struct {
	Id   int            `json:"message_id"`
	From apiMessageFrom `json:"from"`
	Chat apiMessageChat `json:"chat"`
	Text string         `json:"text"`
}

type apiUpdate struct {
	UpdateId int        `json:"update_id"`
	Message  apiMessage `json:"message"`
}

type apiResponse struct {
	Status           bool        `json:"ok"`
	Result           []apiUpdate `json:"result"`
	ErrorDescription string      `json:"description"`
}

type apiSendResponse struct {
	Status           bool   `json:"ok"`
	ErrorDescription string `json:"description"`
}
