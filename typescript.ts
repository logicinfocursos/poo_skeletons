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

interface Product {
    id: number;
    name: string;
    categoryId: number;
    category?: Category;
}

interface Category {
    id: number;
    name: string;
}

abstract class Repositories<T> {
    constructor(public data: T[]) { }

    getAll(): T[] {
        return this.data;
    }

    getById(id: number): T | undefined {
        return this.data.find(item => (item as any).id === id);
    }
}

class ProductRepository extends Repositories<Product> {
    constructor() {
        super(db.products.map(p => ({
            ...p,
            category: db.categories.find(c => c.id === p.categoryId)
        })));
    }
}

class CategoryRepository extends Repositories<Category> {
    constructor() {
        super(db.categories);
    }
}

// Exemplo de uso
function fetchData() {
    const prodRepo = new ProductRepository();

    // Método getAll
    const products = prodRepo.getAll();
    console.log(products);

    // Método getById
    const productIdToFind = 2;
    const productById = prodRepo.getById(productIdToFind);
    console.log(`Product with id ${productIdToFind}:`, productById);

    const catRepo = new CategoryRepository();

    // Método getAll
    const categories = catRepo.getAll();
    console.log(categories);

    // Método getById
    const categoryIdToFind = 1;
    const categoryById = catRepo.getById(categoryIdToFind);
    console.log(`Category with id ${categoryIdToFind}:`, categoryById);
}

fetchData();
