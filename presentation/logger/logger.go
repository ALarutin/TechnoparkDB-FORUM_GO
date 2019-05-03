package logger

import (
	"log"
	"os"
)

var (
	Info    *log.Logger
	Error   *log.Logger
	Trace   *log.Logger
	Debug   *log.Logger
	Warning *log.Logger
	Fatal   *log.Logger
)

var (
	file = "logger.log"
)

func init() {
	_file, err := os.Create(file)
	if err != nil {
		return
	}

	Info = log.New(_file,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Error = log.New(_file,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Trace = log.New(_file,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Debug = log.New(_file,
		"DEBUG: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Warning = log.New(_file,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)

	Fatal = log.New(_file,
		"FATAL: ",
		log.Ldate|log.Ltime|log.Lshortfile,
	)
}

/*В любой пакет нужно испортировать ""github.com/go-park-mail-ru/2019_1_SleeplessNights/log""
log.<log>.Println("commit")*/
