package bean

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type HistoryReleaseDeployBean struct {
	DeployBean
}

// sourceURL, execPath, serverName, serverIP
func (hrdb *HistoryReleaseDeployBean) Execute(stringArgs []string, env *map[string]interface{}) error{
	if len(stringArgs) != 4 {
		return fmt.Errorf("argument error")
	}
	if err := os.MkdirAll(stringArgs[1], os.ModePerm); err != nil {
		return fmt.Errorf("error while execute git clone, err: %s", err)
	}
	sourcePath := hrdb.prepareSourcePath(stringArgs[1], stringArgs[0])
	return hrdb.doDeploy(sourcePath, stringArgs[2], stringArgs[3])
}

func (hrdb *HistoryReleaseDeployBean) prepareSourcePath(execPath, sourceURL string) string{
	// download source(sourceURL) to execPath
	response, err := http.Get(sourceURL)
	if err != nil {
		return ""
	}
	defer response.Body.Close()
	out, err := os.Create(fmt.Sprintf("%s/%s", execPath, "app.jar"))
	defer out.Close()
	io.Copy(out, response.Body)
	return fmt.Sprintf("%s/", execPath)
}
