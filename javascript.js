"use strict"
const db = {
    "products": [
        { "id": 1, "name": "product1", "categoryId": 2 },
        { "id": 2, "name": "product2", "categoryId": 1 },
        { "id": 3, "name": "product3", "categoryId": 2 },
        { "id": 4, "name": "product4", "categoryId": 1 }
    ],
    "categories": [
        { "id": 1, "name": "category1" },
        { "id": 2, "name": "category2" }
    ]
}
class Repositories {
    constructor(data) {
        this.data = data
    }
    getAll() {
        return this.data
    }
    getById(id) {
        return this.data.find(item => item.id === id)
    }
}
class ProductRepository extends Repositories {
    constructor() {
        super(db.products.map(p => (Object.assign(Object.assign({}, p), { category: db.categories.find(c => c.id === p.categoryId) }))))
    }
}
class CategoryRepository extends Repositories {
    constructor() {
        super(db.categories)
    }
}
// Exemplo de uso
function fetchData() {
    const prodRepo = new ProductRepository()
    // Método getAll
    const products = prodRepo.getAll()
    console.log(products)
    // Método getById
    const productIdToFind = 2
    const productById = prodRepo.getById(productIdToFind)
    console.log(`Product with id ${productIdToFind}:`, productById)
    const catRepo = new CategoryRepository()
    // Método getAll
    const categories = catRepo.getAll()
    console.log(categories)
    // Método getById
    const categoryIdToFind = 1
    const categoryById = catRepo.getById(categoryIdToFind)
    console.log(`Category with id ${categoryIdToFind}:`, categoryById)
}
fetchData()
