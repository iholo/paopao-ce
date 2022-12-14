// Copyright 2022 ROC. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package servants

import (
	"github.com/alimy/cfg"
	"github.com/gin-gonic/gin"
	"github.com/rocboss/paopao-ce/internal/servants/admin"
	"github.com/rocboss/paopao-ce/internal/servants/bot"
	"github.com/rocboss/paopao-ce/internal/servants/docs"
	"github.com/rocboss/paopao-ce/internal/servants/localoss"
	"github.com/rocboss/paopao-ce/internal/servants/space"
	"github.com/rocboss/paopao-ce/internal/servants/statick"
	"github.com/rocboss/paopao-ce/internal/servants/web"
)

// RegisterWebServants register all the servants to gin.Engine
func RegisterWebServants(e *gin.Engine) {
	docs.RegisterDocs(e)
	statick.RegisterStatick(e)

	cfg.Be("LocalOSS", func() {
		localoss.RouteLocalOSS(e)
	})

	web.RouteWeb(e)
}

// RegisterAdminServants register all the servants to gin.Engine
func RegisterAdminServants(e *gin.Engine) {
	admin.RouteManager(e)
}

// RegisterSpaceXServants register all the servants to gin.Engine
func RegisterSpaceXServants(e *gin.Engine) {
	space.RouteSpaceX(e)
}

// RegisterBotServants register all the servants to gin.Engine
func RegisterBotServants(e *gin.Engine) {
	bot.RouteBot(e)
}

// RegisterLocalossServants register all the servants to gin.Engine
func RegisterLocalossServants(e *gin.Engine) {
	localoss.RouteLocaloss(e)
}
