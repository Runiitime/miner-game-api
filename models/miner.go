package models

type MinerId struct {
	Id string `json:"id"`
}

type MinerInfo struct {
	Level       string `json:"level"`
	Salary      int64  `json:"salary"`
	Energy      int    `json:"energy"`
	Coal        int64  `json:"coal"`
	BreakTime   int    `json:"breakTime"`
	LevelUp     bool   `json:"levelUp"`
	LevelUpInfo string `json:"levelUpInfo"`
}

type MinerInfoMap = map[string]MinerInfo
