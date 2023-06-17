package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

var quotes = []string{
	"I can do this all day. - Steve Rogers",
	"I'm always angry. - Bruce Banner",
	"That's my secret, Captain. I'm always angry. - Bruce Banner",
	"I am Groot. - Groot",
	"Part of the journey is the end. - Tony Stark",
	"Whatever it takes. - The Avengers",
	"Avengers Assemble. - Steve Rogers",
	"Some people move on, but not us. - Steve Rogers",
	"I am inevitable. - Thanos",
	"Sometimes the best we can do is to start over. - Peggy Carter",
	"We have a Hulk. - Tony Stark",
	"Wakanda Forever! - T'Challa",
	"Nothing goes over my head. My reflexes are too fast. - Drax",
	"I've got red in my ledger. I'd like to wipe it out. - Natasha Romanoff",
	"I don't care. He killed my mom. - Tony Stark",
	"I'm just a kid from Brooklyn. - Steve Rogers",
	"We are Venom. - Venom",
	"I have nothing to prove to you. - Carol Danvers",
	"He may have been your father, boy, but he wasn't your daddy. - Yondu",
	"I don't know the science, but I know how it feels. - Scott Lang",
	"I'm Mary Poppins, y'all! - Yondu",
	"I assure you, brother, the sun will shine on us again. - Loki",
	"This is my bargain, you mewling quim! - Loki",
	"Dread it. Run from it. Destiny arrives all the same. - Thanos",
}

func main() {
	rand.Seed(time.Now().UnixNano())

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintln(w, "\nSorry, You've taken a wrong turn!  - Adapted from Thor \n")
			return
		}
		fmt.Fprintln(w, "\nWith great power comes great responsibility. - Uncle Ben 🕷️ 🕸️ \n")
	})

	mux.HandleFunc("/quote", func(w http.ResponseWriter, r *http.Request) {
		quote := quotes[rand.Intn(len(quotes))]
		fmt.Fprintln(w, "\n"+quote+"\n")
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "\nI told you, I don't want to join your super server health group. – Adapted from Tony Stark\n")
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "\nThe Server is healthy. ❤️  \n")
	})

	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "\nHello, I am Iron Man. - Tony Stark\n")
	})

	fmt.Println("You trust me? - Tony Stark. 📣 I'm listening on :8080\n")
	http.ListenAndServe(":8080", mux)
}
