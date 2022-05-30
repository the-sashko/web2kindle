package script

import (
	"fmt"
	"os/exec"
	"web2kindle/logger"
)

const (
	binBashPath       = "/bin/bash"
	convertScriptPath = "../scripts/convert.sh"
	convertScriptName = "Convert"
)

func RunConvert() {
	_ = doRun(convertScriptPath, convertScriptName)
}

func doRun(scriptPath string, scriptName string) string {
	cmd := exec.Command(binBashPath, scriptPath)

	result, err := cmd.Output()

	if err != nil {
		doError(err, scriptName)
	}

	return string(result)
}

func doError(errorEntity error, scriptName string) {
	errorMessage := fmt.Sprintf("%s Script", scriptName)

	logger.LogError(errorMessage, errorEntity, true)
}
