package dbops

import (
	"api/app/defs"
	"api/app/utils"
	"database/sql"
	"log"
	"time"
)

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid := utils.NewUUID()
	t := time.Now()
	ctime := t.Format("2006-01-02 15:04:05") // M D  y, HH:MM:SS
	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info
	(id, author_id, name, display_ctime) VALUES (?, ?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmtIns.Close()
	return res, nil
}

func GetVideoCredential(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id = ?")
	if err != nil {
		log.Printf("get Video error: %s", err)
		return nil, err
	}

	var aid int
	var name string
	var display_ctime string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &display_ctime)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	stmtOut.Close()
	return &defs.VideoInfo{
		Id:           vid,
		AuthorId:     aid,
		Name:         name,
		DisplayCtime: display_ctime,
	}, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		log.Printf("delete video error: %s", err)
	}
	_, err = stmtDel.Exec(vid)
	defer stmtDel.Close()
	return nil
}
