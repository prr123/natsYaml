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

    cfg, err := RdYaml(yamlFilnam)
    if err != nil {t.Errorf("could not decode yaml file: %v", err)}

    server := cfg.Server
    if server.Addr != "0.0.0.0" {t.Errorf("should be 0.0.0.0, but is %s", server.Addr)}

    PrintNatsConfig(cfg)
}

