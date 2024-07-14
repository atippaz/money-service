package other_db

import (
	"github.com/google/uuid"
)

type DummySpendingInstance interface {
	GetSpending() []DummySpending
}

type DummySpending struct {
	ID   uuid.UUID
	Name string
}

type dummySpendingInstance struct{}

func (d *dummySpendingInstance) GetSpending() []DummySpending {
	spending := []DummySpending{
		{ID: uuid.New(), Name: "sd1"},
		{ID: uuid.New(), Name: "sd2"},
	}
	return spending
}

func GetInstanceDb() DummySpendingInstance {
	return &dummySpendingInstance{}
}
