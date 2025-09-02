package program

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/mahibulhaque/gofast/flags"
	tpl "github.com/mahibulhaque/gofast/template"
	"github.com/mahibulhaque/gofast/template/advanced"
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

	if p.DBDriver != "none" {
		p.createDBDriverMap()

		err = utils.GoGetPackage(projectPath, p.DBDriverMap[p.DBDriver].packageName)

		if err != nil {
			log.Println("Could not install go dependency for chosen driver")

			return err
		}

		err = p.CreatePath(internalDatabasePath, projectPath)

		if err != nil {
			log.Printf("Error creating path: %s", internalDatabasePath)

			return err
		}

		err = p.CreateFileWithInjection(internalDatabasePath, projectPath, "database.go", "database")
		if err != nil {
			log.Printf("Error injecting db.go file: %v", err)
			return err
		}

		if p.DBDriver != "sqlite" {
			err = p.CreateFileWithInjection(internalDatabasePath, projectPath, "database_test.go", "integration-tests")
			if err != nil {
				log.Printf("Error injecting database_test.go file: %v", err)
				return err
			}
		}
	}

	if p.DBDriver != "none" && p.DBDriver != "sqlite" {
		p.createDockerMap()
		p.Docker = p.DBDriver

		err = p.CreateFileWithInjection(root, projectPath, "docker-compose.yml", "db-docker")
		if err != nil {
			log.Printf("Error injecting docker-compose.yml file: %v", err)
			return err
		}
	}

	err = utils.GoGetPackage(projectPath, godotenvPackage)

	if err != nil {
		log.Println("Could not install go dependency")

		return err
	}

	err = p.CreatePath(cmdApiPath, projectPath)
	if err != nil {
		log.Printf("Error creating path: %s", projectPath)
		return err
	}

	err = p.CreateFileWithInjection(cmdApiPath, projectPath, "main.go", "main")
	if err != nil {
		return err
	}

	makeFile, err := os.Create(filepath.Join(projectPath, "Makefile"))
	if err != nil {
		return err
	}

	defer makeFile.Close()

	makeFileTemplate := template.Must(template.New("makefile").Parse(string(framework.MakeTemplate())))
	err = makeFileTemplate.Execute(makeFile, p)
	if err != nil {
		return err
	}

	readmeFile, err := os.Create(filepath.Join(projectPath, "README.md"))
	if err != nil {
		return err
	}
	defer readmeFile.Close()

	// inject readme template
	readmeFileTemplate := template.Must(template.New("readme").Parse(string(framework.ReadmeTemplate())))
	err = readmeFileTemplate.Execute(readmeFile, p)
	if err != nil {
		return err
	}

	err = p.CreatePath(internalServerPath, projectPath)
	if err != nil {
		log.Printf("Error creating path: %s", internalServerPath)
		return err
	}
	if p.AdvancedOptions[string(flags.React)] {
		if err := p.CreateViteReactProject(projectPath); err != nil {
			return fmt.Errorf("failed to set up React project: %w", err)
		}
	}

	// Create .github/workflows folder and inject release.yml and go-test.yml
	if p.AdvancedOptions[string(flags.GoProjectWorkflow)] {
		err = p.CreatePath(githubActionPath, projectPath)
		if err != nil {
			log.Printf("Error creating path: %s", githubActionPath)
			return err
		}

		err = p.CreateFileWithInjection(githubActionPath, projectPath, "release.yml", "releaser")
		if err != nil {
			log.Printf("Error injecting release.yml file: %v", err)
			return err
		}

		err = p.CreateFileWithInjection(githubActionPath, projectPath, "go-test.yml", "go-test")
		if err != nil {
			log.Printf("Error injecting go-test.yml file: %v", err)
			return err
		}

		err = p.CreateFileWithInjection(root, projectPath, ".goreleaser.yml", "releaser-config")
		if err != nil {
			log.Printf("Error injecting .goreleaser.yml file: %v", err)
			return err
		}
	}

	if p.AdvancedOptions[string(flags.Websocket)] {
		p.CreateWebsocketImports(projectPath)
	}

	if p.AdvancedOptions[string(flags.Docker)] {
		Dockerfile, err := os.Create(filepath.Join(projectPath, "Dockerfile"))
		if err != nil {
			return err
		}
		defer Dockerfile.Close()

		// inject Docker template
		dockerfileTemplate := template.Must(template.New("Dockerfile").Parse(string(advanced.Dockerfile())))
		err = dockerfileTemplate.Execute(Dockerfile, p)
		if err != nil {
			return err
		}

		if p.DBDriver == "none" || p.DBDriver == "sqlite" {

			Dockercompose, err := os.Create(filepath.Join(projectPath, "docker-compose.yml"))
			if err != nil {
				return err
			}
			defer Dockercompose.Close()

			// inject DockerCompose template
			dockerComposeTemplate := template.Must(template.New("docker-compose.yml").Parse(string(advanced.DockerCompose())))
			err = dockerComposeTemplate.Execute(Dockercompose, p)
			if err != nil {
				return err
			}
		}
	}

	err = p.CreateFileWithInjection(internalServerPath, projectPath, "routes_test.go", "tests")
	if err != nil {
		return err
	}

	err = p.CreateFileWithInjection(internalServerPath, projectPath, "server.go", "server")
	if err != nil {
		log.Printf("Error injecting server.go file: %v", err)
		return err
	}

	err = p.CreateFileWithInjection(root, projectPath, ".env", "env")
	if err != nil {
		log.Printf("Error injecting .env file: %v", err)
		return err
	}

	// Create gitignore
	gitignoreFile, err := os.Create(filepath.Join(projectPath, ".gitignore"))
	if err != nil {
		return err
	}
	defer gitignoreFile.Close()

	// inject gitignore template
	gitignoreTemplate := template.Must(template.New(".gitignore").Parse(string(framework.GitIgnoreTemplate())))
	err = gitignoreTemplate.Execute(gitignoreFile, p)
	if err != nil {
		return err
	}

	// Create .air.toml file
	airTomlFile, err := os.Create(filepath.Join(projectPath, ".air.toml"))
	if err != nil {
		return err
	}

	defer airTomlFile.Close()

	// inject air.toml template
	airTomlTemplate := template.Must(template.New("airtoml").Parse(string(framework.AirTomlTemplate())))
	err = airTomlTemplate.Execute(airTomlFile, p)
	if err != nil {
		return err
	}

	err = utils.GoTidy(projectPath)
	if err != nil {
		log.Printf("Could not go tidy in new project %v\n", err)
		return err
	}

	err = utils.GoFmt(projectPath)
	if err != nil {
		log.Printf("Could not gofmt in new project %v\n", err)
		return err
	}

	if p.GitOptions != flags.Skip {
		nameSet, err := utils.CheckGitConfig("user.name")
		if err != nil {
			return err
		}

		if !nameSet {
			fmt.Println("user.name is not set in git config.")
			fmt.Println("Please set up git config before trying again.")
			panic("\nGIT CONFIG ISSUE: user.name is not set in git config.\n")
		}
		// Initialize git repo
		err = utils.ExecuteCmd("git", []string{"init"}, projectPath)
		if err != nil {
			log.Printf("Error initializing git repo: %v", err)
			return err
		}

		// Git add files
		err = utils.ExecuteCmd("git", []string{"add", "."}, projectPath)
		if err != nil {
			log.Printf("Error adding files to git repo: %v", err)
			return err
		}

		if p.GitOptions == flags.Commit {
			// Git commit files
			err = utils.ExecuteCmd("git", []string{"commit", "-m", "Initial commit"}, projectPath)
			if err != nil {
				log.Printf("Error committing files to git repo: %v", err)
				return err
			}
		}
	}
	return nil
}

// CreatePath creates the given directory in the projectPath
func (p *Project) CreatePath(pathToCreate string, projectPath string) error {
	path := filepath.Join(projectPath, pathToCreate)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, 0o751)
		if err != nil {
			log.Printf("Error creating directory %v\n", err)
			return err
		}
	}

	return nil
}

// CreateFileWithInjection creates the given file at the
// project path, and injects the appropriate template
func (p *Project) CreateFileWithInjection(pathToCreate string, projectPath string, fileName string, methodName string) error {
	createdFile, err := os.Create(filepath.Join(projectPath, pathToCreate, fileName))
	if err != nil {
		return err
	}

	defer createdFile.Close()

	switch methodName {
	case "main":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Main())))
		err = createdTemplate.Execute(createdFile, p)
	case "server":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.Server())))
		err = createdTemplate.Execute(createdFile, p)
	case "routes":
		routeFileBytes := p.FrameworkMap[p.ProjectType].templater.Routes()
		createdTemplate := template.Must(template.New(fileName).Parse(string(routeFileBytes)))
		err = createdTemplate.Execute(createdFile, p)
	case "releaser":
		createdTemplate := template.Must(template.New(fileName).Parse(string(advanced.Releaser())))
		err = createdTemplate.Execute(createdFile, p)
	case "go-test":
		createdTemplate := template.Must(template.New(fileName).Parse(string(advanced.Test())))
		err = createdTemplate.Execute(createdFile, p)
	case "releaser-config":
		createdTemplate := template.Must(template.New(fileName).Parse(string(advanced.ReleaserConfig())))
		err = createdTemplate.Execute(createdFile, p)
	case "database":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.DBDriverMap[p.DBDriver].templater.Service())))
		err = createdTemplate.Execute(createdFile, p)
	case "db-docker":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.DockerMap[p.Docker].templater.Docker())))
		err = createdTemplate.Execute(createdFile, p)
	case "integration-tests":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.DBDriverMap[p.DBDriver].templater.Tests())))
		err = createdTemplate.Execute(createdFile, p)
	case "tests":
		createdTemplate := template.Must(template.New(fileName).Parse(string(p.FrameworkMap[p.ProjectType].templater.TestHandler())))
		err = createdTemplate.Execute(createdFile, p)
	case "env":
		if p.DBDriver != "none" {

			envBytes := [][]byte{
				tpl.GlobalEnvTemplate(),
				p.DBDriverMap[p.DBDriver].templater.Env(),
			}
			createdTemplate := template.Must(template.New(fileName).Parse(string(bytes.Join(envBytes, []byte("\n")))))
			err = createdTemplate.Execute(createdFile, p)

		} else {
			createdTemplate := template.Must(template.New(fileName).Parse(string(tpl.GlobalEnvTemplate())))
			err = createdTemplate.Execute(createdFile, p)
		}
	}

	if err != nil {
		return err
	}

	return nil
}

func (p *Project) CreateWebsocketImports(appDir string) {
	websocketDependency := []string{"github.com/coder/websocket"}
	if p.ProjectType == flags.Fiber {
		websocketDependency = []string{"github.com/gofiber/contrib/websocket"}
	}

	// Websockets require a different package depending on what framework is
	// choosen. The application calls go mod tidy at the end so we don't
	// have to here
	err := utils.GoGetPackage(appDir, websocketDependency)
	if err != nil {
		log.Fatal(err)
	}

	importsPlaceHolder := string(p.FrameworkMap[p.ProjectType].templater.WebsocketImports())

	importTmpl, err := template.New("imports").Parse(importsPlaceHolder)
	if err != nil {
		log.Fatalf("CreateWebsocketImports failed to create template: %v", err)
	}
	var importBuffer bytes.Buffer
	err = importTmpl.Execute(&importBuffer, p)
	if err != nil {
		log.Fatalf("CreateWebsocketImports failed write template: %v", err)
	}
	newImports := strings.Join([]string{string(p.AdvancedTemplates.TemplateImports), importBuffer.String()}, "\n")
	p.AdvancedTemplates.TemplateImports = newImports
}

func (p *Project) CreateViteReactProject(projectPath string) error {
	if err := checkNpmInstalled(); err != nil {
		return err
	}

	originalDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get current directory: %w", err)
	}

	defer func() {
		if err := os.Chdir(originalDir); err != nil {
			fmt.Fprintf(os.Stderr, "failed to change back to original directory: %v\n", err)
		}
	}()

	// change into the project directory to run vite command
	err = os.Chdir(projectPath)
	if err != nil {
		fmt.Println("failed to change into project directory: %w", err)
	}

	fmt.Println("Installing create-vite (using cache if available)...")
	cmd := exec.Command("npm", "create", "vite@latest", "frontend", "--",
		"--template", "react-ts",
		"--prefer-offline",
		"--no-fund")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to use create-vite: %w", err)
	}

	frontendPath := filepath.Join(projectPath, "frontend")
	if err := os.MkdirAll(frontendPath, 0755); err != nil {
		return fmt.Errorf("failed to create frontend directory: %w", err)
	}

	if err := os.Chdir(frontendPath); err != nil {
		return fmt.Errorf("failed to change to frontend directory: %w", err)
	}

	srcDir := filepath.Join(frontendPath, "src")
	if err := os.MkdirAll(srcDir, 0755); err != nil {
		return fmt.Errorf("failed to create src directory: %w", err)
	}

	if err := os.WriteFile(filepath.Join(frontendPath, "vite.config.ts"), advanced.ReactViteConfigFile(), 0644); err != nil {
		return fmt.Errorf("failed to write main.tsx template: %w", err)
	}

	if err := os.WriteFile(filepath.Join(srcDir, "main.tsx"), advanced.ReactMainFile(), 0644); err != nil {
		return fmt.Errorf("failed to write main.tsx template: %w", err)
	}

	if err := os.WriteFile(filepath.Join(srcDir, "styles.css"), advanced.ReactStylesCssFile(), 0644); err != nil {
		return fmt.Errorf("failed to write main.tsx template: %w", err)
	}

	err = p.CreateFileWithInjection("", projectPath, ".env", "env")
	if err != nil {
		return fmt.Errorf("failed to create global .env file: %w", err)
	}

	// Read from the global `.env` file and create the frontend-specific `.env`
	globalEnvPath := filepath.Join(projectPath, ".env")
	vitePort := "8080" // Default fallback

	// Read the global .env file
	if data, err := os.ReadFile(globalEnvPath); err == nil {
		lines := strings.Split(string(data), "\n")
		for _, line := range lines {
			if strings.HasPrefix(line, "PORT=") {
				vitePort = strings.SplitN(line, "=", 2)[1] // Get the backend port value
				break
			}
		}
	}

	// Use a template to generate the frontend .env file
	frontendEnvContent := fmt.Sprintf("VITE_PORT=%s\n", vitePort)
	if err := os.WriteFile(filepath.Join(frontendPath, ".env"), []byte(frontendEnvContent), 0644); err != nil {
		return fmt.Errorf("failed to create frontend .env file: %w", err)
	}

	return nil
}

func checkNpmInstalled() error {
	cmd := exec.Command("npm", "--version")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("npm is not installed: %w", err)
	}
	return nil
}
