package models

import (
    "log"
    "xyzforum/config"
)

type ReportForum struct {
    Id int `json:"id,omitempty"`
    IdForum int `json:"id_forum,omitempty"`
    IdOwner int `json:"id_owner,omitempty"`
    IdReporter int `json:"id_reporter,omitempty"`
    Problem string `json:"problem,omitempty"`
    Created int64 `json:"created,omitempty"`
    Updated int64 `json:"updated,omitempty"`
    Status int `json:"status,omitempty"`
}

type ReportForumModel struct {}

func (reportForumModel *ReportForumModel) ReportForums() []ReportForum {
    var reportForums []ReportForum
    var temp ReportForum

    rows, err := config.DB.Query("SELECT * FROM report_forums")
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.IdForum, &temp.IdOwner, &temp.IdReporter, &temp.Problem, &temp.Created, &temp.Updated, &temp.Status)
        if err != nil {
            log.Println(err)
        }

        reportForums = append(reportForums, temp)
    }

    return reportForums
}

func (reportForumModel *ReportForumModel) Create(reportForum ReportForum) (int64, error) {
    result, err := config.DB.Exec("INSERT INTO report_forums (id_forum, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportForum.IdForum, reportForum.IdOwner, reportForum.IdReporter, reportForum.Problem, reportForum.Created, reportForum.Updated, reportForum.Status)
    if err != nil {
        log.Println(err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (reportForumModel *ReportForumModel) ById(id int) (ReportForum, error) {
    var reportForum ReportForum
    rows, err := config.DB.Query("SELECT * FROM report_forums WHERE id = ?", id)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&reportForum.Id, &reportForum.IdForum, &reportForum.IdOwner, &reportForum.IdReporter, &reportForum.Problem, &reportForum.Created, &reportForum.Updated, &reportForum.Status)
        if err != nil {
            log.Println(err)
        }
    }

    return reportForum, nil
}

func (reportForumModel *ReportForumModel) ByIdForum(idForum int) []ReportForum {
    var reportForums []ReportForum
    var temp ReportForum

    rows, err := config.DB.Query("SELECT * FROM report_forums WHERE id_forum = ?", idForum)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.IdForum, &temp.IdOwner, &temp.IdReporter, &temp.Problem, &temp.Created, &temp.Updated, &temp.Status)
        if err != nil {
            log.Println(err)
        }

         reportForums = append(reportForums, temp)
    }

    return reportForums
}

func (reportForumModel *ReportForumModel) ByIdOwner(IdOwner int) []ReportForum {
    var reportForums []ReportForum
    var temp ReportForum

    rows, err := config.DB.Query("SELECT * FROM report_forums WHERE id_owner = ?", IdOwner)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.IdForum, &temp.IdOwner, &temp.IdReporter, &temp.Problem, &temp.Created, &temp.Updated, &temp.Status)
        if err != nil {
            log.Println(err)
        }

        reportForums = append(reportForums, temp)
    }

    return reportForums
}

func (reportForumModel *ReportForumModel) ByIdReporter(IdReporter int) []ReportForum {
    var reportForums []ReportForum
    var temp ReportForum

    rows, err := config.DB.Query("SELECT * FROM report_forums WHERE id_reporter = ?", IdReporter)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.IdForum, &temp.IdOwner, &temp.IdReporter, &temp.Problem, &temp.Created, &temp.Updated, &temp.Status)
        if err != nil {
            log.Println(err)
        }

        reportForums = append(reportForums, temp)
    }

    return reportForums
}

//get id owner forum
func (reportForumModel *ReportForumModel) GetIdOwner(id int) int {
    var idOwner int

    err := config.DB.QueryRow("SELECT id_user FROM forums WHERE id = ?", id).Scan(&idOwner)

    if err != nil {
        log.Println(err)
        return 0
    }

    return idOwner
}

func (reportForumModel *ReportForumModel) Update(id int, updated int64) (int64, error) {
    result, err := config.DB.Exec("UPDATE report_forums SET updated = ?, status = 1 WHERE id = ?", updated, id)
    if err != nil {
        log.Println(err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }

    return rowsAffected, nil
}

func (reportForumModel *ReportForumModel) Delete(id int) (int64, error) {
    result, err := config.DB.Exec("DELETE FROM report_forums WHERE id = ?", id)
    if err != nil {
        log.Println(err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }

    return rowsAffected, nil
}