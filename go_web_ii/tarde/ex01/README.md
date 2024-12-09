# Exercício 1: Adicionando um produto
Desta vez, vamos adicionar um produto à slice carregada na memória. Dentro do caminho /products, adicionamos o método POST, para o qual enviaremos o novo produto no corpo da request. Esse método tem algumas restrições, vamos aprender sobre elas:

1. Não é necessário passar o Id, no momento de adicioná-lo, ele deve ser inferido a partir do status da lista de produtos, verificando se eles não estão repetidos, pois deve ser um campo exclusivo.


2.Nenhum dado pode estar vazio, exceto is_published (vazio indica um valor false).


3. O campo code_value deve ser exclusivo para cada produto.


4. Os tipos de dados devem corresponder aos definidos no enunciado do problema.


5. A data de validade deve ter o formato: XX/XX/XXXXXX, e também devemos verificar se o dia, o mês e o ano são valores válidos.


Lembre-se: se uma consulta for mal formulada pelo cliente, o status code será 4XX.


# Exercício 2: Trazendo o produto

Faça uma consulta a um método GET com o ID do produto recém-adicionado. Lembre-se de que a lista de produtos é carregada na memória; se você terminar a execução do programa, esse produto não estará na próxima execução.
