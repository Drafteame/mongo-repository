package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/Drafteame/mgorepo"
	"github.com/Drafteame/mgorepo/driver"
)

func main() {
	d, err := driver.NewTest(&testing.T{})
	if err != nil {
		panic(err)
	}

	if errSeed := Collection(collection, 100, d.Client().Database(d.DbName())); errSeed != nil {
		panic(errSeed)
	}

	searchExample(d)
	searchAndSortExample(d)
	searchOneExample(d)
	getExample(d)
	updateExample(d)
	deleteExample(d)
}

func searchExample(d mgorepo.Driver) {
	fmtPrintln("------ searchExample ------")

	repo := NewUserRepository(d)

	// Search users with age 21
	age := 21

	orders := NewSearchOrders()
	opts := mgorepo.NewSearchOptions(UserSearchFilters{GreaterThanAge: &age}, orders)

	users, errSearch := repo.Search(context.Background(), opts)
	if errSearch != nil {
		panic(errSearch)
	}

	for _, user := range users {
		fmtPrintln(user)
	}
}

func searchAndSortExample(d mgorepo.Driver) {
	fmtPrintln("------ searchAndSortExample ------")

	repo := NewUserRepository(d)

	// Search users with age 21
	age := 21

	orders := NewSearchOrders().Add("age", mgorepo.OrderDesc)
	opts := mgorepo.NewSearchOptions(UserSearchFilters{GreaterThanAge: &age}, orders)

	users, errSearch := repo.Search(context.Background(), opts)
	if errSearch != nil {
		panic(errSearch)
	}

	for _, user := range users {
		fmtPrintln(user)
	}
}

func searchOneExample(d mgorepo.Driver) {
	fmtPrintln("------ searchOneExample ------")

	repo := NewUserRepository(d)

	// Search user with name "name_1"
	name := "name_1"

	orders := NewSearchOrders()
	opts := mgorepo.NewSearchOptions(UserSearchFilters{Name: &name}, orders).
		WithLimit(1)

	user, errSearch := repo.Search(context.Background(), opts)
	if errSearch != nil {
		panic(errSearch)
	}

	fmtPrintln(user)
}

func getExample(d mgorepo.Driver) {
	fmtPrintln("------ getExample ------")

	repo := NewUserRepository(d)

	// Search user with name "name_1"
	name := "name_1"

	orders := NewSearchOrders()
	opts := mgorepo.NewSearchOptions(UserSearchFilters{Name: &name}, orders).
		WithLimit(1)

	users, errSearch := repo.Search(context.Background(), opts)
	if errSearch != nil {
		panic(errSearch)
	}

	fmtPrintln(users)

	userByID, errGet := repo.Get(context.Background(), users[0].ID)
	if errGet != nil {
		panic(errGet)
	}

	fmtPrintln(userByID)
}

func updateExample(d mgorepo.Driver) {
	fmtPrintln("------ updateExample ------")

	repo := NewUserRepository(d)

	name := "name_1"

	orders := NewSearchOrders()
	opts := mgorepo.NewSearchOptions(UserSearchFilters{Name: &name}, orders).
		WithLimit(1)

	users, errSearch := repo.Search(context.Background(), opts)
	if errSearch != nil {
		panic(errSearch)
	}

	newAge := 22

	updateFields := UserUpdateFields{
		Age: &newAge,
	}

	modified, updateErr := repo.Update(context.Background(), users[0].ID, updateFields)
	if updateErr != nil {
		panic(updateErr)
	}

	fmtPrintln("modified:", modified)

	userByID, errGet := repo.Get(context.Background(), users[0].ID)
	if errGet != nil {
		panic(errGet)
	}

	fmtPrintln(userByID)
}

func deleteExample(d mgorepo.Driver) {
	fmtPrintln("------ deleteExample ------")

	repo := NewUserRepository(d)

	name := "name_1"

	orders := NewSearchOrders()
	opts := mgorepo.NewSearchOptions(UserSearchFilters{Name: &name}, orders).
		WithLimit(1)

	users, errSearch := repo.Search(context.Background(), opts)
	if errSearch != nil {
		panic(errSearch)
	}

	deleted, deleteErr := repo.Delete(context.Background(), users[0].ID)
	if deleteErr != nil {
		panic(deleteErr)
	}

	fmtPrintln("deleted:", deleted)

	userByID, errGet := repo.Get(context.Background(), users[0].ID)
	fmtPrintln("userByID:", userByID, "errGet:", errGet)
}

func fmtPrintln(args ...any) {
	_, _ = fmt.Println(args...)
}
