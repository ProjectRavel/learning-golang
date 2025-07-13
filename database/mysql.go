package database

var connection string;

func init(){
	connection = "mysql"
}

func GetDataBase () string {
	return connection
}