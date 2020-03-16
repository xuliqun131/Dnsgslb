package models


func SelectContent(content string) (bool, error) {
	DB = InitDB()
	result, err := DB.Query("select id from contents where content = ? ", content)
	if err != nil {
		return false, err
	}
	if result.Next() {
		return true, nil
	}
	return false, nil
}

func InsertContent(content string) error {
	DB = InitDB()
	result, err := DB.Exec("insert INTO contents(content) values(?)", content)
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

func DelContent(content string) error {
	DB = InitDB()
	_, err := DB.Exec("delete from contents where content = ?", content)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}

func ChangeContent(content string, args string) error {
	DB = InitDB()
	_, err := DB.Exec("update contents set content = ? where content = ?", content, args)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}

