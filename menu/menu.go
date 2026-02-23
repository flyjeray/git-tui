package menu

// rootMenu is the top-level menu tree shown on startup.
var RootMenu = []MenuItem{
	BranchMenuItem,
	RemotesMenuItem,
}

var StartMenu = []MenuLevel{{
	Items:  RootMenu,
	Cursor: 0,
}}
