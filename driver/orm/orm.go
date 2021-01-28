package orm

import (
	"errors"

	"github.com/leguminosa/kururu-engine/driver"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	dialectMap = map[string]func(dsn string) gorm.Dialector{
		"postgres": postgres.Open,
		"mysql":    mysql.Open,
	}
)

func Open(dialect, connectionString string) (driver.OrmItf, error) {
	dialector, ok := dialectMap[dialect]
	if !ok {
		return nil, errors.New("invalid dialect given")
	}

	db, err := gorm.Open(dialector(connectionString), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	return wrap(db), err
}

type (
	gormDB struct {
		*gorm.DB
	}
)

func wrap(db *gorm.DB) driver.OrmItf {
	return &gormDB{DB: db}
}

func (db *gormDB) Select(query string, args ...interface{}) driver.OrmItf {
	return wrap(db.DB.Select(query, args...))
}

func (db *gormDB) Where(query interface{}, args ...interface{}) driver.OrmItf {
	return wrap(db.DB.Where(query, args...))
}

func (db *gormDB) Preload(query string, args ...interface{}) driver.OrmItf {
	return wrap(db.DB.Preload(query, args...))
}

// Specify the model target for given database operation
func (db *gormDB) Model(value interface{}) driver.OrmItf {
	return wrap(db.DB.Model(value))
}

// Insert value to database
func (db *gormDB) Create(value interface{}) driver.OrmItf {
	return wrap(db.DB.Create(value))
}

// Insert value to database or update it if primary key is given
func (db *gormDB) Save(value interface{}) driver.OrmItf {
	return wrap(db.DB.Save(value))
}

func (db *gormDB) Updates(value interface{}) driver.OrmItf {
	return wrap(db.DB.Updates(value))
}

func (db *gormDB) Delete(value interface{}, conds ...interface{}) driver.OrmItf {
	return wrap(db.DB.Delete(value, conds...))
}

func (db *gormDB) Raw(sql string, values ...interface{}) driver.OrmItf {
	return wrap(db.DB.Raw(sql, values...))
}

// Find first matching record by given conditions, ordered by primary key
func (db *gormDB) First(dest interface{}, conds ...interface{}) driver.OrmItf {
	return wrap(db.DB.First(dest, conds...))
}

func (db *gormDB) Find(dest interface{}, conds ...interface{}) driver.OrmItf {
	return wrap(db.DB.Find(dest, conds...))
}

func (db *gormDB) Error() error {
	return db.DB.Error
}

func (db *gormDB) Close() error {
	realDB, err := db.DB.DB()
	if err != nil {
		return err
	}
	return realDB.Close()
}

func (db *gormDB) AutoMigrate(dst ...interface{}) error {
	return db.DB.AutoMigrate(dst...)
}

type (
	gormAssociation struct {
		*gorm.Association
	}
)

func wrapAssociation(association *gorm.Association) driver.OrmAssociationItf {
	return &gormAssociation{Association: association}
}

func (db *gormDB) Association(column string) driver.OrmAssociationItf {
	return wrapAssociation(db.DB.Association(column))
}

func (assoc *gormAssociation) Replace(values ...interface{}) error {
	return assoc.Association.Replace(values...)
}

func (assoc *gormAssociation) Find(out interface{}, conds ...interface{}) error {
	return assoc.Association.Find(out, conds...)
}
