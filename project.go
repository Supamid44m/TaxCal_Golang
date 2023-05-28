package main

import (
	"html/template"
	"net/http"
	"strconv"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}
func main() {

	//port := os.Getenv("PORT")
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/process", procressor)
	//http.HandleFunc("/test", test)
	http.ListenAndServe(":8080", nil)

	//http.ListenAndServe(":"+port, nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func procressor(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	inc, err := strconv.ParseFloat(r.FormValue("income"), 64)
	expen, err := strconv.ParseFloat(r.FormValue("expen"), 64)

	/////ส่วนตัว
	disper, err := strconv.ParseFloat(r.FormValue("disper"), 64)
	diswed, err := strconv.ParseFloat(r.FormValue("diswed"), 64)
	chil1, err := strconv.ParseFloat(r.FormValue("chil1"), 64)
	chil2, err := strconv.ParseFloat(r.FormValue("chil2"), 64)
	f, err := strconv.ParseFloat(r.FormValue("f"), 64)
	m, err := strconv.ParseFloat(r.FormValue("m"), 64)
	fdp, err := strconv.ParseFloat(r.FormValue("fdp"), 64)
	mdp, err := strconv.ParseFloat(r.FormValue("mdp"), 64)
	husdp, err := strconv.ParseFloat(r.FormValue("husdp"), 64)
	/////ส่วนตัว
	//การลงทุน
	rmf, err := strconv.ParseFloat(r.FormValue("rmf"), 64)
	ssf, err := strconv.ParseFloat(r.FormValue("ssf"), 64)
	pvd, err := strconv.ParseFloat(r.FormValue("pvd"), 64)
	nrf, err := strconv.ParseFloat(r.FormValue("nrf"), 64)
	ann, err := strconv.ParseFloat(r.FormValue("ann"), 64)
	h1, err := strconv.ParseFloat(r.FormValue("h1"), 64)
	h2, err := strconv.ParseFloat(r.FormValue("h2"), 64)
	h3, err := strconv.ParseFloat(r.FormValue("h3"), 64)
	h4, err := strconv.ParseFloat(r.FormValue("h4"), 64)
	h5, err := strconv.ParseFloat(r.FormValue("h5"), 64)
	//การลงทุน
	if err != nil {
		// handle the error in some way
	}

	fname := r.FormValue("uname")
	d := struct {
		Income     float64 ///รายได้หลัก
		Fname      string
		Pincome    float64 ///เงินได้สุทธิ
		PincomeInt int64
		Expenses   float64 ///ค่าใช้จ่าย
		Discount   int64   ///ค่าลดหย่อนรวม
		Disper     float64
		Disgua     float64
		Disheal    float64
		Tax        float64
		Testfolat  float64
		Text       string
		Taxint     int
		Intc       int
		Check      float64
		Rmf        float64
		Ssf        float64
		Pvd        float64
		Nrf        float64
		Ann        float64
		H1         float64
		H2         float64
		H3         float64
		H4         float64
		H5         float64
		DisperInt  int64
		DishealInt int64
		DisguaInt  int64
	}{
		Fname:    fname,
		Income:   inc,
		Expenses: expen,
		Disper:   diswed + (chil1 * 30000) + (chil2 * 60000) + f + m + fdp + mdp + husdp + disper, ///ค่าลดหย่อนส่วนตัว
		Rmf:      rmf,
		Ssf:      ssf,
		Pvd:      pvd,
		Nrf:      nrf,
		Ann:      ann,
		H1:       h1,
		H2:       h2,
		H3:       h3,
		H4:       h4,
		H5:       h5,

		///เงินได้สุทธิ = รายได้หลัก-ค่าใช้จ่าย-ค่าลดหย่อนนรวม

	}
	///ลดหย่อนประกัน
	if d.H1 < 100000 {
		d.H1 = d.H1
	}
	if d.H1 > 100000 {
		d.H1 = 100000
	}
	if d.H2 < 25000 {
		d.H2 = d.H2
	}
	if d.H2 > 25000 {
		d.H2 = 25000
	}
	if d.H3 < 7200 {
		d.H3 = d.H3
	}
	if d.H3 > 7200 {
		d.H3 = 7200
	}
	if d.H4 < 15000 {
		d.H4 = d.H4
	}
	if d.H4 > 15000 {
		d.H4 = 15000
	}
	if d.H5 < 100000 {
		d.H5 = d.H5
	}
	if d.H5 > 100000 {
		d.H5 = 100000
	}
	///ลดหย่อนประกัน
	if d.Nrf < 13200 {
		d.Nrf = d.Nrf
	}
	if d.Nrf > 13200 {
		d.Nrf = 13200

	}

	d.Disheal = d.H1 + d.H2 + d.H3 + d.H4 + d.H5
	d.Disgua = d.Rmf + d.Ssf + d.Pvd + d.Nrf + d.Ann
	///RMF
	if d.Rmf <= (d.Pincome*30)/100 && d.Rmf < 500000 {
		d.Rmf = d.Rmf
	}
	if d.Rmf > (d.Pincome*30)/100 || d.Rmf > 500000 {
		d.Rmf = (d.Pincome * 30) / 100

	} ///RMF
	//SSf
	if d.Ssf <= (d.Pincome*30)/100 && d.Ssf < 200000 {
		d.Ssf = d.Ssf
	}
	if d.Ssf > (d.Pincome*30)/100 || d.Ssf > 200000 {
		d.Ssf = (d.Pincome * 30) / 100

	} //ssf
	///pvd
	if d.Pvd <= (d.Pincome*15)/100 && d.Pvd < 500000 {
		d.Pvd = d.Pvd
	}
	if d.Pvd > (d.Pincome*15)/100 || d.Pvd > 500000 {
		d.Pvd = (d.Pincome * 15) / 100

	} ///pvd
	//ann
	if d.Ann <= (d.Pincome*15)/100 && d.Ann < 200000 {
		d.Ann = d.Ann
	}
	if d.Pvd > (d.Pincome*15)/100 || d.Ann > 200000 {
		d.Ann = (d.Pincome * 15) / 100

	}
	//ann
	///รวมลดหน่อย ประกัน
	if d.Disheal > 100000 {
		d.Disheal = 100000
	} else {
		d.Disheal = d.Disheal
	}
	///รวมลดหน่อย ประกัน
	///รวม ลดหย่อน ลงทุน
	if d.Disgua > 100000 {
		d.Disgua = 100000
	} else {
		d.Disgua = d.Disgua
	}

	///รวม ลดหย่อน ลงทุน
	d.Discount = int64(d.Disper) + int64(d.Disgua) + int64(d.Disheal)
	d.Pincome = d.Income - d.Expenses - float64(d.Discount) ///เงินได้สุทธิ = รายได้หลัก-ค่าใช้จ่าย-ค่าลดหย่อนนรวม
	d.PincomeInt = int64(d.Pincome)

	if d.Pincome <= 150000 {
		d.Tax = 0
		d.Taxint = int(d.Tax)
		tpl.ExecuteTemplate(w, "procress.gohtml", d)
	}
	if d.Pincome > 150000 && d.Pincome <= 300000 {
		d.Tax = ((d.Pincome - 150000) * 0.05)
		d.Taxint = int(d.Tax)
		tpl.ExecuteTemplate(w, "procress.gohtml", d)

	}
	if d.Pincome > 300000 && d.Pincome <= 500000 {
		d.Tax = ((d.Pincome - 300000) * 0.10) + 7500
		d.Taxint = int(d.Tax)
		tpl.ExecuteTemplate(w, "procress.gohtml", d)

	}
	if d.Pincome > 500000 && d.Pincome <= 750000 {
		d.Tax = ((d.Pincome - 500000) * 0.15) + 27500
		d.Taxint = int(d.Tax)
		tpl.ExecuteTemplate(w, "procress.gohtml", d)

	}
	if d.Pincome > 750000 && d.Pincome <= 1000000 {
		d.Tax = ((d.Pincome - 750000) * 0.20) + 65000
		d.Taxint = int(d.Tax)
		tpl.ExecuteTemplate(w, "procress.gohtml", d)

	}
	if d.Pincome > 1000000 && d.Pincome <= 2000000 {
		d.Tax = ((d.Pincome - 1000000) * 0.25) + 115000
		d.Taxint = int(d.Tax)
		tpl.ExecuteTemplate(w, "procress.gohtml", d)

	}
	if d.Pincome > 2000000 && d.Pincome <= 5000000 {
		d.Tax = ((d.Pincome - 2000000) * 0.30) + 365000
		d.Taxint = int(d.Tax)
		tpl.ExecuteTemplate(w, "procress.gohtml", d)

	}
	if d.Pincome > 5000000 {
		d.Tax = ((d.Pincome - 5000000) * 0.35) + 1265000
		d.Taxint = int(d.Tax)
		tpl.ExecuteTemplate(w, "procress.gohtml", d)

	}
	//tpl.ExecuteTemplate(w, "procress.gohtml", d)
}
