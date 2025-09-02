package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/mahibulhaque/gofast/internal/flags"
	"github.com/mahibulhaque/gofast/internal/modules"
	"github.com/mahibulhaque/gofast/internal/program"
	"github.com/mahibulhaque/gofast/internal/steps"
	"github.com/mahibulhaque/gofast/internal/tui/components/list"
	"github.com/mahibulhaque/gofast/internal/tui/components/logo"
	"github.com/mahibulhaque/gofast/internal/tui/components/multiInput"
	"github.com/mahibulhaque/gofast/internal/tui/components/multiSelect"
	"github.com/mahibulhaque/gofast/internal/tui/components/spinner"
	"github.com/mahibulhaque/gofast/internal/tui/components/textinput"
	"github.com/mahibulhaque/gofast/internal/tui/styles"
	"github.com/spf13/cobra"
)

func init() {
	var flagFramework flags.Framework
	var flagDBDriver flags.Database
	var advancedFeatures flags.AdvancedFeatures
	var flagGit flags.Git
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringP("name", "n", "", "Name of project to create")
	createCmd.Flags().VarP(&flagFramework, "framework", "f", fmt.Sprintf("Framework to use. Allowed values: %s", strings.Join(flags.AllowedProjectTypes, ", ")))
	createCmd.Flags().VarP(&flagDBDriver, "driver", "d", fmt.Sprintf("Database drivers to use. Allowed values: %s", strings.Join(flags.AllowedDBDrivers, ", ")))
	createCmd.Flags().BoolP("advanced", "a", false, "Get prompts for advanced features")
	createCmd.Flags().Var(&advancedFeatures, "feature", fmt.Sprintf("Advanced feature to use. Allowed values: %s", strings.Join(flags.AllowedAdvancedFeatures, ", ")))
	createCmd.Flags().VarP(&flagGit, "git", "g", fmt.Sprintf("Git to use. Allowed values: %s", strings.Join(flags.AllowedGitsOptions, ", ")))

	RegisterStaticCompletions(createCmd, "framework", flags.AllowedProjectTypes)
	RegisterStaticCompletions(createCmd, "driver", flags.AllowedDBDrivers)
	RegisterStaticCompletions(createCmd, "feature", flags.AllowedAdvancedFeatures)
	RegisterStaticCompletions(createCmd, "git", flags.AllowedGitsOptions)
}

type Options struct {
	ProjectName *textinput.Output
	ProjectType *list.Selection
	DBDriver    *multiInput.Selection
	Advanced    *multiSelect.Selection
	Workflow    *multiInput.Selection
	Git         *multiInput.Selection
}

func createCmdRun(cmd *cobra.Command, args []string) {
	var err error

	theme := styles.CurrentTheme()

	isInteractive := false
	flagName := cmd.Flag("name").Value.String()

	if flagName != "" && !modules.ValidateModuleName(flagName) {
		err = fmt.Errorf("'%s' is not a valid module name. Please choose a different name", flagName)
		cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
	}

	rootDirName := modules.GetRootDir(flagName)
	if rootDirName != "" && doesDirectoryExistAndIsNotEmpty(rootDirName) {
		err = fmt.Errorf("directory '%s' already exists and is not empty. Please choose a different name", rootDirName)
		cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
	}

	flagFramework := flags.Framework(cmd.Flag("framework").Value.String())
	flagDBDriver := flags.Database(cmd.Flag("driver").Value.String())
	flagGit := flags.Git(cmd.Flag("git").Value.String())

	options := Options{
		ProjectName: &textinput.Output{},
		ProjectType: &list.Selection{},
		DBDriver:    &multiInput.Selection{},
		Advanced: &multiSelect.Selection{
			Choices: make(map[string]bool),
		},
		Git: &multiInput.Selection{},
	}

	project := &program.Project{
		ProjectName:     flagName,
		ProjectType:     flagFramework,
		DBDriver:        flagDBDriver,
		FrameworkMap:    make(map[flags.Framework]program.Framework),
		DBDriverMap:     make(map[flags.Database]program.DBDriver),
		AdvancedOptions: make(map[string]bool),
		GitOptions:      flagGit,
	}

	steps := steps.InitSteps(flagFramework, flagDBDriver)
	fmt.Printf("%s\n", logo.Render("0.0.1", true, logo.DefaultOpts()))

	// Advanced option steps:
	flagAdvanced, err := cmd.Flags().GetBool("advanced")
	if err != nil {
		log.Fatal("failed to retrieve advanced flag")
	}

	if flagAdvanced {
		fmt.Println("*** You are in advanced mode ***")
	}

	if project.ProjectName == "" {
		isInteractive = true
		tprogram := tea.NewProgram(textinput.NewTextInputModel(options.ProjectName, "What is the name of your project?", project))
		if _, err := tprogram.Run(); err != nil {
			log.Printf("Name of project contains an error: %v", err)
			cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
		}

		if options.ProjectName.Output != "" && !modules.ValidateModuleName(options.ProjectName.Output) {
			err = fmt.Errorf("'%s' is not a valid module name. Please choose a different name", options.ProjectName.Output)
			cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
		}

		rootDirName = modules.GetRootDir(options.ProjectName.Output)

		if doesDirectoryExistAndIsNotEmpty(rootDirName) {
			err = fmt.Errorf("directory '%s' already exists and is not empty. Please choose a different name", rootDirName)
			cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
		}

		project.ExitCLI(tprogram)

		project.ProjectName = options.ProjectName.Output

		err := cmd.Flag("name").Value.Set(project.ProjectName)

		if err != nil {
			log.Fatal("Failed to set the name flag value", err)
		}

		if project.ProjectType == "" {
			isInteractive = true
			step := steps.Steps["framework"]

			tprogram := tea.NewProgram(list.NewSingleSelectFromStep(step, options.ProjectType, project))

			if _, err := tprogram.Run(); err != nil {
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}

			project.ExitCLI(tprogram)

			step.Field = options.ProjectType.Choice

			project.ProjectType = flags.Framework(strings.ToLower(options.ProjectType.Choice))
			err := cmd.Flag("framework").Value.Set(project.ProjectType.String())
			if err != nil {
				log.Fatal("failed to set the framework flag value", err)
			}
		}

		if project.DBDriver == "" {
			isInteractive = true

			step := steps.Steps["driver"]

			tprogram = tea.NewProgram(multiInput.InitialModelMulti(step.Options, options.DBDriver, step.Headers, project))
			if _, err := tprogram.Run(); err != nil {
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			project.ExitCLI(tprogram)

			/* NOTE: this type casting is always safe since the user interface can only pass strings that can be cast to a flags.Database instance */
			project.DBDriver = flags.Database(strings.ToLower(options.DBDriver.Choice))
			err := cmd.Flag("driver").Value.Set(project.DBDriver.String())
			if err != nil {
				log.Fatal("failed to set the driver flag value", err)
			}
		}

		if flagAdvanced {
			featureFlags := cmd.Flag("feature").Value.String()

			if featureFlags != "" {
				featureFlagValues := strings.Split(featureFlags, ",")

				for _, key := range featureFlagValues {
					project.AdvancedOptions[key] = true
				}
			} else {

				isInteractive = true
				step := steps.Steps["advanced"]
				tprogram = tea.NewProgram(multiSelect.InitialModelMultiSelect(step.Options, options.Advanced, step.Headers, project))

				if _, err := tprogram.Run(); err != nil {
					cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
				}

				project.ExitCLI(tprogram)

				for key, opt := range options.Advanced.Choices {
					project.AdvancedOptions[strings.ToLower(key)] = opt
					err := cmd.Flag("feature").Value.Set(strings.ToLower(key))
					if err != nil {
						log.Fatal("failed to set the feature flag value", err)
					}
				}

				if err != nil {
					log.Fatal("Failed to set the advanced option", err)
				}
			}
		}
		if project.GitOptions == "" {
			isInteractive = true
			step := steps.Steps["git"]
			tprogram = tea.NewProgram(multiInput.InitialModelMulti(step.Options, options.Git, step.Headers, project))
			if _, err := tprogram.Run(); err != nil {
				cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
			}
			project.ExitCLI(tprogram)

			project.GitOptions = flags.Git(strings.ToLower(options.Git.Choice))
			err := cmd.Flag("git").Value.Set(project.GitOptions.String())
			if err != nil {
				log.Fatal("failed to set the git flag value", err)
			}
		}

		currentWorkingDir, err := os.Getwd()
		if err != nil {
			log.Printf("could not get current working directory: %v", err)
			cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
		}
		project.AbsolutePath = currentWorkingDir

		spinner := tea.NewProgram(spinner.NewSpinnerModel())

		wg := sync.WaitGroup{}

		wg.Add(1)

		go func() {
			defer wg.Done()

			if _, err := spinner.Run(); err != nil {
				cobra.CheckErr(err)
			}
		}()

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("The program encountered an unexpected issue and had to exit. The error was:", r)
				if releaseErr := spinner.ReleaseTerminal(); releaseErr != nil {
					log.Printf("Problem releasing terminal: %v", releaseErr)
				}
			}
		}()

		// This calls the templates
		err = project.CreateMainFile()
		if err != nil {
			if releaseErr := spinner.ReleaseTerminal(); releaseErr != nil {
				log.Printf("Problem releasing terminal: %v", releaseErr)
			}
			log.Printf("Problem creating files for project.")
			cobra.CheckErr(textinput.CreateErrorInputModel(err).Err())
		}

		fmt.Println(theme.S().Text.Render("\nNext steps:"))
		fmt.Println(theme.S().Text.Render(fmt.Sprintf("• cd into the newly created project with: `cd %s`\n", modules.GetRootDir(project.ProjectName))))

		if options.Advanced.Choices["React"] {
			fmt.Println(theme.S().Text.Render("• cd into frontend\n"))
			fmt.Println(theme.S().Text.Render("• npm install\n"))
			fmt.Println(theme.S().Text.Render("• npm run dev\n"))
		}
		if isInteractive {
			nonInteractiveCommand := NonInteractiveCommand(cmd.Use, cmd.Flags())
			fmt.Println(theme.S().Text.Render("Tip: Repeat the equivalent with the following non-interactive command:"))
			fmt.Println(theme.S().Text.Italic(false).Render(fmt.Sprintf("• %s\n", nonInteractiveCommand)))
		}

		err = spinner.ReleaseTerminal()
		if err != nil {
			log.Printf("Could not release terminal: %v", err)
			cobra.CheckErr(err)
		}

	}

}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a Go project from pre defined templates",
	Long:  "GoFast is a CLI tool that allows you to focus on the actual Go code, and not the project structure.",

	Run: createCmdRun,
}

// doesDirectoryExistAndIsNotEmpty checks if the directory exists and is not empty
func doesDirectoryExistAndIsNotEmpty(name string) bool {
	if _, err := os.Stat(name); err == nil {
		dirEntries, err := os.ReadDir(name)
		if err != nil {
			log.Printf("could not read directory: %v", err)
			cobra.CheckErr(textinput.CreateErrorInputModel(err))
		}
		if len(dirEntries) > 0 {
			return true
		}
	}
	return false
}
