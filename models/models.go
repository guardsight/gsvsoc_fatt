package models

type Args struct {
	File    string `help:"file to scan" arg:"-f,--file"`
	Url     string `help:"url to scan" arg:"-u,--url"`
	OutFile string `help:"name of directory to save results to" arg:"-o,--outfile"`
	Workers int    `help:"number of threads for scanning" default:"20" arg:"-w,--workers"`
	NoColor bool   `help:"turn off color output" arg:"-n,--nocolor"`
}
