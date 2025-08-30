package program

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/mahibulhaque/gofast/flags"
	"github.com/mahibulhaque/gofast/template/dbdriver"
	"github.com/mahibulhaque/gofast/template/docker"
	"github.com/mahibulhaque/gofast/template/framework"
	"github.com/mahibulhaque/gofast/utils"
)

type Project struct {
	ProjectName       string
	Exit              bool
	AbsolutePath      string
	ProjectType       flags.Framework
	DBDriver          flags.Database
	Docker            flags.Database
	FrameworkMap      map[flags.Framework]Framework
	DBDriverMap       map[flags.Database]DBDriver
	DockerMap         map[flags.Database]Docker
	AdvancedOptions   map[string]bool
	AdvancedTemplates AdvancedTemplates
	GitOptions        flags.Git
	OSCheck           map[string]bool
}

type AdvancedTemplates struct {
	TemplateRoutes  string
	TemplateImports string
}

type Framework struct {
	packageName []string
	templater   Templater
}

type DBDriver struct {
	packageName []string
	templater   DBDriverTemplater
}

type Docker struct {
	packageName []string
	templater   DockerTemplater
}

type Templater interface {
	Main() []byte
	Server() []byte
	Routes() []byte
	TestHandler() []byte
	WebsocketImports() []byte
}

type DBDriverTemplater interface {
	Service() []byte
	Env() []byte
	Tests() []byte
}

type DockerTemplater interface {
	Docker() []byte
}

type WorkflowTemplater interface {
	Releaser() []byte
	Test() []byte
	ReleaserConfig() []byte
}

var (
	chiPackage     = []string{"github.com/go-chi/chi/v5"}
	gorillaPackage = []string{"github.com/gorilla/mux"}
	routerPackage  = []string{"github.com/julienschmidt/httprouter"}
	ginPackage     = []string{"github.com/gin-gonic/gin"}
	fiberPackage   = []string{"github.com/gofiber/fiber/v2"}
	echoPackage    = []string{"github.com/labstack/echo/v4", "github.com/labstack/echo/v4/middleware"}

	mysqlDriver    = []string{"github.com/go-sql-driver/mysql"}
	postgresDriver = []string{"github.com/jackc/pgx/v5/stdlib"}
	sqliteDriver   = []string{"github.com/mattn/go-sqlite3"}
	redisDriver    = []string{"github.com/redis/go-redis/v9"}
	mongoDriver    = []string{"go.mongodb.org/mongo-driver"}

	godotenvPackage = []string{"github.com/joho/godotenv"}
)

const (
	root                 = "/"
	cmdApiPath           = "cmd/api"
	cmdWebPath           = "cmd/web"
	internalServerPath   = "internal/server"
	internalDatabasePath = "internal/db"
	githubActionPath     = ".github/workflows"
)

func (p *Project) CheckOS() {
	p.OSCheck = make(map[string]bool)

	if runtime.GOOS != "windows" {
		p.OSCheck["UnixBased"] = true
	}
	if runtime.GOOS == "linux" {
		p.OSCheck["linux"] = true
	}
	if runtime.GOOS == "darwin" {
		p.OSCheck["darwin"] = true
	}
}

func (p *Project) ExitCLI(tprogram *tea.Program) {
	if p.Exit {
		// logo render here
		if err := tprogram.ReleaseTerminal(); err != nil {
			log.Fatal(err)
		}
		os.Exit(1)
	}
}

func (p *Project) createFrameworkMap() {
	p.FrameworkMap[flags.Chi] = Framework{
		packageName: chiPackage,
		templater:   framework.ChiTemplates{},
	}
	p.FrameworkMap[flags.StandardLibrary] = Framework{
		packageName: []string{},
		templater:   framework.StandardLibTemplate{},
	}

	p.FrameworkMap[flags.Gin] = Framework{
		packageName: ginPackage,
		templater:   framework.GinTemplates{},
	}

	p.FrameworkMap[flags.Fiber] = Framework{
		packageName: fiberPackage,
		templater:   framework.FiberTemplates{},
	}

	p.FrameworkMap[flags.GorillaMux] = Framework{
		packageName: gorillaPackage,
		templater:   framework.GorillaTemplates{},
	}

	p.FrameworkMap[flags.HttpRouter] = Framework{
		packageName: routerPackage,
		templater:   framework.RouterTemplates{},
	}

	p.FrameworkMap[flags.Echo] = Framework{
		packageName: echoPackage,
		templater:   framework.EchoTemplates{},
	}
}

func (p *Project) createDBDriverMap() {

	p.DBDriverMap[flags.MySql] = DBDriver{
		packageName: mysqlDriver,
		templater:   dbdriver.MysqlTemplate{},
	}
	p.DBDriverMap[flags.Postgres] = DBDriver{
		packageName: postgresDriver,
		templater:   dbdriver.PostgresTemplate{},
	}
	p.DBDriverMap[flags.Sqlite] = DBDriver{
		packageName: sqliteDriver,
		templater:   dbdriver.SqliteTemplate{},
	}
	p.DBDriverMap[flags.Mongo] = DBDriver{
		packageName: mongoDriver,
		templater:   dbdriver.MongoTemplate{},
	}
	p.DBDriverMap[flags.Redis] = DBDriver{
		packageName: redisDriver,
		templater:   dbdriver.RedisTemplate{},
	}
}

func (p *Project) createDockerMap() {
	p.DockerMap = make(map[flags.Database]Docker)

	p.DockerMap[flags.MySql] = Docker{
		packageName: []string{},
		templater:   docker.MysqlDockerTemplate{},
	}
	p.DockerMap[flags.Postgres] = Docker{
		packageName: []string{},
		templater:   docker.PostgresDockerTemplate{},
	}
	p.DockerMap[flags.Mongo] = Docker{
		packageName: []string{},
		templater:   docker.MongoDockerTemplate{},
	}
	p.DockerMap[flags.Redis] = Docker{
		packageName: []string{},
		templater:   docker.RedisDockerTemplate{},
	}
}

func (p *Project) CreateMainFile() error {
	if _, err := os.Stat(p.AbsolutePath); os.IsNotExist(err) {
		if err := os.Mkdir(p.AbsolutePath, 0o754); err != nil {
			log.Printf("Could not create directory: %v", err)
			return err
		}
	}

	if p.GitOptions.String() != flags.Skip {

		emailSet, err := utils.CheckGitConfig("user.email")
		if err != nil {
			return err
		}
		if !emailSet {
			fmt.Println("user.email is not set in git config.")
			fmt.Println("Please set up git config before trying again.")
			panic("\nGIT CONFIG ISSUE: user.email is not set in git config.\n")
		}
	}

	p.ProjectName = strings.TrimSpace(p.ProjectName)

	projectPath := filepath.Join(p.AbsolutePath, utils.GetRootDir(p.ProjectName))

	if _, err := os.Stat(projectPath); os.IsNotExist(err) {
		err := os.MkdirAll(projectPath, 0o751)
		if err != nil {
			log.Printf("Error creating root project directory %v\n", err)
			return err
		}
	}

	// Define Operating system
	p.CheckOS()

	// Create the map for our program
	p.createFrameworkMap()

	err := utils.InitGoMod(p.ProjectName, projectPath)
	if err != nil {
		log.Printf("Could not initialize go.mod in new project %v\n", err)
		return err
	}

	// Install the correct package for the selected framework
	if p.ProjectType != flags.StandardLibrary {
		err = utils.GoGetPackage(projectPath, p.FrameworkMap[p.ProjectType].packageName)
		if err != nil {
			log.Println("Could not install go dependency for the chosen framework")
			return err
		}
	}

	return nil
}
