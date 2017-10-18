package store


type Store interface {
	//文件上传
	//path 上传文件的路径
	//fileBytes 上传的内容
	Upload(path string, fileBytes []byte) error
	//文件下载
	//path 文件的路径
	//返回的内容为文件的二进制字节
	Get(path string) ([]byte, error)
	//文件列表
	//prefix 文件路径前缀
	//limit  显示条数
	List(prefix string, limit int) ([]string, error)
	//删除文件
	//path 文件路径
	Delete(path string) error
}