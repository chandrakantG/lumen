package repositories

import (
	"database/sql"
	"fmt"

	"github.com/chandrakantG/lumen/internal/models"
)

type productRepository struct{}
type ProductRepository interface {
	GetRevenueByProduct(db *sql.DB, start, end string) ([]models.ProductRevenueResponse, error)
}

func NewProductRepository() ProductRepository {
	return &productRepository{}
}

func (r *productRepository) GetRevenueByProduct(db *sql.DB, start, end string) ([]models.ProductRevenueResponse, error) {
	var results []models.ProductRevenueResponse
	query := `
		SELECT p.name, SUM(oi.quantity * p.selling_price) AS revenue
		FROM order_items oi
		JOIN orders o ON o.id = oi.order_id
		JOIN products p ON p.id = oi.product_id
		WHERE o.order_date BETWEEN $1 AND $2
		GROUP BY p.name
		ORDER BY revenue DESC`

	rows, err := db.Query(query, start, end)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		var revenue float64
		if err := rows.Scan(&name, &revenue); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		results = append(results, models.ProductRevenueResponse{
			Product: name,
			Revenue: revenue,
		})
	}
	return results, nil
}
