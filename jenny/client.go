package jenny

import (
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/pkg/browser"
)

var client *gojenkins.Jenkins

func (j *Jenkins) login() {
	c, e := gojenkins.CreateJenkins(nil, j.Uri, j.User, j.Password).Init()
	isLogin(e, func() {
		jobsNames = []prompt.Suggest{}
		getAllJobNames(c)
		client = c
	})
}

func isLogin(e error, callback func()) {
	if e != nil {
		color.Red("Something went wrong, check typos and profile info.")
		color.Yellow("For more information type: profile show -u")
		return
	} else {
		callback()
	}
}

func openLink(j Jenkins) {
	l := fmt.Sprintf("%s/%s", j.Uri, j.Project)
	browser.OpenURL(l)
}

func getStatus(c *gojenkins.Jenkins, name string, number int64) {
	getBuild(c, name, number, func(build *gojenkins.Build) {
		setColor(build.GetResult())
	})
}

func getLastStatus(c *gojenkins.Jenkins, name string) {
	getLastBuild(c, name, func(build *gojenkins.Build, isQueue bool) {
		s := ""

		if isQueue {
			s = "QUEUED"
		} else {
			s = build.GetResult()
		}
		setColor(s)
	})
}

func getLogs(c *gojenkins.Jenkins, name string, number int64) {
	getBuild(c, name, number, func(build *gojenkins.Build) {
		color.Yellow(build.GetConsoleOutput())
	})
}

func getLastLogs(c *gojenkins.Jenkins, name string) {
	getLastBuild(c, name, func(build *gojenkins.Build, isQueue bool) {
		if !isQueue {
			s := build.GetConsoleOutput()
			color.Yellow(s)
		}
	})
}

func stopExecution(c *gojenkins.Jenkins, name string, number int64) {
	getBuild(c, name, number, func(build *gojenkins.Build) {
		s, _ := build.Stop()

		if s {
			setColor("FAILURE")
		} else {
			setColor("SUCCESS")
		}
	})
}

func stopLastExecution(c *gojenkins.Jenkins, name string) {
	getLastBuild(c, name, func(build *gojenkins.Build, isQueue bool) {
		s, _ := build.Stop()

		if s {
			setColor("FAILURE")
		} else {
			setColor("SUCCESS")
		}
	})
}

func getBuild(c *gojenkins.Jenkins, name string, number int64, callback func(b *gojenkins.Build)) {
	build, e := c.GetBuild(name, number)
	isLogin(e, func() {
		callback(build)
	})
}

func getLastBuild(c *gojenkins.Jenkins, name string, callback func(b *gojenkins.Build, q bool)) {
	job, e := c.GetJob(name)
	isLogin(e, func() {
		isQueue, _ := job.IsQueued()
		build, _ := job.GetLastBuild()

		callback(build, isQueue)
	})
}

func setColor(s string) {
	d := color.New(color.FgWhite, color.Bold, color.BgRed)
	switch s {
	case "SUCCESS":
		d = color.New(color.FgWhite, color.Bold, color.BgCyan)
	case "FAILURE":
		d = color.New(color.FgWhite, color.Bold, color.BgRed)
	case "UNSTABLE":
		d = color.New(color.FgWhite, color.Bold, color.BgYellow)
	case "QUEUED":
		d = color.New(color.FgBlack, color.Bold, color.BgHiWhite)
	case "":
		d = color.New(color.FgWhite, color.Bold, color.BgHiBlack)
		s = "BUILDING"
	}
	d.Println(s)
}

//TODO
func stopBuild(c *gojenkins.Jenkins) {
	//build, _ := c.getBuild()
}

func getAllJobNames(c *gojenkins.Jenkins) {
	jobs, e := c.GetAllJobNames()
	isLogin(e, func() {
		for _, job := range jobs {
			jobName := []prompt.Suggest{
				{Text: job.Name},
			}
			jobsNames = append(jobsNames, jobName...)
		}
	})
}
