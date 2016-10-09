package dglogs

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"time"
)

func getToday() (year, month, day int) {
	t := time.Now()
	year = t.Year()
	month = int(t.Month())
	day = t.Day()

	return
}

func getLogFileName() (fname string) {
	year, month, day := getToday()
	fname = fmt.Sprintf("./logs/%s%d%02d%02d.log", beego.AppConfig.String("appname"), year, month, day)

	return
}

// logs
func WriteCriticalLog(format string, v ...interface{}) {
	log := logs.NewLogger(10000)
	fname := getLogFileName()
	sfile := fmt.Sprintf("{\"filename\":\"%s\"}", fname)
	log.SetLogger("file", sfile)

	log.Critical(format, v...)
}

func WriteErrorLog(format string, v ...interface{}) {
	log := logs.NewLogger(10000)
	fname := getLogFileName()
	sfile := fmt.Sprintf("{\"filename\":\"%s\"}", fname)
	//log.SetLogger("file", `{"filename":fname}`)
	log.SetLogger("file", sfile)

	log.Error(format, v...)
}

func WriteDebugLog(format string, v ...interface{}) {
	log := logs.NewLogger(10000)
	fname := getLogFileName()
	sfile := fmt.Sprintf("{\"filename\":\"%s\"}", fname)
	//log.SetLogger("file", `{"filename":fname}`)
	log.SetLogger("file", sfile)

	log.Debug(format, v...)
}

func WriteInfoLog(format string, v ...interface{}) {
	log := logs.NewLogger(10000)
	fname := getLogFileName()
	sfile := fmt.Sprintf("{\"filename\":\"%s\"}", fname)
	//log.SetLogger("file", `{"filename":fname}`)
	log.SetLogger("file", sfile)

	log.Info(format, v...)
}