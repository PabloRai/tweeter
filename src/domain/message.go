package domain

type Message struct {
	msg    string
	readed bool
}

func NewMessage(messageToReceive string) *Message {
	message := Message{
		messageToReceive,
		false,
	}
	return &message
}

func (msg *Message) GetRead() bool {
	return msg.readed
}

func (message *Message) GetText() string {
	return message.msg
}

func (msg *Message) Read() {
	msg.readed = true
}
