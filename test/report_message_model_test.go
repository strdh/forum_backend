package test

import (
    "testing"
    "time"
    "xyzforum/config"
    "xyzforum/models"
    "github.com/stretchr/testify/assert"
)

func init() {
    config.InitializeTestDB()
}

//RM stand for ReportMessage

func TestRMReportMessages(t *testing.T) {
    reportMessageModel := models.ReportMessageModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE report_messages")
    if err != nil {
        t.Error(err)
    }

    reportMessage := models.ReportMessage{
        IdMessage: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_messages (id_message, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportMessage.IdMessage, reportMessage.IdOwner, reportMessage.IdReporter, reportMessage.Problem, reportMessage.Created, reportMessage.Updated, reportMessage.Status)
    if err != nil {
        t.Error(err)
    }

    reportMessages := reportMessageModel.ReportMessages()
    assert.Equal(t, 1, len(reportMessages))
    assert.Equal(t, "testproblem", reportMessages[0].Problem)
    assert.Equal(t, 1, reportMessages[0].Status)
}

func TestRMCreate(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE report_messages")
    if err != nil {
        t.Error(err)
    }

    reportMessageModel := models.ReportMessageModel{}

    reportMessage := models.ReportMessage{
        IdMessage: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_messages (id_message, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportMessage.IdMessage, reportMessage.IdOwner, reportMessage.IdReporter, reportMessage.Problem, reportMessage.Created, reportMessage.Updated, reportMessage.Status)
    if err != nil {
        t.Error(err)
    }

    id, err := reportMessageModel.Create(reportMessage)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, int64(2), id)
}

func TestRMById(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE report_messages")
    if err != nil {
        t.Error(err)
    }

    reportMessageModel := models.ReportMessageModel{}

    reportMessage := models.ReportMessage{
        IdMessage: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_messages (id_message, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportMessage.IdMessage, reportMessage.IdOwner, reportMessage.IdReporter, reportMessage.Problem, reportMessage.Created, reportMessage.Updated, reportMessage.Status)
    if err != nil {
        t.Error(err)
    }

    reportMessage, err = reportMessageModel.ById(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, "testproblem", reportMessage.Problem)
    assert.Equal(t, 1, reportMessage.Status)
}

func TestRMByIdOwner(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE report_messages")
    if err != nil {
        t.Error(err)
    }

    reportMessageModel := models.ReportMessageModel{}

    reportMessage := models.ReportMessage{
        IdMessage: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_messages (id_message, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportMessage.IdMessage, reportMessage.IdOwner, reportMessage.IdReporter, reportMessage.Problem, reportMessage.Created, reportMessage.Updated, reportMessage.Status)
    if err != nil {
        t.Error(err)
    }

    reportMessages := reportMessageModel.ByIdOwner(1)
    assert.Equal(t, 1, len(reportMessages))
    assert.Equal(t, "testproblem", reportMessages[0].Problem)
    assert.Equal(t, 1, reportMessages[0].Status)
}

func TestRMByIdReporter(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE report_messages")
    if err != nil {
        t.Error(err)
    }

    reportMessageModel := models.ReportMessageModel{}

    reportMessage := models.ReportMessage{
        IdMessage: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_messages (id_message, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportMessage.IdMessage, reportMessage.IdOwner, reportMessage.IdReporter, reportMessage.Problem, reportMessage.Created, reportMessage.Updated, reportMessage.Status)
    if err != nil {
        t.Error(err)
    }

    reportMessages := reportMessageModel.ByIdReporter(1)
    assert.Equal(t, 1, len(reportMessages))
    assert.Equal(t, "testproblem", reportMessages[0].Problem)
    assert.Equal(t, 1, reportMessages[0].Status)
}

func TestRMUpdate(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE report_messages")
    if err != nil {
        t.Error(err)
    }

    reportMessageModel := models.ReportMessageModel{}

    reportMessage := models.ReportMessage{
        IdMessage: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_messages (id_message, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportMessage.IdMessage, reportMessage.IdOwner, reportMessage.IdReporter, reportMessage.Problem, reportMessage.Created, reportMessage.Updated, reportMessage.Status)
    if err != nil {
        t.Error(err)
    }

    reportMessage, err = reportMessageModel.ById(1)
    if err != nil {
        t.Error(err)
    }

    reportMessage.Problem = "updatedtestproblem"
    reportMessage.Status = 2

    err = reportMessageModel.Update(1, reportMessage.Status)
    if err != nil {
        t.Error(err)
    }

    reportMessage, err = reportMessageModel.ById(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, 2, reportMessage.Status)
}

func TestRMDelete(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE report_messages")
    if err != nil {
        t.Error(err)
    }

    reportMessageModel := models.ReportMessageModel{}

    reportMessage := models.ReportMessage{
        IdMessage: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_messages (id_message, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportMessage.IdMessage, reportMessage.IdOwner, reportMessage.IdReporter, reportMessage.Problem, reportMessage.Created, reportMessage.Updated, reportMessage.Status)
    if err != nil {
        t.Error(err)
    }

    reportMessage, err = reportMessageModel.ById(1)
    if err != nil {
        t.Error(err)
    }

    err = reportMessageModel.Delete(reportMessage.Id)
    if err != nil {
        t.Error(err)
    }

    reportMessages := reportMessageModel.ReportMessages()
    assert.Equal(t, 0, len(reportMessages))
}