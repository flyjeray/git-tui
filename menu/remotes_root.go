package menu

var RemotesMenuItem = MenuItem{
	Label: "Remotes",
	Children: []MenuItem{
		RemotesListItem,
		RemotesAddItem,
		RemotesDeleteItem,
	},
}
