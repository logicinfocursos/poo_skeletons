package main

import "fmt"

type Product struct {
    ID         int
    Name       string
    CategoryID int
    Category   *Category // Pode ser nil se não encontrado
}

type Category struct {
    ID   int
    Name string
}

type ProductRepository struct {
    Data []Product
}

type CategoryRepository struct {
    Data []Category
}

func (r *ProductRepository) GetAll() []Product {
    return r.Data
}

func (r *ProductRepository) GetByID(id int) *Product {
    for _, p := range r.Data {
        if p.ID == id {
            return &p
        }
    }
    return nil
}

func (r *CategoryRepository) GetAll() []Category {
    return r.Data
}

func (r *CategoryRepository) GetByID(id int) *Category {
    for _, c := range r.Data {
        if c.ID == id {
            return &c
        }
    }
    return nil
}

func fetchData() {
    products := []Product{
        {ID: 1, Name: "product1", CategoryID: 2},
        {ID: 2, Name: "product2", CategoryID: 1},
        {ID: 3, Name: "product3", CategoryID: 2},
        {ID: 4, Name: "product4", CategoryID: 1},
    }

    categories := []Category{
        {ID: 1, Name: "category1"},
        {ID: 2, Name: "category2"},
    }

    prodRepo := ProductRepository{Data: products}
    catRepo := CategoryRepository{Data: categories}

    // Exemplo de uso
    fmt.Println("Products:")
    for _, p := range prodRepo.GetAll() {
        fmt.Println(p)
    }

    fmt.Println("\nCategories:")
    for _, c := range catRepo.GetAll() {
        fmt.Println(c)
    }

    // Exemplo de busca por ID
    productIdToFind := 2
    productByID := prodRepo.GetByID(productIdToFind)
    fmt.Printf("\nProduct with id %d: %+v\n", productIdToFind, productByID)

    categoryIdToFind := 1
    categoryByID := catRepo.GetByID(categoryIdToFind)
    fmt.Printf("\nCategory with id %d: %+v\n", categoryIdToFind, categoryByID)
}

func main() {
    fetchData()
}
