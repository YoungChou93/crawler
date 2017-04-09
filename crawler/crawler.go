package crawler

import (
	"errors"
	"fmt"
	"github.com/YoungZhou93/crawler/controller"
	"time"
)

type Crawler struct {
	startTime string
	tasks     map[string]*controller.Controller
}

func NewCrawler() *Crawler {
	starttime := time.Now().Format("2006-01-02 15:04:05")
	newmap := make(map[string]*controller.Controller)
	return &Crawler{startTime: starttime, tasks: newmap}
}

func (c *Crawler) Register(name string, controller *controller.Controller) error {
	if _, exsits := c.tasks[name]; exsits {
		return errors.New("name is exsits")
	}
	c.tasks[name] = controller
	return nil
}

func (c *Crawler) GetTask(name string) (*controller.Controller, error) {
	if controller, exsits := c.tasks[name]; exsits {
		return controller, nil
	}
	return nil, errors.New("task does not exist")
}

func (c *Crawler) ListTask() {
	for name, _ := range c.tasks {
		fmt.Println(name)
	}
}

func (c *Crawler) Print() {
	fmt.Print("crawler start at")
	fmt.Println(c.startTime)
}

func  (c *Crawler) SetTime(time string){
	c.startTime=time
}