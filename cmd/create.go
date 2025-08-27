package cmd

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mahibulhaque/gofast/flags"
)

const logo = `

 ____  _                       _       _   
|  _ \| |                     (_)     | |  
| |_) | |_   _  ___ _ __  _ __ _ _ __ | |_ 
|  _ <| | | | |/ _ \ '_ \| '__| | '_ \| __|
| |_) | | |_| |  __/ |_) | |  | | | | | |_ 
|____/|_|\__,_|\___| .__/|_|  |_|_| |_|\__|
				   | |                     
				   |_|                     

`

var (
	logoStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	tipMsgStyle    = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("190")).Italic(true)
	endingMsgStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
)

func init() {
	// var flagFramework flags.Framework
	// var flagDBDriver flags.Database
	// var advancedFeatures flags.AdvancedFeatures
	// var flagGit flags.Git
	// rootCmd.AddCommand(createCmd)
	//
	// createCmd.Flags().StringP("name", "n", "", "Name of project to create")
	// createCmd.Flags().VarP(&flagFramework, "framework", "f", fmt.Sprintf("Framework to use. Allowed values: %s", strings.Join(flags.AllowedProjectTypes, ", ")))
	// createCmd.Flags().VarP(&flagDBDriver, "driver", "d", fmt.Sprintf("Database drivers to use. Allowed values: %s", strings.Join(flags.AllowedDBDrivers, ", ")))
	// createCmd.Flags().BoolP("advanced", "a", false, "Get prompts for advanced features")
	// createCmd.Flags().Var(&advancedFeatures, "feature", fmt.Sprintf("Advanced feature to use. Allowed values: %s", strings.Join(flags.AllowedAdvancedFeatures, ", ")))
	// createCmd.Flags().VarP(&flagGit, "git", "g", fmt.Sprintf("Git to use. Allowed values: %s", strings.Join(flags.AllowedGitsOptions, ", ")))

	// utils.RegisterStaticCompletions(createCmd, "framework", flags.AllowedProjectTypes)
	// utils.RegisterStaticCompletions(createCmd, "driver", flags.AllowedDBDrivers)
	// utils.RegisterStaticCompletions(createCmd, "feature", flags.AllowedAdvancedFeatures)
	// utils.RegisterStaticCompletions(createCmd, "git", flags.AllowedGitsOptions)
}
