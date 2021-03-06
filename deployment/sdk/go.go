package sdk

import (
	"fmt"
	"github.com/viant/endly"
	"github.com/viant/endly/deployment/deploy"
	"github.com/viant/endly/model"
	"github.com/viant/endly/system/exec"
	"github.com/viant/endly/util"
)

//TODO complete implementation
type goService struct{}

func (s *goService) setSdk(context *endly.Context, request *SetRequest) (*Info, error) {
	var result = &Info{}
	var sdkHome = "/opt/sdk/go"
	var runResponse = &exec.RunResponse{}
	hasGoRoot := endly.Run(context, exec.NewRunRequest(request.Target, false, "ls -al /usr/local/go"), nil) == nil
	if err := endly.Run(context, exec.NewExtractRequest(request.Target, nil, exec.NewExtractCommand("ls -al /opt/sdk/go", "", nil, nil)), runResponse); err == nil {
		if !util.CheckNoSuchFileOrDirectory(runResponse.Output) || !hasGoRoot {
			var request = exec.NewRunRequest(request.Target, false, "export GOROOT='/opt/sdk/go'")

			_ = endly.Run(context, request, nil)
		}
	}
	var extractRequest = exec.NewExtractRequest(request.Target, exec.DefaultOptions(),
		exec.NewExtractCommand("go version", "", nil, nil,
			model.NewExtract("version", "go version go([^\\s]+)", false)),
	)

	extractRequest.SystemPaths = append(extractRequest.SystemPaths, fmt.Sprintf("%v/bin", sdkHome))
	if err := endly.Run(context, extractRequest, runResponse); err != nil {
		return nil, err
	}
	var stdout = runResponse.Stdout()
	if util.CheckCommandNotFound(stdout) || util.CheckNoSuchFileOrDirectory(stdout) {
		return nil, errSdkNotFound
	}
	result.Sdk = "go"
	result.Home = sdkHome
	if version, ok := runResponse.Data["version"]; ok {
		result.Version = version.(string)
	}
	if !deploy.MatchVersion(request.Version, result.Version) {
		return nil, errSdkNotFound
	}
	if result.Version == "" {
		result.Version = request.Version
	}
	return result, nil
}
