package service

import (
    "errors"
    "encoding/json"
    "io/ioutil"
    "net/http"
)

const (
    prefix          string = "https://api.bscscan.com/api?"
    apiKey          string = "&apikey="
    module          string = "module="
    account         string = "account"
    action          string = "&action="
    tokenTX         string = "tokentx"
    contractAddress string ="&address="
    sort            string = "&sort=asc"
    page            string = "&page="
    offset          string = "&offset="
)

type apiResponse struct {
        Status  string      `json:"status"`
        Message string      `json:"message"`
        Result  interface{}  `json:"result"`
}

func RequestContract(key string, address string) ([]Transaction, error) {
    var transactions   [] Transaction

    url := prefix +
        module +
        account +
        action +
        tokenTX +
        contractAddress +
        address +
        sort +
        apiKey +
        key + "&page=1&offset=100"
    response, err := http.Get(url)
    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    data, err := ioutil.ReadAll(response.Body)
    if err != nil {
        return nil, err
    }

    apiRsp := apiResponse{}
    err = json.Unmarshal(data, &apiRsp)
    if err != nil {
        return nil, err
    }

    if apiRsp.Status == "0" {
        return nil, errors.New(apiRsp.Message)
    }

    for _, element := range apiRsp.Result.([]interface{}) {
        txn := element.(map[string]interface{})
        transaction := Transaction{
            From:       txn["from"].(string),
            To:         txn["to"].(string),
            Gas:        txn["gas"].(string),
            GasPrice:   txn["gasPrice"].(string),
        }
        transactions = append(transactions, transaction)
    }

    return transactions, nil
}
