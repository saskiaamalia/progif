package main
 
import (
 	"database/sql"
 	"html/template"
 	"log"
 	"net/http"
 
 _ "github.com/go-sql-driver/mysql"
)
 
var db *sql.DB
var err error
var tpl *template.Template

//Tipe struct untuk schedule
type sched struct {
    ID        	int64
    Tanggal  	string
    Kegiatan 	string
    Tempat  	string
    Keterangan  string
    PIC			string
}

func init() {
  // Membuka database jadwal_mbwg
	db, err = sql.Open("mysql", "root:@tcp(167.205.67.251:3306)/jadwal_mbwg")
	checkErr(err)
	err = db.Ping()
	checkErr(err)
	tpl = template.Must(template.ParseGlob("templates/*"))
}



//Fungsi utama
func main() {
    defer db.Close()
    http.HandleFunc("/", home)
    http.HandleFunc("/viewSched", index)
    http.HandleFunc("/schedForm", schedForm)
    http.HandleFunc("/createSched", createSched)
    http.HandleFunc("/editSched", editSched)
    http.HandleFunc("/deleteSched", deleteSched)
    http.HandleFunc("/updateSched", updateSched)
    log.Println("Server is up on 9091 port")
    log.Fatalln(http.ListenAndServe(":9091", nil))
}

// Fungsi cek error
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// Fungsi untuk menunjukkan tampilan home
func home(w http.ResponseWriter, req *http.Request) {
  tpl.ExecuteTemplate(w, "home.gohtml", nil)
}

//Fungsi untuk menunjukan tampilan data jadwal
func index(w http.ResponseWriter, req *http.Request) {
  rows, e := db.Query(
       `SELECT ID,
        Tanggal,
        Kegiatan,
        Tempat,
        Keterangan,
        PIC
          FROM schedule_mbwg;`)
   checkErr(e)

    schedule_mbwg := make([]sched, 0)
    for rows.Next() {
      schd := sched{}
      rows.Scan(&schd.ID, &schd.Tanggal, &schd.Kegiatan, &schd.Tempat,  &schd.Keterangan, &schd.PIC)
      schedule_mbwg = append(schedule_mbwg, schd)
    }
    log.Println(schedule_mbwg)
    tpl.ExecuteTemplate(w, "index.gohtml", schedule_mbwg)
}

// Fungsi untuk menampilkan form pengisian jadwal
func schedForm(w http.ResponseWriter, req *http.Request) {
  err = tpl.ExecuteTemplate(w, "schedForm.gohtml", nil)
  checkErr(err)
}

// Fungsi untuk membuat jadwal baru
func createSched(w http.ResponseWriter, req *http.Request) {
   if req.Method == http.MethodPost {
        schd := sched{}

        schd.Tanggal = req.FormValue("Tanggal")
        schd.Kegiatan = req.FormValue("Kegiatan")
        schd.Tempat = req.FormValue("Tempat")
        schd.Keterangan = req.FormValue("Keterangan")
        schd.PIC = req.FormValue("PIC")
      
         _, err = db.Exec(
         	"INSERT INTO schedule_mbwg (ID, Tanggal, Kegiatan, Tempat, Keterangan, PIC) VALUES (ID, ?, ?, ?, ?, ?);",
            schd.Tanggal,
            schd.Kegiatan,
            schd.Tempat,
            schd.Keterangan,
            schd.PIC,
       	)
    
        checkErr(err)
        http.Redirect(w, req, "/viewSched", http.StatusSeeOther)
        return
  	}
  http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
}

// Fungsi untuk mengubah jadwal yang sudah ada
func editSched(w http.ResponseWriter, req *http.Request) {
    ID := req.FormValue("ID")
    rows, err := db.Query(
        `SELECT ID,
        	Tanggal,
       		Kegiatan,
        	Tempat,
        	Keterangan,
        	PIC
          FROM schedule_mbwg
        WHERE ID = ` + ID + `;`)
    checkErr(err)
    schd := sched{}
    for rows.Next() {
        rows.Scan(&schd.ID, &schd.Tanggal, &schd.Kegiatan, &schd.Tempat,  &schd.Keterangan, &schd.PIC)
    }
    tpl.ExecuteTemplate(w, "editSched.gohtml", schd)
}

// Fungsi untuk memperbarui jadwal yang sudah ada
func updateSched(w http.ResponseWriter, req *http.Request) {
    _, er := db.Exec(
        "UPDATE schedule_mbwg SET Tanggal = ?, Kegiatan = ?, Tempat = ?, Keterangan = ?, PIC = ? WHERE ID = ?",
        req.FormValue("Tanggal"),
        req.FormValue("Kegiatan"),
        req.FormValue("Tempat"),
        req.FormValue("Keterangan"),
        req.FormValue("PIC"),
        req.FormValue("ID"),
    )
    checkErr(er)
    http.Redirect(w, req, "/viewSched", http.StatusSeeOther)
}

// Fungsi untuk menghapus jadwal
func deleteSched(w http.ResponseWriter, req *http.Request) {
    ID := req.FormValue("ID")
   /* if ID == "" {
       http.Error(w, "Please write ID", http.StatusBadRequest)
           return
    } else {*/
    _, er := db.Exec("DELETE FROM schedule_mbwg WHERE ID = ?;", ID)
    checkErr(er)
    //_, er := rows.Exec(ID)
    //checkErr(er)

    http.Redirect(w, req, "/viewSched", http.StatusSeeOther)
 // }
}
