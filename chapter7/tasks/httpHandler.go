// p. 236
package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// for locking resources
var mu = &sync.Mutex{}

type dollars float32
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("%.2f", d) }

// task 7.12
func styled() map[string]interface{} {
	return map[string]interface{}{
		"table": template.CSS(`
			border-collapse: collapse;
			width: 40vw;
			text-align: left;
		`),
		"th": template.CSS(`
			border: 1px solid #ddd;
			padding: 6px;
			padding-top: 10px;
			padding-bottom: 10px;
			background-color: #4caf50;
			color: white;
		`),
		"td": template.CSS(`
			border: 1px solid #ddd;
			padding: 8px;
		`),
	}
}

const templ = `
<table style={{ styled.table }}>
	<tr>
		<th style={{ styled.th }}>Item</th>
		<th style={{ styled.th }}>Price, $</th>
	</tr>

	{{ range $item, $price := . }}
	<tr
		onmouseover="this.style.backgroundColor='#f7f7f7'"
		onmouseout="this.style.backgroundColor=''"
		>
		<td style={{ styled.td }}>{{ $item }}</td>
		<td style={{ styled.td }}>{{ $price }}</td>
	</tr>
	{{ end }}
</table>
`

func (db database) list(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	// task 7.12
	needTable := r.URL.Query().Get("table")
	// print styled table
	if needTable == "true" {
		t := template.Must(template.New("list").
			Funcs(template.FuncMap{"styled": styled}).
			Parse(templ))

		if err := t.Execute(w, db); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "<p style='color: red'>someting went wrong...</p>")
			fmt.Println(err)
			return
		}
	} else {
		// just print all items
		w.Write([]byte("for styled table enter: '/list?table=true'\n"))
		w.Write([]byte("------------------------------------------\n\n"))
		for item, price := range db {
			fmt.Fprintf(w, "%s: $%s\n", item, price)
		}
	}
}

func (db database) price(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	item := r.URL.Query().Get("item")

	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not found: %s\n", item)
		return
	}

	fmt.Fprintf(w, "$%s\n", price)
}

func (db database) read(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	item := r.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "no item")
		return
	}

	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "not found: %s\n", item)
		return
	}

	fmt.Fprintf(w, "%s - $%s\n", item, price)
}

func (db database) create(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	query := r.URL.Query()
	item, price := query.Get("item"), query.Get("price")

	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no item or no price: [%s] - [%s]\n", item, price)
		return
	}

	priceF, err := strconv.ParseFloat(price, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "can not convert <%v> to price", priceF)
		return
	}

	db[item] = dollars(priceF)
	fmt.Fprintf(w, "item (%s) added!", item)
}

func (db database) update(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	query := r.URL.Query()
	item, newPrice := query.Get("item"), query.Get("price")

	if item == "" || newPrice == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "no item or no price: [%s] - [%s]\n", item, newPrice)
		return
	}

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not found: %s\n", item)
		return
	}

	newPriceF, err := strconv.ParseFloat(newPrice, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "can not convert <%v> to price", newPrice)
		return
	}

	db[item] = dollars(newPriceF)
	fmt.Fprintf(w, "item (%s) updated!", item)
}

func (db database) delete(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	item := r.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "no item\n")
		return
	}

	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item not found: %s\n", item)
		return
	}

	delete(db, item)
	fmt.Fprintf(w, "item <%s> deleted!\n", item)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	db := database{"socks": 5, "hats": 25, "boots": 50}

	// mux := http.NewServeMux()
	// mux.Handle("/list", http.HandlerFunc(db.list))
	// mux.Handle("/price", http.HandlerFunc(db.price))

	// using DefaultServeMux
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	// task 7.11
	http.HandleFunc("/update", db.update)
	http.HandleFunc("/delete", db.delete)
	http.HandleFunc("/read", db.read)
	http.HandleFunc("/create", db.create)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// TODO: write test for this api
