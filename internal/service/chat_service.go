package service

type ChatService interface {
    DummyLogic() string
}

type chatServiceImpl struct{}

func NewChatService() ChatService {
    return &chatServiceImpl{}
}

func (c *chatServiceImpl) DummyLogic() string {
    return "Service Layer Logic"
}
