package main

import (
	"bytes"
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"net/http"
	"time"
)

func main() {

	bot, err := tgbotapi.NewBotAPI("6570561873:AAHC6b6mDii2plHVw8jGbHjn7d6dZVQdDEQ")

	if err != nil {
		log.Println(err)
	}
	for {
		time.Sleep(time.Second * 60)
		msg := tgbotapi.NewMessage(-4008292339, getInvoiceOrderInfoAntrpay())
		bot.Send(msg)
		msg = tgbotapi.NewMessage(-4008292339, getInvoiceOrderInfoPaylama())
		bot.Send(msg)
		msg = tgbotapi.NewMessage(-4008292339, getInvoiceOrderInfoAlfakit())
		bot.Send(msg)
	}
}

func getInvoiceOrderInfoAntrpay() string {
	start := time.Now()
	type bodyStruct struct {
		UUID string `json:"uuid"`
		Type string `json:"type"`
	}
	body := bodyStruct{
		UUID: "be9c70b2-6414-43db-81c4-46d88325be38",
		Type: "invoice",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		println("тело не отправилось: ", err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	urlAntrpay := "https://antrpay.com/api/api/repayment/fetch_order_info"
	reqAnta, _ := http.NewRequest("POST", urlAntrpay, bodyReader)
	reqAnta.Header.Add("Content-Type", "application/json")
	reqAnta.Header.Add("API-Key", "1ufTLuAU1VizXxPVz7lRtSi2pFtE3jQijH3jk5dFSYG8vt91UUCSxOSPdYA6DaRyiE7R5GP7RVDX8w8OsWMcIvRnDetHV3kJnqqtewJEww7I4Wp6WlusxDkcxprCZa3ZofIKudJlmuqoe0d3OG8klbYJpSS1rvPLUfUAOMbhwnQNmHNwTVf8OQ0xMiARx5eoqW62zA8cZCkeT9p0oqZgRrQNySxKjcnbKrONQx3FK6yUmFoHdmow9idJ2hBAOO1D")
	res, err := http.DefaultClient.Do(reqAnta)
	if err != nil {
		return "***\nСервер Antrpay не отвечает, статус 5хх \n@Glav_Control_IT @gl_control @Control_Otc"
	}

	defer res.Body.Close()

	log.Println(res)
	if res.StatusCode == 200 {

		return "***\n" + res.Status + " время ответа от сервера Antrpay - " + time.Since(start).String()
	}
	return "***\n" + res.Status + "\nAntrpay не ответил 200 @Glav_Control_IT @gl_control @Control_Otc.\nВремя ответа от сервера " + string(time.Since(start))
}

func getInvoiceOrderInfoPaylama() string {
	start := time.Now()
	type bodyStruct struct {
		ExternalID string `json:"externalID"`
		OrderType  string `json:"orderType"`
	}
	body := bodyStruct{
		ExternalID: "3ec6c19b-98ea-47f7-90ca-492f172568d7",
		OrderType:  "invoice",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		println("тело не отправилось: ", err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	urlPaylama := "https://admin.paylama.io/api/api/payment/get_order_details"
	reqLama, _ := http.NewRequest("POST", urlPaylama, bodyReader)
	reqLama.Header.Add("Content-Type", "application/json")
	reqLama.Header.Add("API-Key", "AOGPjAAJWHGs5q2VckXgJQejawYk6NFHQbxq0dZxMMfRaEvsTQIaHwW3WfbTZ2Q7HSNs1XumtjtUnBN2gLt3vs8hbjmlJtnq1wgFHfzYEyJDeAkfrOTg7zEHAQJ5nK3b2i7c98jk7ors9MKxhKLinTwG0zOWd37QhlDfw2d2zqNJoe5mRWnySrMukmdIuDw3bxnwzJUM0M9rzk8ukeZxQfNX9yvjnjotptzE7TmeacH3e0y50RJG5Menbu1n7XK9")
	res, err := http.DefaultClient.Do(reqLama)
	if err != nil {
		return "Сервер Paylama не отвечает, статус 5хх \n@Glav_Control_IT @gl_control @Control_Otc"
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		return res.Status + " время ответа от сервера Paylama - " + time.Since(start).String()
	}
	return res.Status + "\nPaylama не ответил 200 @Glav_Control_IT @gl_control @Control_Otc.\nВремя ответа от сервера " + string(time.Since(start))
}

func getInvoiceOrderInfoAlfakit() string {
	start := time.Now()
	type bodyStruct struct {
		UUID string `json:"uuid"`
		Type string `json:"type"`
	}
	body := bodyStruct{
		UUID: "eaeb98e1-4dee-43b8-a4c8-e5eeb9d53f64",
		Type: "invoice",
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		println("тело не отправилось: ", err)
	}
	bodyReader := bytes.NewReader(jsonBody)
	urlAlfakit := "https://api.alfakit.io/api/v1/fetch_order_info"
	reqKit, _ := http.NewRequest("POST", urlAlfakit, bodyReader)
	reqKit.Header.Add("Content-Type", "application/json")
	reqKit.Header.Add("API-Key", "0oIxwI9vYffpUEUp8JuFPfXRu4bTeg93tq8mfyQPD1y28JwaHT3BoxcW5HtgAfSEUgcp9LYjRHzsL2YyqNGMnhRUjTFQftTLacp80GqEW7ej6B1hR5SsAQfBs0pD6GTM67MRw0Fp0knsDXOWvXrK6wybg1I9eDTkxnfoObDRI6O3wtXKyyFc0Pawe3WzdTaawbiOm9fice7PCbQA0XkSLGgTnalhzZfwtcGRsglcWbRMfbu8KlWMQPslwA0y67fC")
	res, err := http.DefaultClient.Do(reqKit)
	if err != nil {
		return "Сервер Alfakit не отвечает, статус 5хх \n@Glav_Control_IT @gl_control @Control_Otc" + "\n***"
	}
	defer res.Body.Close()
	if res.StatusCode == 200 {
		return res.Status + " время ответа от сервера Alfakit - " + time.Since(start).String() + "\n***"
	}
	return res.Status + "\nAlfakit не ответил 200 @Glav_Control_IT @gl_control @Control_Otc.\nВремя ответа от сервера " + time.Since(start).String() + "\n***"
}
