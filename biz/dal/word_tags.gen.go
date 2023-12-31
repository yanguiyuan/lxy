// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/yanguiyuan/lxy/biz/model/dbmodel"
)

func newWordTag(db *gorm.DB, opts ...gen.DOOption) wordTag {
	_wordTag := wordTag{}

	_wordTag.wordTagDo.UseDB(db, opts...)
	_wordTag.wordTagDo.UseModel(&dbmodel.WordTag{})

	tableName := _wordTag.wordTagDo.TableName()
	_wordTag.ALL = field.NewAsterisk(tableName)
	_wordTag.WordID = field.NewInt64(tableName, "word_id")
	_wordTag.TagID = field.NewInt64(tableName, "tag_id")

	_wordTag.fillFieldMap()

	return _wordTag
}

type wordTag struct {
	wordTagDo

	ALL    field.Asterisk
	WordID field.Int64
	TagID  field.Int64

	fieldMap map[string]field.Expr
}

func (w wordTag) Table(newTableName string) *wordTag {
	w.wordTagDo.UseTable(newTableName)
	return w.updateTableName(newTableName)
}

func (w wordTag) As(alias string) *wordTag {
	w.wordTagDo.DO = *(w.wordTagDo.As(alias).(*gen.DO))
	return w.updateTableName(alias)
}

func (w *wordTag) updateTableName(table string) *wordTag {
	w.ALL = field.NewAsterisk(table)
	w.WordID = field.NewInt64(table, "word_id")
	w.TagID = field.NewInt64(table, "tag_id")

	w.fillFieldMap()

	return w
}

func (w *wordTag) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := w.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (w *wordTag) fillFieldMap() {
	w.fieldMap = make(map[string]field.Expr, 2)
	w.fieldMap["word_id"] = w.WordID
	w.fieldMap["tag_id"] = w.TagID
}

func (w wordTag) clone(db *gorm.DB) wordTag {
	w.wordTagDo.ReplaceConnPool(db.Statement.ConnPool)
	return w
}

func (w wordTag) replaceDB(db *gorm.DB) wordTag {
	w.wordTagDo.ReplaceDB(db)
	return w
}

type wordTagDo struct{ gen.DO }

type IWordTagDo interface {
	gen.SubQuery
	Debug() IWordTagDo
	WithContext(ctx context.Context) IWordTagDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IWordTagDo
	WriteDB() IWordTagDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IWordTagDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IWordTagDo
	Not(conds ...gen.Condition) IWordTagDo
	Or(conds ...gen.Condition) IWordTagDo
	Select(conds ...field.Expr) IWordTagDo
	Where(conds ...gen.Condition) IWordTagDo
	Order(conds ...field.Expr) IWordTagDo
	Distinct(cols ...field.Expr) IWordTagDo
	Omit(cols ...field.Expr) IWordTagDo
	Join(table schema.Tabler, on ...field.Expr) IWordTagDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IWordTagDo
	RightJoin(table schema.Tabler, on ...field.Expr) IWordTagDo
	Group(cols ...field.Expr) IWordTagDo
	Having(conds ...gen.Condition) IWordTagDo
	Limit(limit int) IWordTagDo
	Offset(offset int) IWordTagDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IWordTagDo
	Unscoped() IWordTagDo
	Create(values ...*dbmodel.WordTag) error
	CreateInBatches(values []*dbmodel.WordTag, batchSize int) error
	Save(values ...*dbmodel.WordTag) error
	First() (*dbmodel.WordTag, error)
	Take() (*dbmodel.WordTag, error)
	Last() (*dbmodel.WordTag, error)
	Find() ([]*dbmodel.WordTag, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dbmodel.WordTag, err error)
	FindInBatches(result *[]*dbmodel.WordTag, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*dbmodel.WordTag) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IWordTagDo
	Assign(attrs ...field.AssignExpr) IWordTagDo
	Joins(fields ...field.RelationField) IWordTagDo
	Preload(fields ...field.RelationField) IWordTagDo
	FirstOrInit() (*dbmodel.WordTag, error)
	FirstOrCreate() (*dbmodel.WordTag, error)
	FindByPage(offset int, limit int) (result []*dbmodel.WordTag, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IWordTagDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (w wordTagDo) Debug() IWordTagDo {
	return w.withDO(w.DO.Debug())
}

func (w wordTagDo) WithContext(ctx context.Context) IWordTagDo {
	return w.withDO(w.DO.WithContext(ctx))
}

func (w wordTagDo) ReadDB() IWordTagDo {
	return w.Clauses(dbresolver.Read)
}

func (w wordTagDo) WriteDB() IWordTagDo {
	return w.Clauses(dbresolver.Write)
}

func (w wordTagDo) Session(config *gorm.Session) IWordTagDo {
	return w.withDO(w.DO.Session(config))
}

func (w wordTagDo) Clauses(conds ...clause.Expression) IWordTagDo {
	return w.withDO(w.DO.Clauses(conds...))
}

func (w wordTagDo) Returning(value interface{}, columns ...string) IWordTagDo {
	return w.withDO(w.DO.Returning(value, columns...))
}

func (w wordTagDo) Not(conds ...gen.Condition) IWordTagDo {
	return w.withDO(w.DO.Not(conds...))
}

func (w wordTagDo) Or(conds ...gen.Condition) IWordTagDo {
	return w.withDO(w.DO.Or(conds...))
}

func (w wordTagDo) Select(conds ...field.Expr) IWordTagDo {
	return w.withDO(w.DO.Select(conds...))
}

func (w wordTagDo) Where(conds ...gen.Condition) IWordTagDo {
	return w.withDO(w.DO.Where(conds...))
}

func (w wordTagDo) Order(conds ...field.Expr) IWordTagDo {
	return w.withDO(w.DO.Order(conds...))
}

func (w wordTagDo) Distinct(cols ...field.Expr) IWordTagDo {
	return w.withDO(w.DO.Distinct(cols...))
}

func (w wordTagDo) Omit(cols ...field.Expr) IWordTagDo {
	return w.withDO(w.DO.Omit(cols...))
}

func (w wordTagDo) Join(table schema.Tabler, on ...field.Expr) IWordTagDo {
	return w.withDO(w.DO.Join(table, on...))
}

func (w wordTagDo) LeftJoin(table schema.Tabler, on ...field.Expr) IWordTagDo {
	return w.withDO(w.DO.LeftJoin(table, on...))
}

func (w wordTagDo) RightJoin(table schema.Tabler, on ...field.Expr) IWordTagDo {
	return w.withDO(w.DO.RightJoin(table, on...))
}

func (w wordTagDo) Group(cols ...field.Expr) IWordTagDo {
	return w.withDO(w.DO.Group(cols...))
}

func (w wordTagDo) Having(conds ...gen.Condition) IWordTagDo {
	return w.withDO(w.DO.Having(conds...))
}

func (w wordTagDo) Limit(limit int) IWordTagDo {
	return w.withDO(w.DO.Limit(limit))
}

func (w wordTagDo) Offset(offset int) IWordTagDo {
	return w.withDO(w.DO.Offset(offset))
}

func (w wordTagDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IWordTagDo {
	return w.withDO(w.DO.Scopes(funcs...))
}

func (w wordTagDo) Unscoped() IWordTagDo {
	return w.withDO(w.DO.Unscoped())
}

func (w wordTagDo) Create(values ...*dbmodel.WordTag) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Create(values)
}

func (w wordTagDo) CreateInBatches(values []*dbmodel.WordTag, batchSize int) error {
	return w.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (w wordTagDo) Save(values ...*dbmodel.WordTag) error {
	if len(values) == 0 {
		return nil
	}
	return w.DO.Save(values)
}

func (w wordTagDo) First() (*dbmodel.WordTag, error) {
	if result, err := w.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodel.WordTag), nil
	}
}

func (w wordTagDo) Take() (*dbmodel.WordTag, error) {
	if result, err := w.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodel.WordTag), nil
	}
}

func (w wordTagDo) Last() (*dbmodel.WordTag, error) {
	if result, err := w.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodel.WordTag), nil
	}
}

func (w wordTagDo) Find() ([]*dbmodel.WordTag, error) {
	result, err := w.DO.Find()
	return result.([]*dbmodel.WordTag), err
}

func (w wordTagDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*dbmodel.WordTag, err error) {
	buf := make([]*dbmodel.WordTag, 0, batchSize)
	err = w.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (w wordTagDo) FindInBatches(result *[]*dbmodel.WordTag, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return w.DO.FindInBatches(result, batchSize, fc)
}

func (w wordTagDo) Attrs(attrs ...field.AssignExpr) IWordTagDo {
	return w.withDO(w.DO.Attrs(attrs...))
}

func (w wordTagDo) Assign(attrs ...field.AssignExpr) IWordTagDo {
	return w.withDO(w.DO.Assign(attrs...))
}

func (w wordTagDo) Joins(fields ...field.RelationField) IWordTagDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Joins(_f))
	}
	return &w
}

func (w wordTagDo) Preload(fields ...field.RelationField) IWordTagDo {
	for _, _f := range fields {
		w = *w.withDO(w.DO.Preload(_f))
	}
	return &w
}

func (w wordTagDo) FirstOrInit() (*dbmodel.WordTag, error) {
	if result, err := w.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodel.WordTag), nil
	}
}

func (w wordTagDo) FirstOrCreate() (*dbmodel.WordTag, error) {
	if result, err := w.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*dbmodel.WordTag), nil
	}
}

func (w wordTagDo) FindByPage(offset int, limit int) (result []*dbmodel.WordTag, count int64, err error) {
	result, err = w.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = w.Offset(-1).Limit(-1).Count()
	return
}

func (w wordTagDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = w.Count()
	if err != nil {
		return
	}

	err = w.Offset(offset).Limit(limit).Scan(result)
	return
}

func (w wordTagDo) Scan(result interface{}) (err error) {
	return w.DO.Scan(result)
}

func (w wordTagDo) Delete(models ...*dbmodel.WordTag) (result gen.ResultInfo, err error) {
	return w.DO.Delete(models)
}

func (w *wordTagDo) withDO(do gen.Dao) *wordTagDo {
	w.DO = *do.(*gen.DO)
	return w
}
