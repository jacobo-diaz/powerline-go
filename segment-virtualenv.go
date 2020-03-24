package main

import (
	"os"
	"path"
	"fmt"

	pwl "github.com/justjanne/powerline-go/powerline"
)

func segmentVirtualEnv(p *powerline) {
	var env string
	var envIcon string
	var envFg uint8
	var envBg uint8

	if env == "" {
		env, _ = os.LookupEnv("VIRTUAL_ENV")
		envIcon = p.symbolTemplates.VirtualEnvPython
		envFg = p.theme.VirtualEnvPythonFg
		envBg = p.theme.VirtualEnvPythonBg
	}
	if env == "" {
		env, _ = os.LookupEnv("CONDA_ENV_PATH")
		envIcon = p.symbolTemplates.VirtualEnvConda
		envFg = p.theme.VirtualEnvCondaFg
		envBg = p.theme.VirtualEnvCondaBg
	}
	if env == "" {
		env, _ = os.LookupEnv("CONDA_DEFAULT_ENV")
		envIcon = p.symbolTemplates.VirtualEnvConda
		envFg = p.theme.VirtualEnvCondaFg
		envBg = p.theme.VirtualEnvCondaBg
	}
	if env == "" {
		envIcon = ""
		envFg = p.theme.VirtualEnvPythonFg
		envBg = p.theme.VirtualEnvPythonBg
		return
	}
	envName := path.Base(env)

	if envIcon != "" {
		p.appendSegment("venv", pwl.Segment{
			Content: fmt.Sprintf("%s %s", envIcon, envName ) ,
			Foreground: envFg,
			Background: envBg,
		})
	} else {
		p.appendSegment("venv", pwl.Segment{
			Content: envName,
			Foreground: envFg,
			Background: envBg,
		})
	}

}
