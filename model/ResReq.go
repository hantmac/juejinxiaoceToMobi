package model

import "time"

type JueJinResponse struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		ID               string        `json:"_id"`
		Section          []string      `json:"section"`
		Desc             string        `json:"desc"`
		User             string        `json:"user"`
		Price            float64       `json:"price"`
		Title            string        `json:"title"`
		BuyCount         int           `json:"buyCount"`
		ViewCount        int           `json:"viewCount"`
		ContentSize      int           `json:"contentSize"`
		Img              string        `json:"img"`
		Category         string        `json:"category"`
		Tags             []interface{} `json:"tags"`
		SummaryHTML      string        `json:"summaryHtml"`
		CreatedAt        time.Time     `json:"createdAt"`
		UpdatedAt        time.Time     `json:"updatedAt"`
		FinishedAt       string        `json:"finishedAt"`
		IsFinished       bool          `json:"isFinished"`
		IsDeleted        bool          `json:"isDeleted"`
		IsHot            bool          `json:"isHot"`
		IsPublish        int           `json:"isPublish"`
		IsShow           bool          `json:"isShow"`
		Profile          string        `json:"profile"`
		LastSectionCount int           `json:"lastSectionCount"`
		Pv               int           `json:"pv"`
		WechatSignal     string        `json:"wechatSignal"`
		IsTop            bool          `json:"isTop"`
		WechatImg        string        `json:"wechatImg"`
		WechatImgDesc    string        `json:"wechatImgDesc"`
		NavID            []string      `json:"navId"`
		URL              string        `json:"url"`
		IsEditor         bool          `json:"isEditor"`
		IsBuy            bool          `json:"isBuy"`
		UserData         struct {
			Role                string `json:"role"`
			Username            string `json:"username"`
			SelfDescription     string `json:"selfDescription"`
			JobTitle            string `json:"jobTitle"`
			Company             string `json:"company"`
			AvatarHd            string `json:"avatarHd"`
			AvatarLarge         string `json:"avatarLarge"`
//			MobilePhoneVerified bool   `json:"mobilePhoneVerified"`
			IsXiaoceAuthor      string `json:"isXiaoceAuthor"`
			BookletCount        int    `json:"bookletCount"`
			ObjectID            string `json:"objectId"`
			UID                 string `json:"uid"`
		} `json:"userData"`
		ReadLog struct {
			SectionID string `json:"sectionId"`
			Sign      string `json:"sign"`
		} `json:"readLog"`
		Poster string `json:"poster"`
	} `json:"d"`
}
type JueJinRequestBody struct {
	Token    string `json:"token"`
	ClientID string `json:"client_id"`
	Src      string `json:"src"`
	UID      string `json:"uid"`
	ID       string `json:"id"`
}

//type GetSectionResponse struct {
//	S int    `json:"s"`
//	M string `json:"m"`
//	D struct {
//		ID           string    `json:"_id"`
//		Title        string    `json:"title"`
//		IsFree       bool      `json:"isFree"`
//		IsFinished   bool      `json:"isFinished"`
//		User         string    `json:"user"`
//		ViewCount    int       `json:"viewCount"`
//		MetaID       string    `json:"metaId"`
//		Content      string    `json:"content"`
//		ContentSize  int       `json:"contentSize"`
//		HTML         string    `json:"html"`
//		CreatedAt    time.Time `json:"createdAt"`
//		UpdatedAt    time.Time `json:"updatedAt"`
//		IsDeleted    bool      `json:"isDeleted"`
//		Pv           int       `json:"pv"`
//		CommentCount int       `json:"commentCount"`
//		SectionID    string    `json:"sectionId"`
//	} `json:"d"`
//}


type GetSectionResponse struct {
	S int    `json:"s"`
	M string `json:"m"`
	D struct {
		ID           string    `json:"_id"`
		Title        string    `json:"title"`
		IsFree       bool      `json:"isFree"`
		IsFinished   bool      `json:"isFinished"`
		User         string    `json:"user"`
		ViewCount    int       `json:"viewCount"`
		MetaID       string    `json:"metaId"`
		Content      string    `json:"content"`
		ContentSize  int       `json:"contentSize"`
		HTML         string    `json:"html"`
		CreatedAt    time.Time `json:"createdAt"`
		UpdatedAt    time.Time `json:"updatedAt"`
		IsDeleted    bool      `json:"isDeleted"`
		Pv           int       `json:"pv"`
		CommentCount int       `json:"commentCount"`

		SectionID    string    `json:"sectionId"`
	} `json:"d"`
}

