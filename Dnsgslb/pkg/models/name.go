package models


func SelectNames(name string) (bool, error){
	DB := InitDB()
	result, err := DB.Query("select * from names where name = ?", name)
	if err != nil {
		return false, nil
	}
	if result.Next() {
		return true, nil
	}
	return false, nil
}

func InsertNames(domain string, name string) error {
	DB := InitDB()
	result, err := DB.Exec("insert into names(domain_id, name) select " +
		"(select domains.id from domains where domain = ? ) as domain_id, (?) as name", domain, name)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	defer DB.Close()
	return nil
}
