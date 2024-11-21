package utils

import (
	"context"
	"time"

	"github.com/alexliesenfeld/health"
	"gorm.io/gorm"
)

func HealthCheck(db *gorm.DB) health.Checker {
	Checker := health.NewChecker(
		// Set the time-to-live for our cache to 1 second (default).
		health.WithCacheDuration(1*time.Second),

		// Configure a global timeout that will be applied to all checks.
		health.WithTimeout(10*time.Second),

		// A check configuration to see if our database connection is up.
		// The check function will be executed for each HTTP request.
		health.WithCheck(health.Check{
			Name:    "database",      // Unique check name.
			Timeout: 2 * time.Second, // Specific timeout for this check.
			Check: func(ctx context.Context) error {
				sqlDB, err := db.DB() // Get the underlying sql.DB instance
				if err != nil {
					return err // Return error if failed to get sql.DB
				}
				return sqlDB.PingContext(ctx) // Ping the database with context
			},
		}),
	)
	return Checker
}
