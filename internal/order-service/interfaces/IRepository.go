package order

type CacheRepository interface {
	GetCart()
	AddCart()
	DelCart()
}
