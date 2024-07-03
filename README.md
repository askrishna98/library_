# Library Management

Develop a library management system backend in Go that manages details of members, details of books, and borrowing transactions, without using an external database.

- System should allow creating, retrieving , and deleting members,books
- Members should have unique id like - (A000,A001…,AA999.. etc)
- Users should be able to borrow and return books, with functionalities like checking book is available and calculate the penalty if due date is passed
- Unit tests should be implemented

1. ## Data Structures

Member Struct:

- member_id: Unique identifier for members.
- name: Name of the member.
- email: Email address of the member.
- phone: Phone number of the member.
- date: Registration date of the member.
- unique_id: Custom unique identifier formatted like A000, A001, etc.

Book Struct:

- book_id: Unique identifier for books.
- title: Title of the book.
- author: Author of the book.
- category: Category or genre of the book.
- count: Number of copies available.

Book Transaction Struct:

- borrow_id: Unique identifier for each borrowing transaction.
- member_id: ID of the member borrowing the book.
- book_id: ID of the book being borrowed.
- borrow_date: Date when the book was borrowed.
- due_date: Due date for returning the book.
- return_date: Date when the book was returned (if returned).

2. ## Data Storage

The mock database (MockDB) is structured with slices to store:
MockDB struct:

- Books: Slice of Book structs.
- Members: Slice of Member structs.
- BookTransactions: Slice of Transaction structs.

## Service Design Pattern

The functionalities are divided into separate services:

Member Services:

- CreateMember(newMember \*Member) error
- DeleteMember(member \*Member) error
- GetMemberByID(id string) \*Member

Book Services:

- FindBookByAuthor(author string)
- FindBookByCategory(category string)
- FindBookByPrefix(prefix string)
- CreateBook(newBook \*Book)
- DeleteBook(book \*Book)
- CheckBookAvailability(book \*Book) bool

Book Transaction Services:

- BorrowBook(memberID string, bookID int)
- ReturnBook(memberID string, bookID int)
