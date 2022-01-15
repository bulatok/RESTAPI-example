package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"task1/internal/models"
	"task1/internal/store"
)

type Server struct{
	Store store.Store
	Router *mux.Router
}
func Start(conf *Config) error{
	srvr := Server{
		Store: store.NewPSQL(conf.PostgresLink),
		Router: mux.NewRouter(),
	}
	if err := srvr.Store.Open(); err != nil{
		return err
	}

	log.Println("starting listening on port :8080")

	defer srvr.Store.Close()
	srvr.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("Content-Type", "application/json")

		switch r.Method {
		case http.MethodGet:
			data, err := srvr.Store.GetUsers()
			if err != nil{
				log.Printf("%s in GET\n", err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("something went wrong while getting users"))
				return
			}

			w.WriteHeader(http.StatusOK)
			log.Println("successfully showed users")
			w.Write([]byte(data))

		case http.MethodPost:
			var user models.User

			defer r.Body.Close()

			data, _ := io.ReadAll(r.Body)
			if err := json.Unmarshal(data, &user); err != nil{
				log.Printf("%s in POST\n", err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("\nsomething went wrong while reading data"))
				return
			}
			if user.Name != "" && user.Surname != "" && user.PhoneNum != ""{
				if err := srvr.Store.AddUser(user); err != nil{
					log.Printf("%s in POST\n", err)
					w.WriteHeader(http.StatusBadRequest)
					w.Write([]byte("\nsomething went wrong while adding user"))
					return
				}
				log.Println("successfully added user")
				w.WriteHeader(http.StatusOK)
				return
			}

			log.Println("bad json in POST")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("\ninvalid json"))

		case http.MethodDelete:
			type T struct{ID int `json:"id"`}
			var req T

			defer r.Body.Close()
			data, _ := io.ReadAll(r.Body)

			if err := json.Unmarshal(data, &req); err != nil{
				log.Printf("%s in DELETE\n", err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("\nsomething went wrong while parsing json"))
				return
			}

			if req.ID == 0{
				log.Println("bad json in DELETE")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("\ninvalid json"))
				return
			}

			if err := srvr.Store.DeleteByID(req.ID); err != nil{
				log.Printf("%s in DELETE\n", err)
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("\nsomething went wrong while deleting user"))
				return
			}

			log.Println("successfully deleted user")
			w.WriteHeader(http.StatusOK)
		default:
			log.Println("Invalid request method")
			w.WriteHeader(http.StatusFailedDependency)
			w.Write([]byte("invalid request method"))
		}
	})


	return http.ListenAndServe(conf.Port, srvr.Router)
}