// SPDX-License-Identifier: GPL-3.0-or-later

package main

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/netdata/go.d.plugin/agent"
	"github.com/netdata/go.d.plugin/cli"
	"github.com/netdata/go.d.plugin/logger"
	"github.com/netdata/go.d.plugin/pkg/multipath"

	"github.com/jessevdk/go-flags"
	"golang.org/x/net/http/httpproxy"

	_ "github.com/netdata/go.d.plugin/modules"
)

var (
	cd, _               = os.Getwd()
	name                = "go.d"
	userDir             = os.Getenv("NETDATA_USER_CONFIG_DIR")
	stockDir            = os.Getenv("NETDATA_STOCK_CONFIG_DIR")
	varLibDir           = os.Getenv("NETDATA_LIB_DIR")
	lockDir             = os.Getenv("NETDATA_LOCK_DIR")
	watchPath           = os.Getenv("NETDATA_PLUGINS_GOD_WATCH_PATH")
	envLogSeverityLevel = os.Getenv("NETDATA_LOG_SEVERITY_LEVEL")

	version = "unknown"
)

func confDir(opts *cli.Option) multipath.MultiPath {
	if len(opts.ConfDir) > 0 {
		return opts.ConfDir
	}
	if userDir != "" || stockDir != "" {
		return multipath.New(
			userDir,
			stockDir,
		)
	}
	return multipath.New(
		filepath.Join(cd, "/../../../../etc/netdata"),
		filepath.Join(cd, "/../../../../usr/lib/netdata/conf.d"),
	)
}

func modulesConfDir(opts *cli.Option) (mpath multipath.MultiPath) {
	if len(opts.ConfDir) > 0 {
		return opts.ConfDir
	}
	if userDir != "" || stockDir != "" {
		if userDir != "" {
			mpath = append(mpath, filepath.Join(userDir, name))
		}
		if stockDir != "" {
			mpath = append(mpath, filepath.Join(stockDir, name))
		}
		return multipath.New(mpath...)
	}
	return multipath.New(
		filepath.Join(cd, "/../../../../etc/netdata", name),
		filepath.Join(cd, "/../../../../usr/lib/netdata/conf.d", name),
	)
}

func watchPaths(opts *cli.Option) []string {
	if watchPath == "" {
		return opts.WatchPath
	}
	return append(opts.WatchPath, watchPath)
}

func stateFile() string {
	if varLibDir == "" {
		return ""
	}
	return filepath.Join(varLibDir, "god-jobs-statuses.json")
}

func init() {
	// https://github.com/netdata/netdata/issues/8949#issuecomment-638294959
	if v := os.Getenv("TZ"); strings.HasPrefix(v, ":") {
		_ = os.Unsetenv("TZ")
	}
}

func main() {
	opts := parseCLI()

	if opts.Version {
		fmt.Printf("go.d.plugin, version: %s\n", version)
		return
	}

	if envLogSeverityLevel != "" {
		logger.SetSeverityByName(envLogSeverityLevel)
	}

	if opts.Debug {
		logger.SetSeverity(logger.DEBUG)
	}

	a := agent.New(agent.Config{
		Name:              name,
		ConfDir:           confDir(opts),
		ModulesConfDir:    modulesConfDir(opts),
		ModulesSDConfPath: watchPaths(opts),
		VnodesConfDir:     confDir(opts),
		StateFile:         stateFile(),
		LockDir:           lockDir,
		RunModule:         opts.Module,
		MinUpdateEvery:    opts.UpdateEvery,
	})

	a.Debugf("plugin: name=%s, version=%s", a.Name, version)
	if u, err := user.Current(); err == nil {
		a.Debugf("current user: name=%s, uid=%s", u.Username, u.Uid)
	}

	cfg := httpproxy.FromEnvironment()
	a.Infof("env HTTP_PROXY '%s', HTTPS_PROXY '%s'", cfg.HTTPProxy, cfg.HTTPSProxy)

	a.Run()
}

func parseCLI() *cli.Option {
	opt, err := cli.Parse(os.Args)
	if err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	return opt
}
