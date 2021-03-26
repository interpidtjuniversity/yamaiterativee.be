package command

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os/exec"
	"strings"
)

type gitCmdConfig struct {
	dir string
	env []string
	out io.Writer
}

func fetch(ctx context.Context, workingDir, upstream string, refspec ...string) error {
	args := append([]string{"fetch", "--tags", upstream}, refspec...)
	if err := execGitCmd(ctx, args, gitCmdConfig{dir: workingDir}); err != nil &&
		!strings.Contains(strings.ToLower(err.Error()), "couldn't find remote ref") {
		return errors.Wrap(err, fmt.Sprintf("git fetch --tags %s %s", upstream, refspec))
	}
	return nil
}

// runs a `git` command with the supplied arguments.
func execGitCmd(ctx context.Context, args []string, config gitCmdConfig) error {
	c := exec.CommandContext(ctx, "git", args...)

	if config.dir != "" {
		c.Dir = config.dir
	}
	//c.Env = append(env(), config.env...)
	return nil

}