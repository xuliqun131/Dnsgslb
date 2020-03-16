package models


func SelectNameType(name string, types string) (bool, error) {
	DB := InitDB()
	result, err:= DB.Query("select * from nametypes where " +
		"name_id = (select id from names where name = ?) and type_value = " +
		"(select value from types where type = ?)", name, types)
	if err != nil {
		return false, err
	}
	if result.Next() {
		return true, nil
	}
	return false, nil

}

func InsertNameType(name string, domain string, types string, ttl int, persistence int) error {
	DB := InitDB()
	//result, err := DB.Exec("insert into nametypes(name_id, type_value, ttl, persistence)" +
	//	"select (select id from names where name= ?) as name_id, " +
	//		"(select value from types where type = ?) as type_value, (?) as ttl, (?) as persistence", name, types, ttl, persistence)
	result, err := DB.Exec("insert into nametypes(name_id, type_value, ttl, persistence)" +
		"select(select names.id from names join domains on names.domain_id = domains.id where name = ? and domains.domain = ?)as name_id," +
		"(select value from types where type = ?) as type_value, (?) as ttl, (?) as persistence", name, domain, types, ttl, persistence)
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

func ChangeNameType(types string, name string) error {
	DB := InitDB()
	_, err := DB.Exec("update nametypes, names, types set type_value = " +
		"(select value from types where type = ?) " +
		"where name_id = (select id from names where name = ?)", types, name)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}
