package driver

type (
	OrmItf interface {
		Select(query string, args ...interface{}) OrmItf
		Where(query interface{}, args ...interface{}) OrmItf
		Preload(query string, args ...interface{}) OrmItf
		Model(value interface{}) OrmItf
		Create(value interface{}) OrmItf
		Save(value interface{}) OrmItf
		Updates(value interface{}) OrmItf
		Delete(value interface{}, conds ...interface{}) OrmItf
		Raw(sql string, values ...interface{}) OrmItf
		First(dest interface{}, conds ...interface{}) OrmItf
		Find(dest interface{}, conds ...interface{}) OrmItf
		Error() error
		Close() error
		AutoMigrate(dst ...interface{}) error

		Association(column string) OrmAssociationItf
	}
	OrmAssociationItf interface {
		Replace(values ...interface{}) error
		Find(out interface{}, conds ...interface{}) error
	}
)
