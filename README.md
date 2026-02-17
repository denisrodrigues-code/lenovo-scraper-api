# 💻 Lenovo Scraper API

[![Go](https://img.shields.io/badge/Go-1.26-blue?logo=go&logoColor=white)](https://golang.org/)  


API em **Golang** que realiza **web scraping** de notebooks Lenovo do site de teste [webscraper.io](https://webscraper.io/test-sites/e-commerce/static/computers/laptops), retornando informações detalhadas como preço, variações de armazenamento, avaliações, reviews, imagens e URLs.

---

## 📌 Funcionalidades

- Rastreia todos os notebooks Lenovo do site de teste  
- Coleta dados detalhados:  
  - Nome, marca, descrição  
  - Preço base e variações de armazenamento  
  - Avaliações e número de reviews  
  - Imagem e URL do produto  
- Ordena produtos do mais barato para o mais caro  
- Exposição via **API RESTful**  
- Código modular e organizado: `controller`, `scraper`, `service`, `utils`, `model`

---

## 🛠️ Tecnologias

- Golang  
- GoQuery (para parsing HTML)  
- HTTP Client nativo (`net/http`)  

---

## 🏗 Estrutura do Projeto

```
lenovo-scraper/
├── cmd/
│   └── api/
│       └── main.go
├── controller/
│   └── product_controller.go
├── scraper/
│   ├── list_scraper.go
│   └── detail_scraper.go
├── model/
│   └── product.go
├── service/
│   └── product_service.go
├── utils/
│   └── http_client.go
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Instalação

```bash
# Clonar o repositório
git clone https://github.com/denisrodrigues-code/lenovo-scraper-api.git
cd lenovo-scraper-api

# Baixar dependências
go mod tidy

# Executar API
go run main.go
A API estará disponível em: http://localhost:8080/lenovo
```
A API estará disponível em: http://localhost:8080/lenovo

## ⚙️ Endpoints

| Método | Endpoint | Descrição                                                                                   |
| ------ | -------- | ------------------------------------------------------------------------------------------- |
| GET    | /lenovo  | Retorna todos os notebooks Lenovo, ordenados do mais barato para o mais caro em JSON        |

## 📝 Exemplo de Resposta
 ```
[
  {
    "name": "Lenovo V110-15IAP",
    "brand": "Lenovo",
    "description": "Lenovo V110-15IAP, 15.6\" HD, Celeron N3350 1.1GHz, 4GB, 128GB SSD, Windows 10 Home",
    "rating": 5,
    "reviews": 12,
    "image": "https://webscraper.io/images/product1.jpg",
    "url": "https://webscraper.io/test-sites/e-commerce/static/product/1",
    "base_price": 321.94,
    "storages": [
      {
        "size": "128GB",
        "price": 321.94
      },
      {
        "size": "256GB",
        "price": 349.99
      }
    ]
  }
]
```
## 🌟 Melhorias Futuras

- Tratar erros de parsing e requisições de forma robusta
- Criar testes automatizados
- Dockerizar a Aplicação

