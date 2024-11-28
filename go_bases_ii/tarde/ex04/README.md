# Exercício 4 - Cálculo de estatísticas

Os professores de uma universidade na Colômbia precisam calcular algumas estatísticas de notas para os alunos de um curso. Para isso, eles precisam gerar uma função que indique o tipo de cálculo que desejam realizar (mínimo, máximo ou médio) e que retorne outra função e uma mensagem (caso o cálculo não esteja definido) que possa receber um número N de inteiros e retorne o cálculo indicado na função anterior. Exemplo:

const (

   minimum = "minimum"

   average = "average"

   maximum = "maximum"

)


...


minFunc, err := operation(minimum)

averageFunc, err := operation(average)

maxFunc, err := operation(maximum)

 

...


minValue := minFunc(2, 3, 3, 4, 10, 2, 4, 5)

averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)

maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)
