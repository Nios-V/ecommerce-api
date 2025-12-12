package database

import (
	"fmt"
	"log"

	"github.com/Nios-V/ecommerce/api/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.AppConfig.DBHost,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBPort,
		config.AppConfig.DBSSLMode,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established")
}

func CreateEnums(db *gorm.DB) error {
	sqlOrderEnum := `
		DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'order_status_type') THEN
				CREATE TYPE order_status_type AS ENUM ('PENDING', 'PROCESSING', 'SHIPPED', 'DELIVERED', 'CANCELLED');
			END IF;
		END $$;
	`
	if err := db.Exec(sqlOrderEnum).Error; err != nil {
		return err
	}

	sqlPaymentMethodEnum := `
		DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_method_type') THEN
				CREATE TYPE payment_method_type AS ENUM ('CREDIT_CARD', 'DEBIT_CARD', 'BANK_TRANSFER');
			END IF;
		END $$;
	`
	if err := db.Exec(sqlPaymentMethodEnum).Error; err != nil {
		return err
	}

	sqlPaymentStatusEnum := `
		DO $$ BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'payment_status_type') THEN
				CREATE TYPE payment_status_type AS ENUM ('PENDING', 'SUCCESS', 'FAILED');
			END IF;
		END $$;
	`
	if err := db.Exec(sqlPaymentStatusEnum).Error; err != nil {
		return err
	}
	return nil
}

func Migrate(models ...interface{}) {
	if err := CreateEnums(DB); err != nil {
		log.Fatalf("failed to create enums: %v", err)
	}

	err := DB.AutoMigrate(models...)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}
	log.Println("Database migration completed")
}
