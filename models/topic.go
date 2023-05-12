package models

import (
    "log"
    "errors"
    "xyzforum/config"
)

type Topic struct {
    Id int `json:"id,omitempty"`
    Topic string `json:"topic,omitempty"`
    Created int `json:"created,omitempty"`
}

type TopicModel struct {}

func (topicModel *TopicModel) GetTopics() []Topic {
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

func (topicModel *TopicModel) CreateTopic(topic Topic) (int64, error) {
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

func (topicModel *TopicModel) GetTopicById(id int) (Topic, error) {
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

func (topicModel *TopicModel) UpdateTopic(topic Topic, id int) (Topic, error) {
    _, err := config.DB.Exec("UPDATE topics SET topic = ? WHERE id = ?", topic.Topic, id)
    if err != nil {
        log.Println(err)
        return topic, errors.New("Update has failed: " + err.Error())
    }

    return topic, nil
}

func (topicModel *TopicModel) DeleteTopic(id int) error {
    _, err := config.DB.Exec("DELETE FROM topics WHERE id = ?", id)
    if err != nil {
        log.Println(err)
        return errors.New("Delete has failed: "+ err.Error())
    }

    return nil
}