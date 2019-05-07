package main

/**
一个类只需要实现了接口需要的所有函数，则就说这个类实现了该接口
 */
type File struct {
	//。。。
}

func (f *File) Read(buf []byte) (n int, err error) {
	return 0, nil
}

func (f *File) Write(buf []byte) (n int, err error) {
	return 0, nil
}

func (f *File) Seek(off int64, whence int) (pos int64, err error) {
	return 0, nil
}
func (f *File) Close() error {
	return nil
}

type IFile interface {
	Read(buf []byte) (n int, err error)
	Write(buf []byte) (n int, err error)
	Seek(off int64, whence int) (pos int64, err error)
	Close() error
}

type IReader interface {
	Read(buf []byte) (n int, err error)
}

type IWrite interface {
	Write(buf []byte) (n int, err error)
}
type IClose interface {
	Close() error
}

/**
津管File类并没有从这些接口继承，甚至不用不知道这些接口的存在，但是File类实现了这些接口，就可以进行赋值
 */
var file1 IFile = new(File)
var file2 IReader = new(File)
var file3 IWrite = new(File)
var file4 IClose = new(File)

//接口查询
func main() {
	/*if file2, ok := file1.(*File); ok {

	}*/
}
