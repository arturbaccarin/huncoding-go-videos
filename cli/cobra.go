package main

import (
	"fmt"
	"regexp"

	"github.com/spf13/cobra"
)

func cobraMain() {
	var rootCmd = &cobra.Command{Use: "huncodingCli"}

	var nome, email, senha string

	var cmd = &cobra.Command{
		Use:   "create",
		Short: "Cria um novo usuário",
		Run: func(cmd *cobra.Command, args []string) {
			if nome == "" {
				fmt.Println("Nome não pode estar vazio")
				return
			}

			emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
			if !emailRegex.MatchString(email) {
				fmt.Println("email invalido. por favor, insira um email valido")
				return
			}

			if len(senha) < 6 {
				fmt.Println("a senha deve ter pelo menos 6 caracteres")
				return
			}

			fmt.Printf("nome: %s\nemail: %s\nsenha: %s\n", nome, email, senha)
		},
	}

	cmd.Flags().StringVarP(&nome, "nome", "n", "", "Nome do usuário")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Email do usuário")
	cmd.Flags().StringVarP(&senha, "senha", "s", "", "Senha do usuário")

	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}
