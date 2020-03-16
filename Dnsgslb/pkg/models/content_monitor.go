package models

func SelectContentMonitor(content string, monitor string) (bool, error) {
	DB := InitDB()
	result, err := DB.Query("select id from content_monitors where " +
		"content_id = (select id from contents where content = ?) and " +
		"monitor_id = (select id from monitors where monitor = ?)", content, monitor)
	if err != nil {
		return false, nil
	}
	if result.Next() {
		return true, nil
	}
	return false, nil
}


func InsertContentMonitor(content string, monitor string) error{
	DB := InitDB()
	result, err := DB.Exec("insert ignore into content_monitors(content_id, monitor_id) " +
		"select (select id from contents where content = ?) as content_id, " +
		"(select id from monitors where monitor = ?) as monitor_id", content, monitor)
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

func ChangeContentMonitor(content string, monitor string) error {
	DB := InitDB()
	_, err := DB.Exec("update content_monitors, contents, monitors set monitor_id = " +
		"(select id from monitors where monitor = ?) " +
		"where content_id = (select id from contents where content = ?)", monitor, content)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}
