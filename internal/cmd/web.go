package cmd

import (
	"github.com/go-macaron/cache"
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/gzip"
	"github.com/urfave/cli"
	"gopkg.in/macaron.v1"
	"yama.io/yamaIterativeE/internal/conf"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/iteration/pipeline"
	"yama.io/yamaIterativeE/internal/iteration/stage"
	"yama.io/yamaIterativeE/internal/route"
)

var Web = cli.Command{
	Name:  "web",
	Usage: "Start web server",
	Description: `yamaIterativeE web server is the only thing you need to run,
and it takes care of all the other things for you`,
	Action: runWeb,
	Flags: []cli.Flag{
		stringFlag("port, p", "6000", "Temporary port number to prevent conflict"),
		stringFlag("config, c", "", "Custom configuration file path"),
	},
}

// newMacaron initializes Macaron instance.
func newMacaron() *macaron.Macaron {
	m := macaron.New()
	if !conf.Server.DisableRouterLog {
		m.Use(macaron.Logger())
	}
	m.Use(macaron.Recovery())
	if conf.Server.EnableGzip {
		m.Use(gzip.Gziper())
	}
	// Register custom middleware first to make it possible to override files under "public".
	m.Use(macaron.Static(
		"public",
		macaron.StaticOptions{
			SkipLogging: conf.Server.DisableRouterLog,
		},
	))

	renderOpt := macaron.RenderOptions{
		Directory:         "templates",
		IndentJSON:        macaron.Env != macaron.PROD,
	}
	m.Use(macaron.Renderer(renderOpt))

	m.Use(cache.Cacher(cache.Options{
		Adapter:       conf.Cache.Adapter,
		AdapterConfig: conf.Cache.Host,
		Interval:      conf.Cache.Interval,
	}))
	m.Use(captcha.Captchaer(captcha.Options{
		SubURL: conf.Server.Subpath,
	}))

	return m
}

func runWeb(c *cli.Context) error {
	// init database
	route.GlobalInit("")

	m := newMacaron()

	m.Group("", func() {
		m.Get("/", route.Home)

		m.Group("/iteration", func() {
			m.Get("/ping", stage.Ping)
			m.Group("/:iterationId", func() {
				m.Get("/info", stage.IterInfo)
				m.Group("/envType", func() {
					m.Group("/:envType", func() {
						m.Get("", pipeline.IterPipelineInfo)
						m.Get("/info", stage.IterEnvInfo)
					})
				})
				m.Group("/action", func() {
					m.Group("/envType", func() {
						m.Get("/:envType", stage.IterActionInfo)
					})
					m.Group("/:actionId", func() {
						m.Group("/stage", func() {
							m.Get("", pipeline.IterActionState)
							m.Get("/:stageId", stage.IterStageInfo)
							m.Get("/:stageId/state", pipeline.IterStageState)
							m.Group("/step", func() {
								m.Get("/:stepId", stage.StageExecInfo)
								m.Get("/:stepId/state", pipeline.IterStepState)
							})
						})
					})
					m.Group("/new", func() {
						m.Get("/:pipelineId", pipeline.StartPipeline)
					})
				})

			})
		})



	},
		context.Contexter(),
	)

	m.Run()
	//err := http.ListenAndServe("192.168.0.102:6000", m)


	return nil
}