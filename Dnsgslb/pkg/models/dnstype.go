package models

import "Dnsgslb/pkg/api/types"

func SelectType(res *types.DnsType) (map[int]types.DnsType, error) {
	DB = InitDB()
	listtype := make(map[int]types.DnsType)
	result, err := DB.Query("select * from types")
	if err != nil {
		return listtype, err
	}
	defer DB.Close()
	var count int
	for result.Next() {
		err := result.Scan(&res.Type, &res.Description, &res.Value)
		if err != nil {
			return listtype, err
		}
		listtype[count] = *res
		count ++
	}
	return listtype, err
}

func InsertType(types string, description string) error{
	DB = InitDB()
	result, err := DB.Exec("insert INTO types(type, description) values(?,?)", types, description)
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

func DelType(typeid int) error {
	DB = InitDB()
	_, err := DB.Exec("delete from types where value = ?", typeid)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}

func ChangeType(types string, description string, args int) error {
	DB = InitDB()
	_, err := DB.Exec("update types set type = ?, description= ? where value = ?", types, description, args)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}