# ⏱️ CLI country timer 📌

## Por meio deste projeto você poderá verificar o horário de diversos pontos importantes do mundo, em breve você poderá também verificar o clima e as principais notícias.

---

## 📚 Índice

- [📚 Disclaimer](#-disclaimer)
- [⚙️ Projeto](#-projeto)
- [📸 Imagens](#-imagens)
- [🌐 Tecnologias](#-tecnologias-e-bibliotecas)
- [💻 Como executar](#-como-executar)

---

## 📚 disclaimer

### o projeto surgiu com um intuito muito simples, aprender novas tecnologias que pudessem fomentar meu aprendizado, contríbuir para o fomento do meu portifólio e também ajudar e dar dicas para quem está tentando começar na programação. Este é meu primeiro projeto CLI e almejo que seja bem quisto pela comunidade e possa se tornar um projeto que evolua junto ao meu aprendizado, adicionando novas features e proporcionando para as pessoas uma experiência agradável.

---

## ⚙️ Projeto

O projeto foi desenvolvido utilizando Golang como linguagem de programação, frameworks como Bubbletea e APIs como World Time API, sendo extremamente leve e rápido poderá ser executado em basicamente qualquer dispositivo com acesso ao terminal e conexão à internet.

---

## 📸 Imagens

![primeira tabela](https://private-user-images.githubusercontent.com/222334300/693985165-a1bd896e-f9e5-4de7-94b5-39233ce059b8.png?jwt=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3OTM3Mzg5NjUsIm5iZiI6MTc5MzczODY2NSwicGF0aCI6Ii8yMjIzMzQzMDAvNjkzOTg1MTY1LWExYmQ4OTZlLWY5ZTUtNGRlNy05NGI1LTM5MjMzY2UwNTliOC5wbmc_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjYxMTAzJTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI2MTEwM1QyMDQ0MjVaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT0wZmE3YjA3NTc2MjlkZmEwNzBlZjQxOGY5YWQxMDNkYWU0YWQ2YTgxOTA0M2EzODNhMDQ2YWVkYTAzN2E5ZDgyJlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCZyZXNwb25zZS1jb250ZW50LXR5cGU9aW1hZ2UlMkZwbmcifQ.A-JdrxRGVCtsIqGc0P9oYPnuC4Lids61E_ThzryEV_k)

![segunda tabela](https://private-user-images.githubusercontent.com/222334300/708339601-6634ee4a-b329-40a3-a013-14ee77454906.png?jwt=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3OTQ4Njg4NzIsIm5iZiI6MTc5NDg2ODU3MiwicGF0aCI6Ii8yMjIzMzQzMDAvNzA4MzM5NjAxLTY2MzRlZTRhLWIzMjktNDBhMy1hMDEzLTE0ZWU3NzQ1NDkwNi5wbmc_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjYxMTE2JTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI2MTExNlQyMjM2MTJaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT00YjJjMTY1ZGM1ZmI1NDk3NGM1ZmEwZTk5YWRhYjdiYWNkYzg0YjkwZDkzYWZhZDUzMTMyN2FkYzhkNjE3ZDYxJlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCZyZXNwb25zZS1jb250ZW50LXR5cGU9aW1hZ2UlMkZwbmcifQ.PO893mQ-xfFDWPGTxYKm5-w95tzuw9rOCyoQo8gZ52Q)

![terceira tabela](https://private-user-images.githubusercontent.com/222334300/708339603-4f5621ae-695f-481e-8447-3f68c053017d.png?jwt=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJnaXRodWIuY29tIiwiYXVkIjoicmF3LmdpdGh1YnVzZXJjb250ZW50LmNvbSIsImtleSI6ImtleTUiLCJleHAiOjE3OTQ4Njg4NzIsIm5iZiI6MTc5NDg2ODU3MiwicGF0aCI6Ii8yMjIzMzQzMDAvNzA4MzM5NjAzLTRmNTYyMWFlLTY5NWYtNDgxZS04NDQ3LTNmNjhjMDUzMDE3ZC5wbmc_WC1BbXotQWxnb3JpdGhtPUFXUzQtSE1BQy1TSEEyNTYmWC1BbXotQ3JlZGVudGlhbD1BS0lBVkNPRFlMU0E1M1BRSzRaQSUyRjIwMjYxMTE2JTJGdXMtZWFzdC0xJTJGczMlMkZhd3M0X3JlcXVlc3QmWC1BbXotRGF0ZT0yMDI2MTExNlQyMjM2MTJaJlgtQW16LUV4cGlyZXM9MzAwJlgtQW16LVNpZ25hdHVyZT01YjI1NTJjZWY0Y2M4MDdhNGRjNTkxNmE4MjRiN2QxZTEzYzk4MjRmMzljN2U0YWNiMzlkZmY1YzBmNWVmZGQ5JlgtQW16LVNpZ25lZEhlYWRlcnM9aG9zdCZyZXNwb25zZS1jb250ZW50LXR5cGU9aW1hZ2UlMkZwbmcifQ.KH5wv_KfpcQX84FN99GK3fFInNo67kgh9x9AOvW58IM)

---

## 🌐 Tecnologias e bibliotecas

- [💻 Golang](https://go.dev/dl/)
- [📦 Bubbletea](https://github.com/charmbracelet/bubbletea)
- [📦 World Time API](https://time.now/developer/)
- [📦 Fmt](https://pkg.go.dev/cmd/fmt)
- [📦 Time](https://pkg.go.dev/time)

---

## 💻 Como executar

Para compilar e gerar o executável do projeto:

```bash
go build -o projectCLI main.go
```

### Logo em seguida será criado um binário quevocê deve criar um atalho, assim facilitando o ciclo de abertura do prompt de comando
---