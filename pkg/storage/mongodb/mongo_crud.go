package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"app/pkg"
)

type insertOneResult struct {
	ID interface{} `json:"id" bson:"id"`
}

type insertManyResult struct {
	IDS []interface{} `json:"ids" bson:"ids"`
}

type updateResult struct {
	MatchedCount  int64       `json:"matched_count" bson:"matched_count"`
	ModifiedCount int64       `json:"modified_count" bson:"modified_count"`
	UpsertedCount int64       `json:"upserted_count" bson:"upserted_count"`
	UpsertedID    interface{} `json:"upserted_id" bson:"upserted_id"`
}

type deleteResult struct {
	DeletedCount int64 `json:"deleted_count" bson:"deleted_count"`
}

func (c *connection) Ping(client *mongo.Client) {
	// Send a ping to confirm a successful connection
	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		log.Panic(err)
	}
	pkg.Logger.Debug("Pinged your deployment. You successfully connected to MongoDB!")
}

/* ------------------------------- Insert One ------------------------------- */
// document := struct {
// 	Title   string
// 	Message string
// 	Data    interface{}
// }{
// 	Title:   "Test Log",
// 	Message: "Test our log system by kasra's code",
// 	Data: struct {
// 		Name string `json:"name" bson:"name"`
// 	}{"Kasra"},
// }
// result := mongodb.Conn.InsertOne("app", "log", document)
// res, _ := pkg.Marshal(result)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func (c *connection) InsertOne(database, collection string, document interface{}) insertOneResult {
	client := c.Connect()
	defer c.Disconnect(client)

	coll := client.Database(database).Collection(collection)

	result, err := coll.InsertOne(context.TODO(), document)
	if err != nil {
		log.Panic(err)
	}

	return insertOneResult{ID: result.InsertedID}
}

/* ------------------------------- Insert Many ------------------------------ */
// documents := []interface{}{
// 	struct {
// 		Title   string
// 		Message string
// 	}{Title: "Insert Many Title", Message: "My Insert Many Message"},
// 	struct {
// 		Name string
// 		Age  int
// 	}{Name: "Kasra", Age: 37},
// }
// result := mongodb.Conn.InsertMany("app", "log", documents)
// res, _ := pkg.Marshal(result)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func (c *connection) InsertMany(database, collection string, documents []interface{}, opts ...*options.InsertManyOptions) insertManyResult {
	client := c.Connect()
	defer c.Disconnect(client)

	coll := client.Database(database).Collection(collection)

	results, err := coll.InsertMany(context.TODO(), documents, opts...)
	if err != nil {
		log.Panic(err)
	}

	return insertManyResult{IDS: results.InsertedIDs}
}

/* -------------------------------- Find One -------------------------------- */
// filter := bson.D{{Key: "id", Value: "641196ac5986aae6482be366"}}
// opts := options.FindOne()
// result := mongodb.Conn.FindOne("app", "log", filter, opts)
// res, _ := pkg.Marshal(result)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func (c *connection) FindOne(database, collection string, filter bson.D, opts ...*options.FindOneOptions) bson.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := c.Connect()
	defer c.Disconnect(client)

	c.sanitizeFilter(&filter)

	coll := client.Database(database).Collection(collection)

	var result bson.M
	err := coll.FindOne(ctx, filter, opts...).Decode(&result)
	if err != nil {
		log.Panic(err)
	}

	return result
}

/* ---------------------------------- Find ---------------------------------- */
// filter := bson.D{{Key: "priority", Value: 0}}
// opts := options.Find()
// opts.SetSort(bson.D{{Key: "created_at", Value: -1}})
// results := mongodb.Conn.Find("app", "log", filter, opts)
// res, _ := pkg.Marshal(results)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func (c *connection) Find(database, collection string, filter bson.D, opts ...*options.FindOptions) []bson.M {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := c.Connect()
	defer c.Disconnect(client)

	c.sanitizeFilter(&filter)

	coll := client.Database(database).Collection(collection)

	cursor, err := coll.Find(context.TODO(), filter, opts...)
	if err != nil {
		log.Panic(err)
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Panic(err)
	}

	return results
}

/* ------------------------------ Update By ID ------------------------------ */
// result := mongodb.Conn.UpdateByID("app", "log", "641196ac5986aae6482be366",
// 	bson.D{
// 		{Key: "$set", Value: bson.D{
// 			{Key: "title", Value: "new title"},
// 		}},
// 	})
// res, _ := pkg.Marshal(result)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func (c *connection) UpdateByID(database, collection, id string, update interface{}, opts ...*options.UpdateOptions) updateResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := c.Connect()
	defer c.Disconnect(client)

	coll := client.Database(database).Collection(collection)

	docID, isValidId := c.checkId(id)
	if !isValidId {
		message := fmt.Sprintf("hex string `%s` is not a valid ObjectID", id)
		log.Println(message)
		return updateResult{
			MatchedCount:  0,
			ModifiedCount: 0,
			UpsertedCount: 0,
			UpsertedID:    nil,
		}
	}

	result, err := coll.UpdateByID(ctx, docID, update, opts...)
	if err != nil {
		log.Panic(err)
	}

	return updateResult{
		MatchedCount:  result.MatchedCount,
		ModifiedCount: result.ModifiedCount,
		UpsertedCount: result.UpsertedCount,
		UpsertedID:    result.UpsertedID,
	}
}

/* ------------------------------- Update One ------------------------------- */
// filter := bson.D{{Key: "id", Value: "641196ac5986aae6482be366"}}
// result := mongodb.Conn.UpdateOne("app", "log", filter,
// 	bson.D{
// 		{Key: "$set", Value: bson.D{
// 			{Key: "title", Value: "new title 2"},
// 		}},
// 	},
// )
// res, _ := pkg.Marshal(result)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func (c *connection) UpdateOne(database, collection string, filter bson.D, update interface{}, opts ...*options.UpdateOptions) updateResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := c.Connect()
	defer c.Disconnect(client)

	c.sanitizeFilter(&filter)

	coll := client.Database(database).Collection(collection)

	result, err := coll.UpdateOne(ctx, filter, update, opts...)
	if err != nil {
		log.Panic(err)
	}

	return updateResult{
		MatchedCount:  result.MatchedCount,
		ModifiedCount: result.ModifiedCount,
		UpsertedCount: result.UpsertedCount,
		UpsertedID:    result.UpsertedID,
	}
}

/* -------------------------------- DeleteOne ------------------------------- */
// filter := bson.D{{Key: "id", Value: "6411d0a4db46166dde06ba4e"}}
// result := mongodb.Conn.DeleteOne("app", "log", filter)
// res, _ := pkg.Marshal(result)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func (c *connection) DeleteOne(database, collection string, filter bson.D, opts ...*options.DeleteOptions) deleteResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := c.Connect()
	defer c.Disconnect(client)

	c.sanitizeFilter(&filter)

	coll := client.Database(database).Collection(collection)

	result, err := coll.DeleteOne(ctx, filter, opts...)
	if err != nil {
		log.Panic(err)
	}

	return deleteResult{DeletedCount: result.DeletedCount}
}

/* ------------------------------- Replace One ------------------------------ */
// filter := bson.D{{Key: "id", Value: "6411d0b7349067432e1d1eb1"}}
// replacement := bson.D{{Key: "title", Value: "new title 2"}}
// result := mongodb.Conn.ReplaceOne("app", "log", filter, replacement)
// res, _ := pkg.Marshal(result)
// fmt.Println(string(res))
/* -------------------------------------------------------------------------- */
func (c *connection) ReplaceOne(database, collection string, filter bson.D, replacement interface{}, opts ...*options.ReplaceOptions) updateResult {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client := c.Connect()
	defer c.Disconnect(client)

	c.sanitizeFilter(&filter)

	coll := client.Database(database).Collection(collection)

	result, err := coll.ReplaceOne(ctx, filter, replacement, opts...)
	if err != nil {
		log.Panic(err)
	}

	return updateResult{
		MatchedCount:  result.MatchedCount,
		ModifiedCount: result.ModifiedCount,
		UpsertedCount: result.UpsertedCount,
		UpsertedID:    result.UpsertedID,
	}
}
