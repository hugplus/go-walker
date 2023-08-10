package migrate

import (
	"bytes"
	"fmt"
	"strconv"
	"text/template"
	"time"

	"github.com/hugplus/go-walker/common/utils"
	"github.com/hugplus/go-walker/core"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configYml string
	generate  bool
	goAdmin   bool
	host      string
	StartCmd  = &cobra.Command{
		Use:     "migrate",
		Short:   "Initialize the database",
		Example: "go-admin migrate -c config/settings.yml",
		Run: func(cmd *cobra.Command, args []string) {
			run()
		},
	}
)

// fixme 在您看不见代码的时候运行迁移，我觉得是不安全的，所以编译后最好不要去执行迁移
func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "resources/config.dev.yaml", "Start server with provided configuration file")
	StartCmd.PersistentFlags().BoolVarP(&generate, "generate", "g", false, "generate migration file")
	StartCmd.PersistentFlags().BoolVarP(&goAdmin, "goAdmin", "a", false, "generate go-admin migration file")
	StartCmd.PersistentFlags().StringVarP(&host, "domain", "d", "*", "select tenant host")
}

func run() {

	if !generate {
		fmt.Println(`start init`)
		//1. 读取配置

		v := viper.New()
		v.SetConfigFile(configYml)
		v.SetConfigType("yaml")
		err := v.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
		v.WatchConfig()

		if err = v.Unmarshal(&core.Cfg); err != nil {
			fmt.Println(err)
		}
		//3. 初始化数据库链接
		core.Init()
		//4. 数据库迁移
		fmt.Println("数据库迁移开始")
		//_ = migrateModel()
		fmt.Println(`数据库基础数据初始化成功`)

	} else {
		fmt.Println(`generate migration file`)
		_ = genFile()
	}
}

// func migrateModel() error {
// 	if host == "" {
// 		host = "master"
// 	}
// 	db := core.Db(host)

// 	//初始化数据库时候用
// 	db.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")

// 	// err := db.Debug().AutoMigrate(&models.Migration{})
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	migration.Migrate.SetDb(db.Debug())
// 	migration.Migrate.Migrate()
// 	return err
// }

func genFile() error {
	t1, err := template.ParseFiles("template/migrate.template")
	if err != nil {
		return err
	}
	m := map[string]string{}
	m["GenerateTime"] = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	m["Package"] = "version_local"
	if goAdmin {
		m["Package"] = "version"
	}
	var b1 bytes.Buffer
	err = t1.Execute(&b1, m)
	if goAdmin {
		utils.FileCreate(b1, "./cmd/migrate/migration/version/"+m["GenerateTime"]+"_migrate.go")
	} else {
		utils.FileCreate(b1, "./cmd/migrate/migration/version-local/"+m["GenerateTime"]+"_migrate.go")
	}
	return nil
}
