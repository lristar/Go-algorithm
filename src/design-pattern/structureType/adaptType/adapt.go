package adaptType

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

var injectTagName = "inject"

type IMediaPlayer interface {
	Play(types string)
}

type IMediaAdapter interface {
	PlayM()
}

type MediaAdapter struct {
	sync.Mutex
	MediaList map[string]interface{}
	Go        IMediaAdapter `inject:"IMedia"`
}

func (m *MediaAdapter) AddMedia(key string, v interface{}) {
	m.Lock()
	m.MediaList[key] = v
	m.Unlock()
}

func (m *MediaAdapter) DelMedia(key string) {
	m.Lock()
	delete(m.MediaList, key)
	m.Unlock()
}

func (m *MediaAdapter)InitGo(types string)error{
	v, ok := m.MediaList[types]
	if !ok {
		fmt.Println("适配器不支持这个")
		return errors.New("适配器不支持该类型")
	}
	ma := reflect.ValueOf(m)
	elemValue := ma.Elem()
	elemValue.FieldByName("Go").Set(reflect.ValueOf(v))
	if err:=recover();err!=nil{
		return err.(error)
	}
	return nil
}

type MediaPlayer struct {
	mediaAdapter *MediaAdapter
}

func (m *MediaPlayer) CreateMediaAdapter(mad *MediaAdapter) {
	m.mediaAdapter = mad
}


func (m *MediaPlayer) Player(types string) {
	if err:=m.mediaAdapter.InitGo(types);err!=nil{
		fmt.Println("初始化失败")
		return
	}
	m.mediaAdapter.Go.PlayM()
}
