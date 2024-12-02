# Exercício 3 - Registro de clientes
O mesmo estudo do exercício anterior solicita uma funcionalidade para poder registrar novos dados de clientes. Os dados necessários são:

File
Name
ID
Phone number
Adress

- Atividade 1: Antes de registrar um cliente, é necessário verificar se o cliente já existe. Para fazer isso, você precisa ler os dados de uma matriz. Caso ele se repita, você precisa tratar o erro adequadamente, como vimos até agora. Esse erro deve:
1.- gerar um panic;

2.- console iniciar a mensagem: “Error: client already exists”, e continuar com a execução do programa normalmente.

- Atividade 2: Depois de tentar verificar se o cliente a ser registrado já existe, desenvolva uma função para validar se todos os dados a serem registrados para um cliente contêm um valor diferente de zero. Essa função deve retornar pelo menos dois valores. Um deles deverá ser do tipo erro, caso um valor zero seja inserido como parâmetro (lembre-se dos valores zero de cada tipo de dado, por exemplo: 0, “”, nil).

- Atividade 3: Antes de encerrar a execução, mesmo que ocorram panics, as seguintes mensagens devem ser impressas no console: “End of execution” e “Several errors were detected at runtime”. Use o defer  para atender a esse requisito..

Requisitos gerais:

- Use a recover para recuperar o valor de qualquer pânico que possa ocorrer.
- Lembre-se de realizar as validações necessárias para cada retorno que possa conter um valor de erro.
- Gere um erro, personalizando-o de acordo com sua preferência usando uma das funções Go (execute também a validação relevante para o caso de erro retornado).
