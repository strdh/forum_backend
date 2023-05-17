package models

import (
    "log"
    "errors"
    "xyzforum/config"
)

type ReportMessage struct {
    Id int `json:"id,omitempty"`
    IdMessage int `json:"id_message,omitempty"`
    IdOwner int `json:"id_owner,omitempty"`
    IdReporter int `json:"id_reporter,omitempty"`
    Problem string `json:"problem,omitempty"`
    Created int64 `json:"created,omitempty"`
    Updated int64 `json:"updated,omitempty"`
    Status int `json:"status,omitempty"`
}

type ReportMessageModel struct {}

func (reportMessageModel *ReportMessageModel) ReportMessages() []ReportMessage {
    var reportMessages []ReportMessage
    var temp ReportMessage

    rows, err := config.DB.Query("SELECT * FROM report_messages")
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.IdMessage, &temp.IdOwner, &temp.IdReporter, &temp.Problem, &temp.Created, &temp.Updated, &temp.Status)
        if err != nil {
            log.Println(err)
        }

        reportMessages = append(reportMessages, temp)
    }

    return reportMessages
}

func (reportMessageModel *ReportMessageModel) Create(reportMessage ReportMessage) (int64, error) {
    result, err := config.DB.Exec("INSERT INTO report_messages (id_message, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportMessage.IdMessage, reportMessage.IdOwner, reportMessage.IdReporter, reportMessage.Problem, reportMessage.Created, reportMessage.Updated, reportMessage.Status)
    if err != nil {
        log.Println(err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (reportMessageModel *ReportMessageModel) ById(id int) (ReportMessage, error) {
    var reportMessage ReportMessage
    rows, err := config.DB.Query("SELECT * FROM report_messages WHERE id = ?", id)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&reportMessage.Id, &reportMessage.IdMessage, &reportMessage.IdOwner, &reportMessage.IdReporter, &reportMessage.Problem, &reportMessage.Created, &reportMessage.Updated, &reportMessage.Status)
        if err != nil {
            log.Println(err)
        }

        return reportMessage, nil
    } else {
        return reportMessage, errors.New("ReportMessage not found")
    }
}

func (reportMessageModel *ReportMessageModel) ByIdMessage(idMessage int) []ReportMessage {
    var reportMessages []ReportMessage
    var temp ReportMessage
    rows, err := config.DB.Query("SELECT * FROM report_messages WHERE id_message = ?", idMessage)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.IdMessage, &temp.IdOwner, &temp.IdReporter, &temp.Problem, &temp.Created, &temp.Updated, &temp.Status)
        if err != nil {
            log.Println(err)
        }

        reportMessages = append(reportMessages, temp)
    } 

    return reportMessages
}

func (reportMessageModel *ReportMessageModel) ByIdOwner(IdOwner int) []ReportMessage {
    var reportMessages []ReportMessage
    var temp ReportMessage
    rows, err := config.DB.Query("SELECT * FROM report_messages WHERE id_owner = ?", IdOwner)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.IdMessage, &temp.IdOwner, &temp.IdReporter, &temp.Problem, &temp.Created, &temp.Updated, &temp.Status)
        if err != nil {
            log.Println(err)
        }

        reportMessages = append(reportMessages, temp)
    }

    return reportMessages
}

func (reportMessageModel *ReportMessageModel) ByIdReporter(IdReporter int) []ReportMessage {
    var reportMessages []ReportMessage
    var temp ReportMessage
    rows, err := config.DB.Query("SELECT * FROM report_messages WHERE id_reporter = ?", IdReporter)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.IdMessage, &temp.IdOwner, &temp.IdReporter, &temp.Problem, &temp.Created, &temp.Updated, &temp.Status)
        if err != nil {
            log.Println(err)
        }

        reportMessages = append(reportMessages, temp)
    }

    return reportMessages
}

//Get forum id of message and id user of message
func (reportMessageModel *ReportMessageModel) GetIdFO(id int) (int, int) {
    var idForum int
    var idOwner int

    err := config.DB.QueryRow("SELECT id_forum, id_user FROM forum_messages WHERE id = ?", id).Scan(&idForum, &idOwner)
    if err != nil {
        log.Println(err)
        return 0, 0
    }

    return idForum, idOwner
}

func (reportMessageModel *ReportMessageModel) Update(id int, status int) error {
    _, err := config.DB.Exec("UPDATE report_messages SET status = ? WHERE id = ?", status, id)
    if err != nil {
        log.Println(err)
        return err
    }

    return nil
}

func (reportMessageModel *ReportMessageModel) Delete(id int) error {
    _, err := config.DB.Exec("DELETE FROM report_messages WHERE id = ?", id)
    if err != nil {
        log.Println(err)
        return err
    }

    return nil
}