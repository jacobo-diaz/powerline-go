package main

import (
	pwl "github.com/justjanne/powerline-go/powerline"
	"os"
)

func segmentSSHHostname(p *powerline) {
	var hostPrompt string
	var foreground, background uint8
	if *p.args.ColorizeHostname {
		hostName := getHostName()
		hostPrompt = hostName

		hash := getMd5(hostName)
		background = hash[0]
		foreground = p.theme.HostnameColorizedFgMap[background]
	} else {
		if *p.args.Shell == "bash" {
			hostPrompt = "\\h"
		} else if *p.args.Shell == "zsh" {
			hostPrompt = "%m"
		} else {
			hostPrompt = getHostName()
		}

		foreground = p.theme.HostnameFg
		background = p.theme.HostnameBg
	}

	sshClient, _ := os.LookupEnv("SSH_CLIENT")
	if sshClient != "" {
		p.appendSegment("host", pwl.Segment{
			Content:    hostPrompt,
			Foreground: foreground,
			Background: background,
		})
	}
}
