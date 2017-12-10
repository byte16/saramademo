package main

import (
	"github.com/pkg/errors"
)

//1. FetchAccount(id) 从Redis读取账户实例信息
//2. updateAccount(id, data) 更新指定账户信息
//3. ToAccount(map) 将从Redis读到的账户信息转换为模型数据，return *BankAccount object.

type BankAccount struct {
	Id      string
	Name    string
	Balance int
}

func FetchAccount(id string) {

}

func updateAccount(id string, data map[string]interface{}) (*BankAccount, error) {
	if data == nil {
		return nil, errors.New("unable to update accout with nil data map.")
	}
	name, _ := data["Name"].(string)
	balance, _ := data["Balance"].(int)

	return &BankAccount{
		Id:      id,
		Name:    name,
		Balance: balance,
	}, nil
}

func ToAccount(data map[string]interface{}) {

}
