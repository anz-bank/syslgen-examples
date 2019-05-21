package todos

var ID int64 = 5

var PostList = Posts{
	Post{Body: "post1 body", ID: 1, Title: "Todo post1 from serviceimpl", UserID: 1},
	Post{Body: "post2 body", ID: 2, Title: "Todo post2 from serviceimpl", UserID: 2},
	Post{Body: "post3 body", ID: 3, Title: "Todo post3 from serviceimpl", UserID: 3},
	Post{Body: "post4 body", ID: 4, Title: "Todo post4 from serviceimpl", UserID: 4},
	Post{Body: "post1 duplicate body", ID: 1, Title: "Todo post1 from serviceimpl", UserID: 4},
}

var Todos = []Todo{
	{Completed: true, ID: 1, Title: "Todo1 from serviceimpl", UserID: 1},
	{Completed: true, ID: 2, Title: "Todo2 from serviceimpl", UserID: 2},
	{Completed: true, ID: 3, Title: "Todo3 from serviceimpl", UserID: 3},
	{Completed: true, ID: 4, Title: "Todo4 from serviceimpl", UserID: 4},
	{Completed: true, ID: 1, Title: "Todo1 duplicate from serviceimpl", UserID: 5},
}
