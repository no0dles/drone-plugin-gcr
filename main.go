package main

import (
	"os"
	"fmt"
	"strings"
	"os/exec"
	"time"
)

func main() {
	token := GetParameter("GOOGLE_TOKEN", "")
	repo := GetParameter("PLUGIN_REPO", "")

	registry := GetParameter("PLUGIN_REGISTRY", "gcr.io")
	dockerfile := GetParameter("PLUGIN_DOCKERFILE", "Dockerfile")
	buildPath := GetParameter("PLUGIN_BUILDPATH", ".")
	tagList := GetParameter("PLUGIN_TAGS", "latest")
	cacheFrom := GetParameter("PLUGIN_CACHE_FROM", "")

	token = strings.TrimSpace(token)

	if IsEmpty(token) {
		fmt.Println("Missing GOOGLE_TOKEN secret")
		os.Exit(1)
	}

	if IsEmpty(repo) {
		fmt.Println("Missing repo parameter")
		os.Exit(1)
	}

	StartDocker()

	tags := strings.Split(tagList, ",")
	args := []string{"build", "-f", dockerfile}
	if !IsEmpty(cacheFrom) {
		args = append(args, "--cache-from", fmt.Sprintf("%v/%v:%v", registry, repo, cacheFrom))
	}
	for _, tag := range tags {
		args = append(args, "-t", fmt.Sprintf("%v/%v:%v", registry, repo, tag))
	}
	args = append(args, buildPath)

	fmt.Printf("Building docker image %v:%v\n", repo, tagList)
	ExecuteCommand("docker", args...)

	fmt.Printf("Log in to %v\n", registry)
	ExecuteCommand("docker", "login", "-u", "_json_key", "-p", token, registry)

	for _, tag := range tags {
		image := fmt.Sprintf("%v/%v:%v", registry, repo, tag)
		fmt.Printf("Pushing image %v\n", image)
		ExecuteCommand("docker", "push", image)
	}
}

func GetParameter(name string, defaultValue string) string {
	value := os.Getenv(name)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

func IsEmpty(parameter string) bool {
	return len(parameter) == 0
}

func StartDocker() {
	cmd := exec.Command("/usr/local/bin/dockerd")
	go func() {
		cmd.Run()
	}()

	for i := 0; i < 15; i++ {
		cmd := exec.Command("docker", "info")
		err := cmd.Run()
		if err == nil {
			break
		}
		time.Sleep(time.Second * 1)
	}
}

func ExecuteCommand(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	var outStream OutStream
	var errStream ErrorStream

	cmd.Stdout = outStream
	cmd.Stderr = errStream

	err := cmd.Run()

	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
		os.Exit(1)
	}
}

type OutStream struct{}

func (out OutStream) Write(p []byte) (int, error) {
	os.Stdout.WriteString(string(p))
	return len(p), nil
}

type ErrorStream struct{}

func (out ErrorStream) Write(p []byte) (int, error) {
	os.Stderr.WriteString(string(p))
	return len(p), nil
}
