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

# APIs

### 1. API Endpoint: POST /api/members

#### Purpose

Create a new member in the system.

#### Request Body (JSON)

```json
{
  "Name": "ASWINKRISHNA P",
  "Email": "aswin@example.com",
  "Phone": "9865272720"
}
```

#### Response Body (JSON)

```json
{
  "ID": "A001",
  "Name": "ASWINKRISHNA P",
  "Email": "aswin@example.com",
  "Phone": "9865272720",
  "Date": "2024-01-01"
}
```

- responds with the newly created member details including a unique ID.

### 2. API Endpoint: GET /api/members/:id

#### Purpose

Retrieve details of a specific member by their ID.

#### Request

- Endpoint: `/api/members/:id`
- Method: GET

#### Response Body (JSON)

```json
{
  "Member_id": "A001",
  "Name": "ASWIN KRISHNA P",
  "Email": "aswin@example.com",
  "Phone": "123-456-7890",
  "Date": "2024-01-01"
}
```

### 3. API Endpoint: POST /api/books

#### Purpose

Create a new book in the system.

#### Request Body (JSON)

```json
{
  "Title": "Book 2",
  "Author": "Author1",
  "Category": "Fiction",
  "Count": 2
}
```

#### Response Body (JSON)

```json
{
  "id": 1,
  "Title": "Book 2",
  "Author": "Author1",
  "Category": "Fiction",
  "Count": 2
}
```

- responds with the newly created book details including a unique ID (`id`).

### 4. API Endpoint: GET /api/books?category=CATEGORY&author=AUTHORNAME&prefix=ANYPREFIX

#### Purpose

Retrieve a list of books based on optional filtering parameters.

#### Request

- Endpoint: `/api/books`
- Method: GET
- Query Parameters:
  - `category`: Optional, Filters books by category.
  - `author`: Optional, Filters books by author.
  - `prefix`: Optional, Filters books by title prefix.

#### Response Body (JSON)

```json
[
  {
    "Book_id": 0,
    "Title": "To Kill a Mockingbird",
    "Author": "Harper Lee",
    "Category": "Fiction",
    "Count": 5
  }
]
```

### 5. API Endpoint: POST /api/borrow

#### Purpose

TO borrow Book

#### Request Body (JSON)

```json
{
  "member_id": "A003",
  "book_id": 6
}
```

#### Response Body (JSON)

### 6. API Endpoint: PATCH /api/return

#### Purpose

To return Book

#### Request Body (JSON)

```json
{
  "member_id": "A003",
  "book_id": 6
}
```

#### Response Body (JSON)
