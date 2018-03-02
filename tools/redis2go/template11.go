package main

const template11 string = `/// -------------------------------------------------------------------------------
/// THIS FILE IS ORIGINALLY GENERATED BY redis2go.exe.
/// PLEASE DO NOT MODIFY THIS FILE.
/// -------------------------------------------------------------------------------

package {{packagename}}

import (
	"errors"
	"fmt"
	go_redis_orm "github.com/fananchong/go-redis-orm.v2"
	"github.com/garyburd/redigo/redis"
)

type RD_{{classname}} struct {
	Key {{key_type}} {{rediskey}}
	{{fields_def}}

    __dirtyData map[string]interface{}
	__isLoad bool
	__dbKey string
}

func NewRD_{{classname}}(key {{key_type}}) *RD_{{classname}} {
	return &RD_{{classname}} {
		Key: key,
		__dbKey: {{func_dbkey}},
		__dirtyData: make(map[string]interface{}),
	}
}

func (this *RD_{{classname}}) Load(dbName string) error {
	if this.__isLoad == true {
		return errors.New("alreay load!")
	}
	db := go_redis_orm.GetDB(dbName)
	val, err := redis.Values(db.Do("HGETALL", this.__dbKey))
	if err != nil {
		return err
	}
	if err := redis.ScanStruct(val, this); err != nil {
		return err
	}
	this.__isLoad = true
	return nil
}

func (this *RD_{{classname}}) Save(dbName string) error {
	if len(this.__dirtyData) == 0 {
		return nil
	}
	db := go_redis_orm.GetDB(dbName)
	if _, err := db.Do("HMSET", redis.Args{}.Add(this.__dbKey).AddFlat(this.__dirtyData)...); err != nil {
    	return err
	}
	this.__dirtyData = make(map[string]interface{})
	return nil
}

func (this *RD_{{classname}}) Delete(dbName string) error {
	db := go_redis_orm.GetDB(dbName)
	_, err := db.Do("DEL", this.__dbKey)
	return err
}

func (this *RD_{{classname}}) IsLoad() bool {
	return this.__isLoad
}

{{func_get}}

{{func_set}}`

const getFuncString = `func (this *RD_{{classname}}) Get{{field_name_upper}}() {{field_type}} {
	return this.{{field_name_lower}}
}`

const setFuncString = `func (this *RD_{{classname}}) Set{{field_name_upper}}(value {{field_type}}) {
	this.{{field_name_lower}} = value
	this.__dirtyData["{{field_name_lower}}"] = value
}`

const dbkeyFuncString_int = `"{{classname}}:" + fmt.Sprintf("%d", key)`

const dbkeyFuncString_str = `"{{classname}}:" + key`
