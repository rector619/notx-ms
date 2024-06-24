package models

import (
	"fmt"
	"time"

	"github.com/SineChat/notification-ms/pkg/repository/storage/mongodb"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (NotificationRecord) CollectionName() string {
	return "notification_record"
}

type NotificationRecord struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name         string             `bson:"name" json:"name"`
	Data         string             `bson:"data" json:"data"`
	Attempts     int                `bson:"attempts" json:"attempts"`
	Sent         bool               `bson:"sent" json:"sent"`
	Abandoned    bool               `bson:"abandoned" json:"abandoned"`
	AttemptAgain int                `bson:"attempt_again" json:"attempt_again"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	DeletedAt    time.Time          `bson:"deleted_at" json:"-"`
	Deleted      bool               `bson:"deleted" json:"-"`
}

func (n *NotificationRecord) GetSomeUnsentNotifications(db *mongodb.Database, limit int) ([]NotificationRecord, error) {
	details := []NotificationRecord{}
	filter := bson.M{"sent": false, "abandoned": false, "attempt_again": bson.M{"$gt": 0, "$lte": int(time.Now().Unix())}}
	err := db.SelectAllFromDbWithLimit("-_id", limit, &n, filter, &details)
	if err != nil {
		return details, err
	}
	return details, nil
}

func (n *NotificationRecord) CreateNotificationRecord(db *mongodb.Database) error {
	err := db.CreateOneRecord(&n)
	if err != nil {
		return fmt.Errorf("notification record creation failed: %v", err.Error())
	}
	return nil
}

func (n *NotificationRecord) UpdateAllFields(db *mongodb.Database) error {
	return db.SaveAllFields(&n)
}
