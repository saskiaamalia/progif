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
 
/*func init() {
 tpl = template.Must(template.ParseGlob("templates/*"))
}*/
 
/*func main() {
 log.Fatalln(http.ListenAndServe(":8080", nil))
}*/

//create user struct
type sched struct {
    ID        	int64
    Tanggal  	string
    Kegiatan 	string
    Tempat  	string
    Keterangan  string
    PIC			string
}

func init() {
	db, err = sql.Open("mysql", "root:pass@tcp(127.0.0.1:3307)/jadwal")
	checkErr(err)
	err = db.Ping()
	checkErr(err)
	tpl = template.Must(template.ParseGlob("templates/*"))
}



//create handler
func main() {
    defer db.Close()
    http.HandleFunc("/", index)
    http.HandleFunc("/schedForm", schedForm)
    http.HandleFunc("/createSched", createSched)
    http.HandleFunc("/editSched", editSched)
    http.HandleFunc("/deleteSched", deleteSched)
    http.HandleFunc("/updateSched", updateSched)
    log.Println("Server is up on 8180 port")
    log.Fatalln(http.ListenAndServe(":8180", nil))
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

//handler function
func index(w http.ResponseWriter, req *http.Request) {
  rows, e := db.Query(
       `SELECT ID,
        Tanggal,
        Kegiatan,
        Tempat,
        Keterangan,
        PIC
          FROM schedule;`)
   checkErr(e)

    schedule := make([]sched, 0)
    for rows.Next() {
      schd := sched{}
      rows.Scan(&schd.ID, &schd.Tanggal, &schd.Kegiatan, &schd.Tempat,  &schd.Keterangan, &schd.PIC)
      schedule = append(schedule, schd)
    }
    log.Println(schedule)
    tpl.ExecuteTemplate(w, "index.gohtml", schedule)
}

// sched form handle function
func schedForm(w http.ResponseWriter, req *http.Request) {
  err = tpl.ExecuteTemplate(w, "schedForm.gohtml", nil)
  checkErr(err)
}

//create sched handle function
func createSched(w http.ResponseWriter, req *http.Request) {
   if req.Method == http.MethodPost {
        schd := sched{}
        schd.Tanggal = req.FormValue("Tanggal")
        schd.Kegiatan = req.FormValue("Kegiatan")
        schd.Tempat = req.FormValue("Tempat")
        schd.Keterangan = req.FormValue("Keterangan")
        schd.PIC = req.FormValue("PIC")
       /* bPass, e := bcrypt.GenerateFromPassword([]byte(req.FormValue("password")), bcrypt.MinCost)
        if e != nil {
                   http.Error(w, e.Error(), http.StatusInternalServerError)
                   return
                }
        usr.Password = bPass
        _, e = db.Exec(*/
         _, err = db.Exec(
         	"INSERT INTO schedule (Tanggal, Kegiatan, Tempat, Keterangan, PIC) VALUES (?, ?, ?, ?, ?)",
            schd.Tanggal,
            schd.Kegiatan,
            schd.Tempat,
            schd.Keterangan,
            schd.PIC,
       	)
       /* if e != nil {
                   http.Error(w, e.Error(), http.StatusInternalServerError)
                   return
                }*/
        checkErr(err)
        http.Redirect(w, req, "/", http.StatusSeeOther)
        tpl.ExecuteTemplate(w, "editSched.gohtml", schd)
        return
  	}
    http.Error(w, "Method Not Supported", http.StatusMethodNotAllowed)
}

// edit sched handle function
func editSched(w http.ResponseWriter, req *http.Request) {
    ID := req.FormValue("ID")
    rows, err := db.Query(
        `SELECT ID,
        	Tanggal,
       		Kegiatan,
        	Tempat,
        	Keterangan,
        	PIC
          FROM schedule
        WHERE ID = ` + ID + `;`)
    checkErr(err)
    schd := sched{}
    for rows.Next() {
        rows.Scan(&schd.ID, &schd.Tanggal, &schd.Kegiatan, &schd.Tempat,  &schd.Keterangan, &schd.PIC)
    }
    tpl.ExecuteTemplate(w, "editSched.gohtml", schd)
}

// update sched handle function
func updateSched(w http.ResponseWriter, req *http.Request) {
    _, er := db.Exec(
        "UPDATE schedule SET Tanggal = ?, Kegiatan = ?, Tempat = ?, Keterangan = ?, PIC = ? WHERE ID = ? ",
        req.FormValue("Tanggal"),
        req.FormValue("Kegiatan"),
        req.FormValue("Tempat"),
        req.FormValue("Keterangan"),
        req.FormValue("PIC"),
        req.FormValue("ID"),
    )
    checkErr(er)
    http.Redirect(w, req, "/", http.StatusSeeOther)
}

// delete sched handler function
func deleteSched(res http.ResponseWriter, req *http.Request) {
    ID := req.FormValue("ID")
    if ID == "" {
       http.Error(res, "Please write ID", http.StatusBadRequest)
           return
    }
    _, er := db.Exec("DELETE FROM schedule WHERE ID = ?", ID)
    checkErr(er)
    http.Redirect(res, req, "/", http.StatusSeeOther)
}
