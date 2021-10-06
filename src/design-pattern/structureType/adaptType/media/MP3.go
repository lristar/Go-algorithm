package media

import (
	"fmt"
	"reflect"
)

type IMp3 interface {
	PlayMp3()
}

type Mp3 struct {
	Music string
}

func (m *Mp3)PlayM(){
	m.PlayMp3()
}

func (m *Mp3) PlayMp3() {
	fmt.Println("music:", m.Music, "types:", reflect.TypeOf(m))
}
