package gocmds

import "github.com/mahibulhaque/gofast/internal/executor"

// InitGoMod initializes go.mod with the given project name
// in the selected directory
func InitGoMod(projectName string, appDir string) error {
	if err := executor.ExecuteCmd("go",
		[]string{"mod", "init", projectName},
		appDir); err != nil {
		return err
	}

	return nil
}

// GoGetPackage runs "go get" for a given package in the
// selected directory
func GoGetPackage(appDir string, packages []string) error {
	for _, packageName := range packages {
		if err := executor.ExecuteCmd("go",
			[]string{"get", "-u", packageName},
			appDir); err != nil {
			return err
		}
	}

	return nil
}

// GoFmt runs "gofmt" in a selected directory using the
// simplify and overwrite flags
func GoFmt(appDir string) error {
	if err := executor.ExecuteCmd("gofmt",
		[]string{"-s", "-w", "."},
		appDir); err != nil {
		return err
	}

	return nil
}

// GoModReplace runs "go mod edit -replace" in the selected
// replace_payload e.g: github.com/gocql/gocql=github.com/scylladb/gocql@v1.14.4
func GoModReplace(appDir string, replace string) error {
	if err := executor.ExecuteCmd("go",
		[]string{"mod", "edit", "-replace", replace},
		appDir,
	); err != nil {
		return err
	}

	return nil
}

func GoTidy(appDir string) error {
	err := executor.ExecuteCmd("go", []string{"mod", "tidy"}, appDir)
	if err != nil {
		return err
	}
	return nil
}
