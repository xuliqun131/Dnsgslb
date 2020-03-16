package models

import "Dnsgslb/pkg/api/types"

func Selectdomainname(domainid int) (string, error) {
	DB = InitDB();
	var name string;
	err := DB.QueryRow("select domain from domains where id = ? ", domainid).Scan(&name)
	if err != nil {
		return name, err
	}
	return name, nil
}

func SelectDomain(res *types.Domains) (map[int]types.Domains, error) {
	DB = InitDB()
	listdomain := make(map[int]types.Domains)
	result, err := DB.Query("select * from domains")
	if err != nil {
		return listdomain, err
	}
	defer DB.Close()
	var count int
	for result.Next() {
		err := result.Scan(&res.Id, &res.Domain)
		if err != nil {
			return listdomain, err
		}
		listdomain[count] = *res
		count ++
	}
	return listdomain, err
}

func InsertDomain(domain string) error{
	DB = InitDB()
	result, err := DB.Exec("insert INTO domains(domain) values(?)", domain)
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

func DelDomain(domainid int) error {
	DB = InitDB()
	_, err := DB.Exec("delete from domains where id = ?", domainid)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}

func ChangeDoamin(domain string, args int) error {
	DB = InitDB()
	_, err := DB.Exec("update domains set domain = ? where id = ?", domain, args)
	if err != nil {
		return err
	}
	defer DB.Close()
	return nil
}