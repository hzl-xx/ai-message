package common

type Params struct {
	Type string `json:"type"`
	Sentry Sentry `json:"sentry"`
	Common Common `json:"common"`
	Mail Mail `json:"mail"`
}

type Common struct {
	Type string `json:"type"`
	Message string `json:"message"`
}

type Sentry struct {
	ProjectName string `json:"projectName"`
	Level string `json:"level"`
	Time string `json:"time"`
	Message string `json:"message"`
	Href string `json:"href"`
	Type string `json:"type"`
}

type Mail struct {
	Title string `json:"title"`
	From string `json:"from"`
	To string `json:"to"`
	Message string `json:"message"`
	Password string `json:"password"`
}
