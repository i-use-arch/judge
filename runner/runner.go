package runner

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/i-use-arch/judge/dbconn"
)

// Runner lets you run file
type Runner struct {
	Client *dbconn.Client
}

func (r *Runner) Run(id uint64) error {

	fileContent, err := r.Client.GetSubmission(id)

	if err != nil {
		return err
	}

	dir, err := ioutil.TempDir("submissions", fmt.Sprintf("%d", id))

	if err != nil {
		return err
	}

	outFile, err := ioutil.TempFile(dir, "submission")

	if err != nil {
		return err
	}

	_, err = outFile.WriteString(fileContent)

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// cmd := exec.CommandContext(ctx, "go", "run", "test/test.go")
	name := filepath.Base(outFile.Name())
	fmt.Printf("running file %s\n", name)
	cmd := exec.CommandContext(ctx, "python", name)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Println("running job")
	cmd.Run()
	fmt.Println("done running job")

	cmd = exec.CommandContext(context.Background(), "rm", "-rf", dir)
	cmd.Run()

	fmt.Println("done cleaning up")
	foobar(cancel, cmd)

	return nil
}

func foobar(cancel context.CancelFunc, cmd *exec.Cmd) {

}
