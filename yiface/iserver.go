package yiface

type IServer interface {
    Start()
    Stop()
    Serve()
}