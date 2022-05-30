package config

type TelegramType struct {
	Token  string `json:"token"`
	ChatID int    `json:"chat_id"`
}

type SmtpType struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Sender   string `json:"sender"`
}

type credentialsType struct {
	TelegramMainBot TelegramType `json:"telegram_main_bot"`
	TelegramLogBot  TelegramType `json:"telegram_log_bot"`
	SMTP            SmtpType     `json:"smtp"`
}

type configType struct {
	ProjectName string `json:"project_name"`
	Email       string `json:"email"`
	TestUrl     string `json:"test_url"`
}
