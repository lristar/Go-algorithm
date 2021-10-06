package media

import (
	"fmt"
	"reflect"
)

type IMp4 interface {
	PlayMp4()
}

type Mp4 struct {
	Music string
}

func (m *Mp4)PlayM(){
	m.PlayMp4()
}

func (m *Mp4) PlayMp4() {
	fmt.Println("music:", m.Music, "types:", reflect.TypeOf(m))
}
