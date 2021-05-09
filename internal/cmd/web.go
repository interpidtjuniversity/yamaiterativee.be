package cmd

import (
	"github.com/go-macaron/binding"
	"github.com/go-macaron/cache"
	"github.com/go-macaron/captcha"
	"github.com/go-macaron/gzip"
	"github.com/urfave/cli"
	"gopkg.in/macaron.v1"
	"yama.io/yamaIterativeE/internal/conf"
	"yama.io/yamaIterativeE/internal/context"
	"yama.io/yamaIterativeE/internal/form"
	"yama.io/yamaIterativeE/internal/home/application"
	"yama.io/yamaIterativeE/internal/home/iterations"
	"yama.io/yamaIterativeE/internal/home/server"
	"yama.io/yamaIterativeE/internal/home/workbench"
	"yama.io/yamaIterativeE/internal/iteration/env"
	"yama.io/yamaIterativeE/internal/iteration/pipeline"
	"yama.io/yamaIterativeE/internal/iteration/stage"
	"yama.io/yamaIterativeE/internal/registry/consul"
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
		stringFlag("config, c", "", "Changeable configuration file path"),
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
	// init global resource
	//resource.InitResource()
	// init application config(usage in create new application)
	//application.InitConfig()

	m := newMacaron()

	m.Group("", func() {
		m.Get("/", route.Home)

		m.Group("/home", func() {
			m.Group("/workbench", func() {
				m.Group("/newiteration", func() {
					m.Get("/allusers", workbench.GetAllOwners)
					m.Get("/ownerrepos/:ownerName", workbench.GetOwnerApplications)
					m.Post("/new", workbench.NewIteration)
				})
			})

			m.Group("/application", func() {
				m.Group("/newapplication", func() {
					m.Post("/new", binding.BindIgnErr(form.Application{}), application.NewApplication)
					m.Get("/allusers", application.GetAllUsers)

					m.Group("/optionconfig", func() {
						m.Get("/javaspring", application.GetJavaSpringConfig)
					})
				})
			})

			m.Group("/iterations", func() {
				m.Group("/user/:username", func() {
					m.Get("/all", iterations.GetUserAllIterations)
				})
			})

			m.Group("/server", func() {
				m.Group("/newserver", func() {
					m.Post("/new", server.NewServer)
				})
			})
		})

		m.Group("/iteration", func() {
			m.Group("/:iterationId", func() {
				m.Get("/info", env.IterInfo)
				m.Group("/envType", func() {
					m.Group("/:envType", func() {
						m.Get("", pipeline.IterPipelineInfo)
						m.Get("/info", env.IterEnvInfo)
					})
				})
				m.Group("/action", func() {
					m.Group("/envType", func() {
						m.Get("/:envType", env.IterActionInfo)
					})
					m.Group("/:actionId", func() {
						m.Get("/state", pipeline.IterActionState)
						m.Group("/stage", func() {
							m.Get("/:stageId", stage.IterStageInfo)
							m.Get("/:stageId/state", pipeline.IterStageState)
							m.Group("/:stageId/step", func() {
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

		m.Group("/v1", func() {
			m.Group("/status", func() {
				m.Get("/leader", consul.Leader)
			})

			m.Group("/agent", func() {
				m.Group("/service", func() {
					m.Put("/register", consul.Register)
					m.Put("/deregister/:service", consul.DeRegister)
				})
				m.Group("/check", func() {

				})
				m.Get("/self", consul.Ping)
			})

			m.Group("/catalog", func() {
				m.Get("/services", consul.GetServices)
			})

			m.Get("/health/service/:service", consul.GetServiceInstances)
		})

	},
		context.Contexter(),
	)

	m.Run()
	//err := http.ListenAndServe("192.168.0.102:6000", m)


	return nil
}