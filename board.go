package gochan

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const BoardsEndPoint = "https://a.4cdn.org/boards.json"

type CoolDown struct {
	Threads uint16 `json:"threads"`
	Replies uint16 `json:"replies"`
	Images  uint16 `json:"images"`
}

type Board struct {
	Board           string        `json:"board"`
	Title           string        `json:"title"`
	WsBoard         ChanBool      `json:"ws_board"`
	PerPage         uint8         `json:"per_page"`
	Pages           uint8         `json:"pages"`
	MaxFileSize     int64         `json:"max_filesize"`
	MaxWebmFileSize int64         `json:"max_webm_filesize"`
	MaxCommentChars uint16        `json:"max_comment_chars"`
	MaxWebmDuration time.Duration `json:"max_webm_duration"`
	BumpLimit       uint16        `json:"bump_limit"`
	ImageLimit      uint16        `json:"image_limit"`
	CoolDowns       CoolDown      `json:"cooldowns"`
	MetaDescription string        `json:"meta_description"`
	Spoilers        ChanBool      `json:"spoilers"`
	CustomSpoilers  uint16        `json:"custom_spoilers"`
	IsArchived      ChanBool      `json:"is_archived"`
	TrollFlags      ChanBool      `json:"troll_flags"`
	CountryFlags    ChanBool      `json:"country_flags"`
	UserIDs         ChanBool      `json:"user_ids"`
	Oekaki          ChanBool      `json:"oekaki"`
	TextOnly        ChanBool      `json:"text_only"`
	ForcedAnon      ChanBool      `json:"forced_anon"`
	WebmAudio       ChanBool      `json:"webm_audio"`
	RequireSubject  ChanBool      `json:"require_subject"`
	MinImageWidth   uint32        `json:"min_image_width"`
	MinImageHeight  uint32        `json:"min_image_height"`
}

type Boards struct {
	Boards []Board `json:"boards"`
}

func GetBoards() ([]Board, error) {
	resp, err := http.Get(BoardsEndPoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	boards := Boards{}
	if err := json.Unmarshal(body, &boards); err != nil {
		return nil, err
	}

	return boards.Boards, nil
}
