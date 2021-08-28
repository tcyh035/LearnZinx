package yiface

type IRequest interface {
	GetConnection() IConnection

	GetData() []byte
}
