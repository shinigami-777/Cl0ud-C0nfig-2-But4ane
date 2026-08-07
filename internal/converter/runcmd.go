package converter

import (
	"fmt"
	"strings"
)

func ConvertRuncmd(runcmd []any) (*File, *Unit) {
	if len(runcmd) == 0 {
		return nil, nil
	}

	var scriptLines []string
	scriptLines = append(scriptLines, "#!/bin/bash", "set -euo pipefail")

	for _, cmd := range runcmd {
		switch v := cmd.(type) {
		case string:
			scriptLines = append(scriptLines, v)
		case []any:
			var parts []string
			for _, part := range v {
				if s, ok := part.(string); ok {
					parts = append(parts, fmt.Sprintf("%q", s))
				}
			}
			scriptLines = append(scriptLines, strings.Join(parts, " "))
		}
	}

	file := &File{
		Path: "/etc/cc-runcmd.sh",
		Mode: IntPtr(0755),
		Contents: FileContents{
			Inline: strings.Join(scriptLines, "\n") + "\n",
		},
	}

	unitContents := `[Unit]
Description=cloud-config runcmd
Wants=network-online.target
After=network-online.target

[Service]
Type=oneshot
RemainAfterExit=yes
ExecStart=/bin/bash /etc/cc-runcmd.sh

[Install]
WantedBy=multi-user.target
`
	unit := &Unit{
		Name:     "cc-runcmd.service",
		Enabled:  BoolPtr(true),
		Contents: unitContents,
	}

	return file, unit
}
