package pattern

import "fmt"

// IDBconnection : Struct can be assigned any type in runtime
type IDBconnection interface {
	Connect()
}

type DBconnection struct {
	Db IDBconnection // Compatible to accept any type in runtime
}

// DBConnect : Receiver function for Struct
func (con DBconnection) DBConnect() {
	con.Db.Connect()
}

// MySqlConnection : Implementing first behavior to connect to MySQL
type MySqlConnection struct {
	ConnectionString string
}

func (con MySqlConnection) Connect() {
	fmt.Println("Connecting to MySQL DB with connection string : ", con.ConnectionString)
}

// PostgreSqlConnection : Implementing second behavior to connect to PostgreSql
type PostgreSqlConnection struct {
	ConnectionString string
}

func (con PostgreSqlConnection) Connect() {
	fmt.Println("Connecting to PostgreSql DB with connection string : ", con.ConnectionString)
}

// MongoDbConnection : Implementing third behavior to connect to MongoDB
type MongoDbConnection struct {
	ConnectionString string
}

func (con MongoDbConnection) Connect() {
	fmt.Println("Connecting to MongoDB with connection string : ", con.ConnectionString)
}

func GetStrategyName(s string) string {
	return "Code for : -- " + s + " Pattern --"
}
