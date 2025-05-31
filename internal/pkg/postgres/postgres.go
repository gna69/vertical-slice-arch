package postgres

type Client struct {
}

func NewClient() *Client {
	return new(Client)
}

func (c *Client) Exec() {}
