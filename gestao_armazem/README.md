
### Empresa Acme, Lda

# Membros do grupo
1. Manuel Afonso
2. Novais David

## Instrução para rodar o programa 

1. Para gerar o  executavel entra no diretório/pasta da aplicação a partir da linha de comando:  
```
  go build -o ACME main.go
```
1.2 Para gerar o executavel em um diretorio/pasta diferente, a partir linha de comando:
  ```
 go build -o ~/[o_caminho_onde_pretendes_colocar_o_executavel]/[nome_do_executavel] main.go
 ex.: go build -o ~/Documents/n/teste main.go   
```
2. para executar a aplicação deve executar o seguinte comando:
```
  ./[nome_do_executavel] [funcionalidade] -i [identificador do produto] -q [quantidade do produto a expedir]
```
```
  Ex.: ./teste expedirMercadoria -i 003 -q 50
```
2.1 Ou 
```
   ./[nome_do_executavel] [funcionalidade] --idProduto [identificador do produto] --quantidade [quantidade do produto a expedir]
```
```
  ./teste expedirMercadoria --idProduto 003 --quantidade 50
```
