package store

import (
	"database/sql"

	"github.com/Laura-2950/API-Go-Products.git/API-Go-Products/internal/domain"
)

type SqlStore struct {
	DB *sql.DB
}

func (s *SqlStore) Read(id int) (*domain.Product, error) {
	var productReturn domain.Product

	query := "SELECT * FROM products WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&productReturn.ID, &productReturn.Name, &productReturn.Quantity, &productReturn.CodeValue, &productReturn.IsPublished, &productReturn.Expiration, &productReturn.Price)
	if err != nil {
		return nil, err
	}
	return &productReturn, nil
}

func (s *SqlStore) ReadAll() ([]domain.Product, error) {
	var productsReturn []domain.Product
	var productReturn domain.Product

	query := "SELECT * FROM products;"
	rows, err := s.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&productReturn.ID, &productReturn.Name, &productReturn.Quantity, &productReturn.CodeValue, &productReturn.IsPublished, &productReturn.Expiration, &productReturn.Price)
		if err != nil {
			return nil, err
		}
		productsReturn = append(productsReturn, productReturn)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return productsReturn, nil
}

func (s *SqlStore) Create(product domain.Product) (*domain.Product, error) {
	query := "INSERT INTO products (name, quantity, code_value, is_published, expiration, price) VALUES (?, ?, ?, ?, ?, ?);"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return &domain.Product{}, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price)
	if err != nil {
		return &domain.Product{}, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return &domain.Product{}, err
	}

	lid, _ := res.LastInsertId()
	product.ID = int(lid)
	return &product, nil
}

func (s *SqlStore) Update(product domain.Product) (*domain.Product, error) {
	stmt, err := s.DB.Prepare("UPDATE products SET name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? WHERE id = ?")
	if err != nil {
		return &domain.Product{}, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price, product.ID)
	if err != nil {
		return &domain.Product{}, err
	}
	return &product, nil
}

func (s *SqlStore) Delete(id int) error {
	stmt := "DELETE FROM products WHERE id = ?"
	_, err := s.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *SqlStore) Exist(codeValue string) bool {
	var exist bool
	var id int

	query := "SELECT id FROM products WHERE code_value = ?;"
	row := s.DB.QueryRow(query, codeValue)
	err := row.Scan(&id)
	if err != nil {
		return exist
	}

	if id > 0 {
		exist = true
	}

	return exist

}

func (s *SqlStore) ExistId(id int) bool {
	var exist bool
	query := "SELECT * FROM products WHERE id = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			exist = true
		}
	} else {
		exist = false
	}
	return exist
}
