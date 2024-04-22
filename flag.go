package nruntime

import (
	"flag"
	"fmt"
	"os"
)

type BootFlags struct {
	CmdPrintHelp    *bool
	CmdPrintVersion *bool
	OptWorkDir      *string
	OptEnvPrefix    *string
}

type BootOptions struct {
	WorkDir   string
	EnvPrefix string
}

func InterceptFlags() *BootOptions {
	// Init built-in flags
	f := newBootFlags()

	// Parse
	flag.Parse()

	// Intercept help command
	if *f.CmdPrintHelp {
		fmt.Printf("%s. Available Commands and Options:\n\n", AppName())
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Intercept version command
	if *f.CmdPrintVersion {
		fmt.Printf("%s\n"+
			"  Version         : %s\n"+
			"  Build Signature : %s\n",
			AppName(), AppVersion(), BuildSignature())
		os.Exit(0)
	}

	// Parse built-in boot options
	return &BootOptions{
		WorkDir:   *f.OptWorkDir,
		EnvPrefix: *f.OptEnvPrefix,
	}
}

func newBootFlags() BootFlags {
	return BootFlags{
		CmdPrintHelp:    flag.Bool("help", false, "Command: Show available commands and options"),
		CmdPrintVersion: flag.Bool("version", false, "Command: Show version"),
		OptWorkDir:      flag.String("dir", ".", "Option: Set working directory"),
		OptEnvPrefix:    flag.String("env-prefix", "", "Option: Set environment variable prefix"),
	}
}
