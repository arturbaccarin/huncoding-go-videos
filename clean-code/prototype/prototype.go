package prototype

type Usuario struct {
	Nome     string
	Idade    int
	Endereco *Endereco
}

type Endereco struct {
	Rua    string
	CEP    string
	Numero int
}

func main() {
	felipe := Usuario{
		Nome:  "Felipe",
		Idade: 20,
		Endereco: &Endereco{
			Rua:    "Rua 1",
			CEP:    "12345-678",
			Numero: 1,
		},
	}

	marcos := felipe.DeepCopy()

	/*
		// TODO: Criar um usuario Marcos a partir do objeto felipe, mudando apenas o nome
		marcos := felipe
		marcos.Nome = "Marcos"
		fmt.Println(marcos, felipe)

		// TODO: Alterar o endereco de Marcos sem alterar o endereco de Felipe
		// marcos.Endereco.Rua = "TEST TEST" //Shallow Copy, o novo objeto aponta para a referência do antigo alterando os dois

		marcos.Endereco = &Endereco{ // Deep copy - crie um objeto em um novo endereço de memória
			Rua:    felipe.Endereco.Rua,
			CEP:    felipe.Endereco.CEP,
			Numero: felipe.Endereco.Numero,
		}

		fmt.Println(marcos.Endereco, felipe.Endereco)
	*/

}

func (u *Usuario) DeepCopy() *Usuario {
	return &Usuario{
		Nome:     u.Nome,
		Endereco: u.DeepCopyAddress(),
	}
}

func (u *Usuario) DeepCopyAddress() *Endereco {
	return &Endereco{
		Rua:    u.Endereco.Rua,
		CEP:    u.Endereco.CEP,
		Numero: u.Endereco.Numero,
	}
}
