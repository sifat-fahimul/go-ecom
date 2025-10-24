package repo

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgUrl      string  `json:"imageUrl"`
}

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(id int) (*Product, error)
	List() ([]*Product, error)
	Delete(id int) error
	Update(p Product) (*Product, error)
}

type productRepo struct {
	productList []*Product
}

func NewProductRepo() ProductRepo {
	repo := &productRepo{}

	generateInitialProducts(repo)
	return repo
}

func (r *productRepo) Create(p Product) (*Product, error) {
	p.ID = len(r.productList) + 1
	r.productList = append(r.productList, &p)
	return &p, nil
}
func (r *productRepo) Get(id int) (*Product, error) {
	for _, product := range r.productList {
		if product.ID == id {
			return product, nil
		}
	}

	return nil, nil
}
func (r *productRepo) List() ([]*Product, error) {
	return r.productList, nil
}
func (r *productRepo) Delete(id int) error {
	var tempList []*Product
	for _, p := range r.productList {
		if p.ID != id {
			tempList = append(tempList, p)

		}
	}
	r.productList = tempList

	return nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	for idx, p := range r.productList {
		if p.ID == product.ID {
			r.productList[idx] = &product
			return r.productList[idx], nil
		}
	}
	return nil, nil
}

func generateInitialProducts(r *productRepo) {
	prod1 := &Product{
		ID:          1,
		Title:       "Apple",
		Description: "this is apple phone",
		Price:       2300,
		ImgUrl:      "dasfdjfl",
	}
	r.productList = append(r.productList, prod1)
}
