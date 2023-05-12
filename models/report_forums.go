package models

import (
    "log"
    "errors"
    "xyzforum/config"
)

type ReportForum struct {
    Id int `json:"id,omitempty"`
    IdForum int `json:"id_forum,omitempty"`
    IdForumOwner int `json:"id_forum_owner,omitempty"`
    IdUserReporter int `json:"id_user_reporter,omitempty"`
    Problem string `json:"problem,omitempty"`
    Created int `json:"created,omitempty"`
    Updated int `json:"updated,omitempty"`
    Status int `json:"status,omitempty"`
}

type ReportForumModel struct {}

func (reportForumModel *ReportForumModel) GetReportForums() []ReportForum {
    var reportForums []ReportForum
    var temp ReportForum

    rows, err := config.DB.Query("SELECT * FROM report_forums")
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.IdForum, &temp.IdForumOwner, &temp.IdUserReporter, &temp.Problem, &temp.Created, &temp.Updated, &temp.Status)
        if err != nil {
            log.Println(err)
        }

        reportForums = append(reportForums, temp)
    }

    return reportForums
}

func (reportForumModel *ReportForumModel) CreateReportForum(reportForum ReportForum) (int64, error) {
    result, err := config.DB.Exec("INSERT INTO report_forums (id_forum, id_forum_owner, id_user_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportForum.IdForum, reportForum.IdForumOwner, reportForum.IdUserReporter, reportForum.Problem, reportForum.Created, reportForum.Updated, reportForum.Status)
    if err != nil {
        log.Println(err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (reportForumModel *ReportForumModel) GetReportForumById(id int) (ReportForum, error) {
    var reportForum ReportForum
    rows, err := config.DB.Query("SELECT * FROM report_forums WHERE id = ?", id)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&reportForum.Id, &reportForum.IdForum, &reportForum.IdForumOwner, &reportForum.IdUserReporter, &reportForum.Problem, &reportForum.Created, &reportForum.Updated, &reportForum.Status)
        if err != nil {
            log.Println(err)
        }
    }

    return reportForum, nil
}

func (reportForumModel *ReportForumModel) GetReportForumByIdForum(idForum int) (ReportForum, error) {
    var reportForum ReportForum
    rows, err := config.DB.Query("SELECT * FROM report_forums WHERE id_forum = ?", idForum)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&reportForum.Id, &reportForum.IdForum, &reportForum.IdForumOwner, &reportForum.IdUserReporter, &reportForum.Problem, &reportForum.Created, &reportForum.Updated, &reportForum.Status)
        if err != nil {
            log.Println(err)
            return reportForum, errors.New("Report Forum not found")
        }
    }

    return reportForum, nil
}

func (reportForumModel *ReportForumModel) GetReportForumByIdForumOwner(idForumOwner int) (ReportForum, error) {
    var reportForum ReportForum
    rows, err := config.DB.Query("SELECT * FROM report_forums WHERE id_forum_owner = ?", idForumOwner)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&reportForum.Id, &reportForum.IdForum, &reportForum.IdForumOwner, &reportForum.IdUserReporter, &reportForum.Problem, &reportForum.Created, &reportForum.Updated, &reportForum.Status)
        if err != nil {
            log.Println(err)
        }
    }

    return reportForum, nil
}

func (reportForumModel *ReportForumModel) GetReportForumByIdUserReporter(idUserReporter int) (ReportForum, error) {
    var reportForum ReportForum
    rows, err := config.DB.Query("SELECT * FROM report_forums WHERE id_user_reporter = ?", idUserReporter)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&reportForum.Id, &reportForum.IdForum, &reportForum.IdForumOwner, &reportForum.IdUserReporter, &reportForum.Problem, &reportForum.Created, &reportForum.Updated, &reportForum.Status)
        if err != nil {
            log.Println(err)
        }
    }

    return reportForum, nil
}

func (reportForumModel *ReportForumModel) UpdateReportForum(id int, updated int) (int64, error) {
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

func (reportForumModel *ReportForumModel) DeleteReportForum(id int) (int64, error) {
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