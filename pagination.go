package hillchartsapi

type PaginationOpts struct {
	Limit int
	Page  int
}

type PaginatedList struct {
	HasMore    bool
	TotalCount int
}
