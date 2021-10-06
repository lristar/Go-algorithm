package structureType

import (
	"Go-algorithm/src/design-pattern/structureType/adaptType"
	"Go-algorithm/src/design-pattern/structureType/adaptType/media"
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

