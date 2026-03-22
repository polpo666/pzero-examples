package cmd

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/polpo666/pzero/core/configcenter"
	"github.com/polpo666/pzero/core/configcenter/subscriber"
	"github.com/polpo666/pzero/core/swaggerv2"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"

	"simpleapi/internal/config"
	"simpleapi/internal/handler"
	"simpleapi/internal/middleware"
	"simpleapi/internal/svc"
	"simpleapi/plugins"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "simpleapi server",
	Long:  "simpleapi server",
	Run: func(cmd *cobra.Command, args []string) {
		cc := configcenter.MustNewConfigCenter[config.Config](configcenter.Config{
			Type: "yaml",
		}, subscriber.MustNewFsnotifySubscriber(cmd.Flag("config").Value.String(), subscriber.WithUseEnv(true)))

		// set up logger
		logx.Must(logx.SetUp(cc.MustGetConfig().Log.LogConf))

		// print banner
		printBanner(cc.MustGetConfig().Banner)
		// print version
		printVersion()

		// create service context
		svcCtx := svc.NewServiceContext(cc)

		// create rest server
		restServer := rest.MustNewServer(svcCtx.ConfigCenter.MustGetConfig().Rest.RestConf)

		// register auto generated routes
		handler.RegisterHandlers(restServer, svcCtx)
		// register swagger routes
		swaggerv2.RegisterRoutes(restServer)
		// register middleware
		middleware.Register(restServer)

		// load plugins
		plugins.LoadPlugins(restServer, svcCtx)

		group := service.NewServiceGroup()
		group.Add(restServer)

		logx.Infof("Starting rest server at %s:%d...", cc.MustGetConfig().Rest.Host, cc.MustGetConfig().Rest.Port)
		group.Start()
	},
}

func printBanner(c config.BannerConf) {
	figure.NewColorFigure(c.Text, c.FontName, c.Color, true).Print()
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
