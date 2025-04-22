package main

import (
	_ "taylor-music-back/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"taylor-music-back/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
