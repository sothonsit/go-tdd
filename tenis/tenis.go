package tenis

type TenisScore struct{}

func Tenis() TenisScore {
	return TenisScore{}
}

func (t TenisScore) AGetPoint() (int, int) {
	return 15, 0
}
