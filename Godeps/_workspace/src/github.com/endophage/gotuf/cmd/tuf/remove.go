package main

import (
	"errors"

	"github.com/endophage/gotuf"
	"github.com/flynn/go-docopt"
)

func init() {
	register("remove", cmdRemove, `
usage: tuf remove [--expires=<days>] [--all] [<path>...]

Remove target file(s).

Options:
  --all              Remove all target files.
  --expires=<days>   Set the targets manifest to expire <days> days from now.
`)
}

func cmdRemove(args *docopt.Args, repo *tuf.Repo) error {
	paths := args.All["<path>"].([]string)
	if len(paths) == 0 && !args.Bool["--all"] {
		return errors.New("either specify some paths or set the --all flag to remove all targets")
	}
	if arg := args.String["--expires"]; arg != "" {
		expires, err := parseExpires(arg)
		if err != nil {
			return err
		}
		return repo.RemoveTargetsWithExpires(expires, paths...)
	}
	return repo.RemoveTargets(paths)
}
