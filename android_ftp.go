package main

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"goftp.io/server/core"
	"goftp.io/server/driver/file"
	"log"
	"os"
	"path/filepath"
	"time"
)

type Result1 struct {
	Code    int    `json:"code"`
	Message string `json:"msg"`
}

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
			Name:     "root", // FTP 账号
			Password: "root", // FTP 密码
		},
		Port: Port,
	}
	// start ftp server
	s := core.NewServer(opt)
	var n Notification
	s.RegisterNotifer(n)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}

type Notification struct {
}

func (n Notification) BeforeLoginUser(conn *core.Conn, userName string) {
	fmt.Println("BeforeLoginUser")
}

func (n Notification) BeforePutFile(conn *core.Conn, dstPath string) {
	fmt.Println("BeforePutFile")
}

func (n Notification) BeforeDeleteFile(conn *core.Conn, dstPath string) {
	fmt.Println("BeforeDeleteFile")
}

func (n Notification) BeforeChangeCurDir(conn *core.Conn, oldCurDir, newCurDir string) {
	fmt.Println("BeforeChangeCurDir")
}

func (n Notification) BeforeCreateDir(conn *core.Conn, dstPath string) {
	fmt.Println("BeforeCreateDir")
}

func (n Notification) BeforeDeleteDir(conn *core.Conn, dstPath string) {
	fmt.Println("BeforeDeleteDir")
}

func (n Notification) BeforeDownloadFile(conn *core.Conn, dstPath string) {
	fmt.Println("BeforeDownloadFile")
}

func (n Notification) AfterUserLogin(conn *core.Conn, userName, password string, passMatched bool, err error) {
	fmt.Println("AfterUserLogin")
}

func (n Notification) AfterFileDeleted(conn *core.Conn, dstPath string, err error) {
	fmt.Println("AfterFileDeleted")
}

func (n Notification) AfterFileDownloaded(conn *core.Conn, dstPath string, size int64, err error) {
	fmt.Println("AfterFileDownloaded")
}

func (n Notification) AfterCurDirChanged(conn *core.Conn, oldCurDir, newCurDir string, err error) {
	fmt.Println("AfterCurDirChanged")
}

func (n Notification) AfterDirCreated(conn *core.Conn, dstPath string, err error) {
	fmt.Println("AfterDirCreated")
}

func (n Notification) AfterDirDeleted(conn *core.Conn, dstPath string, err error) {
	fmt.Println("AfterDirDeleted")
}

func (n Notification) AfterFilePut(conn *core.Conn, dstPath string, size int64, err error) {
	fmt.Println("AfterFilePut")
	//con, err := net.Dial("tcp4", "127.0.0.1:26001")
	//if err != nil {
	//	fmt.Println("net dial err: ", err)
	//	return
	//}
	//defer con.Close()
	//var res Result1
	//res.Code = 1
	//res.Message = "文件传输完成"
	//
	//data, err := json.Marshal(res)
	//if err != nil {
	//	fmt.Println("json err: ", err)
	//	return
	//}
	//_, err = con.Write(data)
	//if err != nil {
	//	fmt.Println("json err: ", err)
	//	return
	//}

	c, err := ftp.Dial("10.0.0.3:21", ftp.DialWithTimeout(5*time.Second))
	if err != nil {
		fmt.Println("ftp.Dial: ", err)
		return
	}
	error := c.Login("root", "root")
	if error != nil {
		fmt.Println("ftp.Dial: ", error)
		return
	}

	fmt.Println("os.Open: ", dstPath)
	open, err := os.Open("./static" + dstPath)
	if err != nil {
		fmt.Println("os.Open: ", error)
		return
	}
	fmt.Println(open.Name())
	_, f := filepath.Split(open.Name())
	err1 := c.Stor(f, open)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
}
