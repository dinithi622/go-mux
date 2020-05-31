// model.go

package main

import (
    "database/sql"
)

type product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float64 `json:"price"`
}

type article struct {
    ID    int     `json:"id"`
    Product_ID    int     `json:"product_ID"`
    Article_name  string  `json:"article_name"`
}

func (p *product) getProduct(db *sql.DB) error {
    return db.QueryRow("SELECT name, price FROM products WHERE id=$1",
        p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) updateProduct(db *sql.DB) error {
    _, err :=
        db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3",
            p.Name, p.Price, p.ID)

    return err
}

func (p *product) deleteProduct(db *sql.DB) error {
    _, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

    return err
}

func (p *product) createProduct(db *sql.DB) error {
    err := db.QueryRow(
        "INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
        p.Name, p.Price).Scan(&p.ID)

    if err != nil {
        return err
    }

    return nil
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
    rows, err := db.Query(
        "SELECT id, name,  price FROM products LIMIT $1 OFFSET $2",
        count, start)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    products := []product{}

    for rows.Next() {
        var p product
        if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
            return nil, err
        }
        products = append(products, p)
    }

    return products, nil
}

func (p *article) createArticle(db *sql.DB) error {
    err := db.QueryRow(
        "INSERT INTO articles(product_ID, article_name) VALUES($1, $2) RETURNING id",
        p.Product_ID, p.Article_name).Scan(&p.ID)

    if err != nil {
        return err
    }

    return nil
}

func (p *article) getArticle(db *sql.DB) error {
    return db.QueryRow("SELECT article_name FROM articles WHERE id=$1",
        p.ID).Scan(&p.Article_name)
}

func getArticles(db *sql.DB, start, count int) ([]article, error) {
    rows, err := db.Query(
        "SELECT id, article_name FROM articles LIMIT $1 OFFSET $2",
        count, start)

    if err != nil {
        return nil, err
    }

    defer rows.Close()

    articles := []article{}

    for rows.Next() {
        var p article
        if err := rows.Scan(&p.ID, &p.Article_name); err != nil {
            return nil, err
        }
        articles = append(articles, p)
    }

    return articles, nil
}