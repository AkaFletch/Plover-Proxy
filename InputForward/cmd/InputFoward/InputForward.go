package main

import (
	"log/slog"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
	version "github.com/AkaFletch/Plover-Proxy/InputForward/m/v2/internal/version"
)

func main() {
	slog.Info("Starting InputForward", "version", version.Version)
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.CtrlC {
			return true, nil
		}
		slog.Info(key.String())
		return false, nil
	})
}
