package test

import (
    "testing"
    "time"
    "xyzforum/config"
    "xyzforum/models"
    "github.com/joho/godotenv"
    "github.com/stretchr/testify/assert"
)

func init() {
    err := godotenv.Load("../.env")
    if err != nil {
        panic(err)
    }
    config.InitializeTestDB()
}

//RF stand for ReportForum

func TestRFReportForums(t *testing.T) {
    reportForumModel := models.ReportForumModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE report_forums")
    if err != nil {
        t.Error(err)
    }

    reportForum := models.ReportForum{
        IdForum: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_forums (id_forum, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportForum.IdForum, reportForum.IdOwner, reportForum.IdReporter, reportForum.Problem, reportForum.Created, reportForum.Updated, reportForum.Status)
    if err != nil {
        t.Error(err)
    }

    reportForums := reportForumModel.ReportForums()
    assert.Equal(t, 1, len(reportForums))
    assert.Equal(t, "testproblem", reportForums[0].Problem)
    assert.Equal(t, 1, reportForums[0].Status)
}

func TestRFCreate(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE report_forums")
    if err != nil {
        t.Error(err)
    }

    reportForumModel := models.ReportForumModel{}

    reportForum := models.ReportForum{
        IdForum: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = reportForumModel.Create(reportForum)
    if err != nil {
        t.Error(err)
    }

    reportForums := reportForumModel.ReportForums()
    assert.Equal(t, 1, len(reportForums))
    assert.Equal(t, "testproblem", reportForums[0].Problem)
    assert.Equal(t, 1, reportForums[0].Status)
}

func TestRFById(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE report_forums")
    if err != nil {
        t.Error(err)
    }

    reportForumModel := models.ReportForumModel{}

    reportForum := models.ReportForum{
        IdForum: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_forums (id_forum, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportForum.IdForum, reportForum.IdOwner, reportForum.IdReporter, reportForum.Problem, reportForum.Created, reportForum.Updated, reportForum.Status)
    if err != nil {
        t.Error(err)
    }

    reportForum, err = reportForumModel.ById(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, "testproblem", reportForum.Problem)
    assert.Equal(t, 1, reportForum.Status)
}

func TestRFByIdForum(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE report_forums")
    if err != nil {
        t.Error(err)
    }

    reportForumModel := models.ReportForumModel{}

    reportForum := models.ReportForum{
        IdForum: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_forums (id_forum, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportForum.IdForum, reportForum.IdOwner, reportForum.IdReporter, reportForum.Problem, reportForum.Created, reportForum.Updated, reportForum.Status)
    if err != nil {
        t.Error(err)
    }

    reportForums := reportForumModel.ByIdForum(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, 1, len(reportForums))
    assert.Equal(t, "testproblem", reportForums[0].Problem)
    assert.Equal(t, 1, reportForums[0].Status)
}

func TestRFByIdOwner(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE report_forums")
    if err != nil {
        t.Error(err)
    }

    reportForumModel := models.ReportForumModel{}

    reportForum := models.ReportForum{
        IdForum: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_forums (id_forum, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportForum.IdForum, reportForum.IdOwner, reportForum.IdReporter, reportForum.Problem, reportForum.Created, reportForum.Updated, reportForum.Status)
    if err != nil {
        t.Error(err)
    }

    reportForums := reportForumModel.ByIdOwner(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, 1, len(reportForums))
    assert.Equal(t, "testproblem", reportForums[0].Problem)
    assert.Equal(t, 1, reportForums[0].Status)
}

func TestRFByIdReporter(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE report_forums")
    if err != nil {
        t.Error(err)
    }

    reportForumModel := models.ReportForumModel{}

    reportForum := models.ReportForum{
        IdForum: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_forums (id_forum, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportForum.IdForum, reportForum.IdOwner, reportForum.IdReporter, reportForum.Problem, reportForum.Created, reportForum.Updated, reportForum.Status)
    if err != nil {
        t.Error(err)
    }

    reportForums := reportForumModel.ByIdReporter(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, 1, len(reportForums))
    assert.Equal(t, "testproblem", reportForums[0].Problem)
    assert.Equal(t, 1, reportForums[0].Status)
}

func TestRFUpdate(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE report_forums")
    if err != nil {
        t.Error(err)
    }

    reportForumModel := models.ReportForumModel{}

    reportForum := models.ReportForum{
        IdForum: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_forums (id_forum, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportForum.IdForum, reportForum.IdOwner, reportForum.IdReporter, reportForum.Problem, reportForum.Created, reportForum.Updated, reportForum.Status)
    if err != nil {
        t.Error(err)
    }

    dataUpdate := time.Now().Unix()

    _, err = reportForumModel.Update(1, dataUpdate)
    if err != nil {
        t.Error(err)
    }

    reportForum, err = reportForumModel.ById(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, dataUpdate, reportForum.Updated)
}

func TestRFDelete(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE report_forums")
    if err != nil {
        t.Error(err)
    }

    reportForumModel := models.ReportForumModel{}

    reportForum := models.ReportForum{
        IdForum: 1,
        IdOwner: 1,
        IdReporter: 1,
        Problem: "testproblem",
        Created: time.Now().Unix(),
        Updated: time.Now().Unix(),
        Status: 1,
    }

    _, err = config.DB.Exec("INSERT INTO report_forums (id_forum, id_owner, id_reporter, problem, created, updated, status) VALUES (?, ?, ?, ?, ?, ?, ?)", reportForum.IdForum, reportForum.IdOwner, reportForum.IdReporter, reportForum.Problem, reportForum.Created, reportForum.Updated, reportForum.Status)
    if err != nil {
        t.Error(err)
    }

    result, err := reportForumModel.Delete(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, int64(1), result)
}