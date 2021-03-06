/*
 * Copyright (c) 2018. David Hanson. All Rights Reserved.
 *
 */

package registry

import (
	"github.com/djdoeslinux/choreobot/command/counter"
	"github.com/djdoeslinux/choreobot/command/scoreboard"
	"github.com/djdoeslinux/choreobot/command/turing_test"
	"github.com/djdoeslinux/choreobot/core"
	"github.com/djdoeslinux/choreobot/meter"
	"github.com/djdoeslinux/choreobot/moderator"
	"github.com/djdoeslinux/choreobot/straw_poll"
	"github.com/djdoeslinux/choreobot/user"
	"github.com/jinzhu/gorm"
)

func AutoMigrate(db *gorm.DB) {
	var models []interface{}
	models = append(models, core.Models...)
	models = append(models, counter.Models...)
	models = append(models, scoreboard.Models...)
	models = append(models, meter.Models...)
	models = append(models, moderator.Models...)
	models = append(models, straw_poll.Models...)
	models = append(models, turing_test.Models...)
	models = append(models, user.Models...)

	db.AutoMigrate(models...)
}
