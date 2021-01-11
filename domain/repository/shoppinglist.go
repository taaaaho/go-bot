package repository

type ShoppingList interface {
	Delete(content string) bool
	DeleteAll() (int, error)
	Add(content string) error
	List() ([]List, error)
	BroadCast() error
	AddPoint(user string, point int) error
	GetPoint(user string) int
}

type List struct {
	Contents string `firestore:"goods"`
}

const (
	ReplayDelete string = "削除したよ：%s"
)
