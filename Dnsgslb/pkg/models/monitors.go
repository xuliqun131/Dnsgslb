package models

import "Dnsgslb/pkg/api/types"

func ListMonitor(res *types.Monitors) (map[int]types.Monitors, error) {
	DB := InitDB()
	listmonitor := make(map[int]types.Monitors)
	result, err := DB.Query("select monitor from monitors")
	if err != nil {
		return listmonitor, err
	}
	defer DB.Close()
	var count int
	for result.Next() {
		err := result.Scan(&res.Monitor)
		if err != nil {
			return listmonitor, err
		}
		listmonitor[count] = *res
		count ++
	}
	return listmonitor, err
}

func ListallMonitor(res *types.Monitors) (map[int]types.Monitors, error) {
	DB := InitDB();
	listmonitor := make(map[int]types.Monitors)
	result, err := DB.Query("select * from monitors");
	if err != nil {
		return listmonitor, err
	}
	defer DB.Close()
	var count int
	for result.Next() {
		err := result.Scan(&res.Id, &res.Monitor, &res.Type, &res.Content, &res.Port, &res.Interval, &res.Timeout)
		if err != nil {
			return listmonitor, err
		}
		listmonitor[count] = *res
		count ++
	}
	return listmonitor, nil
}

func SelcetMonitor(monitor string, res *types.Monitors) (*types.Monitors, error) {
	DB := InitDB()

	err := DB.QueryRow("select * from monitors where monitor = ? ", monitor).Scan(&res.Id, &res.Monitor, &res.Type, &res.Content, &res.Port, &res.Interval, &res.Timeout)
	if err != nil {
		return res, err
	}
	defer DB.Close()
	return res,err
}

func InsertMonitor(monitor string, types string, mocontent string, port string, interval int, timeout int) error{
	DB = InitDB()
	result, err := DB.Exec("insert INTO monitors(monitor,monitor_type,monitor_content,monitor_port,monitor_interval,monitor_timeout) values(?, ?, ?, ?, ?, ?)", monitor, types, mocontent, port, interval, timeout)
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

func DelMonitor(monitorid int) error {
	DB = InitDB()
	_, err := DB.Exec("delete from monitors where id = ?", monitorid)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}

func ChangeMonitor(monitor string, types string, mocontent string, port string, interval int, timeout int, args int) error {
	DB = InitDB()
	_, err := DB.Exec("update monitors set monitor = ?, monitor_type= ?, monitor_content = ?, monitor_port = ?, monitor_interval = ?, monitor_timeout = ? where id = ?", monitor, types, mocontent, port, interval, timeout, args)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}