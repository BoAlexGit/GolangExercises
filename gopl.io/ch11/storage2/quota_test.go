// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
// quota_test2.go
// Тест, который заменяет отправление электронной почты
// функцией простого уведомления, которая просто регистрирует уведомляемого
//пользователя и содержимое сообщения:
//!+test
package storage

import (
	"strings"
	"testing"
)

func TestCheckQuotaNotifiesUser(t *testing.T) {
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}

	const user = "joe@example.org"
	usage[user] = 980000000 // simulate a 980MB-used condition

	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called")
	}
	if notifiedUser != user {
		t.Errorf("wrong user (%s) notified, want %s",
			notifiedUser, user)
	}
	const wantSubstring = "98% of your quota"
	if !strings.Contains(notifiedMsg, wantSubstring) {
		t.Errorf("unexpected notification message <<%s>>, "+
			"want substring %q", notifiedMsg, wantSubstring)
	}
}

//!-test

/*
// Имеется одна проблема: после возврата из тестовой функции CheckQuota
// больше не работает так, как надо, потому что все еще использует поддельную
// тестовуюреализацию notifyUsers. (При обновлении глобальных переменных
// всегда имеется риск такого рода.) Мы должны изменить тест, чтобы
// он восстанавливал предыдущее значение так, чтобы последующие тесты
// не наблюдали никакой замены, и должны сделать это на всех путях выполнения,
// включая сбои и аварийные ситуации в тестах.
// Это естественным образом приводит к применению defer.
//!+defer
func TestCheckQuotaNotifiesUser(t *testing.T) {
	// Save and restore original notifyUser.
	saved := notifyUser
	defer func() { notifyUser = saved }()

// Установка поддельной функции для notifyUser.
	// Install the test's fake notifyUser.
	var notifiedUser, notifiedMsg string
	notifyUser = func(user, msg string) {
		notifiedUser, notifiedMsg = user, msg
	}
	// ...rest of test...
}
//!-defer
*/
