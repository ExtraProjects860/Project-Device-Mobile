# Project-Device-Mobile

Este projeto é um trabalho acadêmico focado na aplicação de conceitos de Mobile, UI/UX, API e Micro-Serviços. O objetivo principal foi desenvolver uma solução de aplicativo para facilitar o uso de credenciais de clientes de uma farmácia por meio de um crachá virtual. Além disso, promover algumas funcionalidades extras como: Visualização/Gerenciamento de produtos, lista de desejos, usuários e o próprio crachá com código de barras requerido pela parte interessada.

A problemática foi identificada pelos custos com a confecção de crachás, gerenciamento e tamanho da farmácia para poder arcar com todos esses tópicos.

A solução desenvolvida visa oferecer uma plataforma confiável, intuitiva e segura para registro, visualização e flexibilidade através de um dispositivos móvel onde quer que o usuário esteja.  O projeto alinha-se aos objetivos acadêmicos de desenvolver habilidades em resolução de problemas, aplicar tecnologia da informação e promover inovação e responsabilidade social.

## Tecnologias Utilizadas:

- **Visual Studio Code:** Utilizado como ambiente de desenvolvimento integrado (IDE) para escrever e depurar o código.

- **Figma:** Ferramenta para design de interfaces e prototipagem rápida.

- **Lucidchart:** Utilizado para criação da documentação UML.

- **Android Studio:** Utilizado para emulação de dispositivos móveis para construção do frontend do projeto.

- **PostgresSQL:** Banco de dados utilizado para alocação, administração e design de arquitetura de dados; 

## Softwares Organizacionais:

- **Trello:** Utilizado para organização e planejamento do projeto.

- **Slack:** Plataforma para comunicação em tempo real, facilitando reuniões e discussões, além de planejamento e acompanhamento do progresso.

- **Github:** Serviço para controle de versão e gerenciamento do código, proporcionando acesso colaborativo.

- **Microsoft Word:** Utilizado para criação da documentação do projeto.

## Linguagens de Programação e Frameworks:

- **Go/Gin:** Utilizado para API performance devido a ser uma linguagem rápida compilada e de fácil utilização. Além disso, o mini-framework Gin popular no uso da linguagem e possibilitando a integração de itens em módulos necessários para o projeto;
- **Python/FastApi:** Utilizada devido a sua popularidade possuindo fóruns e documentação abrangente, junto de uma gama gigante de compatibilidade com bibliotecas, sendo usada para micro-serviço de envio de email e seu framework assíncrono de alta performance FastAPI;
- **React Native, Javascript, NativeWind:** A camada de interface foi construída em React Native, utilizando JavaScript e Nativewind para estilização. Essa combinação proporcionou um desenvolvimento ágil e uma interface responsiva, mantendo consistência visual e fluidez na navegação. Além disso, o React Native permitiu a criação de um aplicativo multiplataforma, compatível com diferentes sistemas operacionais móveis;
- **Docker:** Utilizado para o ambiente de containers para evitar problemas de incompatibilidade;
- **Ngrok:** Foi utilizado para expor a API local de forma segura e temporária, permitindo a comunicação entre o aplicativo mobile e o servidor durante a fase de desenvolvimento e testes. Essa abordagem foi essencial para simular o ambiente de produção e validar a integração entre as partes do sistema;
- **Expo.go:** Aplicativo utilizado para testes diretamente em smartphones físicos para validar implementações em ambientes reais.
- **Swagger:** Interface integrada a API para testes de requisições e integrações;

## Funcionalidades:

- Exibição de interface diferenciada entre usuários de permissão Usuário e Administrador;
- Login com opção de lembrar do acesso do usuário, incluindo página para troca de senha;
- Página inicial com Crachá exibindo dados do usuário, acesso ao menu e informações da farmácia;
- Acesso ao Crachá mesmo sem acesso a internet, porém apenas logado e limitando o resto do uso do aplicativo;
- Navegação por menus facilitada por meio de menu lateral, incluindo troca de tema para escuro e claro e sendo salvo no aplicativo, troca de senha logado, produtos, usuários e lista de desejo;
- Cadastramento, Atualização, Deleção e Visualização de itens por meio de formulários simples e diretos;
- Busca e ordenação de itens através de barras de pesquisa e botões;
- Envio de email para troca de senha com token;
- Tratamento de erros, avisos e procedimentos concluídos através de modais e loadings;

## Diagrama do Banco de Dados:

Visualize abaixo o modelo atual de dados:

![Diagrama do Banco de Dados](https://github.com/ExtraProjects860/Project-Device-Mobile/raw/dev/backend/api/schemas/diagram-db-sql.png)

## UI/UX do Projeto:

#### Página de Login

![Página de Login](https://github.com/ExtraProjects860/Project-Device-Mobile/raw/dev/images/Screenshot%20from%202025-11-01%2003-22-07.png)

#### Página de Alteração Senha 1 

![Página de Alteração Senha 1](https://github.com/ExtraProjects860/Project-Device-Mobile/raw/dev/images/Screenshot%20from%202025-11-01%2003-22-14.png)

#### Página de Alteração Senha 2

![Página de Alteração Senha 2](https://github.com/ExtraProjects860/Project-Device-Mobile/raw/dev/images/Screenshot%20from%202025-11-01%2003-22-18.png)

#### Página Home

![Página Home](https://github.com/ExtraProjects860/Project-Device-Mobile/raw/dev/images/Screenshot%20from%202025-11-01%2003-22-33.png)

#### Página de Produtos

![Página de Produtos](https://github.com/ExtraProjects860/Project-Device-Mobile/raw/dev/images/Screenshot%20from%202025-11-01%2003-22-47.png)

#### Página de Lista de Desejos

![Página de Lista de Desejos](https://github.com/ExtraProjects860/Project-Device-Mobile/blob/dev/images/Screenshot%20from%202025-11-01%2003-22-56.png)

#### Página de Usuários

![Página de Usuários](https://github.com/ExtraProjects860/Project-Device-Mobile/raw/dev/images/Screenshot%20from%202025-11-01%2003-23-01.png)

#### Modal de Cadastro de Usuário

![Modal de Cadastro de Usuário](https://github.com/ExtraProjects860/Project-Device-Mobile/raw/dev/images/Screenshot%20from%202025-11-01%2003-23-13.png)

#### Modal de Cadastro de Logout

![Modal de Cadastro de Logout](https://github.com/ExtraProjects860/Project-Device-Mobile/raw/dev/images/Screenshot%20from%202025-11-01%2003-23-24.png)

## Equipe:

Este Projeto está sendo desenvolvido pela equipe:

- [Davi Nascimento](https://github.com/DaviRodrigues) - Desenvolvedor FullStack e Gerente do Projeto;
- [Evando Machado](https://github.com/Evandro206) - Desenvolvedor FullStack e UI/UX;
- [Paulo Otavio](https://github.com/PauloOtavioDuarteRocha) - Desenvolvedor Frontend, Coodernador Auxiliar do Projeto;
- [Ewerton Ribeiro](https://github.com/Notrewell) - Documentação e Testes;

## Licença
Este projeto está sob a licença [MIT License]. Consulte o arquivo [LICENSE](LICENSE) para mais detalhes.
