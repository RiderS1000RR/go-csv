package main

import (
	"archive/zip"
	"encoding/csv"
	"log"
	"os"
)

func main() {
	// 1. สร้างไฟล์ ZIP
	zipFile, err := os.Create("example.zip")
	if err != nil {
		log.Fatalf("ไม่สามารถสร้าง zip file ได้: %v", err)
	}
	defer zipFile.Close()

	// 2. สร้าง ZIP writer
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 3. สร้างไฟล์ CSV ข้างใน ZIP
	csvFileInZip, err := zipWriter.Create("example.csv")
	if err != nil {
		log.Fatalf("ไม่สามารถสร้างไฟล์ csv ใน zip ได้: %v", err)
	}

	// 4. สร้าง CSV writer ใช้กับ csvFileInZip
	csvWriter := csv.NewWriter(csvFileInZip)

	// 5. เขียนข้อมูล
	header := []string{"Name", "Age", "Email"}
	data := [][]string{
		{"Alice", "30", "alice@example.com"},
		{"Bob", "25", "bob@example.com"},
		{"Charlie", "35", "charlie@example.com"},
	}

	// เขียน header
	if err := csvWriter.Write(header); err != nil {
		log.Fatalf("เขียน header ไม่สำเร็จ: %v", err)
	}

	// เขียนข้อมูล
	for _, record := range data {
		if err := csvWriter.Write(record); err != nil {
			log.Fatalf("เขียน record ไม่สำเร็จ: %v", err)
		}
	}

	// 6. Flush ข้อมูล CSV เข้าตัว zip
	csvWriter.Flush()

	if err := csvWriter.Error(); err != nil {
		log.Fatalf("เกิดข้อผิดพลาดระหว่าง flush csv: %v", err)
	}

	log.Println("สร้าง ZIP ที่มี CSV เรียบร้อยแล้ว!")
}
