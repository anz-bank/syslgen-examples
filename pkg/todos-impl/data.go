package todosimpl

import "github.com/anz-bank/syslgen-examples/todos"

// PostList ...
var PostList = todos.Posts{
	todos.Post{Body: "post1 body", ID: 1, Title: "Todo post1 from serviceimpl", UserID: 1},
	todos.Post{Body: "post2 body", ID: 2, Title: "Todo post2 from serviceimpl", UserID: 2},
	todos.Post{Body: "post3 body", ID: 3, Title: "Todo post3 from serviceimpl", UserID: 3},
	todos.Post{Body: "post4 body", ID: 4, Title: "Todo post4 from serviceimpl", UserID: 4},
	todos.Post{Body: "post1 duplicate body", ID: 1, Title: "Todo post1 from serviceimpl", UserID: 4},
}

// Todos ...
var Todos = []todos.Todo{
	{Completed: true, ID: 1, Title: "Todo1 from serviceimpl", UserID: 1},
	{Completed: true, ID: 2, Title: "Todo2 from serviceimpl", UserID: 2},
	{Completed: true, ID: 3, Title: "Todo3 from serviceimpl", UserID: 3},
	{Completed: true, ID: 4, Title: "Todo4 from serviceimpl", UserID: 4},
}
