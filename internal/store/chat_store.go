package store

type ChatStore interface {
    SaveMessage(msg string) error
}

type chatStoreImpl struct{}

func NewChatStore() ChatStore {
    return &chatStoreImpl{}
}

func (cs *chatStoreImpl) SaveMessage(msg string) error {
    // Dummy store method
    return nil
}
