package controller

import (
	"encoding/json" // package untuk enkode dan mendekode json menjadi struct dan sebaliknya
	"fmt"
	// "strconv" // package yang digunakan untuk mengubah string menjadi tipe int

	"log"
	"net/http" // digunakan untuk mengakses objek permintaan dan respons dari api

	"go-postgres-crud/models" //models package dimana Music didefinisikan

	"github.com/gorilla/mux" // digunakan untuk mendapatkan parameter dari router
	_ "github.com/lib/pq"    // postgres golang driver
)

type response struct {
	ID      string  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

type Response struct {
	Status  int           `json:"status"`
	Message string        `json:"message"`
	Data    []models.Music `json:"data"`
}

// Tambah Music
func AddMusic(w http.ResponseWriter, r *http.Request) {

	// create an empty user of type models.User
	// kita buat empty music dengan tipe models.Music
	var music models.Music

	// decode data json request ke music
	err := json.NewDecoder(r.Body).Decode(&music)

	if err != nil {
		log.Fatalf("Tidak bisa mendecode dari request body.  %v", err)
	}

	// panggil modelsnya lalu insert music
	insertID := models.AddMusic(music)

	// format response objectnya
	res := response{
		ID:      insertID,
		Message: "Data music telah ditambahkan",
	}

	// kirim response
	json.NewEncoder(w).Encode(res)
}

// AmbilMusic mengambil single data dengan parameter id
func GetMusic(w http.ResponseWriter, r *http.Request) {
	// kita set headernya
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// dapatkan id music dari parameter request, keynya adalah "id"
	params := mux.Vars(r)
	id 	:= params["id"]

	// memanggil models ambilsatumusic dengan parameter id yg nantinya akan mengambil single data
	music, err := models.GetOneMusic(string(id))

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data music. %v", err)
	}

	// kirim response
	json.NewEncoder(w).Encode(music)
}

// Ambil semua data music
func GetAllMusic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// memanggil models GetAllMusic
	musics, err := models.GetAllMusic()

	if err != nil {
		log.Fatalf("Tidak bisa mengambil data. %v", err)
	}

	var response Response
	response.Status = 1
	response.Message = "Success"
	response.Data = musics

	// kirim semua response
	json.NewEncoder(w).Encode(response)
} 

func UpdateMusic(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)
	id := params["id"]

	// konversikan ke int yang sebelumnya adalah string
	// id, err := strconv.Atoi(params["id"])

	// if err != nil {
	// 	log.Fatalf("Tidak bisa mengubah dari string ke int.  %v", err)
	// }

	// buat variable music dengan type models.Music
	var music models.Music

	// decode json request ke variable music
	err := json.NewDecoder(r.Body).Decode(&music)

	if err != nil {
		log.Fatalf("Tidak bisa decode request body.  %v", err)
	}

	// panggil updatemusic untuk mengupdate data
	updatedRows := models.UpdateMusic(string(id), music)

	// ini adalah format message berupa string
	msg := fmt.Sprintf("Music telah berhasil diupdate. Jumlah yang diupdate %v rows/record", updatedRows)

	// ini adalah format response message
	res := response{
		ID:      string(id),
		Message: msg,
	}

	// kirim berupa response
	json.NewEncoder(w).Encode(res)
}

func DeleteMusic(w http.ResponseWriter, r *http.Request) {

	// kita ambil request parameter idnya
	params := mux.Vars(r)
	id := params["id"]

	// panggil fungsi hapusmusic , dan convert int ke int64
	deletedRows := models.DeleteMusic(string(id))


	// ini adalah format message berupa string
	msg := fmt.Sprintf("music sukses di hapus. Total data yang dihapus %v", deletedRows)

	// ini adalah format reponse message
	res := response{
		ID:      string(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}
