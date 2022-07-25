package db

import (
	"github.com/mbolis/yello/model"
)

var allLists model.Lists = model.Lists{
	Lists: []model.List{},
}

func GetAllLists() (model.Lists, error) {
	return allLists, nil
}

func AddList(l *model.List) error {
	var maxId uint
	for _, lst := range allLists.Lists {
		if lst.Id > maxId {
			maxId = lst.Id
		}
	}

	l.Id = maxId + 1
	allLists.Lists = append(allLists.Lists, *l)
	return nil
}
