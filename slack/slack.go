package slack

import (
    "github.com/ahmedmansourxyz/EpicGamesWebScraper/scraper"
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
)
type SlackRequestBody struct {
    Text string `json:"text"`
}

func SendSlackMessage(webhookUrl string, games []scraper.Game) error {

    for _, game := range games {
        Text := "Title: " + game.Title + "\nPrice: " + game.Price + "\nCurrent: " + game.Status + "\nStart: " + game.Start + "\nEnd: " + game.End

        slackBody, _ := json.Marshal(SlackRequestBody{Text: Text})
        req, err := http.NewRequest(http.MethodPost, webhookUrl, bytes.NewBuffer(slackBody))
        if err != nil {
            return err
        }
        req.Header.Add("Content-Type", "application/json")

        client := &http.Client{}
        resp, err := client.Do(req)
        if err != nil {
            return err
        }

        buf := new(bytes.Buffer)
        buf.ReadFrom(resp.Body)
        if buf.String() != "ok" {
            return errors.New("Non-ok response returned from Slack")
        }
    }
    return nil
}
