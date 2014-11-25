package main

import (
    "./golog"
)

func main () {
    log.Exec("notice", "hello")
    log.Exec("err", "fuck")
    log.Exec("warn", "fuck")
    log.Exec("info", "fuck")
}

