package store


type Store interface {
	//文件上传
	//fileName 上传文件的路径
	//fileBytes 上传的内容
	Upload(fileName string, fileBytes []byte) error
	//文件下载
	//fileName 文件的路径
	//返回的内容为文件的二进制字节
	Get(fileName string) ([]byte, error)
	//文件列表
	//prefix 文件路径前缀
	//limit  显示条数
	List(prefix string, limit int) ([]string, error)
	//删除文件
	//fileName 文件路径
	Delete(fileName string) error
}