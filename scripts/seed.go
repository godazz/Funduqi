package main

import (
	"context"
	"fmt"
	"log"

	"github.com/godazz/Funduqi/db"
	"github.com/godazz/Funduqi/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	roomStore  db.RoomStore
	hotelStore db.HotelStore
	ctx        = context.Background()
)

func seedHotel(name string, location string, rating int) {
	hotel := types.Hotel{
		Name:     name,
		Location: location,
		Rating:   rating,
		Rooms:    []primitive.ObjectID{},
	}
	rooms := []types.Room{
		{
			Size:    "small",
			Seaside: false,
			Price:   69.69,
		},
		{
			Size:    "normal",
			Seaside: true,
			Price:   169.69,
		},
		{
			Size:    "kingsize",
			Seaside: true,
			Price:   269.69,
		},
	}
	insertedHotel, err := hotelStore.InsertHotel(ctx, &hotel)
	if err != nil {
		log.Fatal(err)
	}

	for _, room := range rooms {
		room.HotelID = insertedHotel.ID
		insertedRoom, err := roomStore.InsertRoom(ctx, &room)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(insertedRoom)
	}
}

func main() {
	fmt.Println("--- Seeding the database ---")
	seedHotel("Hilton", "Egypt", 4)
	seedHotel("Ramses", "Iraq", 3)
	seedHotel("Bazel", "Syria", 1)
}

func init() {
	var err error
	client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}

	hotelStore = db.NewMongoHotelStore(client)
	roomStore = db.NewMongoRoomStore(client, hotelStore)
}
