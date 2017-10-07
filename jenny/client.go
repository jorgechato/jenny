package jenny

import (
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"github.com/pkg/browser"
	"strings"
)

var client *gojenkins.Jenkins

func (j *Jenkins) login() {
	c, e := gojenkins.CreateJenkins(nil, j.Uri, j.User, j.Password).Init()
	isLogin(e, func() {
		jobsNames = []prompt.Suggest{}
		getAllJobNames(c, j)
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

func openLink(c *gojenkins.Jenkins, name string) {
	job, e := c.GetJob(name)
	isLogin(e, func() {
		uri := job.GetDetails()
		browser.OpenURL(uri.URL)
	})
}

func describeJob(c *gojenkins.Jenkins, name string) {
	job, e := c.GetJob(name)
	isLogin(e, func() {
		buildsIds, _ := job.GetAllBuildIds()
		max := 10
		if len(buildsIds) < 10 {
			max = len(buildsIds)
		}

		fmt.Println("BUILD ID\tSTATUS\t\t\tDATE\t\tTIME")

		for _, buildId := range buildsIds[0:max] {
			getBuild(c, name, buildId.Number, func(build *gojenkins.Build) {
				result := build.GetResult()
				time := build.GetTimestamp().Format("01/02 15:04:05")
				duration := float64(build.GetDuration()) / 1000.0

				fmt.Printf("%d\t\t%s\t\t%s\t\t%fs\n", buildId.Number, result, time, duration)
			})
		}
	})
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

func getAllJobNames(c *gojenkins.Jenkins, j *Jenkins) {
	jobs, e := c.GetAllJobNames()
	isLogin(e, func() {
		for _, job := range jobs {
			jn, jp := strings.ToUpper(job.Name), strings.ToUpper(j.Project)
			if strings.Contains(jn, jp) {
				jobName := []prompt.Suggest{
					{Text: job.Name},
				}
				jobsNames = append(jobsNames, jobName...)
			}
		}
	})
}

func build(c *gojenkins.Jenkins, name string) {
	c.BuildJob(name)
}
