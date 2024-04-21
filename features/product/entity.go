package product

type Product struct {
	ID uint
	NameProduct string
	Price int
	UserID uint
}

// interface untuk Data Layer
type ProductDataInterface interface {
	Create(userIdLogin int, input Product) error
	GetProductById(idProduct int) (*Product, error)
	GetAllProduct()([]Product, error)
	Update(userIdLogin, idProduct int, input Product) error
	Delete(userIdLogin, idProduct int) error
}

// interface untuk Service Layer
type ProductServiceInterface interface {
	Create(userIdLogin int, input Product) error
	GetProductById(idProduct int) (*Product, error)
	GetAllProduct()([]Product, error)
	Update(userIdLogin, idProduct int, input Product) error
	Delete(userIdLogin, idProduct int) error
}
