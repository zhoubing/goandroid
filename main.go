package main

import (
	"log"

	"goftp.io/server/core"
	"goftp.io/server/driver/file"
)

func main() {
	Name := "FTP Server"
	rootPath := "./static" //FTP根目录
	Port := 2121           //FTP 端口
	var perm = core.NewSimplePerm("test", "test")

	// Server options without hostname or port
	opt := &core.ServerOpts{
		Name: Name,
		Factory: &file.DriverFactory{
			RootPath: rootPath,
			Perm:     perm,
		},
		Auth: &core.SimpleAuth{
			Name:     "username", // FTP 账号
			Password: "Password", // FTP 密码
		},
		Port: Port,
	}
	// start ftp server
	s := core.NewServer(opt)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
