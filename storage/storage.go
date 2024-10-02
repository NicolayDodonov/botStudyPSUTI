package storage

type Storage interface {
	Save(string, *UserInfo) error
	Get(int) (string, error)
	Print(string) (string, error)
}

type UserInfo struct {
	Username        string
	UserID          int
	TypeApplication string
}
