package db

import (
	"fmt"

	"github.com/yusukemisa/go-vue-pay-grpc/backend-api/domain"
)

//SelectAllItems - select all items
func SelectAllItems() (items domain.Items, err error) {
	stmt, err := Conn.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	for stmt.Next() {
		item := &domain.Item{}
		if err := stmt.Scan(&item.ID, &item.Name, &item.Description, &item.Amount); err != nil {
			continue
		}
		items = append(items, *item)
	}
	return
}

//SelectItem - select item
func SelectItem(paramID int64) (item domain.Item, err error) {
	stmt, err := Conn.Prepare(fmt.Sprintf("SELECT * FROM items WHERE id = ? LIMIT 1"))
	if err != nil {
		return
	}
	defer stmt.Close()
	var _id int64
	var _name string
	var _description string
	var _amount int64
	if err = stmt.QueryRow(paramID).Scan(&_id, &_name, &_description, &_amount); err != nil {
		return
	}
	item.ID = _id
	item.Name = _name
	item.Description = _description
	item.Amount = _amount
	return
}
