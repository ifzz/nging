// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"

	"time"
)

type Slice_DbSync []*DbSync

func (s Slice_DbSync) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_DbSync) RangeRaw(fn func(m *DbSync) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

// DbSync 数据表同步方案
type DbSync struct {
	base    factory.Base
	objects []*DbSync

	Id                   uint   `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Name                 string `db:"name" bson:"name" comment:"方案名" json:"name" xml:"name"`
	SourceAccountId      uint   `db:"source_account_id" bson:"source_account_id" comment:"源数据库账号ID" json:"source_account_id" xml:"source_account_id"`
	DsnSource            string `db:"dsn_source" bson:"dsn_source" comment:"同步源" json:"dsn_source" xml:"dsn_source"`
	DestinationAccountId uint   `db:"destination_account_id" bson:"destination_account_id" comment:"目标数据库账号ID" json:"destination_account_id" xml:"destination_account_id"`
	DsnDestination       string `db:"dsn_destination" bson:"dsn_destination" comment:"目标数据库" json:"dsn_destination" xml:"dsn_destination"`
	Tables               string `db:"tables" bson:"tables" comment:"要同步的表" json:"tables" xml:"tables"`
	SkipTables           string `db:"skip_tables" bson:"skip_tables" comment:"要跳过的表" json:"skip_tables" xml:"skip_tables"`
	AlterIgnore          string `db:"alter_ignore" bson:"alter_ignore" comment:"要忽略的列、索引、外键" json:"alter_ignore" xml:"alter_ignore"`
	Drop                 uint   `db:"drop" bson:"drop" comment:"删除待同步数据库中多余的字段、索引、外键 " json:"drop" xml:"drop"`
	MailTo               string `db:"mail_to" bson:"mail_to" comment:"发送邮件" json:"mail_to" xml:"mail_to"`
	Created              uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	Updated              int    `db:"updated" bson:"updated" comment:"更新时间" json:"updated" xml:"updated"`
}

// - base function

func (a *DbSync) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *DbSync) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *DbSync) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *DbSync) Context() echo.Context {
	return a.base.Context()
}

func (a *DbSync) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *DbSync) SetNamer(namer func(string) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *DbSync) Namer() func(string) string {
	return a.base.Namer()
}

func (a *DbSync) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *DbSync) Param() *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam()
	}
	return a.base.Param()
}

// - current function

func (a *DbSync) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *DbSync) Objects() []*DbSync {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *DbSync) NewObjects() factory.Ranger {
	return &Slice_DbSync{}
}

func (a *DbSync) InitObjects() *[]*DbSync {
	a.objects = []*DbSync{}
	return &a.objects
}

func (a *DbSync) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *DbSync) Short_() string {
	return "db_sync"
}

func (a *DbSync) Struct_() string {
	return "DbSync"
}

func (a *DbSync) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a.Short_()))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *DbSync) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *DbSync) Get(mw func(db.Result) db.Result, args ...interface{}) error {
	base := a.base
	err := a.Param().SetArgs(args...).SetRecv(a).SetMiddleware(mw).One()
	a.base = base
	return err
}

func (a *DbSync) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param().SetArgs(args...).SetPage(page).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (a *DbSync) GroupBy(keyField string, inputRows ...[]*DbSync) map[string][]*DbSync {
	var rows []*DbSync
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string][]*DbSync{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*DbSync{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (a *DbSync) KeyBy(keyField string, inputRows ...[]*DbSync) map[string]*DbSync {
	var rows []*DbSync
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]*DbSync{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (a *DbSync) AsKV(keyField string, valueField string, inputRows ...[]*DbSync) map[string]interface{} {
	var rows []*DbSync
	if len(inputRows) > 0 {
		rows = inputRows[0]
	} else {
		rows = a.Objects()
	}
	r := map[string]interface{}{}
	for _, row := range rows {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (a *DbSync) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	return a.Param().SetArgs(args...).SetOffset(offset).SetSize(size).SetRecv(recv).SetMiddleware(mw).List()
}

func (a *DbSync) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	err = DBI.Fire("creating", a, nil)
	if err != nil {
		return
	}
	pk, err = a.Param().SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *DbSync) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	a.Updated = int(time.Now().Unix())
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Setter(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *DbSync) Setter(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	return a.Param().SetArgs(args...).SetMiddleware(mw)
}

func (a *DbSync) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *DbSync) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {
	kvset["updated"] = int(time.Now().Unix())
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Setter(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *DbSync) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param().SetArgs(args...).SetSend(a).SetMiddleware(mw).Upsert(func() error {
		a.Updated = int(time.Now().Unix())
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint(v)
		}
	}
	if err == nil {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *DbSync) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param().SetArgs(args...).SetMiddleware(mw).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *DbSync) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param().SetArgs(args...).SetMiddleware(mw).Count()
}

func (a *DbSync) Reset() *DbSync {
	a.Id = 0
	a.Name = ``
	a.SourceAccountId = 0
	a.DsnSource = ``
	a.DestinationAccountId = 0
	a.DsnDestination = ``
	a.Tables = ``
	a.SkipTables = ``
	a.AlterIgnore = ``
	a.Drop = 0
	a.MailTo = ``
	a.Created = 0
	a.Updated = 0
	return a
}

func (a *DbSync) AsMap() map[string]interface{} {
	r := map[string]interface{}{}
	r["Id"] = a.Id
	r["Name"] = a.Name
	r["SourceAccountId"] = a.SourceAccountId
	r["DsnSource"] = a.DsnSource
	r["DestinationAccountId"] = a.DestinationAccountId
	r["DsnDestination"] = a.DsnDestination
	r["Tables"] = a.Tables
	r["SkipTables"] = a.SkipTables
	r["AlterIgnore"] = a.AlterIgnore
	r["Drop"] = a.Drop
	r["MailTo"] = a.MailTo
	r["Created"] = a.Created
	r["Updated"] = a.Updated
	return r
}

func (a *DbSync) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint(value)
		case "name":
			a.Name = param.AsString(value)
		case "source_account_id":
			a.SourceAccountId = param.AsUint(value)
		case "dsn_source":
			a.DsnSource = param.AsString(value)
		case "destination_account_id":
			a.DestinationAccountId = param.AsUint(value)
		case "dsn_destination":
			a.DsnDestination = param.AsString(value)
		case "tables":
			a.Tables = param.AsString(value)
		case "skip_tables":
			a.SkipTables = param.AsString(value)
		case "alter_ignore":
			a.AlterIgnore = param.AsString(value)
		case "drop":
			a.Drop = param.AsUint(value)
		case "mail_to":
			a.MailTo = param.AsString(value)
		case "created":
			a.Created = param.AsUint(value)
		case "updated":
			a.Updated = param.AsInt(value)
		}
	}
}

func (a *DbSync) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint(vv)
		case "Name":
			a.Name = param.AsString(vv)
		case "SourceAccountId":
			a.SourceAccountId = param.AsUint(vv)
		case "DsnSource":
			a.DsnSource = param.AsString(vv)
		case "DestinationAccountId":
			a.DestinationAccountId = param.AsUint(vv)
		case "DsnDestination":
			a.DsnDestination = param.AsString(vv)
		case "Tables":
			a.Tables = param.AsString(vv)
		case "SkipTables":
			a.SkipTables = param.AsString(vv)
		case "AlterIgnore":
			a.AlterIgnore = param.AsString(vv)
		case "Drop":
			a.Drop = param.AsUint(vv)
		case "MailTo":
			a.MailTo = param.AsString(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "Updated":
			a.Updated = param.AsInt(vv)
		}
	}
}

func (a *DbSync) AsRow() map[string]interface{} {
	r := map[string]interface{}{}
	r["id"] = a.Id
	r["name"] = a.Name
	r["source_account_id"] = a.SourceAccountId
	r["dsn_source"] = a.DsnSource
	r["destination_account_id"] = a.DestinationAccountId
	r["dsn_destination"] = a.DsnDestination
	r["tables"] = a.Tables
	r["skip_tables"] = a.SkipTables
	r["alter_ignore"] = a.AlterIgnore
	r["drop"] = a.Drop
	r["mail_to"] = a.MailTo
	r["created"] = a.Created
	r["updated"] = a.Updated
	return r
}

func (a *DbSync) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return factory.BatchValidate(a.Short_(), kvset)
}

func (a *DbSync) Validate(field string, value interface{}) error {
	return factory.Validate(a.Short_(), field, value)
}
