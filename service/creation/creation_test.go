package creation

import (
	"io/ioutil"
	"testing"
)

func TestGetAudio(t *testing.T) {
	//ssml := `<!--ID=B7267351-473F-409D-9765-754A8EBCDE05;Version=1|{\"VoiceNameToIdMapItems\":[{\"Id\":\"5f55541d-c844-4e04-a7f8-1723ffbea4a9\",\"Name\":\"Microsoft Server Speech Text to Speech Voice (zh-CN, XiaoxiaoNeural)\",\"ShortName\":\"zh-CN-XiaoxiaoNeural\",\"Locale\":\"zh-CN\",\"VoiceType\":\"StandardVoice\"}]}-->\n<!--ID=5B95B1CC-2C7B-494F-B746-CF22A0E779B7;Version=1|{\"Locales\":{\"zh-CN\":{\"AutoApplyCustomLexiconFiles\":[{}]}}}-->\n<speak version=\"1.0\" xmlns=\"http://www.w3.org/2001/10/synthesis\" xmlns:mstts=\"http://www.w3.org/2001/mstts\" xmlns:emo=\"http://www.w3.org/2009/10/emotionml\" xml:lang=\"zh-CN\"><voice name=\"zh-CN-XiaoxiaoNeural\"><mstts:express-as style=\"\"><prosody rate=\"0%\" contour=\"\"><say-as interpret-as=\"address\">陕西西安</say-as></prosody></mstts:express-as></voice></speak>`
	c := &Creation{}
	data, err := c.GetAudio("测试文本", "zh-CN-YunxiNeural", "0%", "general", "1.0",
		"default", "50%", "audio-24khz-160kbitrate-mono-mp3")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(len(data))
	err = ioutil.WriteFile("creation.mp3", data, 6666)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAuthToken(t *testing.T) {
	token, err := GetToken()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)
}

func TestVoices(t *testing.T) {
	token, err := GetToken()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(token)

	b, err := GetVoices(token)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(b))
}
