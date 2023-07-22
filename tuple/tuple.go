package tuple

type T2[U1, U2 any] struct {
	First  U1
	Second U2
}

func Make2[U1, U2 any](u1 U1, u2 U2) T2[U1, U2] {
	return T2[U1, U2]{
		First:  u1,
		Second: u2,
	}
}
