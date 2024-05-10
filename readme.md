# Car Management API 🚗

## Propósito 🎯

A **Car Management API** é uma interface de programação de aplicações projetada para facilitar a gestão de veículos em um sistema de inventário de carros. Esta API permite que os usuários visualizem, atualizem, e adicionem novos registros de carros, fornecendo uma maneira eficiente e simplificada de gerenciar dados de carros em um banco de dados.

## Tecnologias Utilizadas 🛠️

Esta API é construída utilizando várias tecnologias modernas e eficazes no desenvolvimento de software:

- **Go (Golang)**: Uma linguagem de programação estaticamente tipada conhecida por sua simplicidade e eficiência. Utilizamos Go para construir toda a lógica de backend da nossa API.
- **GORM**: Um ORM poderoso para Go, que nos ajuda a interagir de maneira mais eficiente com nosso banco de dados, facilitando operações como buscar, inserir, e atualizar registros.
- **PostgreSQL**: Um sistema de gerenciamento de banco de dados relacional, escolhido por sua robustez e conformidade com ACID, garantindo transações seguras e confiáveis.
- **Gin-Gonic/Gorilla Mux** (opcional, dependendo do router escolhido): Frameworks para criar rotas HTTP de maneira mais conveniente e com melhor desempenho em Go.

## Funcionalidades 🌟

- **Listar Carros**: Obtenha uma lista de todos os carros disponíveis no banco de dados.
- **Buscar Carro por ID**: Consulte detalhes específicos de um carro utilizando seu ID.
- **Buscar Carro por Placa**: Encontre facilmente um carro através de sua placa de licença.
- **Buscar Carro por Marca**: Visualize todos os carros disponíveis de uma determinada marca.
- **Atualizar Informações de um Carro**: Modifique dados de um carro existente, como cor, preço ou quilometragem.
- **Adicionar Novo Carro**: Insira um novo carro no sistema, com todas as informações relevantes, como marca, modelo, ano, etc.

