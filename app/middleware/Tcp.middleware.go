package middleware

func (Mr MiddlewareRegister) TcpMiddlewareRegister() map[string][]MakeMiddleware{
	return map[string][]MakeMiddleware{
		"@begin":{isTcp},
	}
}