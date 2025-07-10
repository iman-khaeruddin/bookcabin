package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type Voucher struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	CrewName     string    `json:"crew_name"`
	CrewID       string    `json:"crew_id"`
	FlightNumber string    `json:"flight_number"`
	FlightDate   string    `json:"flight_date"`
	AircraftType string    `json:"aircraft_type"`
	Seat1        string    `json:"seat1"`
	Seat2        string    `json:"seat2"`
	Seat3        string    `json:"seat3"`
}

var db *gorm.DB

func main() {
	dsn := os.Getenv("DB_DSN")
	db = connectToDBWithRetry(dsn, 10, 3*time.Second)
	db.AutoMigrate(&Voucher{})

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "PUT", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}))

	r.POST("/api/check", checkVoucher)
	r.POST("/api/generate", generateVoucher)
	r.Run(":8080")
}

func connectToDBWithRetry(dsn string, maxRetries int, delay time.Duration) *gorm.DB {
	var database *gorm.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("✅ Connected to DB")
			return database
		}

		fmt.Printf("⏳ Waiting for DB... (%d/%d) %v\n", i+1, maxRetries, err)
		time.Sleep(delay)
	}
	panic("❌ Failed to connect to DB after retries: " + err.Error())
}

type CheckRequest struct {
	FlightNumber string `json:"flightNumber" binding:"required"`
	Date         string `json:"date" binding:"required"`
}

type CheckResponse struct {
	Exists bool `json:"exists"`
}

func checkVoucher(c *gin.Context) {
	var req CheckRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var count int64
	db.Model(&Voucher{}).
		Where("flight_number = ? AND flight_date = ?", req.FlightNumber, req.Date).
		Count(&count)

	c.JSON(http.StatusOK, CheckResponse{Exists: count > 0})
}

type GenerateRequest struct {
	Name         string `json:"name" binding:"required"`
	ID           string `json:"id" binding:"required"`
	FlightNumber string `json:"flightNumber" binding:"required"`
	Date         string `json:"date" binding:"required"`
	Aircraft     string `json:"aircraft" binding:"required"`
}

type GenerateResponse struct {
	Success bool     `json:"success"`
	Seats   []string `json:"seats"`
}

func generateVoucher(c *gin.Context) {
	var req GenerateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check for existing assignment
	var existing int64
	db.Model(&Voucher{}).
		Where("flight_number = ? AND flight_date = ?", req.FlightNumber, req.Date).
		Count(&existing)
	if existing > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vouchers already generated for this flight and date."})
		return
	}

	// Generate random seats
	seats := generateRandomSeats(req.Aircraft)

	if len(seats) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unknown Aircraft."})
		return
	}

	// Save to DB
	voucher := Voucher{
		CrewName:     req.Name,
		CrewID:       req.ID,
		FlightNumber: req.FlightNumber,
		FlightDate:   req.Date,
		AircraftType: req.Aircraft,
		Seat1:        seats[0],
		Seat2:        seats[1],
		Seat3:        seats[2],
	}
	db.Create(&voucher)

	c.JSON(http.StatusOK, GenerateResponse{
		Success: true,
		Seats:   seats,
	})
}

func generateRandomSeats(aircraft string) []string {
	layout := map[string]struct {
		Rows []int
		Cols []string
	}{
		"ATR":            {Rows: rangeInts(1, 18), Cols: []string{"A", "C", "D", "F"}},
		"Airbus 320":     {Rows: rangeInts(1, 32), Cols: []string{"A", "B", "C", "D", "E", "F"}},
		"Boeing 737 Max": {Rows: rangeInts(1, 32), Cols: []string{"A", "B", "C", "D", "E", "F"}},
	}

	cfg, ok := layout[aircraft]
	if !ok {
		return []string{}
	}

	seatSet := make(map[string]struct{})
	var seats []string
	for len(seats) < 3 {
		row := cfg.Rows[randInt(0, len(cfg.Rows))]
		col := cfg.Cols[randInt(0, len(cfg.Cols))]
		seat := fmt.Sprintf("%d%s", row, col)
		if _, exists := seatSet[seat]; !exists {
			seatSet[seat] = struct{}{}
			seats = append(seats, seat)
		}
	}
	return seats
}

func rangeInts(start, end int) []int {
	arr := make([]int, end-start+1)
	for i := range arr {
		arr[i] = start + i
	}
	return arr
}

func randInt(min, max int) int {
	return min + int(time.Now().UnixNano())%(max-min)
}
