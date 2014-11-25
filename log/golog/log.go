package log 

import (
    "fmt"
)

type log struct{
    width string
}


func (l *log) Notice(message string, args interface{}) {
    p := fmt.Sprintf("\033[96m%"+l.width+"s \033[90m:\033[0m ", message)
    fmt.Println(p, args)
}

func (l *log) Warn(message string, args interface{}) {
    p := fmt.Sprintf("\033[33m%"+l.width+"s \033[90m:\033[0m ", message)
    fmt.Println(p, args)
}

func (l *log) Err(message string, args interface{}) {
    p := fmt.Sprintf("\033[31m%"+l.width+"s \033[90m:\033[0m ", message)
    fmt.Println(p, args)
}

func (l *log) Debug(message string, args interface{}) {
    p := fmt.Sprintf("\033[36m%"+l.width+"s \033[90m:\033[0m ", message)
    fmt.Println(p, args)
}

func Exec(message string, args interface{}) {
    l := log{width:"14"}
    switch message{
        case "notice":
            l.Notice(message, args)
        case "warn":
            l.Warn(message, args)
        case "err":
            l.Err(message, args)
        case "debug":
        default:
            l.Debug(message, args)
    }
}
