package main

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// establish connection to db
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGODB_ATLAS_URI")))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	if err != nil {
		log.Print(err)
		return
	}

	/////////////////////////////////////////////////////////////////////////////
	// example one -- how to update V of K:setting-2
	arr_one := bson.A{
		bson.D{
			{"K", "setting-1"},
			{"V", 1234}},
		bson.D{
			{"K", "setting-2"},
			{"V", "hello"}}, // update this field
		bson.D{
			{"K", "setting-3"},
			{"V", true}}}
	example_one := bson.D{
		{"deviceID", "example_one"},
		{"preferences", arr_one},
	}
	log.Printf("Example One Is %+v", example_one)
	// insert the sample doc
	resultInsert, err := client.Database("tutorial").Collection("foobar").InsertOne(ctx, example_one)
	if err != nil {
		log.Fatal(err)
		return
	}
	// update the sample doc
	filter_one := bson.D{{"_id", resultInsert.InsertedID.(primitive.ObjectID)}, {"preferences.K", "setting-2"}}
	update_one := bson.D{{"$set", bson.D{{"preferences.$", "\U0001f60d\U0001f60d\U0001f60d"}}}}
	_, err = client.Database("tutorial").Collection("foobar").UpdateOne(ctx, filter_one, update_one)
	if err != nil {
		log.Fatal(err)
	}
	// find modified doc or show in atlas..
	var result map[string]interface{}
	err = client.Database("tutorial").Collection("foobar").FindOne(ctx, bson.D{{"_id", resultInsert.InsertedID.(primitive.ObjectID)}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Example One Updated %+v", result)

	/////////////////////////////////////////////////////////////////////////////
	// example two -- update V of setting-2b
	arr_two := bson.A{
		bson.D{
			{"K", "setting-1"},
			{"V", 270172701727017}},
		bson.D{
			{"K", "setting-2"},
			{"V", bson.D{
				{"setting-2a", "GMT-7"},
				{"setting-2b", 1234}}}}, // update this field
	}
	example_two := bson.D{
		{"deviceID", "example_two"},
		{"preferences", arr_two},
	}
	log.Printf("Example Two Is %+v", example_two)
	// insert the sample doc
	resultInsert, err = client.Database("tutorial").Collection("foobar").InsertOne(ctx, example_two)
	if err != nil {
		log.Fatal(err)
		return
	}
	// update the sample doc
	filter_two := bson.D{{"_id", resultInsert.InsertedID.(primitive.ObjectID)}, {"preferences.K", "setting-2"}}
	update_two := bson.D{{"$set", bson.D{{"preferences.$.setting-2b", "\U0001f929\U0001f929\U0001f929"}}}} // nested update
	_, err = client.Database("tutorial").Collection("foobar").UpdateOne(ctx, filter_two, update_two)
	if err != nil {
		log.Fatal(err)
	}
	// find modified doc or show in atlas..
	err = client.Database("tutorial").Collection("foobar").FindOne(ctx, bson.D{{"_id", resultInsert.InsertedID.(primitive.ObjectID)}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Example Two Updated %+v", result)

	/////////////////////////////////////////////////////////////////////////////
	// example three -- update V of setting-3 / setting-3b / k1
	arr_three := bson.A{
		bson.D{
			{"K", "setting-1"},
			{"V", 270172701727017}},
		bson.D{
			{"K", "setting-2"},
			{"V", bson.D{
				{"setting-2a", "helloworld"},
				{"setting-2b", 270172701727017}}}},
		bson.D{
			{"K", "setting-3"},
			{"V", bson.D{
				{"setting-3a", "foobarfoobar"},
				{"setting-3b", bson.A{
					bson.D{
						{"K", "k1"}, // update this value
						{"V", -2701727017},
					},
					bson.D{
						{"K", "k2"},
						{"V", true},
					},
					bson.D{
						{"K", "k3"},
						{"V", false},
					},
				}}}}},
	}
	example_three := bson.D{
		{"deviceID", "example_three"},
		{"preferences", arr_three},
	}
	log.Printf("Example three Is %+v", example_three)
	// insert the sample doc
	resultInsert, err = client.Database("tutorial").Collection("foobar").InsertOne(ctx, example_three)
	if err != nil {
		log.Fatal(err)
		return
	}
	// update the sample doc
	filter_three := bson.D{{"_id", resultInsert.InsertedID.(primitive.ObjectID)}}
	update_three := bson.D{{"$set", bson.D{{"preferences.$[myFirstFilter].V.setting-3b.$[mySecondFilter].K", "\U0001f525\U0001f525\U0001f525"}}}}
	arrayFilters := []interface{}{
		map[string]interface{}{"myFirstFilter.K": "setting-3"},
		map[string]interface{}{"mySecondFilter.K": "k1"},
	}
	options_three := options.Update().SetArrayFilters(options.ArrayFilters{Filters: arrayFilters})
	log.Printf("options_three  %+v", options_three.ArrayFilters)
	_, err = client.Database("tutorial").Collection("foobar").UpdateOne(ctx, filter_three, update_three, options_three)
	if err != nil {
		log.Fatal(err)
	}
	// find modified doc or show in atlas..
	err = client.Database("tutorial").Collection("foobar").FindOne(ctx, bson.D{{"_id", resultInsert.InsertedID.(primitive.ObjectID)}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Printf("Example three Updated %+v", result)
}
