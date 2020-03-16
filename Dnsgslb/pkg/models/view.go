package models

import "Dnsgslb/pkg/api/types"

func ListView(res *types.Views) (map[int]types.Views, error){
	DB := InitDB()
	listview := make(map[int]types.Views)
	result, err := DB.Query("select view from views")
	if err != nil {
		return listview, err
	}
	defer DB.Close()
	var count int
	for result.Next() {
		err := result.Scan(&res.View)
		if err != nil {
			return listview, err
		}
		listview[count] = *res
		count ++
	}
	return listview, nil
}

func InsertView(views string, rule string) error {
	DB = InitDB()
	result, err := DB.Exec("insert INTO views(view, rule) values(?,?)", views, rule)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}

func DelView(view string) error {
	DB = InitDB()
	_, err := DB.Exec("delete from views where view = ?", view)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}

func ChangeView(view string, rule string, args string) error {
	DB = InitDB()
	_, err := DB.Exec("update views set view = ?, rule = ? where view = ?", view, rule, args)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}