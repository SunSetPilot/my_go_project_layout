package core

//
//import (
//	"context"
//
//	"gorm.io/driver/mysql"
//	"gorm.io/gorm"
//
//)
//
//var (
//	MyDB *DBProxy
//)
//
//type DBProxy struct {
//	db *gorm.DB
//}
//
//func (p *DBProxy) NewRequest(ctx context.Context) *gorm.DB {
//	return p.db.Session(&gorm.Session{Context: ctx})
//}
//
//func InitMysql() error {
//	var err error
//	MyDB = new(DBProxy)
//	dsn := global.ProjectConfig.Load().(conf.Config).DataConf.MysqlDSN
//	MyDB.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	if err != nil {
//		return err
//	}
//	return nil
//}
//
//func ExecSQL(sql string) error {
//	err := MyDB.db.Exec(sql).Error
//	return err
//}
