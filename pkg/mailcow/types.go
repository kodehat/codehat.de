package mailcow

type Mailbox struct {
	Active   int    `json:"active"`
	Domain   string `json:"domain"`
	Username string `json:"username"`
}
