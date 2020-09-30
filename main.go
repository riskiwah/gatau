package main

import (
	"gatau/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", handlers.ShowCommitHash)
	// r.GET("/", handlers.ShowHostname)
	r.Run()
}

// func commitHash() (string, error) {
// 	cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
// 	var stdout, stderr bytes.Buffer
// 	cmd.Stdout = &stdout
// 	cmd.Stderr = &stderr
// 	if err := cmd.Run(); err != nil {
// 		return "", fmt.Errorf("%w; %s", err, stderr.String())
// 	}
// 	return stdout.String(), nil
// }

// func ShowCommitHash(c *gin.Context) {
// 	commit, err := commitHash()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Printf("Commit : %s\n", commit)

// }

// func main() {
// 	r := gin.Default()
// 	r.GET("/", ShowCommitHash)
// 	r.Run()
// }

// func main() {
// 	commit, err := commitHash()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Commit : %s\n", commit)
// }
