package scyna_setup

type client struct {
	code   string
	secret string
}

func NewClient(code, secret string) *client {
	return &client{code: code, secret: secret}
}

func (c *client) UseEndpoint(url string) *client {
	/*TODO*/
	return c
}

func (c *client) Build() {
	/*TODO*/
}
