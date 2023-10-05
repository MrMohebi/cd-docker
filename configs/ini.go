package configs

import (
	"cd-docker/common"
	"gopkg.in/ini.v1"
)

var IniData *ini.File

func IniSetup() {
	var err error
	IniData, err = ini.Load("_configFiles/config.ini")
	common.IsErr(err, true, "Error loading .ini file")
}

func IniGet(section string, key string) string {
	if IniData == nil {
		IniSetup()
	}
	return IniData.Section(section).Key(key).String()
}

func IniSave() error {
	return IniData.SaveTo("_configFiles/config.ini")
}
