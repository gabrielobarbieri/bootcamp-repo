# Exercício 5 - Calcular a quantidade de alimentos

Um abrigo de animais precisa calcular a quantidade de alimentos que precisa comprar para seus animais de estimação. No momento, eles só têm tarântulas, hamsters, cães e gatos, mas esperam poder abrigar muito mais animais. Eles têm os seguintes dados:

Cão 10 kg de comida.
Gato 5 kg de comida.
 Hamster 250 g de comida.
Tarântula 150 g de comida.
É solicitado que você:

### Implemente uma função Animal que receba como parâmetro um valor de texto com o animal especificado e retorne uma função e uma mensagem (caso o animal não exista).
### Uma função para cada animal que calcule a quantidade de comida com base na quantidade do tipo de animal especificado.


Exemplo:

const (

   dog    = "dog"

   cat    = "cat"

)




animalDog, msg := animal(dog)

animalCat, msg := animal(cat)




var amount float64

amount += animalDog(10)

amount += animalCat(10)
