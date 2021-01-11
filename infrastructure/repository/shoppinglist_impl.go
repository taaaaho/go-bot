package repository

import (
	"fmt"

	"github.com/taaaaho/go-bot/domain/repository"
	"github.com/taaaaho/go-bot/pkg/firestore"
	"google.golang.org/api/iterator"
)

const (
	ShoppingCollection string = "list"
	ShoppingField      string = "goods"

	PointCollection string = "point"
	PointField      string = "count"
)

type shoppingList struct {
	// contents string `firestore:"goods"`
}

type firestoreContent struct {
	Content string `firestore:"goods"`
}

// Delete ... delete shopping list by name
func (r *shoppingList) Delete(content string) bool {
	contextForBot := firestore.Initialze()
	iter := contextForBot.Client.
		Collection(ShoppingCollection).
		Where(ShoppingField, "==", content).
		Documents(contextForBot.Context)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			return false
		}
		doc.Ref.Delete(contextForBot.Context)
		return true
	}
}

func (r *shoppingList) DeleteAll() (int, error) {
	var count int
	contextForBot := firestore.Initialze()
	iter := contextForBot.Client.
		Collection(ShoppingCollection).
		Documents(contextForBot.Context)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			return count, nil
		}
		count++
		doc.Ref.Delete(contextForBot.Context)
	}
}

// Add ... add shopping list
func (r *shoppingList) Add(content string) error {
	contextForBot := firestore.Initialze()

	collection := contextForBot.Client.Collection(ShoppingCollection)
	doc := collection.NewDoc()

	if content != "" {
		_, err := doc.Set(contextForBot.Context, firestoreContent{content})
		if err != nil {
			return fmt.Errorf("Error when adding item: %v", err)
		}
	}
	return nil
}

func (r *shoppingList) List() ([]repository.List, error) {
	contextForBot := firestore.Initialze()
	iter := contextForBot.Client.Collection(ShoppingCollection).Documents(contextForBot.Context)
	c := make([]repository.List, 0)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("Error when get list %v", err)
		}

		ent := repository.List{}
		if err := doc.DataTo(&ent); err != nil {
			return nil, fmt.Errorf("Error parse data %v", err)
		}

		c = append(c, ent)
	}

	return c, nil
}

func (r *shoppingList) BroadCast() error {
	contextForBot := firestore.Initialze()
	iter := contextForBot.Client.
		Collection(ShoppingCollection).
		Documents(contextForBot.Context)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			return fmt.Errorf("Error when delete all list %v", err)
		}
		doc.Ref.Delete(contextForBot.Context)
	}
}

type goodPoint struct {
	Point int `firestore:"count"`
}

func (r *shoppingList) AddPoint(user string, point int) error {
	contextForBot := firestore.Initialze()

	// ユーザーIDに紐づくポイント取得 / 存在しない場合のエラーは握り潰す
	doc, _ := contextForBot.Client.Collection(PointCollection).Doc(user).Get(contextForBot.Context)

	// 存在しない場合はゼロ値がセットされる
	var d goodPoint
	doc.DataTo(&d)

	_, err := contextForBot.Client.Collection(PointCollection).Doc(user).Set(contextForBot.Context, map[string]interface{}{
		"count": d.Point + point,
	})

	if err != nil {
		return err
	}

	return nil
}

func (r *shoppingList) GetPoint(user string) int {
	contextForBot := firestore.Initialze()

	// ユーザーIDに紐づくポイント取得
	doc, _ := contextForBot.Client.Collection(PointCollection).Doc(user).Get(contextForBot.Context)

	// 存在しない場合はゼロ値がセットされる
	var d goodPoint
	doc.DataTo(&d)

	return d.Point
}

// NewShoppingList ... return impl repository
func NewShoppingList() repository.ShoppingList {
	return &shoppingList{}
}
