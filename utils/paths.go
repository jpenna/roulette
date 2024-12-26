package utils

import "os"

const (
	DataDir = "data"

	TmpDir       = DataDir + "/tmp"
	TmpVerifyImg = TmpDir + "/ocr.jpeg"
	OutputDir    = DataDir + "/numbers"

	DebugImgFolder  = "debug"
	ResultImgFolder = "result"
	FailedImgFolder = "failed"
)

func init() {
	os.MkdirAll(TmpDir, 0755)
	os.MkdirAll(OutputDir, 0755)
	os.MkdirAll(DebugImgFolder, 0755)
	os.MkdirAll(ResultImgFolder, 0755)
	os.MkdirAll(FailedImgFolder, 0755)
}
