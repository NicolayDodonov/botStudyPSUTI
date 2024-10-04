package storage

type Storage interface {
	Save(string, *UserInfo) error
	Print(string) (string, error)
}

type UserInfo struct {
	Username        string
	TypeApplication string
}

const (
	Tg = "telegram"
	Vk = "vk"
)
