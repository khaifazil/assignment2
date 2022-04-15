package main

func convertDate(month string, date int) (int, error) {
	// make array of months in accumalated days not including current month's days
	// add date given by user.
	var (
		arrOfMonths  = [12]string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
		daysInMonths = [12]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	)

	for i, m := range arrOfMonths {
		if month == m {
			return daysInMonths[i] + date, nil
		}
	}

	return 0, nil
}
