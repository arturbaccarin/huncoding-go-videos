package liskovsubstitutionprinciple

type Quadrilatero interface {
	GetArea() float64

	GetWidth() float64
	GetHeight() float64

	SetWidth(width float64)
	SetHeight(height float64)
}

type Quadrado struct {
	width, height float64
}

type Retangulo struct {
	width, height float64
}

func (r *Retangulo) GetArea() float64 {
	return r.width * r.height
}

func (r *Retangulo) GetWidth() float64 {
	return r.width
}

func (r *Retangulo) GetHeight() float64 {
	return r.height
}

func (r *Retangulo) SetWidth(width float64) {
	r.width = width
}

func (r *Retangulo) SetHeight(height float64) {
	r.height = height
}

func (r *Quadrado) GetArea() float64 {
	// devemos usar r.height ou r.width aqui?
	return r.width * r.height
}

func (r *Quadrado) GetWidth() float64 {
	return r.width
}

func (r *Quadrado) GetHeight() float64 {
	return r.height
}

func (r *Quadrado) SetWidth(width float64) {
	r.height = width
	r.width = width
}

func (r *Quadrado) SetHeight(height float64) {
	panic("para trocar a altura do quadrado, chame o SetWidth")
}
