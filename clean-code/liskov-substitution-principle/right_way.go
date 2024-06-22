package liskovsubstitutionprinciple

type Area interface {
	GetArea() float64
}

type RetanguloType interface {
	Area

	GetWidth() float64
	GetHeight() float64

	SetWidth(width float64)
	SetHeight(height float64)
}

type QuadradoType interface {
	Area

	GetSize() float64
	SetSize(size float64)
}

type Retangulo2 struct {
	width, height float64
}

type Quadrado2 struct {
	size float64
}
