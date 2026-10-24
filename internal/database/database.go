package database

import (
	"fmt"
	"log"

	"shopProject/internal/config"
	"shopProject/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=Asia/Tehran",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connection established successfully.")

	// Pre-migration: Handle transition from old schema to new schema
	preMigrationSQL := `
	DO $$
	DECLARE
		r RECORD;
	BEGIN
		-- 1. Drop foreign key constraints on products referencing categories
		FOR r IN (
			SELECT constraint_name
			FROM information_schema.table_constraints
			WHERE table_schema = 'public'
			  AND table_name = 'products'
			  AND constraint_type = 'FOREIGN KEY'
			  AND constraint_name ILIKE '%categor%'
		) LOOP
			EXECUTE 'ALTER TABLE public.products DROP CONSTRAINT IF EXISTS ' || quote_ident(r.constraint_name);
		END LOOP;
	END $$;

	-- 2. Drop old categories table if it has old columns (e.g. 'name') to rebuild with single category TEXT column
	DO $$
	BEGIN
		IF EXISTS (
			SELECT 1 FROM information_schema.columns 
			WHERE table_schema = 'public' AND table_name = 'categories' AND column_name = 'name'
		) THEN
			DROP TABLE public.categories CASCADE;
		END IF;
	END $$;

	-- 3. Create categories table with single column category TEXT
	CREATE TABLE IF NOT EXISTS public.categories (
		category TEXT PRIMARY KEY
	);

	-- 4. Add category TEXT and rating columns to products
	ALTER TABLE public.products ADD COLUMN IF NOT EXISTS category TEXT;
	ALTER TABLE public.products ADD COLUMN IF NOT EXISTS rating NUMERIC DEFAULT 4.5;

	-- 5. Drop NOT NULL on legacy category_id if present
	DO $$ BEGIN
		IF EXISTS (
			SELECT 1 FROM information_schema.columns 
			WHERE table_schema = 'public' AND table_name = 'products' AND column_name = 'category_id'
		) THEN
			ALTER TABLE public.products ALTER COLUMN category_id DROP NOT NULL;
		END IF;
	END $$;
	`
	if err := db.Exec(preMigrationSQL).Error; err != nil {
		log.Printf("Warning during pre-migration: %v\n", err)
	}

	// Auto Migration
	err = db.AutoMigrate(
		&models.User{},
		&models.OTPCode{},
		&models.Category{},
		&models.Product{},
		&models.Comment{},
		&models.CartItem{},
	)
	if err != nil {
		return nil, fmt.Errorf("auto migration failed: %w", err)
	}

	log.Println("Database migration completed.")
	DB = db
	return db, nil
}
