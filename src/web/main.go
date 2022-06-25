package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	_ "github.com/gin-gonic/gin"
	"github.com/h0ldnack/journalDandy/lib"
)

////////////////////////////////////////////////////////////////////////////////
// func main() {															  //
// 	tmpl := template.Must(template.ParseFiles("forms.html"))				  //
// 																			  //
// 	//db, err := sql.Open("sqlite3", file)									  //
// 																			  //
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {		  //
// 		if r.Method != http.MethodPost {									  //
// 			tmpl.Execute(w, nil)											  //
// 			return															  //
// 		}																	  //
// 		entry := Entry{														  //
// 			Datetime:  gTime(),												  //
// 			Gratitude: r.FormValue("gratitude"),							  //
// 			Note:      r.FormValue("note"),									  //
// 		}																	  //
// 																			  //
// 		// do something with details										  //
// 		content, err := json.Marshal(entry)									  //
// 		if err != nil {														  //
// 			panic(err)														  //
// 		}																	  //
// 		os.WriteFile("/tmp/testes.json", content, 0777)						  //
// 		fmt.Printf("%v\n", entry)											  //
// 		tmpl.Execute(w, struct{ Success bool }{true})						  //
// 	})																		  //
// 																			  //
// 	http.ListenAndServe(":8180", nil)										  //
// }																		  //
////////////////////////////////////////////////////////////////////////////////

func daily() {
	tmpl := template.Must(template.ParseFiles("daily.html"))

	http.HandleFunc("/daily", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		s1 := jd.Step{
			Tracked: r.FormValue("t1"),
			Grade:   r.FormValue("g1"),
		}
		s2 := jd.Step{
			Tracked: r.FormValue("t2"),
			Grade:   r.FormValue("g2"),
		}
		s3 := jd.Step{
			Tracked: r.FormValue("t3"),
			Grade:   r.FormValue("g3"),
		}
		s4 := jd.Step{
			Tracked: r.FormValue("t4"),
			Grade:   r.FormValue("g4"),
		}
		s5 := jd.Step{
			Tracked: r.FormValue("t5"),
			Grade:   r.FormValue("g5"),
		}

		ratings := [5]jd.Step{s1, s2, s3, s4, s5}

		daily := jd.Daily{
			Datetime: time.Now().Format(time.RFC3339),
			Ratings:  ratings,
		}
		fmt.Printf("%v Hi\n", daily)
		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8083", nil)
}
func main() {
	daily()
}
