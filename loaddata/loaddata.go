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
		{Title: "Wuthering Heights", Author: "Emily Bronte", Category: "Gothic", Count: 3},
		{Title: "The Picture of Dorian Gray", Author: "Oscar Wilde", Category: "Classic", Count: 5},
		{Title: "The Adventures of Sherlock Holmes", Author: "Arthur Conan Doyle", Category: "Mystery", Count: 6},
		{Title: "The Road", Author: "Cormac McCarthy", Category: "Post-apocalyptic", Count: 4},
		{Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Category: "Fantasy", Count: 7},
		{Title: "Les Misérables", Author: "Victor Hugo", Category: "Historical Fiction", Count: 5},
		{Title: "Slaughterhouse-Five", Author: "Kurt Vonnegut", Category: "Science Fiction", Count: 3},
		{Title: "One Hundred Years of Solitude", Author: "Gabriel García Márquez", Category: "Magical Realism", Count: 4},
		{Title: "The Handmaid's Tale", Author: "Margaret Atwood", Category: "Dystopian", Count: 6},
		{Title: "Gone with the Wind", Author: "Margaret Mitchell", Category: "Historical", Count: 5},
		{Title: "The Alchemist", Author: "Paulo Coelho", Category: "Philosophical", Count: 8},
		{Title: "Heart of Darkness", Author: "Joseph Conrad", Category: "Literary", Count: 3},
		{Title: "The Brothers Karamazov", Author: "Fyodor Dostoevsky", Category: "Psychological", Count: 4},
		{Title: "Catch-22", Author: "Joseph Heller", Category: "Satirical", Count: 6},
		{Title: "The Martian", Author: "Andy Weir", Category: "Science Fiction", Count: 5},
		{Title: "Dracula", Author: "Bram Stoker", Category: "Gothic Horror", Count: 4},
		{Title: "Moby-Dick", Author: "Herman Melville", Category: "Adventure", Count: 7},
		{Title: "A Tale of Two Cities", Author: "Charles Dickens", Category: "Historical Fiction", Count: 5},
		{Title: "The Kite Runner", Author: "Khaled Hosseini", Category: "Contemporary", Count: 6},
		{Title: "The Shining", Author: "Stephen King", Category: "Horror", Count: 3},
		{Title: "The Name of the Wind", Author: "Patrick Rothfuss", Category: "Fantasy", Count: 5},
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
