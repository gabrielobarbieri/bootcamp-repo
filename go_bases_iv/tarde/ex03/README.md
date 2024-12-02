# Exercício 3 - Registro de Clientes

## Overview
Este exercício visa desenvolver uma funcionalidade para registrar novos clientes. Os dados necessários para o registro incluem:

- File
- Name
- ID
- Phone number
- Address

## Atividades

### Atividade 1: Verificação de Cliente Existente
Antes de registrar um cliente, é necessário verificar se ele já existe. Para isso, siga os passos:

- Leia os dados de uma matriz.
- Caso o cliente já exista, trate o erro adequadamente:
  - Gere um panic.
  - Exiba a mensagem no console: “Error: client already exists” e continue a execução do programa normalmente.

### Atividade 2: Validação de Dados
Desenvolver uma função que valide se todos os dados do cliente contêm valores diferentes de zero. A função deve:

- Retornar pelo menos dois valores: um deles deve ser do tipo erro, que será gerado caso algum valor zero seja inserido (considerando os valores zero pertinentes a cada tipo: 0, "", nil, etc.).

### Atividade 3: Mensagens Finais
Antes de encerrar a execução, independentemente de ocorrerem panics, imprima no console as seguintes mensagens:

- “End of execution”
- “Several errors were detected at runtime”

Utilize defer para atender a este requisito.

## Requisitos Gerais

- Use recover para capturar o valor de qualquer panic que possa ocorrer.
- Realize as validações necessárias para cada retorno que possa conter um valor de erro.
- Gere um erro personalizado conforme sua preferência, utilizando uma das funções disponíveis no Go, e valide adequadamente em caso de erro retornado.
