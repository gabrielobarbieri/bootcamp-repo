# Exercício 2 - Imprimir dados


Em seguida, vamos criar um arquivo "customers.txt" com informações sobre os clientes do estúdio.

Agora que o arquivo existe, o pânico não deve ser acionado.

Criamos o arquivo "customers.txt" e adicionamos as informações do cliente.
Estendemos o código do ponto um para que possamos ler esse arquivo e imprimir os dados que ele contém. Caso não seja possível ler o arquivo, um "panic" deve ser iniciado.
Lembre-se de que sempre que a execução terminar, independentemente do resultado, devemos ter um "defer" para indicar que a execução foi concluída. Lembre-se também de fechar os arquivos ao final de seu uso.