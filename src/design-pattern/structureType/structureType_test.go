package structureType

import (
	"Go-algorithm/src/design-pattern/structureType/adaptType"
	"Go-algorithm/src/design-pattern/structureType/adaptType/media"
	"Go-algorithm/src/design-pattern/structureType/agentType"
	"Go-algorithm/src/design-pattern/structureType/agentType/person"
	"testing"
)

func Test_Adapt(t *testing.T) {
	madapt := adaptType.MediaAdapter{MediaList: make(map[string]interface{})}
	madapt.AddMedia("MP3", &media.Mp3{Music: "稻香"})
	madapt.AddMedia("MP4", &media.Mp4{Music: "七里香"})
	mplayer := adaptType.MediaPlayer{}
	mplayer.CreateMediaAdapter(&madapt)
	mplayer.Player("MP6")
}

func Test_proxy(t *testing.T) {
	proxy:=agentType.NewProxy()
	proxy.AddServers("Apple",&person.Server{
		ProductName:    "Apple",
		ProductBalance: 1.99,
		ProductNum:     50,
		Values:         0,
	})
	proxy.AddServers("orange",&person.Server{
		ProductName:    "orange",
		ProductBalance: 3.99,
		ProductNum:     10,
		Values:         0,
	})
	proxy.AddServers("banana",&person.Server{
		ProductName:    "banana",
		ProductBalance: 0.99,
		ProductNum:     100,
		Values:         0,
	})
	proxy.ShowProduct()
	// 创建客户
	proxy.AddClient(person.NewClient(60.00))
	proxy.SellProduct("banana",10)
	proxy.SellProduct("orange",2)
	proxy.SellProduct("orange",8)
	proxy.SellProduct("Apple",10)
}

