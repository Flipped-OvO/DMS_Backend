package util

import (
	"backend/util/setting"
	"log"
	"os"
	"path"
	"strconv"
)

func LogInitialization() {
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)

	dir := path.Dir(setting.Config.Runtime.LogPath)
	exists, err := PathExists(dir)
	if err != nil {
		log.Println(err)
	}
	if !exists {
		err = os.MkdirAll(dir, 0777)
		if err != nil {
			log.Println(err)
		}
	}
	file, err := os.OpenFile(setting.Config.Runtime.LogPath+"error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)

	log.SetOutput(file)
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


func FloatToString(num float64) string {
	return strconv.FormatFloat(num, 'f', 0, 64)
}

func ParseFloat(num string) float64 {
	result, _ := strconv.ParseFloat(num, 10)
	return result
}