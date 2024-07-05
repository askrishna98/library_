package loaddata

import (
	"github.com/askrishna98/library_/models"
	service "github.com/askrishna98/library_/services"
)

func LoadData(DB *models.MockDB, IdGenerator *service.IdGenerator) {
	members := []models.Member{
		{Name: "Alice Johnson", Email: "alice@example.com", Phone: "123-456-7890", Date: "2024-01-01"},
		{Name: "Bob Smith", Email: "bob@example.com", Phone: "234-567-8901", Date: "2024-01-02"},
		{Name: "Charlie Brown", Email: "charlie@example.com", Phone: "345-678-9012", Date: "2024-01-03"},
		{Name: "David Wilson", Email: "david@example.com", Phone: "456-789-0123", Date: "2024-01-04"},
		{Name: "Eva Green", Email: "eva@example.com", Phone: "567-890-1234", Date: "2024-01-05"},
		{Name: "Frank Harris", Email: "frank@example.com", Phone: "678-901-2345", Date: "2024-01-06"},
		{Name: "Grace Lee", Email: "grace@example.com", Phone: "789-012-3456", Date: "2024-01-07"},
		{Name: "Hannah Taylor", Email: "hannah@example.com", Phone: "890-123-4567", Date: "2024-01-08"},
		{Name: "Ian Thomas", Email: "ian@example.com", Phone: "901-234-5678", Date: "2024-01-09"},
		{Name: "Jane White", Email: "jane@example.com", Phone: "012-345-6789", Date: "2024-01-10"},
	}

	books := []models.Book{
		{Title: "To Kill a Mockingbird", Author: "Harper Lee", Category: "Fiction", Count: 5},
		{Title: "1984", Author: "George Orwell", Category: "Science Fiction", Count: 3},
		{Title: "Pride and Prejudice", Author: "Jane Austen", Category: "Romance", Count: 4},
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Category: "Classic", Count: 6},
		{Title: "Moby Dick", Author: "Herman Melville", Category: "Adventure", Count: 2},
		{Title: "War and Peace", Author: "Leo Tolstoy", Category: "Historical", Count: 7},
		{Title: "The Catcher in the Rye", Author: "J.D. Salinger", Category: "Fiction", Count: 5},
		{Title: "The Hobbit", Author: "J.R.R. Tolkien", Category: "Fantasy", Count: 8},
		{Title: "Crime and Punishment", Author: "Fyodor Dostoevsky", Category: "Psychological", Count: 3},
		{Title: "Brave New World", Author: "Aldous Huxley", Category: "Dystopian", Count: 4},
		{Title: "The Odyssey", Author: "Homer", Category: "Epic", Count: 6},
		{Title: "Jane Eyre", Author: "Charlotte Bronte", Category: "Gothic", Count: 3},
		{Title: "The Divine Comedy", Author: "Dante Alighieri", Category: "Poetry", Count: 2},
		{Title: "Frankenstein", Author: "Mary Shelley", Category: "Horror", Count: 5},
		{Title: "Anna Karenina", Author: "Leo Tolstoy", Category: "Literary", Count: 4},
	}

	memberServices := service.GetInstanceOfMemberService(DB, IdGenerator)
	bookServices := service.GetInstanceOfBookService(DB, IdGenerator)

	for _, member := range members {
		memberServices.CreateMember(&member)
	}

	for _, book := range books {
		bookServices.CreateBook(&book)
	}
}
