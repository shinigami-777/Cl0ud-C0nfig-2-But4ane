package main

func Convert(cc *CloudConfig) *ButaneConfig {
	bc := &ButaneConfig{
		Variant: "flatcar",
		Version: "1.1.0",
	}

	users := ConvertUsers(cc.Users)
	if len(users) > 0 {
		bc.Passwd = &Passwd{
			Users: users,
		}
	}

	files := ConvertFiles(cc.WriteFiles)
	scriptFile, unit := ConvertRuncmd(cc.Runcmd)
	if scriptFile != nil {
		files = append(files, *scriptFile)
	}

	if len(files) > 0 {
		bc.Storage = &Storage{
			Files: files,
		}
	}

	if unit != nil {
		bc.Systemd = &Systemd{
			Units: []Unit{*unit},
		}
	}

	return bc
}
