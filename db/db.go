package db

const (
	DBNAME     = "Funduqi-db"
	TestDBNAME = "Funduqi-db-test"
	DBURI      = "mongodb://localhost:27017"
)

type Store struct {
	User  UserStore
	Hotel HotelStore
	Room  RoomStore
}
