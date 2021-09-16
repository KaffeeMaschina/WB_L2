package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
 1. Реализовать вспомогательные функции для сериализации объектов доменной области в JSON.
 2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
 3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
 4. Реализовать middleware для логирования запросов

Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
 1. Реализовать все методы.
 2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
 3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
 4. Код должен проходить проверки go vet и golint.
*/

func DeleteEvent(userID, eventID int, date string) error {
	dirName := fmt.Sprintf("user_%d", userID)
	fileName := fmt.Sprintf("%s/%s%v.txt", dirName, date, eventID)
	if isUserExist := CheckUser(dirName); !isUserExist {
		return fmt.Errorf("cannot delete user: user directory %s does not exist", dirName)
	}
	if isEventExist := CheckEvent(dirName, date, eventID); !isEventExist {
		return fmt.Errorf("cannot delete event: event %s does not exist", dirName)
	}
	if err := os.Remove(fileName); err != nil {
		return err
	}

	files, err := os.ReadDir(dirName)
	if err != nil {
		return err
	}
	if len(files) == 0 {
		if err = os.Remove(dirName); err != nil {
			return err
		}
	} else {
		for _, f := range files {
			if strings.Contains(f.Name(), date) {
				id, err := GetID(f.Name(), date)
				if err != nil {
					return err
				}
				if id > eventID {
					err := os.Rename(fileName, fmt.Sprintf("%s/%s%v.txt", dirName, date, id-1))
					if err != nil {
						return err
					}

				}
			}
		}
	}
	return nil
}
func CheckEvent(dirName, date string, eventID int) bool {
	fileName := fmt.Sprintf("%s/%s%v.txt", dirName, date, eventID)
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		fmt.Printf("Event %s does not exist", date)
		return false
	}

	return true
}
func CheckUser(dirName string) bool {
	if _, err := os.Stat(dirName); os.IsNotExist(err) {
		fmt.Printf("user directory %s does not exist", dirName)
		return false
	}
	return true
}
func CreateUser(dirName string) error {
	if isUserExist := CheckUser(dirName); !isUserExist {
		err := os.Mkdir(dirName, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}
func CreateEvent(userID int, date, details string) error {
	var err error
	dirName := fmt.Sprintf("user_%d", userID)
	if isUserExist := CheckUser(dirName); !isUserExist {
		err = CreateUser(dirName)
		if err != nil {
			return err
		}
	}
	files, err := os.ReadDir(dirName)
	if err != nil {
		return err
	}
	max := -1
	for _, f := range files {
		if strings.Contains(f.Name(), date) {
			if temp, _ := GetID(f.Name(), date); temp > max {
				max = temp
			}
		}
	}
	if max == -1 {
		max = 0
	}
	id := max + 1
	fileName := fmt.Sprintf("%s/%s%v.txt", dirName, date, id)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write([]byte(details))

	return nil
}
func GetID(name, date string) (int, error) {
	noSuffix := strings.TrimSuffix(name, ".txt")
	id := strings.Split(noSuffix, date)[1]
	return strconv.Atoi(id)
}
func UpdateEvent(userID, eventID int, details, date string) error {
	dirName := fmt.Sprintf("user_%d", userID)
	fileName := fmt.Sprintf("%s/%s%v.txt", dirName, date, eventID)
	if isUserExist := CheckUser(dirName); !isUserExist {
		return fmt.Errorf("cannot update user: user directory %s does not exist", dirName)
	}
	if isEventExist := CheckEvent(dirName, date, eventID); !isEventExist {
		return fmt.Errorf("cannot update event: event %s does not exist", dirName)
	}
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Truncate(0)
	file.Write([]byte(details))
	return nil

}
func getEventsForDay(userID int, date string) ([]byte, error) {
	dirName := fmt.Sprintf("user_%d", userID)

	files, err := os.ReadDir(dirName)
	if err != nil {
		return nil, err
	}
	data := make([]byte, 0)
	for _, f := range files {
		if strings.Contains(f.Name(), date) {
			fileData, err := os.ReadFile(fmt.Sprintf("%s/%s", dirName, f.Name()))
			if err != nil {
				return nil, err
			}
			data = append(fileData, '\n')

		}
	}
	return data, nil
}
func getEventsForWeek(userID int, date string) ([]byte, error) {

	day, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, err
	}
	data := make([]byte, 0)
	for i := 0; i < 7; i++ {
		day = day.AddDate(0, 0, i)
		dayData, err := getEventsForDay(userID, day.Format("2006-01-02"))
		if err != nil {
			return nil, err
		}
		data = append(dayData, '\n')
	}

	return data, nil
}
func getEventsForMonth(userID int, date string) ([]byte, error) {

	day, err := time.Parse("2006-01-02", date)
	month := day.Month()
	if err != nil {
		return nil, err
	}
	data := make([]byte, 0)
	for i := 0; ; i++ {
		day = day.AddDate(0, 0, i)
		if day.Month() != month {
			break
		}
		dayData, err := getEventsForDay(userID, day.Format("2006-01-02"))
		if err != nil {
			return nil, err
		}
		data = append(dayData, '\n')
	}

	return data, nil
}
func main() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	http.HandleFunc("/create_event", createEventHandler)
	http.HandleFunc("/update_event", updateEventHandler)
	http.HandleFunc("/delete_event", deleteEventHandler)
	http.HandleFunc("/events_for_day", eventsForDayHandler)
	http.HandleFunc("/events_for_week", eventsForWeekHandler)
	http.HandleFunc("/events_for_month", eventsForMonthHandler)

	http.Handle("/", loggingMiddleware(http.DefaultServeMux))

	http.ListenAndServe("localhost:8080", nil)

	<-done
	log.Printf("server stopped\n")
}
func createEventHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(r.Form.Get("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	date := r.Form.Get("date")
	details := r.Form.Get("details")

	err = CreateEvent(userID, date, details)
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	response := map[string]string{"result": "Event created successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
func updateEventHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	userID, err := strconv.Atoi(r.Form.Get("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	date := r.Form.Get("date")
	details := r.Form.Get("details")
	eventID, err := strconv.Atoi(r.Form.Get("event_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = UpdateEvent(userID, eventID, details, date)
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
	}
	response := map[string]string{"result": "Event updated successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
func deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	userID, err := strconv.Atoi(r.Form.Get("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	date := r.Form.Get("date")
	eventID, err := strconv.Atoi(r.Form.Get("event_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	err = DeleteEvent(userID, eventID, date)
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	response := map[string]string{"result": "Event deleted successfully"}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
func eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	queryParams := r.URL.Query()
	userID, err := strconv.Atoi(queryParams.Get("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	date := queryParams.Get("date")
	events, err := getEventsForDay(userID, date)
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(events)
}
func eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	queryParams := r.URL.Query()
	userID, err := strconv.Atoi(queryParams.Get("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	date := queryParams.Get("date")
	events, err := getEventsForWeek(userID, date)
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(events)
}
func eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	queryParams := r.URL.Query()
	userID, err := strconv.Atoi(queryParams.Get("user_id"))
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	date := queryParams.Get("date")
	events, err := getEventsForMonth(userID, date)
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(events)
}
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
