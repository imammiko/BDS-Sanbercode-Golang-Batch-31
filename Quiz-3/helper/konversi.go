package helper

func KonversiKetebalan(halaman int) string {
	if halaman > 201 {
		return "tebal"
	} else if halaman >= 101 && halaman <= 200 {
		return "sedang"
	} else {
		return "tipis"
	}
}
