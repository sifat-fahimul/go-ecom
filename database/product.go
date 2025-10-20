package database

var productList []Product

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

func Store(p Product) Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return p
}

func List() []Product {
	return productList
}

func Get(id int) *Product {
	for _, product := range productList {
		if product.ID == id {
			return &product
		}
	}

	return nil

}

func Update(product Product) *Product {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product
			return &productList[idx]
		}
	}
	return nil
}

func Delete(id int) {
	var tempList []Product
	for _, p := range productList {
		if p.ID != id {
			tempList = append(tempList, p)

		}
	}
	productList = tempList
}

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Apple",
		Description: "this is apple phone",
		Price:       2300,
		ImgUrl:      "dasfdjfl",
	}
	productList = append(productList, prod1)
}
