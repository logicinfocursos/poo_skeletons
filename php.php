<?php

$db = [
    "products" => [
        [ "id" => 1, "name" => "product1", "categoryId" => 2 ],
        [ "id" => 2, "name" => "product2", "categoryId" => 1 ],
        [ "id" => 3, "name" => "product3", "categoryId" => 2 ],
        [ "id" => 4, "name" => "product4", "categoryId" => 1 ]
    ],
    "categories" => [
        [ "id" => 1, "name" => "category1" ],
        [ "id" => 2, "name" => "category2" ]
    ]    
];

class Product {
    public $id;
    public $name;
    public $categoryId;

    public function __construct($id, $name, $categoryId) {
        $this->id = $id;
        $this->name = $name;
        $this->categoryId = $categoryId;
    }
}

class Category {
    public $id;
    public $name;

    public function __construct($id, $name) {
        $this->id = $id;
        $this->name = $name;
    }
}

interface IRepository {
    public function getAll();
}

abstract class BaseRepository implements IRepository {
    public $data;

    public function __construct($data) {
        $this->data = $data;
    }

    public function getAll() {
        return $this->data;
    }

    public function getById($id) {
        foreach ($this->data as $item) {
            if ($item->id === $id) {
                return $item;
            }
        }
        return null;
    }
}

class ProductRepository extends BaseRepository {
    public function __construct($db) {
        parent::__construct(array_map(function($p) use ($db) {
            return new Product($p["id"], $p["name"], $p["categoryId"]);
        }, $db["products"]));
    }
}

class CategoryRepository extends BaseRepository {
    public function __construct($db) {
        parent::__construct(array_map(function($c) {
            return new Category($c["id"], $c["name"]);
        }, $db["categories"]));
    }
}

$productRepository = new ProductRepository($db);
print_r($productRepository->getAll());

$categoryRepository = new CategoryRepository($db);
print_r($categoryRepository->getAll());
