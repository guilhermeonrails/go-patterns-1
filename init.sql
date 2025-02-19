-- Cria a tabela "itens"
CREATE TABLE IF NOT EXISTS itens (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    codigo VARCHAR(50) NOT NULL UNIQUE,
    descricao VARCHAR(255),
    preco DECIMAL(10,2) NOT NULL,
    quantidade INTEGER NOT NULL
);

-- Insere 50 itens com dados exemplares
INSERT INTO itens (nome, codigo, descricao, preco, quantidade) VALUES
('Teclado Mecânico', 'TEC001', 'Teclado mecânico com retroiluminação', 150.00, 20),
('Mouse Óptico', 'MOU002', 'Mouse óptico sem fio', 80.00, 50),
('Monitor LED 24"', 'MON003', 'Monitor LED Full HD de 24 polegadas', 700.00, 15),
('Impressora Laser', 'IMP004', 'Impressora a laser com duplex', 1200.00, 8),
('Notebook Ultra Fino', 'NOT005', 'Notebook ultrafino com 8GB RAM', 3500.00, 10),
('Smartphone Android', 'SMA006', 'Smartphone Android com 64GB de armazenamento', 1200.00, 25),
('Tablet 10"', 'TAB007', 'Tablet de 10 polegadas', 900.00, 18),
('Câmera Digital', 'CAM008', 'Câmera digital com sensor de 20MP', 850.00, 12),
('Headset Gamer', 'HEA009', 'Headset com som surround', 300.00, 30),
('Caixa de Som Bluetooth', 'SOM010', 'Caixa de som portátil Bluetooth', 250.00, 40),
('Roteador Wi-Fi', 'ROT011', 'Roteador com alta velocidade e cobertura', 350.00, 22),
('Pen Drive 64GB', 'PEN012', 'Pen drive USB 3.0 de 64GB', 120.00, 100),
('HD Externo 1TB', 'HDX013', 'Disco rígido externo de 1TB', 400.00, 16),
('SSD 500GB', 'SSD014', 'SSD de 500GB para desktops', 600.00, 14),
('Placa de Vídeo GTX 1660', 'GPU015', 'Placa de vídeo GTX 1660 para jogos', 2500.00, 7),
('Processador Intel i5', 'CPU016', 'Processador Intel i5 de 10ª geração', 900.00, 10),
('Memória RAM 8GB', 'RAM017', 'Módulo de memória RAM 8GB DDR4', 300.00, 40),
('Fonte de Alimentação 500W', 'FON018', 'Fonte de alimentação 500W com certificação 80 Plus', 250.00, 20),
('Gabinete ATX', 'GAB019', 'Gabinete ATX com ventiladores inclusos', 350.00, 15),
('Cooler para CPU', 'COL020', 'Cooler para processador com LED', 120.00, 30),
('Monitor Curvo 27"', 'MON021', 'Monitor curvo de 27 polegadas', 1300.00, 10),
('Teclado sem Fio', 'TEC022', 'Teclado sem fio compacto', 180.00, 25),
('Mouse Gamer', 'MOU023', 'Mouse gamer com alta precisão', 150.00, 35),
('Cadeira Gamer', 'CHA024', 'Cadeira gamer ergonômica', 800.00, 5),
('Mesa para Computador', 'MES025', 'Mesa ampla para setup gamer', 600.00, 8),
('Monitor LED 21"', 'MON026', 'Monitor LED de 21 polegadas', 500.00, 12),
('Tablet Android', 'TAB027', 'Tablet Android com 32GB', 750.00, 20),
('Notebook Gamer', 'NOT028', 'Notebook gamer com placa dedicada', 5000.00, 6),
('Smartwatch', 'SMA029', 'Smartwatch com monitoramento de saúde', 450.00, 30),
('Câmera de Segurança', 'CAM030', 'Câmera de segurança com resolução 1080p', 350.00, 18),
('Projetor Portátil', 'PRO031', 'Projetor portátil com alta luminosidade', 1500.00, 4),
('Microfone Condensador', 'MIC032', 'Microfone condensador para estúdio', 400.00, 15),
('Lâmpada LED', 'LAM033', 'Lâmpada LED de alta eficiência', 50.00, 100),
('Switch Gerenciável', 'SWI034', 'Switch gerenciável de 24 portas', 800.00, 7),
('Roteador Mesh', 'ROT035', 'Sistema de roteador mesh para cobertura total', 1200.00, 9),
('Impressora Multifuncional', 'IMP036', 'Impressora multifuncional com scanner', 1100.00, 10),
('Scanner de Documentos', 'SCA037', 'Scanner de alta resolução', 650.00, 8),
('Cabo HDMI 2m', 'CAB038', 'Cabo HDMI 2 metros de alta velocidade', 30.00, 150),
('Cabo USB-C', 'CAB039', 'Cabo USB-C para carregamento rápido', 25.00, 200),
('Carregador Portátil', 'CAR040', 'Carregador portátil de 10000mAh', 150.00, 60),
('Fone de Ouvido In-Ear', 'FON041', 'Fone de ouvido in-ear com cancelamento de ruído', 120.00, 40),
('Fone de Ouvido Over-Ear', 'FON042', 'Fone de ouvido over-ear com alta fidelidade', 300.00, 35),
('Estabilizador de Voltagem', 'EST043', 'Estabilizador para proteger equipamentos eletrônicos', 220.00, 25),
('No-break 600VA', 'NOB044', 'No-break de 600VA para proteção contra quedas de energia', 400.00, 10),
('Câmera Action', 'CAM045', 'Câmera de ação resistente a impactos', 800.00, 15),
('Drone Fotográfico', 'DRO046', 'Drone com câmera 4K', 3500.00, 5),
('Leitor de Cartões', 'LEI047', 'Leitor de cartões multi-formato', 200.00, 30),
('Teclado Gamer RGB', 'TEC048', 'Teclado gamer com iluminação RGB', 250.00, 20),
('Mouse Sem Fio', 'MOU049', 'Mouse sem fio ergonômico', 130.00, 40),
('Hub USB 4 Portas', 'HUB050', 'Hub USB com 4 portas e suporte para USB 3.0', 100.00, 50);

--------------------------------------------------

-- Cria a tabela "categorias"
CREATE TABLE IF NOT EXISTS categoria (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL,
    codigo VARCHAR(50) NOT NULL UNIQUE,
    descricao VARCHAR(300)
);

-- Insere 5 categorias de produtos
INSERT INTO categoria (nome, codigo, descricao) VALUES
('Eletrônicos', 'ELEC', 'Produtos eletrônicos em geral.'),
('Periféricos', 'PERI', 'Acessórios e periféricos para computadores.'),
('Informática', 'INFO', 'Componentes e equipamentos de informática.'),
('Acessórios', 'ACCS', 'Diversos acessórios para dispositivos.'), 
('Eletrodomésticos', 'ELET', 'Eletrodomésticos para uso residencial.');
