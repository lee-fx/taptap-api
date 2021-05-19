package dbops

import (
	"api/app/defs"
	"api/app/utils"
)

func AddNewComments(vid string, aid int, comment string) error {
	uuid := utils.NewUUID()
	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(uuid, vid, aid, comment)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.login_name, comments.content FROM comments INNER JOIN users ON comments.author_id=users.id
					WHERE comments.id=? And comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)
	if err != nil {
		return nil, err
	}
	var res []*defs.Comment
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var id, name, comment string
		if err := rows.Scan(&id, &name, &comment); err != nil {
			return res, err
		}
		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: comment}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}
