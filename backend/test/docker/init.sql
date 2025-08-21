-- Criação tabela de testes
CREATE TABLE tb_test (
    id SERIAL PRIMARY KEY,
    word VARCHAR(10) NOT NULL,
    num NUMERIC(10, 2) NOT NULL,
    created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Inserção de dados de exemplo
INSERT INTO tb_test (word, num) VALUES
    ('teste_01', 123),
    ('teste_02', 456),
    ('teste_03', 0.1);