package book

type Service interface {
	FindAll() ([]Books, error)
	FindById(id int) (Books, error)
	CreateBook(bookRequest BookRequest) (Books, error)
	Update(id int, bookRequest BookRequest) (Books, error)
	Delete(id int) (Books, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func(s *service) FindAll() ([]Books, error){
	return s.repository.FindAll()
}

func(s *service) FindById(id int) (Books, error){
	books, err := s.repository.FindById(id)
	return books, err
}

func(s *service) CreateBook(bookRequest BookRequest) (Books, error){
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := Books{
		Title: bookRequest.Title,
		Price: int(price),
		Description: bookRequest.Description,
		Rating: int(rating),
	}

	newBook, err := s.repository.CreateBook(book)

	return newBook, err
}

func(s *service) Update(id int, bookRequest BookRequest) (Books, error){
	book, _ := s.repository.FindById(id)

	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book.Title = bookRequest.Title
	book.Price = int(price)
	book.Rating = int(rating)
	book.Description = bookRequest.Description

	updatedBook , err := s.repository.Update(book)
	return updatedBook, err
}

func (s *service) Delete(id int) (Books, error){
	book, _ := s.repository.FindById(id)
	deletedBook , err := s.repository.Delete(book)

	return deletedBook, err
}