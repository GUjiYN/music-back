package main

import (
	_ "taylor-music-back/internal/packed"

	_ "taylor-music-back/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"taylor-music-back/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
