package main

import (
	"context"
	"fmt"
	"log"

	"github.com/godazz/Funduqi/db"
	"github.com/godazz/Funduqi/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	fmt.Println("--- Seeding the database ---")

	ctx := context.Background()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}

	hotelStore := db.NewMongoHotelStore(client, db.DBNAME)
	roomStore := db.NewRoomStore(client, db.DBNAME)

	hotel := types.Hotel{
		Name:     "Ramses Hilton",
		Location: "Cairo",
	}
	rooms := []types.Room{
		{
			Type:      types.SingleRoomType,
			BasePrice: 69.69,
		},
		{
			Type:      types.DoubleRoomType,
			BasePrice: 169.69,
		},
		{
			Type:      types.SeaSideRoomType,
			BasePrice: 269.69,
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
