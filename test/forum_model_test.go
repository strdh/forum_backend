package test

import (
    "testing"
    "time"
    "xyzforum/config"
    "xyzforum/models"
    "github.com/google/uuid"
    "github.com/stretchr/testify/assert"
)

func init() {
    config.InitializeTestDB()
}

func TestForumForums(t *testing.T) {
    forumModel := models.ForumModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE forums")
    if err != nil {
        t.Error(err)
    }

    slug := uuid.New()
    binarySlug, _ := slug.MarshalBinary()

    forum := models.Forum{
        IdUser: 1,
        Title: "testtitle",
        Slug: binarySlug,
        Description: "testdescription",
        ActiveUsers: 1,
        Messages: 1,
        Status: 1,
        Created: time.Now().Unix(),
    }

    _, err = config.DB.Exec("INSERT INTO forums (id_user, title, slug, description, active_users, messages, status, created) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", forum.IdUser, forum.Title, forum.Slug, forum.Description, forum.ActiveUsers, forum.Messages, forum.Status, forum.Created)
    if err != nil {
        t.Error(err)
    }

    forums := forumModel.Forums()
    assert.Equal(t, 1, len(forums))
    assert.Equal(t, "testtitle", forums[0].Title)
    assert.Equal(t, "testdescription", forums[0].Description)
    assert.Equal(t, 1, forums[0].Status)
    assert.Equal(t, 1, forums[0].ActiveUsers)
}

func TestForumCreate(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE forums")
    if err != nil {
        t.Error(err)
    }

    forumModel := models.ForumModel{}

    slug := uuid.New()
    binarySlug, _ := slug.MarshalBinary()

    forum := models.Forum{
        IdUser: 1,
        Title: "testtitle",
        Slug: binarySlug,
        Description: "testdescription",
        ActiveUsers: 1,
        Messages: 1,
        Status: 1,
        Created: time.Now().Unix(),
    }

    id, err := forumModel.Create(forum)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, int64(1), id)
}

func TestForumById1(t *testing.T) {
    forumModel := models.ForumModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE forums")
    if err != nil {
        t.Error(err)
    }

    slug := uuid.New()
    binarySlug, _ := slug.MarshalBinary()

    forum := models.Forum{
        IdUser: 1,
        Title: "testtitle",
        Slug: binarySlug,
        Description: "testdescription",
        ActiveUsers: 1,
        Messages: 1,
        Status: 1,
        Created: time.Now().Unix(),
    }

    _, err = config.DB.Exec("INSERT INTO forums (id_user, title, slug, description, active_users, messages, status, created) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", forum.IdUser, forum.Title, forum.Slug, forum.Description, forum.ActiveUsers, forum.Messages, forum.Status, forum.Created)
    if err != nil {
        t.Error(err)
    }

    forum, err = forumModel.ById(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, "testtitle", forum.Title)
    assert.Equal(t, "testdescription", forum.Description)
    assert.Equal(t, 1, forum.Status)
    assert.Equal(t, 1, forum.ActiveUsers)
}

func TestForumById2(t *testing.T) {
    forumModel := models.ForumModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE forums")
    if err != nil {
        t.Error(err)
    }

    slug := uuid.New()
    binarySlug, _ := slug.MarshalBinary()

    forum := models.Forum{
        IdUser: 1,
        Title: "testtitle",
        Slug: binarySlug,
        Description: "testdescription",
        ActiveUsers: 1,
        Messages: 1,
        Status: 1,
        Created: time.Now().Unix(),
    }

    _, err = config.DB.Exec("INSERT INTO forums (id_user, title, slug, description, active_users, messages, status, created) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", forum.IdUser, forum.Title, forum.Slug, forum.Description, forum.ActiveUsers, forum.Messages, forum.Status, forum.Created)
    if err != nil {
        t.Error(err)
    }

    forum, err = forumModel.ById(2)
    assert.Equal(t, models.Forum{}, forum)
    assert.Equal(t, "Forum not found", err.Error())
}

func TestUpdateForum(t *testing.T) {
    forumModel := models.ForumModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE forums")
    if err != nil {
        t.Error(err)
    }

    slug := uuid.New()
    binarySlug, _ := slug.MarshalBinary()

    forum := models.Forum{
        IdUser: 1,
        Title: "testtitle",
        Slug: binarySlug,
        Description: "testdescription",
        ActiveUsers: 1,
        Messages: 1,
        Status: 1,
        Created: time.Now().Unix(),
    }

    _, err = config.DB.Exec("INSERT INTO forums (id_user, title, slug, description, active_users, messages, status, created) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", forum.IdUser, forum.Title, forum.Slug, forum.Description, forum.ActiveUsers, forum.Messages, forum.Status, forum.Created)
    if err != nil {
        t.Error(err)
    }

    forum.Title = "testtitleupdate"
    forum.Description = "testdescriptionupdate"
    

    result, err := forumModel.Update(forum, 1)
    if err != nil {
        t.Error(err)
    }

    forum, err = forumModel.ById(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, int64(1), result)
    assert.Equal(t, "testtitleupdate", forum.Title)
    assert.Equal(t, "testdescriptionupdate", forum.Description)
}

func TestDeleteForum(t *testing.T) {
    forumModel := models.ForumModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE forums")
    if err != nil {
        t.Error(err)
    }

    slug := uuid.New()
    binarySlug, _ := slug.MarshalBinary()

    forum := models.Forum{
        IdUser: 1,
        Title: "testtitle",
        Slug: binarySlug,
        Description: "testdescription",
        ActiveUsers: 1,
        Messages: 1,
        Status: 1,
        Created: time.Now().Unix(),
    }

    _, err = config.DB.Exec("INSERT INTO forums (id_user, title, slug, description, active_users, messages, status, created) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", forum.IdUser, forum.Title, forum.Slug, forum.Description, forum.ActiveUsers, forum.Messages, forum.Status, forum.Created)
    if err != nil {
        t.Error(err)
    }

    result, err := forumModel.Delete(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, int64(1), result)
}

