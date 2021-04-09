package main

import (
	"flag"
	"fmt"
	"github.com/pelletier/go-toml"
	"os/exec"
)

var index int

func init() {
	flag.IntVar(&index, "i", 0, "用户索引数")
}

func main() {
	flag.Parse()
	repo, user, email, password := parse(index)
	remoteUrl := repoUrl(repo, user, password)
	setup(remoteUrl, user, email)
	fmt.Printf("switch %s done!", user)
}

func parse(index int) (repo string, user string, email string, password string) {
	config, _ := toml.LoadFile("./config.toml")
	repo = config.Get("repo").(string)
	user = config.Get("users").([]interface{})[index].(string)
	email = config.Get("emails").([]interface{})[index].(string)
	password = config.Get("passwords").([]interface{})[index].(string)
	return
}

func repoUrl(repo string, user string, password string) (url string) {
	url = fmt.Sprintf("http://%s:%s@%s", user, password, repo)
	return
}

func setup(remoteUrl string, user string, email string) {
	cmd := exec.Command("git", "config", "user.name", user)
	err := cmd.Run()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
	cmd = exec.Command("git", "config", "user.email", email)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
	cmd = exec.Command("git", "remote", "set-url", "origin", remoteUrl)
	err = cmd.Run()
	if err != nil {
		fmt.Println("Execute Command failed:" + err.Error())
		return
	}
}
