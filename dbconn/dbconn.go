package dbconn

import (
	"context"
	"flag"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoStr = flag.String("mongostr", "", "connection string for mongodb")
)

// Client represents a connection to the MongoDB server.
type Client struct {
	c   *mongo.Client
	ctx context.Context
}

type Submission struct {
}

func (c *Client) GetSubmission(id uint64) (string, error) {
	filter := bson.M{
		"_id": id,
	}
	cur, err := c.c.Database("Code4Trees").Collection("Submissions").Find(c.ctx, filter)
	if err != nil {
		return "", err
	}
	cur.Next(c.ctx)
	raw := cur.Current

	res, err := raw.LookupErr("code")

	if err != nil {
		return "", err
	}

	str, ok := res.StringValueOK()

	if !ok {
		return "", fmt.Errorf("unable to find string key")
	}

	return str, nil

}

func MakeClient(ctx context.Context) (*Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(*mongoStr))
	if err != nil {
		return nil, err
	}
	newCtx, _ := context.WithTimeout(ctx, 10*time.Second)
	err = client.Connect(newCtx)

	return &Client{c: client, ctx: ctx}, err
}
