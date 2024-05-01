package main

import (
    "encoding/json"
    "fmt"
)

type Product struct {
    ID         int
    Name       string
    CategoryID int
    Category   *Category
}

type Category struct {
    ID   int
    Name string
}

type Repository interface {
    GetAll() []interface{}
    GetByID(int) interface{}
}

type ProductRepository struct {
    data []Product
}

type CategoryRepository struct {
    data []Category
}

var db = struct {
    Products   []Product
    Categories []Category
}{
    Products: []Product{
        {ID: 1, Name: "product1", CategoryID: 2},
        {ID: 2, Name: "product2", CategoryID: 1},
        {ID: 3, Name: "product3", CategoryID: 2},
        {ID: 4, Name: "product4", CategoryID: 1},
    },
    Categories: []Category{
        {ID: 1, Name: "category1"},
        {ID: 2, Name: "category2"},
    },
}

func (r *ProductRepository) GetAll() []interface{} {
    var products []interface{}
    for _, product := range r.data {
        products = append(products, product)
    }
    return products
}

func (r *ProductRepository) GetByID(id int) interface{} {
    for _, product := range r.data {
        if product.ID == id {
            return product
        }
    }
    return nil
}

func (r *CategoryRepository) GetAll() []interface{} {
    var categories []interface{}
    for _, category := range r.data {
        categories = append(categories, category)
    }
    return categories
}

func (r *CategoryRepository) GetByID(id int) interface{} {
    for _, category := range r.data {
        if category.ID == id {
            return category
        }
    }
    return nil
}

func NewProductRepository() *ProductRepository {
    var products []Product
    for _, product := range db.Products {
        product := product // Create a new variable for each iteration
        for i := range db.Categories {
            if product.CategoryID == db.Categories[i].ID {
                product.Category = &db.Categories[i]
            }
        }
        products = append(products, product)
    }
    return &ProductRepository{data: products}
}

func NewCategoryRepository() *CategoryRepository {
    return &CategoryRepository{data: db.Categories}
}

func printJSON(v interface{}) {
    b, _ := json.MarshalIndent(v, "", "  ")
    fmt.Println(string(b))
}

func main() {
    prodRepo := NewProductRepository()

    products := prodRepo.GetAll()
    printJSON(products)

    productById := prodRepo.GetByID(2)
    fmt.Println("Product with id 2:")
    printJSON(productById)

    catRepo := NewCategoryRepository()

    categories := catRepo.GetAll()
    printJSON(categories)

    categoryById := catRepo.GetByID(1)
    fmt.Println("Category with id 1:")
    printJSON(categoryById)
}
