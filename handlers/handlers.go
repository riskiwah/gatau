package handlers

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

func commitHash() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("%w; %s", err, stderr.String())
	}
	return stdout.String(), nil
}

// exported ShowToIndex
func ShowCommitHash(c *gin.Context) {
	commit, err := commitHash()
	if err != nil {
		panic(err)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"commit": commit,
	})

	// fmt.Printf("Commit : %s\n", commit)
	// c.JSON(200, commit)

}

// exported ShowHostname
func ShowHostname(c *gin.Context) {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"hostname": hostname,
	})
}

// some notes :
// https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
// https://github.com/maguec/micro-leaderboard
