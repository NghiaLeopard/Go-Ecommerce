package utils

func PageCount(limit int64, totalCount int64) int64 {
	if limit == 0 {
		return 0
	}

	totalPage := totalCount / limit

	if totalCount%limit > 0 {
		return totalPage + 1
	}

	return totalPage
}
