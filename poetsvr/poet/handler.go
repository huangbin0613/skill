package poet

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"poetsvr/data"
	"poetsvr/schemago"
)

var SessionMap map[string]*IdiomGameSessionInfo

func init() {
	//load data
	SessionMap = make(map[string]*IdiomGameSessionInfo)
}

func Intent(writer http.ResponseWriter, request *http.Request) {
	//fmt.Println(request.Method, request.URL)

	bts, err := ioutil.ReadAll(request.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var req CommonReq
	// 1. 解析参数
	if err := json.Unmarshal(bts, &req); err != nil {
		fmt.Println("json Unmarshal err", err)
	}
	//fmt.Println("req", ToJson(req))

	// 1.1 解析slots为map
	slots := make(map[string]string)
	for _, slot := range req.Slots {
		if len(slot.NormalizeValue) > 0 {
			slots[slot.SlotName] = slot.NormalizeValue
		} else {
			slots[slot.SlotName] = slot.SlotValue
		}
	}
	//fmt.Println(slots)

	// 2. 定义返回
	var rsp schemago.SkillAnswer

	// 3. 执行意图获得结果

	switch req.IntentName {
	case "根据名字点诗":
		name := slots["诗歌名字"]
		fmt.Println(name)
		// todo 寻找这个古诗的内容 mysql : xorm

		fmt.Println(data.Poet{})

		author := "李白"
		poet := "床前明月光"
		rsp.Complete = true
		rsp.AnswerType = "text"
		text := fmt.Sprintf("好的，为你播放%v, %v %v %v。", name, name, author, poet)

		rsp.TextInfo = &schemago.TextInfo{
			ShortAnswer: text,
		}

	case "开始成语接龙":
		idiom := "半斤八两" //todo 随机搞一个成语  redis
		session := &IdiomGameSessionInfo{
			Round:     1,
			LastIdiom: idiom,
			BotIdiom:  []string{idiom},
		}
		// 存起来 todo 存redis
		SessionMap[req.SessionId] = session
		rsp.Complete = false
		rsp.AnswerType = "text"
		rsp.TextInfo = &schemago.TextInfo{
			ShortAnswer: fmt.Sprintf("好的，开始成语接龙，%v", idiom),
		}
	case "成语接龙":
		session := SessionMap[req.SessionId]
		if session == nil {
			fmt.Println("错误：session找不到")
			return
		}
		text := slots["成语"]
		// 1. 校验第一个字是否为上一次成语的结尾的拼音。
		// 2. 校验text 是否为正确的成语
		if 1 == 1 {
			// 用户是对的
			//session.UserIdiom.Add(text)
			session.UserIdiom = append(session.UserIdiom, text)
			// 重新出一个新成语, 去重
			newIdiom := "为所欲为"
			session.LastIdiom = newIdiom
			session.BotIdiom = append(session.BotIdiom, newIdiom)
			session.Round++
			// 更新保存
			SessionMap[req.SessionId] = session

			rsp.Complete = false
			rsp.AnswerType = "text"
			rsp.TextInfo = &schemago.TextInfo{
				ShortAnswer: fmt.Sprintf("%v", newIdiom),
			}
			fmt.Println(ToJson(session))
		} else {
			session.ErrCount++
			// todo 提示用户错误
		}
	case "退出成语接龙":
		rsp.Complete = true
		rsp.AnswerType = "text"
		rsp.TextInfo = &schemago.TextInfo{
			ShortAnswer: fmt.Sprintf("好的，已退出"),
		}
		delete(SessionMap, req.SessionId)
	default:

	}
	writer.Write([]byte(ToJson(rsp)))

}
