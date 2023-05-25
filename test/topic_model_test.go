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

func TestTopicTopics(t *testing.T) {
    topicModel := models.TopicModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE topics")
    if err != nil {
        t.Error(err)
    }

    topic := models.Topic{
        Topic: "testtopic",
        Created: time.Now().Unix(),
    }

    _, err = config.DB.Exec("INSERT INTO topics (topic, created) VALUES (?, ?)", topic.Topic, topic.Created)
    if err != nil {
        t.Error(err)
    }

    topics := topicModel.Topics()
    assert.Equal(t, 1, len(topics))
    assert.Equal(t, "testtopic", topics[0].Topic)
}

func TestTopicCreate(t *testing.T) {
    _, err := config.DB.Exec("TRUNCATE TABLE topics")
    if err != nil {
        t.Error(err)
    }

    topicModel := models.TopicModel{}

    topic := models.Topic{
        Topic: "testtopic",
        Created: time.Now().Unix(),
    }

    _, err = topicModel.Create(topic)
    if err != nil {
        t.Error(err)
    }

    topics := topicModel.Topics()
    assert.Equal(t, 1, len(topics))
    assert.Equal(t, "testtopic", topics[0].Topic)
}

func TestTopicById(t *testing.T) {
    topicModel := models.TopicModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE topics")
    if err != nil {
        t.Error(err)
    }

    topic := models.Topic{
        Topic: "testtopic",
        Created: time.Now().Unix(),
    }

    _, err = config.DB.Exec("INSERT INTO topics (topic, created) VALUES (?, ?)", topic.Topic, topic.Created)
    if err != nil {
        t.Error(err)
    }

    topic, err = topicModel.ById(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, "testtopic", topic.Topic)
}

func TestTopicUpdate(t *testing.T) {
    topicModel := models.TopicModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE topics")
    if err != nil {
        t.Error(err)
    }

    topic := models.Topic{
        Topic: "testtopic",
        Created: time.Now().Unix(),
    }

    _, err = config.DB.Exec("INSERT INTO topics (topic, created) VALUES (?, ?)", topic.Topic, topic.Created)
    if err != nil {
        t.Error(err)
    }

    topic, err = topicModel.ById(1)
    if err != nil {
        t.Error(err)
    }

    topic.Topic = "testtopicupdate"
    result, err := topicModel.Update(topic, 1)
    if err != nil {
        t.Error(err)
    }

    topic, err = topicModel.ById(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, "testtopicupdate", topic.Topic)
    assert.Equal(t, int64(1), result)
}

func TestTopicDelete(t *testing.T) {
    topicModel := models.TopicModel{}

    _, err := config.DB.Exec("TRUNCATE TABLE topics")
    if err != nil {
        t.Error(err)
    }

    topic := models.Topic{
        Topic: "testtopic",
        Created: time.Now().Unix(),
    }

    _, err = config.DB.Exec("INSERT INTO topics (topic, created) VALUES (?, ?)", topic.Topic, topic.Created)
    if err != nil {
        t.Error(err)
    }

    result, err := topicModel.Delete(1)
    if err != nil {
        t.Error(err)
    }

    assert.Equal(t, int64(1), result)
}