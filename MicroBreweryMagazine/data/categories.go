package data

import (
	"encoding/json"
	"io"

	"github.com/mateuszlesko/MicroBreweryIoT/MicroBreweryMagazine2/helpers"
)

type Category struct {
	Category_id   int    `json:"id"`
	Category_name string `json:"name"`
}

/// ***********CATEGORY********************
func (c *Category) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}
func ToJSON(c []Category, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}
func (c *Category) FromJSON(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(c) // pass reference to myself, map json to struct
}

/// **********CATEGORY DB LOGIC********************

func SelectCategories() ([]Category, error) {
	err, db := helpers.OpenConnection()

	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM categories")
	if err != nil {
		return nil, err
	}

	var categories []Category

	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.Category_id, &category.Category_name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	defer rows.Close()
	defer db.Close()
	return categories, nil
}

func SelectCategoryWhereID(id int) (*Category, error) {
	err, db := helpers.OpenConnection()
	defer db.Close()
	if err != nil {
		return nil, err
	}
	var category Category

	if err := db.QueryRow("SELECT * FROM categories WHERE category_id=$1;", id).Scan(&category.Category_id, &category.Category_name); err != nil {
		return nil, err
	}

	return &category, nil
}

func UpdateCategory(category Category) (*Category, error) {
	err, db := helpers.OpenConnection()
	if err != nil {
		return nil, err
	}
	smt, err := db.Prepare(`update categories set category_name=$1 where category_id=$2`)
	if err != nil {
		return nil, err
	}
	if _, err := smt.Exec(category.Category_name, category.Category_id); err != nil {
		return nil, err
	}
	defer smt.Close()
	defer db.Close()
	return &category, nil
}

func InsertCategory(name string) error {
	err, db := helpers.OpenConnection()
	if err != nil {
		return err
	}
	smt, err := db.Prepare(`insert into categories(category_name) values($1)`)
	if err != nil {
		return err
	}
	_, err = smt.Exec(name)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	defer smt.Close()
	defer db.Close()
	return nil
}

func DeleteCategory(id int) error {
	err, db := helpers.OpenConnection()
	defer db.Close()
	if err != nil {
		return err
	}
	smt, err := db.Prepare(`delete from categories where category_id=$1`)
	if err != nil {
		return err
	}
	defer smt.Close()
	if err != nil {
		return err
	}
	if _, err := smt.Exec(id); err != nil {
		return err
	}

	return nil
}
