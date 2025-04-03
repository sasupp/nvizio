package helper

import "xtrinio.com/model"

const (
	monetaryItemType = "monetaryItemType"
)

func CommonUnit(items []model.Item) string {
	for _, item := range items {
		if item.Concept.Type == monetaryItemType {
			for _, fact := range item.Facts {
				if !fact.IsNil {
					return fact.UnitShort
				}
			}
		}
	}
	return ""
}

func CommonScale(items []model.Item) (int, string) {
	for _, item := range items {
		if item.Facts != nil {
			for _, fact := range item.Facts {
				if fact.Decimals != 0 {
					decimals := fact.Decimals
					if decimals < 0 {
						decimals *= -1
					}
					if decimals <= 4 {
						if decimals == 2 {
							return -3, "%.1f"
						}
						return -3, "%.0f"
					}
					if decimals <= 7 {
						if decimals == 5 {
							return -6, "%.1f"
						}
						return -6, "%.0f"
					}
					if decimals == 8 {
						return -9, "%.1f"
					}
					return -9, "%.0f"
				}
			}
		}
	}
	return 0, "%.0f"
}
