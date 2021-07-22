package goseaweed

type Seaweed interface {
	UploadFile(objectName string, content []byte)(string, error)
	GetFile(objectName string) ([]byte, error)
	RemoveFile(objectName string) (error,int)
}