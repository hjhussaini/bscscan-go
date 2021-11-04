package service

type Transaction struct {
    BlockNumber         string  `json:"blockNumber,omitempty"`
    TimeStamp           string  `json:"timeStamp,omitempty"`
    Hash                string  `json:"hash,omitempty"`
    Nonce               string  `json:"nonce,omitempty"`
    BlockHash           string  `json:"blockHash,omitempty"`
    From                string  `json:"from,omitempty"`
    ContractAddress     string  `json:"contractAddress,omitempty"`
    To                  string  `json:"to,omitempty"`
    Value               string  `json:"value,omitempty"`
    TokenName           string  `json:"tokenName,omitempty"`
    TokenSymbol         string  `json:"tokenSymbol,omitempty"`
    TokenDecimal        string  `json:"tokenDecimal,omitempty"`
    TransactionIndex    string  `json:"transactionIndex,omitempty"`
    Gas                 string  `json:"gas,omitempty"`
    GasPrice            string  `json:"gasPrice,omitempty"`
    GasUsed             string  `json:"gasUsed,omitempty"`
    CumulativeGasUsed   string  `json:"cumulativeGasUsed,omitempty"`
    Input               string  `json:"input,omitempty"`
    Confirmations       string  `json:"confirmations,omitempty"`
}
