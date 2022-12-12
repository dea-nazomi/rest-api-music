package models

import (
	"database/sql"
	"fmt"
	"go-postgres-crud/config"
	"log"


	"github.com/satori/go.uuid"
	_ "github.com/lib/pq" // postgres golang driver
)

// Music schema dari tabel Music
// kita coba dengan jika datanya null
// jika return datanya ada yg null, silahkan pake NullString, contohnya dibawah
// Penulis       config.NullString `json:"penulis"`
type Music struct {
	ID       		string `json:"id"`
	Name    		string `json:"name"`
	Album       	string `json:"album"`
	Album_art 		string `json:"album_art"`
	Singer 			string `json:"singer"`
	Publish_date    string `json:"publish_date"`
	Created_at 		string `json:"created_at"`
	Updated_at 		string `json:"updated_at"`

}

func AddMusic(music Music) string {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	//generate uuid
	myuuid:= uuid.NewV4()

	// kita buat insert query
	// mengembalikan nilai id akan mengembalikan id dari music yang dimasukkan ke db
	sqlStatement := `INSERT INTO music_list (music_id, music_name, music_album, music_album_art, music_singer, music_publish_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING music_id`

	// id yang dimasukkan akan disimpan di id ini
	var id string

	// Scan function akan menyimpan insert id didalam id id
	err := db.QueryRow(sqlStatement, myuuid, music.Name, music.Album, music.Album_art, music.Singer, music.Publish_date).Scan(&id)
	if err != nil {
		log.Fatalf("Tidak Bisa mengeksekusi query. %v", err)
	}

	fmt.Printf("Insert data single record %v", id)

	// return insert id
	return id
}


func GetAllMusic() ([]Music, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var musics []Music

	// kita buat select query
	sqlStatement := `SELECT * FROM music_list`

	// mengeksekusi sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// kita tutup eksekusi proses sql qeurynya
	defer rows.Close()

	// kita iterasi mengambil datanya
	for rows.Next() {
		var music Music

		// kita ambil datanya dan unmarshal ke structnya
		err = rows.Scan(&music.ID, &music.Name, &music.Album, &music.Album_art, &music.Singer, &music.Publish_date, &music.Created_at, &music.Updated_at)

		if err != nil {
			log.Fatalf("tidak bisa mengambil dataa. %v", err)
		}

		// masukkan kedalam slice musics
		musics = append(musics, music)

	}

	// return empty music atau jika error
	return musics, err
}

// mengambil satu music
func GetOneMusic(id string) (Music, error) {
	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	var music Music

	// buat sql query
	sqlStatement := `SELECT * FROM music_list WHERE music_id=$1`

	// eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&music.ID, &music.Name, &music.Album, &music.Album_art, &music.Singer, &music.Publish_date, &music.Created_at, &music.Updated_at)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("Tidak ada data yang dicari!")
		return music, nil
	case nil:
		return music, nil
	default:
		log.Fatalf("tidak bisa mengambil data. %v", err)
	}

	return music, err
}

// update user in the DB
func UpdateMusic(id string, music Music) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// kita buat sql query create
	sqlStatement := `UPDATE music_list SET music_name=$2, music_album=$3, music_album_art=$4, music_singer=$5, music_publish_date=$6, music_updated_at=$7 WHERE music_id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id, music.Name, music.Album, music.Album_art, music.Singer, music.Publish_date, music.Updated_at)

	if err != nil {
		log.Fatalf("Tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa banyak row/data yang diupdate
	rowsAffected, err := res.RowsAffected()

	//kita cek
	if err != nil {
		log.Fatalf("Error ketika mengecheck rows/data yang diupdate. %v", err)
	}

	fmt.Printf("Total rows/record yang diupdate %v\n", rowsAffected)

	return rowsAffected
}

func DeleteMusic(id string) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	// kita tutup koneksinya di akhir proses
	defer db.Close()

	// buat sql query
	sqlStatement := `DELETE FROM music_list WHERE music_id=$1`

	// eksekusi sql statement
	res, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("tidak bisa mengeksekusi query. %v", err)
	}

	// cek berapa jumlah data/row yang di hapus
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("tidak bisa mencari data. %v", err)
	}

	fmt.Printf("Total data yang terhapus %v", rowsAffected)

	return rowsAffected
}
