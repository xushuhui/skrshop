package main

import (
	"flag"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	"gopkg.in/yaml.v2"

	pb "skrshop/api/helloworld/v1"
	"skrshop/internal/conf"

	"skrshop/internal/service"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
	// flagconf is the config flag.
	flagconf string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
}

func newApp(logger log.Logger, hs *http.Server, gs *grpc.Server, greeter *service.GreeterService) *kratos.App {
	pb.RegisterGreeterServer(gs, greeter)
	pb.RegisterGreeterHTTPServer(hs, greeter)
	return kratos.New(
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
			gs,
		),
	)
}

func main() {

	flag.Parse()
	logger := log.NewStdLogger()

	config := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
		config.WithDecoder(func(kv *config.KeyValue, v map[string]interface{}) error {
			return yaml.Unmarshal(kv.Value, v)
		}),
	)

	if err := config.Load(); err != nil {
		panic(err)
	}

	var bc conf.Bootstrap
	if err := config.Scan(&bc); err != nil {
		panic(err)
	}

	app, err := initApp(bc.Server, bc.Data, logger)
	if err != nil {
		panic(err)
	}

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
func initDb() {
	//client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	//if err != nil {
	//	l.Fatalf("failed opening connection to sqlite: %v", err)
	//}
	//defer client.Close()
	//// Run the auto migration tool.
	//if err := client.Schema.Create(context.Background()); err != nil {
	//	l.Fatalf("failed creating schema resources: %v", err)
	//}
}
