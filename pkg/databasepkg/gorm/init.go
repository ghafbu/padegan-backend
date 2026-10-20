package gorm

type DBName string

const (
	PostgresDB DBName = "postgres"
	MysqlDB    DBName = "mysql"
	SqliteDB   DBName = "sqlite"
)

//func Init(cfg *config.App, dbName DBName) contract.DB {
//	switch dbName {
//	case Postgres:
//		return postgres.Init(cfg)
//	case MysqlDB:
//		return mysql.Init(cfg)
//	//case SqliteDB:
//	//	return mysql.InitMigrator(cfg)
//	default:
//		panic("don`t support database!")
//	}
//}
