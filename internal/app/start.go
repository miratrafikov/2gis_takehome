// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
package app

import (
	"applicationDesignTest/internal/config"
	"applicationDesignTest/internal/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var Orders = []model.Order{}

var Availability = []model.RoomAvailability{
	{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 1), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 2), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 3), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 4), Quota: 1},
	{HotelID: "reddison", RoomID: "lux", Date: date(2024, 1, 5), Quota: 0},
}

func Start(ctx context.Context, cfg config.Config) {
	mux := http.NewServeMux()
	mux.HandleFunc("/orders", createOrder)

	LogInfo("Server listening on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if errors.Is(err, http.ErrServerClosed) {
		LogInfo("Server closed")
	} else if err != nil {
		LogErrorf("Server failed: %s", err)
		os.Exit(1)
	}
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var newOrder model.Order
	json.NewDecoder(r.Body).Decode(&newOrder)

	daysToBook := daysBetween(newOrder.From, newOrder.To)

	unavailableDays := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDays[day] = struct{}{}
	}

	for _, dayToBook := range daysToBook {
		for i, availability := range Availability {
			if !availability.Date.Equal(dayToBook) || availability.Quota < 1 {
				continue
			}
			availability.Quota -= 1
			Availability[i] = availability
			delete(unavailableDays, dayToBook)
		}
	}

	if len(unavailableDays) != 0 {
		http.Error(w, "Hotel room is not available for selected dates", http.StatusInternalServerError)
		LogErrorf("Hotel room is not available for selected dates:\n%v\n%v", newOrder, unavailableDays)
		return
	}

	Orders = append(Orders, newOrder)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newOrder)

	LogInfo("Order successfully created: %v", newOrder)
}

func daysBetween(from time.Time, to time.Time) []time.Time {
	if from.After(to) {
		return nil
	}

	days := make([]time.Time, 0)
	for d := toDay(from); !d.After(toDay(to)); d = d.AddDate(0, 0, 1) {
		days = append(days, d)
	}

	return days
}

func toDay(timestamp time.Time) time.Time {
	return time.Date(timestamp.Year(), timestamp.Month(), timestamp.Day(), 0, 0, 0, 0, time.UTC)
}

func date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}

var logger = log.Default()

func LogErrorf(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	logger.Printf("[Error]: %s\n", msg)
}

func LogInfo(format string, v ...any) {
	msg := fmt.Sprintf(format, v...)
	logger.Printf("[Info]: %s\n", msg)
}
