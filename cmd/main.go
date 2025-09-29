package main

import (
	"context"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	dsn := "host=localhost user=analyzer password=test123 dbname=kurs port=5432 sslmode=disable TimeZone=Europe/Berlin"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("got error != nil")
	}

	ctx := context.Background()

	db.AutoMigrate(&Product{})

	// create
	err = gorm.G[Product](db).Create(ctx, &Product{Code: "D42", Price: 100})
	err = gorm.G[Product](db).Create(ctx, &Product{Code: "ABC", Price: 300})

	// Read
	product, err := gorm.G[Product](db).Where("Code = ?", "ABC").First(ctx)
	fmt.Printf("First Product:\n Code: %s, Price: %d\n", product.Code, product.Price)
	products, err := gorm.G[Product](db).Where("price BETWEEN ? and ?", 100, 400).Find(ctx)
	for _, i := range products {
		fmt.Printf("Products: Code: %s, Price: %d\n", i.Code, i.Price)
	}

	// Update
	_, err = gorm.G[Product](db).Where("id = ?", product.ID).Update(ctx, "Price", 200)
	// Update multiple fields
	//_, err = gorm.G[Product](db).Where("id = ?", product.ID).Updates(ctx, map[string]interface{}{"Price": 500, "Code": "F42"})
	_ = db.Model(&products).Updates(map[string]interface{}{"Price": 500, "Code": "D42"})

	// Delete
	_, err = gorm.G[Product](db).Where("price BETWEEN ? and ?", 100, 400).Delete(ctx)

}
