package book

import "console/system/databases"

func SaveBookInfo()  {
	databases.MasterDB("book")
}