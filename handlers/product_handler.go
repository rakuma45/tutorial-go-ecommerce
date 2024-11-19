package handlers

import (
	"rakuma45/tutorial-go-ecommerce/models"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var products []models.Product
		var wg sync.WaitGroup

		// Menambahkan satu goroutine ke WaitGroup
		wg.Add(1)

		// Memulai goroutine untuk melakukan operasi yg butuh waktu lama
		go func() {
			defer wg.Done() // Menandakan bahwa goroutine telah selesai
			db.Find(&products)
		}()

		// Menunggu goroutine selesai
		wg.Wait()

		ctx.JSON(200, products)
	}
}

func GetProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Memeriksa apakah datanya ada dengan memeriksa ID nya
		id := ctx.Param("id")
		var product models.Product
		// Mengambil hasil pertama dari tabel database 
		if err := db.First(&product, id).Error; err != nil {
			ctx.JSON(404, gin.H{"message": "Product not found"})
			return
		}
		ctx.JSON(200, product)
	}
}

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input models.Product
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Create(&input)
		ctx.JSON(201, input)
	}
}

func UpdateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Memeriksa apakah datanya ada dengan memeriksa ID nya
		id := ctx.Param("id")
		var product models.Product
		// Mengambil hasil pertama dari tabel database 
		if err := db.First(&product, id).Error; err != nil {
			ctx.JSON(404, gin.H{"message": "Product not found"})
			return
		}

		// Mengupdate data sesuai ID yang dicek
		var input models.Product
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{"message": "Invalid input"})
			return
		}

		db.Model(&product).Updates(input)
		ctx.JSON(200, product)
	}
}

func DeleteProduct(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Memeriksa apakah datanya ada dengan memeriksa ID nya
		id := ctx.Param("id")
		var product models.Product
		// Mengambil hasil pertama dari tabel database 
		if err := db.First(&product, id).Error; err != nil {
			ctx.JSON(404, gin.H{"message": "Product not found"})
			return
		}
		
		db.Delete(&product)
		ctx.JSON(200, gin.H{"message": "Product deleted"})
	}
}