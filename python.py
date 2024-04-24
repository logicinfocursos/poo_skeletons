db = {
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

class Repositories:
    def __init__(self, data):
        self.data = data

    def get_all(self):
        return self.data

    def get_by_id(self, id):
        return next((item for item in self.data if item["id"] == id), None)

class ProductRepository(Repositories):
    def __init__(self):
        super().__init__([{
            **p,
            "category": next((c for c in db["categories"] if c["id"] == p["categoryId"]), None)
        } for p in db["products"]])

class CategoryRepository(Repositories):
    def __init__(self):
        super().__init__(db["categories"])

# Exemplo de uso
def fetch_data():
    prod_repo = ProductRepository()

    # Método get_all
    products = prod_repo.get_all()
    print(products)

    # Método get_by_id
    product_id_to_find = 2
    product_by_id = prod_repo.get_by_id(product_id_to_find)
    print(f"Product with id {product_id_to_find}:", product_by_id)

    cat_repo = CategoryRepository()

    # Método get_all
    categories = cat_repo.get_all()
    print(categories)

    # Método get_by_id
    category_id_to_find = 1
    category_by_id = cat_repo.get_by_id(category_id_to_find)
    print(f"Category with id {category_id_to_find}:", category_by_id)

fetch_data()
