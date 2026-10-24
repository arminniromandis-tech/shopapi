package database

import (
	"log"

	"shopProject/internal/models"

	"gorm.io/gorm"
)

// SeedData inserts sample categories and products
func SeedData(db *gorm.DB) {
	log.Println("Checking and seeding initial store data...")

	// 1. Categories (فقط مردانه، زنانه، بچه گانه، و کفش)
	categories := []models.Category{
		{Category: "بچه گانه"},
		{Category: "مردانه"},
		{Category: "زنانه"},
		{Category: "کفش"},
	}

	for _, cat := range categories {
		var existing models.Category
		if err := db.Where("category = ?", cat.Category).First(&existing).Error; err != nil {
			db.Create(&cat)
		}
	}

	// 2. Products (13 products)
	// - 5 مورد با درصد تخفیف خوب و IsSpecialOffer = true
	// - 1 مورد بدون تخفیف (DiscountPercent = 0, IsSpecialOffer = false)
	// - سایر موارد با تخفیف‌های معمول
	// - قیمت‌ها به تومان
	products := []models.Product{
		// محصولات بچگانه
		{
			Category:        "بچه گانه",
			Name:            "لباس بچگانه مدل ۱",
			ImageURL:        "https://ajdtyyslbskcwyseajbw.supabase.co/storage/v1/object/public/shop/kids1.png",
			Description:     "پوشاک بچگانه برای استفاده روزمره.",
			Price:           380000,
			DiscountPercent: 35,
			IsSpecialOffer:  true, // پیشنهاد شگفت انگیز ۱
			Rating:          4.5,
		},
		{
			Category:        "بچه گانه",
			Name:            "لباس بچگانه مدل ۲",
			ImageURL:        "https://ajdtyyslbskcwyseajbw.supabase.co/storage/v1/object/public/shop/kids2.png",
			Description:     "مدل دیگری از پوشاک بچگانه برای تکمیل استایل کودک.",
			Price:           420000,
			DiscountPercent: 15,
			IsSpecialOffer:  false,
			Rating:          4.2,
		},

		// محصولات مردانه
		{
			Category:        "مردانه",
			Name:            "پیراهن مردانه",
			ImageURL:        "https://ajdtyyslbskcwyseajbw.supabase.co/storage/v1/object/public/shop/men1shirt.png",
			Description:     "پیراهن مردانه برای استفاده روزمره.",
			Price:           650000,
			DiscountPercent: 40,
			IsSpecialOffer:  true, // پیشنهاد شگفت انگیز ۲
			Rating:          4.7,
		},
		{
			Category:        "مردانه",
			Name:            "هودی مردانه",
			ImageURL:        "https://ajdtyyslbskcwyseajbw.supabase.co/storage/v1/object/public/shop/men2hodi.png",
			Description:     "هودی مردانه برای استایل روزمره.",
			Price:           890000,
			DiscountPercent: 45,
			IsSpecialOffer:  true, // پیشنهاد شگفت انگیز ۳
			Rating:          4.8,
		},
		{
			Category:        "مردانه",
			Name:            "لباس آستین بلند مردانه",
			ImageURL:        "https://ajdtyyslbskcwyseajbw.supabase.co/storage/v1/object/public/shop/men3astinBoland.png",
			Description:     "لباس مردانه آستین بلند مناسب استفاده روزمره.",
			Price:           520000,
			DiscountPercent: 20,
			IsSpecialOffer:  false,
			Rating:          4.3,
		},
		{
			Category:        "مردانه",
			Name:            "جوراب مردانه",
			ImageURL:        "https://ajdtyyslbskcwyseajbw.supabase.co/storage/v1/object/public/shop/men4sokes.png",
			Description:     "جوراب مردانه برای استفاده روزانه.",
			Price:           95000,
			DiscountPercent: 0, // بدون تخفیف (دقیقاً یک محصول بدون تخفیف)
			IsSpecialOffer:  false,
			Rating:          4.1,
		},
		{
			Category:        "مردانه",
			Name:            "تیشرت مردانه مدل ۱",
			ImageURL:        "https://ajdtyyslbskcwyseajbw.supabase.co/storage/v1/object/public/shop/men5Tshirt.png",
			Description:     "تیشرت مردانه برای استایل روزمره.",
			Price:           480000,
			DiscountPercent: 10,
			IsSpecialOffer:  false,
			Rating:          4.4,
		},
		{
			Category:        "مردانه",
			Name:            "پیراهن خاکستری مردانه",
			ImageURL:        "https://ajdtyyslbskcwyseajbw.supabase.co/storage/v1/object/public/shop/men6grayShirt.png",
			Description:     "پیراهن مردانه با رنگ خاکستری.",
			Price:           590000,
			DiscountPercent: 25,
			IsSpecialOffer:  false,
			Rating:          4.6,
		},
		{
			Category:        "مردانه",
			Name:            "تیشرت مردانه مدل ۲",
			ImageURL:        "https://ajdtyyslbskcwyseajbw.supabase.co/storage/v1/object/public/shop/men7tShirt.png",
			Description:     "مدل دیگری از تیشرت مردانه برای استفاده روزمره.",
			Price:           450000,
			DiscountPercent: 15,
			IsSpecialOffer:  false,
			Rating:          4.2,
		},
		{
			Category:        "مردانه",
			Name:            "کیف روزمره",
			ImageURL:        "https://ajdtyyslbskcwyseajbw.supabase.co/storage/v1/object/public/shop/men8bag.png",
			Description:     "کیف مناسب استفاده روزمره.",
			Price:           780000,
			DiscountPercent: 30,
			IsSpecialOffer:  false,
			Rating:          4.5,
		},

		// محصولات زنانه
		{
			Category:        "زنانه",
			Name:            "لباس زنانه مدل ۱",
			ImageURL:        "https://ajdtyyslbskcwyseajbw.supabase.co/storage/v1/object/public/shop/woman1.png",
			Description:     "پوشاک زنانه برای تکمیل استایل روزمره.",
			Price:           850000,
			DiscountPercent: 50,
			IsSpecialOffer:  true, // پیشنهاد شگفت انگیز ۴
			Rating:          4.9,
		},
		{
			Category:        "زنانه",
			Name:            "لباس زنانه مدل ۲",
			ImageURL:        "https://ajdtyyslbskcwyseajbw.supabase.co/storage/v1/object/public/shop/woman2.png",
			Description:     "مدل دیگری از پوشاک زنانه.",
			Price:           920000,
			DiscountPercent: 35,
			IsSpecialOffer:  true, // پیشنهاد شگفت انگیز ۵
			Rating:          4.7,
		},
		{
			Category:        "زنانه",
			Name:            "لباس زنانه مدل ۳",
			ImageURL:        "https://ajdtyyslbskcwyseajbw.supabase.co/storage/v1/object/public/shop/woman3.png",
			Description:     "پوشاک زنانه با مدل متفاوت.",
			Price:           740000,
			DiscountPercent: 20,
			IsSpecialOffer:  false,
			Rating:          4.3,
		},
	}

	for _, p := range products {
		var existing models.Product
		if err := db.Where("image_url = ?", p.ImageURL).First(&existing).Error; err != nil {
			db.Create(&p)
		} else {
			// Update values if already existed
			existing.Category = p.Category
			existing.Name = p.Name
			existing.Description = p.Description
			existing.Price = p.Price
			existing.DiscountPercent = p.DiscountPercent
			existing.IsSpecialOffer = p.IsSpecialOffer
			existing.Rating = p.Rating
			db.Save(&existing)
		}
	}

	log.Println("Database seeded successfully with categories and products.")
}
