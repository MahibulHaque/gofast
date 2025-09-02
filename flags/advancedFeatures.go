package flags

import (
	"fmt"
	"strings"
)

type AdvancedFeatures []string

const (
	GoProjectWorkflow string = "githubaction"
	Websocket         string = "websocket"
	React             string = "react"
	Docker            string = "docker"
)

var AllowedAdvancedFeatures = []string{string(React), string(GoProjectWorkflow), string(Websocket), string(Docker)}

func (f AdvancedFeatures) String() string {
	return strings.Join(f, ",")
}

func (f *AdvancedFeatures) Type() string {
	return "AdvancedFeatures"
}

func (f *AdvancedFeatures) Set(value string) error {
	for _, advancedFeature := range AllowedAdvancedFeatures {
		if advancedFeature == value {
			*f = append(*f, advancedFeature)
			return nil
		}
	}

	return fmt.Errorf("advanced Feature to use. Allowed values: %s", strings.Join(AllowedAdvancedFeatures, ", "))
}
