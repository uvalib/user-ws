package logger

import (
    "log"
)

func Log( msg string ) {
    log.Printf( "USERINFO: %s", msg )
}