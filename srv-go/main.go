package main

import (
	"svr/app/cfg"
	"svr/app/util/logx"
	"svr/cmd/bbs/hdl"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
)

func init() {
	cfg.InitBase()
	//cfg.InitEmail()
	cfg.InitCache()
	cfg.InitDB()
}

func main() {
	engine := jet.New("./views", ".jet.html")
	if cfg.Debug {
		engine.Reload(true)
		engine.Debug(true)
	}

	app := fiber.New(fiber.Config{
		Views:       engine,
		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,
		Prefork:     true,
		//CaseSensitive: true,
		//StrictRouting: true,
		ErrorHandler: hdl.ErrInternal,
	})
	routers(app)
	logx.Fatal(app.Listen(cfg.PortBBS))
}

func routers(app *fiber.App) {
	app.Static("/assets", "./views/assets")
	app.Get("/robots.txt", hdl.Robots)
	app.Get("/sitemap.xml", hdl.Sitemap)

	app.Get("/", hdl.Index)
	app.Get("/sign", hdl.SignUp)
	app.Post("/api/sign/up", hdl.SignUp)
	app.Post("/api/sign/in", hdl.SignIn)

	app.Get("/user/profile", hdl.UserProfile)
	app.Get("/user/setting", hdl.UserSetting)
	app.Use(hdl.ErrNotFound)
}
