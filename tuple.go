package fn

type Tuple[T1 any, T2 any] struct {
	V1 T1
	V2 T2
}

type Tuple3[T1 any, T2 any, T3 any] struct {
	V1 T1
	V2 T2
	V3 T3
}

type Tuple4[T1 any, T2 any, T3 any, T4 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T3
}

type Tuple5[T1 any, T2 any, T3 any, T4 any, T5 any] struct {
	V1 T1
	V2 T2
	V3 T3
	V4 T4
	V5 T3
}
