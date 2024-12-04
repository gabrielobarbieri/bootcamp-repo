package main

import (
  "bufio"
  "fmt"
  "os"

  "github.com/google/uuid"
)

type Aluno struct {
  matricula string
  nome string
  telefone string
  email string
} 

type Alunos struct {
  listaAlunos map[string]Aluno 
} 

func (a *Alunos) Add(uuid string, novoAluno Aluno) {
  a.listaAlunos[uuid] = novoAluno
} 

func (a *Alunos) Search(matricula string) *Aluno {
  aluno, exists := a.listaAlunos[matricula]
  
  if !exists {
    fmt.Println("Aluno não encontrado!")
    return nil
  }

  return &aluno
}

func (a *Alunos) UpdateAluno(aluno *Aluno, nome string, telefone string, email string) {
    aluno.nome = nome 
    aluno.telefone = telefone 
    aluno.email = email

    a.listaAlunos[aluno.matricula] = *aluno
    fmt.Println("Aluno alterado com sucesso!")
}

func (a *Alunos) DeleteAluno(matricula string) {
  delete(a.listaAlunos, matricula)  
}

func main() {
  reader := bufio.NewReader(os.Stdin)
  alunos := Alunos{listaAlunos: make(map[string]Aluno)}

  fmt.Println("=== Menu de opcoes ===")
  for {
    fmt.Println("Selecione uma opcao a seguir:") 
    fmt.Println("1- Incluir aluno") 
    fmt.Println("2- Listar todos os alunos") 
    fmt.Println("3- Pesquisar aluno por matrícula") 
    fmt.Println("4- Alterar aluno") 
    fmt.Println("5- Excluir aluno") 
    fmt.Println("6- Sair") 

    input, err := reader.ReadString('\n')
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error reading input:", err)
      return
    }

    input = input[:len(input)-1]

    switch input {
      case "1":
        novoAluno := &Aluno{}
        
        novoAluno.matricula = uuid.New().String()

        fmt.Println("Digite o nome:") 
        nome, _ := reader.ReadString('\n')
        novoAluno.nome = nome[:len(nome) - 1]
      
        fmt.Println("Digite o telefone:") 
        telefone, _ := reader.ReadString('\n')
        novoAluno.telefone = telefone[:len(telefone) - 1]
      
        fmt.Println("Digite o email:") 
        email, _ := reader.ReadString('\n')
        novoAluno.email = email[:len(email) - 1]

        alunos.Add(novoAluno.matricula, *novoAluno)
        fmt.Println("Aluno incluído com sucesso!")

      case "2":
        fmt.Println("Lista de alunos: ")
        if len(alunos.listaAlunos) == 0 {
          fmt.Println("Nenhum aluno cadastrado.")
        } else {
          for matricula, aluno := range alunos.listaAlunos {
            fmt.Printf("Matrícula: %s, Nome: %s, Telefone: %s, Email: %s\n", matricula, aluno.nome, aluno.telefone, aluno.email)
          }
        }

      case "3":
        fmt.Println("Digite a matricula: ")
        matricula, _ := reader.ReadString('\n')
        matricula = matricula[:len(matricula) - 1]

        aluno := alunos.Search(matricula) 
        
        fmt.Printf("Matrícula: %s, Nome: %s, Telefone: %s, Email: %s\n", matricula, aluno.nome, aluno.telefone, aluno.email)
      case "4": 
        fmt.Println("Digite a matricula: ")
        matricula, _ := reader.ReadString('\n')
        matricula = matricula[:len(matricula) - 1]

        aluno := alunos.Search(matricula)        
    
        if aluno != nil {
          fmt.Println("Digite o novo nome: ")
          nome, _ := reader.ReadString('\n')
          nome = nome[:len(nome) - 1]
          
          fmt.Println("Digite o novo telefone: ")
          telefone, _ := reader.ReadString('\n')
          telefone = telefone[:len(telefone) - 1]

          fmt.Println("Digite o novo email: ")
          email, _ := reader.ReadString('\n')
          email = email[:len(email) - 1]

          alunos.UpdateAluno(aluno, nome, telefone, email) 
        } else {
          fmt.Println("Aluno não encontrado!")
        }
        


      case "5":
        fmt.Println("Digite a matrícula: ")
        matricula, _ := reader.ReadString('\n')
        matricula = matricula[:len(matricula) - 1]
  
        alunos.DeleteAluno(matricula)       
        fmt.Println("Aluno deletado com sucesso!")

      case "6":
        fmt.Println("saindo...")
        os.Exit(0)
        
      default:
        fmt.Println("Opção inválida. Tente novamente!")
    }
  }
}
