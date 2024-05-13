// Ниже реализован сервис бронирования номеров в отеле. В предметной области
// выделены два понятия: Order — заказ, который включает в себя даты бронирования
// и контакты пользователя, и RoomAvailability — количество свободных номеров на
// конкретный день.
//
// Задание:
// - провести рефакторинг кода с выделением слоев и абстракций
// - применить best-practices там где это имеет смысл
// - исправить имеющиеся в реализации логические и технические ошибки и неточности
package main

import (
	"applicationDesignTest/internal/app"
	"applicationDesignTest/internal/config"
	"context"
	"log"
)

func main() {
	ctx := context.Background()
	cfg := getConfig()
	app.Start(ctx, cfg)
}

func getConfig() config.Config {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}
	return cfg
}
