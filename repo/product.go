package repo

type Product struct {
	ID          int    `json:"_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Price       string `json:"price"`
	ImgUrl      string `json:"imageUrl"`
}

var productList = make([]Product, 0)

type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productId int) (*Product, error)
	List() *[]Product
	Delete(productId int) error
	Update(p Product) error
}

type productRepo struct {
	product *[]Product
}

func NewProductRepo() ProductRepo {
	return &productRepo{}
}
func (r productRepo) Create(p Product) (*Product, error) {
	productList = append(productList, p)
	return nil, nil
}
func (r productRepo) Get(productId int) (*Product, error) {
	for _, product := range productList {
		if product.ID == productId {
			return &product, nil
		}
	}
	return nil, nil
}
func (r productRepo) List() *[]Product {
	return &productList
}
func (r productRepo) Delete(productId int) error {
	var tmp []Product
	for _, product := range productList {
		if product.ID != productId {
			tmp = append(tmp, product)
		}
	}
	productList = tmp
	return nil
}
func (r productRepo) Update(p Product) error {
	for i, product := range productList {
		if product.ID == p.ID {
			productList[i] = p
			// fmt.Println(productList[i])
			return nil
		}
	}
	return nil
}
