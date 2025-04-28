package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// 1. สร้างไฟล์ใหม่
	file, err := os.Create("example.csv")
	if err != nil {
		log.Fatalf("ไม่สามารถสร้างไฟล์ได้: %v", err)
	}
	defer file.Close()

	// 2. สร้าง writer สำหรับเขียนข้อมูล csv
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// 3. เขียนข้อมูลลงไป
	header := []string{"Name", "Age", "Email"}
	data := [][]string{
		{"Alice", "30", "alice@example.com"},
		{"Bob", "25", "bob@example.com"},
		{"Charlie", "35", "charlie@example.com"},
	}

	// เขียน header
	if err := writer.Write(header); err != nil {
		log.Fatalf("ไม่สามารถเขียน header: %v", err)
	}

	// เขียนข้อมูล
	for _, record := range data {
		if err := writer.Write(record); err != nil {
			log.Fatalf("ไม่สามารถเขียน record: %v", err)
		}
	}

	log.Println("สร้างไฟล์ CSV สำเร็จ!")
}
