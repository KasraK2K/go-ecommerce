package mongodb

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* -------------------------------- Check ID -------------------------------- */
// Check string id to be a valid id and convert it to ObjectID
/* -------------------------------------------------------------------------- */
func (c *connection) checkId(id string) (primitive.ObjectID, bool) {
	// Check id is valid or not
	// isValidId := primitive.IsValidObjectID(id)
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, false
	}
	// Change id to ObjectId
	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.NilObjectID, false
	}
	return docID, true
}

/* ----------------------------- Sanitize Filter ---------------------------- */
// This filter have been created to get bson.D and change id to ObjectID
/* -------------------------------------------------------------------------- */
func (c *connection) sanitizeFilter(filter *bson.D) {
	reservedFilter := *filter
	for index, filterItem := range reservedFilter {
		if filterItem.Key == "id" || filterItem.Key == "_id" {
			docID, isValidId := c.checkId(filterItem.Value.(string))
			if !isValidId {
				message := fmt.Sprintf("hex string `%s` is not a valid ObjectID", filterItem.Value)
				log.Println(message)
				return
			}
			reservedFilter[index].Key = "_id"
			reservedFilter[index].Value = docID
		}
	}
}
