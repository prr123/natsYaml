package natsyaml

import (
	"testing"
	"os"
)



func TestRdYaml(t *testing.T) {

	yamlFilnam := "../yaml/natsConfig.yaml"

	yamlData, err := os.ReadFile(yamlFilnam)
	if err != nil {t.Errorf("could not read yaml config: %v", err)}
	if len(yamlData) == 0 {t.Errorf("no yaml data in file!")}
}
