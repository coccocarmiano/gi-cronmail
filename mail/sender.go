package mail

type MailClient interface {
	Send(string, string) error
}
