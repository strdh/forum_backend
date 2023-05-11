package models

import (
    "log"
    "errors"
    "xyzforum/config"
)

type Forum struct {
    Id int `json:"id,omitempty"`
    IdUser int `json:"id,omitempty"`
    Title string `json:"title,omitempty"`
    Slug string `json:"slug,omitempty"`
    Description string `json:"description,omitempty"`
    ActiveUsers int `json:"active_users,omitempty"`
    Messages int `json:"messages,omitempty"`
    Status int `json:"status,omitempty"`
    Created int `json:"created,omitempty"`
}

type ForumModel struct {}

func (forumModel *ForumModel) GetForums() []Forum {
    var forums []Forum
    var temp Forum

    rows, err := config.DB.Query("SELECT * FROM forums")
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.IdUser, &temp.Title, &temp.Slug, &temp.Description, &temp.ActiveUsers, &temp.Messages, &temp.Status, &temp.Created)
        if err != nil {
            log.Println(err)
        }

        forums = append(forums, temp)
    }

    return forums
}

func (forumModel *ForumModel) CreateForum(forum Forum) (int64, error) {
    result, err := config.DB.Exec("INSERT INTO forums (id_user, title, description, slug) VALUES (?, ?, ?, ?)", forum.IdUser, forum.Title, forum.Slug, forum.Description)
    if err != nil {
        log.Println(err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (forumModel *ForumModel) GetForumById(id int) (Forum, error) {
    var forum Forum
    rows, err := config.DB.Query("SELECT * FROM forums WHERE id = ?", id)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&forum.Id, &forum.IdUser, &forum.Title, &forum.Slug, &forum.Description, &forum.ActiveUsers, &forum.Messages, &forum.Status, &forum.Created)
        if err != nil {
            log.Println(err)
        }

        return forum, nil
    }

    return forum, errors.New("Forum not found")
}

func (forumModel *ForumModel) UpdateForum(forum Forum, id int) (Forum, error) {
    _, err := config.DB.Exec("UPDATE forums SET title = ?, description = ? WHERE id = ?", forum.Title, forum.Description, id)
    if err != nil {
        log.Println(err)
        return forum, errors.New("Update has failed: " + err.Error())
    }

    return forum, nil
}

func (forumModel *ForumModel) DeleteForum(id int) error {
    _, err := config.DB.Exec("DELETE FROM forums WHERE id = ?", id)
    if err != nil {
        log.Println(err)
        return errors.New("Delete has failed: " + err.Error())
    }

    return nil
}