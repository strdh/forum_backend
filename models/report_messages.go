package models

import (
    "log"
    "errors"
    "xyzforum/config"
)

type ReportMessage struct {
    Id int `json:"id,omitempty"`
    IdMessage int `json:"id_message,omitempty"`
    IdMessageOwner int `json:"id_message_owner,omitempty"`
    IdUserReporter int `json:"id_user_reporter,omitempty"`
    Problem string `json:"problem,omitempty"`
    Created int `json:"created,omitempty"`
    Updated int `json:"updated,omitempty"`
    Status int `json:"status,omitempty"`
}

type ReportMessageModel struct {}

func (reportMessageModel *ReportMessageModel) GetReportMessages() []ReportMessage {
    var reportMessages []ReportMessage
    var temp ReportMessage

    rows, err := config.DB.Query("SELECT * FROM report_messages")
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.IdMessage, &temp.IdMessageOwner, &temp.IdUserReporter, &temp.Problem, &temp.Created, &temp.Updated, &temp.Status)
        if err != nil {
            log.Println(err)
        }

        reportMessages = append(reportMessages, temp)
    }

    return reportMessages
}

func (reportMessageModel *ReportMessage) CreateReportMessage(reportMessage ReportMessage) (int64, error) {
    result, err := config.DB.Exec("INSERT INTO report_messages (id_message, id_message_owner, id_user_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportMessage.IdMessage, reportMessage.IdMessageOwner, reportMessage.IdUserReporter, reportMessage.Problem, reportMessage.Created, reportMessage.Updated, reportMessage.Status)
    if err != nil {
        log.Println(err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (reportMessageModel *ReportMessageModel) GetReportMessageById(id int) (ReportMessage, error) {
    var reportMessage ReportMessage
    rows, err := config.DB.Query("SELECT * FROM report_messages WHERE id = ?", id)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&reportMessage.Id, &reportMessage.IdMessage, &reportMessage.IdMessageOwner, &reportMessage.IdUserReporter, &reportMessage.Problem, &reportMessage.Created, &reportMessage.Updated, &reportMessage.Status)
        if err != nil {
            log.Println(err)
        }

        return reportMessage, nil
    } else {
        return reportMessage, errors.New("ReportMessage not found")
    }
}

func (reportMessageModel *ReportMessageModel) GetReportMessageByIdMessage(idMessage int) (ReportMessage, error) {
    var reportMessage ReportMessage
    rows, err := config.DB.Query("SELECT * FROM report_messages WHERE id_message = ?", idMessage)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&reportMessage.Id, &reportMessage.IdMessage, &reportMessage.IdMessageOwner, &reportMessage.IdUserReporter, &reportMessage.Problem, &reportMessage.Created, &reportMessage.Updated, &reportMessage.Status)
        if err != nil {
            log.Println(err)
        }

        return reportMessage, nil
    } else {
        return reportMessage, errors.New("ReportMessage not found")
    }
}

func (reportMessageModel *ReportMessageModel) GetReportMessageByIdMessageOwner(idMessageOwner int) (ReportMessage, error) {
    var reportMessage ReportMessage
    rows, err := config.DB.Query("SELECT * FROM report_messages WHERE id_message_owner = ?", idMessageOwner)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&reportMessage.Id, &reportMessage.IdMessage, &reportMessage.IdMessageOwner, &reportMessage.IdUserReporter, &reportMessage.Problem, &reportMessage.Created, &reportMessage.Updated, &reportMessage.Status)
        if err != nil {
            log.Println(err)
        }

        return reportMessage, nil
    } else {
        return reportMessage, errors.New("ReportMessage not found")
    }
}

func (reportMessageModel *ReportMessageModel) GetReportMessageByIdUserReporter(idUserReporter int) (ReportMessage, error) {
    var reportMessage ReportMessage
    rows, err := config.DB.Query("SELECT * FROM report_messages WHERE id_user_reporter = ?", idUserReporter)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&reportMessage.Id, &reportMessage.IdMessage, &reportMessage.IdMessageOwner, &reportMessage.IdUserReporter, &reportMessage.Problem, &reportMessage.Created, &reportMessage.Updated, &reportMessage.Status)
        if err != nil {
            log.Println(err)
        }

        return reportMessage, nil
    } else {
        return reportMessage, errors.New("ReportMessage not found")
    }
}

func (reportMessageModel *ReportMessageModel) UpdateReportMessage(id int, reportMessage ReportMessage) error {
    _, err := config.DB.Exec("UPDATE report_messages SET status = ? WHERE id = ?", reportMessage.Status, id)
    if err != nil {
        log.Println(err)
        return err
    }

    return nil
}

func (reportMessageModel *ReportMessageModel) DeleteReportMessage(id int) error {
    _, err := config.DB.Exec("DELETE FROM report_messages WHERE id = ?", id)
    if err != nil {
        log.Println(err)
        return err
    }

    return nil
}