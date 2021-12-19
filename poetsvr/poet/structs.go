package poet

type SlotInfo struct {
	SlotName       string
	SlotValue      string
	NormalizeValue string
}

type CommonReq struct {
	RequestId  string
	SessionId  string
	Query      string
	SkillId    string
	SkillName  string
	IntentId   uint64
	IntentName string
	Slots      []*SlotInfo //map: name -> value
	ClientIp   uint32
	Timestamp  int64
	Signature  string
	ThirdApiId uint64
}

// 成语接龙的session状态
type IdiomGameSessionInfo struct {
	Round     int
	ErrCount  int
	LastIdiom string
	UserIdiom []string
	BotIdiom  []string
}
