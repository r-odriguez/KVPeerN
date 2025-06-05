//go:build windows
// +build windows

package main

import "os"

var LOCALAPPDATA = os.Getenv("LOCALAPPDATA")

var CONFIG_DIR_PATH = LOCALAPPDATA + "\\VPeer-N\\"
