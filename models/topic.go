package models

import (
    "log"
    "errors"
    "xyzforum/config"
)

type Topic struct {
    Id int `json:"id,omitempty"`
    Topic string `json:"topic,omitempty"`
    Created int64 `json:"created,omitempty"`
}

type TopicModel struct {}

func (topicModel *TopicModel) Topics() []Topic {
    var topics []Topic
    var temp Topic

    rows, err := config.DB.Query("SELECT * FROM topics")
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&temp.Id, &temp.Topic, &temp.Created)
        if err != nil {
            log.Println(err)
        }

        topics = append(topics, temp)
    }

    return topics
}

func (topicModel *TopicModel) Create(topic Topic) (int64, error) {
    result, err := config.DB.Exec("INSERT INTO topics (topic, created) VALUES (?, ?)", topic.Topic, topic.Created)
    if err != nil {
        log.Println(err)
    }

    id, err := result.LastInsertId()
    if err != nil {
        return 0, err
    }

    return id, nil
}

func (topicModel *TopicModel) ById(id int) (Topic, error) {
    var topic Topic
    rows, err := config.DB.Query("SELECT * FROM topics WHERE id = ?", id)
    if err != nil {
        log.Println(err)
    }
    defer rows.Close()

    if rows.Next() {
        err := rows.Scan(&topic.Id, &topic.Topic, &topic.Created)
        if err != nil {
            log.Println(err)
        }

        return topic, nil
    }

    return topic, errors.New("Topic not found")
}

func (topicModel *TopicModel) Update(topic Topic, id int) (int64, error) {
    result, err := config.DB.Exec("UPDATE topics SET topic = ? WHERE id = ?", topic.Topic, id)
    if err != nil {
        log.Println(err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }

    return rowsAffected, nil
}

func (topicModel *TopicModel) Delete(id int) (int64, error) {
    result, err := config.DB.Exec("DELETE FROM topics WHERE id = ?", id)
    if err != nil {
        log.Println(err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return 0, err
    }

    return rowsAffected, nil
}