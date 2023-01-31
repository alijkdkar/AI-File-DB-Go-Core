package handeler

import (
	"net/http"

	"github.com/alijkdkar/AI-File-DB-Go-Core/pkg/entity"
	"github.com/alijkdkar/AI-File-DB-Go-Core/pkg/utility"
)

func GetHashFile(w http.ResponseWriter, r *http.Request) {
	redisDB, err := utility.GetInstance()
	if r.Method == "GET" {

		fileID := r.URL.Query().Get("fileID")

		if err != nil {
			panic("redis is down")
		}

		if realFileName, err := redisDB.Get(fileID); err == nil {
			filePath, er := redisDB.Get("FilePath")

			if er != nil {
				panic("redis Get Key Error")
			}
			f1 := entity.File{}
			f1.Load(realFileName)
			f1.SaveHash(filePath, fileID)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{'mission:complited'}"))
		w.WriteHeader(http.StatusAccepted)

	}
}

func GetUnHashFile(w http.ResponseWriter, r *http.Request) {
	redisDB, err := utility.GetInstance()
	if r.Method == "GET" {

		fileID := r.URL.Query().Get("fileID")

		if err != nil {
			panic("redis is down")
		}

		if realFileName, err := redisDB.Get(fileID); err == nil {
			filePath, er := redisDB.Get("FilePath")

			if er != nil {
				panic("redis Get Key Error")
			}
			f1 := entity.File{}
			f1.Load(realFileName)
			f1.Save(filePath, fileID)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{'mission:complited'}"))
		w.WriteHeader(http.StatusAccepted)

	}

}
