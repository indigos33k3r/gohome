package camera

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/barnybug/gohome/config"
)

type Motion struct {
	Conf config.CameraNodeConf
}

func (self *Motion) GotoPreset(preset int) error {
	return nil
}

func (self *Motion) Snapshot() (io.ReadCloser, error) {
	return nil, nil
}

func (self *Motion) Video(filename string, duration time.Duration) error {
	return nil
}

func (self *Motion) Ir(b bool) error {
	return nil
}

var DetectCommands = map[bool]string{
	true:  "detection/start",
	false: "detection/pause",
}

func (self *Motion) Detect(b bool) (err error) {
	url := fmt.Sprintf("%s/%s", self.Conf.Url, DetectCommands[b])
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}
