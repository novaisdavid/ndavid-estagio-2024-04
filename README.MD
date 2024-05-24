Projeto firmação (Execução do software)

1- Acesse o diretorio main ( que está dentro do diretorio firmacao)

2- No diretorio main, via terminal execute o comando: " go build . "

3- Ainda no terminal para matricular um formando execute o seguinte comando: " ./main  matricular --nome <nome do formando> --email <email do formando> --telefone <número de telefone do formando> --curso <curso para matricula> --regime <online ou presencial> "
Ex: ./main  matricular --nome Tito --email tito@gmail.com --telefone 923456789 --curso cc --regime online
    OBS.:  Se a último curso feito pelo formando estiver dentro do prazo de 15 dias após a última a formação, a sua nova matricula receberá um desconto de 25%.

4- Para iniciar um curso, deverá executar o seguinte comando: " ./main Iniciarcurso --nomeC <nome do curso> "
Ex: ./main Iniciarcurso --nomeC Ingles
  OBS: Uma informação de lembrete é enviada dois dias antes do inicio do curso aos formandos.

5- Para concluir um curso, deverá executar o seguinte comando: " ./main Concluircurso --nomeC <nome do curso> "
Ex: ./main Concluircurso --nomeC Ingles
  OBS: Na conclusão de um curso, é enviada uma informação de certificados emitidos.

6- Para avaliar um curso, deverá executar o seguinte comando:  " ./main Avaliarcurso --idcurso <ID do curso> --idFormando <ID do formando> --avaliacao <nota da avaliacao de 0 à 100> "
Ex:  ./main Avaliarcurso --idcurso 57 --idFormando 90 --avaliacao 50
  OBS: Caso a avaliação receber um valor abaixo dos 60, o formador receberá uma penalização de 2% a cada 1% abaixo dos 60.
Informações Adicionais:  
  -Depois de executar o comando para efectuação da matricula, os dados serão armazenados em um ficheiro txt com nome "MatriculaFormando.txt", no diretorio "matricula".
  -Depois de executar o comando para avaliação do curso, os dados da avaliação serão armazenados em um ficheiro txt com nome "CursosAvaliados.txt", no diretorio "matricula".
