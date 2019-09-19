// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// storage1.go
// See page 311.

// Package storage is part of a hypothetical cloud storage server.
//!+main
package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

var usage = make(map[string]int64)

func bytesInUse(username string) int64 { return usage[username] }
// Настройка отправителя электронных писем.
// Примечание: никогда не помещайте пароль в исходный текст!
// Email sender configuration.
// NOTE: never put passwords in source code!
const sender = "notifications@example.com"
const password = "correcthorsebatterystaple"
const hostname = "smtp.example.com"
// Внимание, вы использовали %d байт хранилища,
//%d%% вашей квоты.
const template = `Warning: you are using %d bytes of storage,
%d%% of your quota.`


func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000 // 1GB
	percent := 100 * used / quota
	if percent < 90 {
		return // OK
	}
	msg := fmt.Sprintf(template, used, percent)
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":587", auth, sender,
		[]string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s) failed: %s", username, err)
	}
}

//!-main
