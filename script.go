package main

import (
	"context"
	"database/sql"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	aabi "github.com/inscription/abi"
	"math/big"
	"os"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var inscriber = make(map[string][]int64)
var fraudToken = make(map[int64]string)

//var tokenTrack = make(map[int64][]string)

const createContractBlock = 18856232
const endSearchBlock = 18889471
const step = 500
const rpcUrl = "https://eth-mainnet.g.alchemy.com/v2/<api key>"

const contract = "0x8c578a6e31fc94b1facd58202be53a8385bacbf7"
const trimLeft = "0x000000000000000000000000"

var myABI *abi.ABI
var abiCaller *aabi.Erc721gpCaller
var myDB *sql.DB
var recordBlock int64

func main() {
	db, err := sql.Open("sqlite3", "identifier.sqlite")
	if err != nil {
		panic(err)
	}
	myDB = db
	defer db.Close()

	a, err := aabi.Erc721gpMetaData.GetAbi()
	if err != nil {
		panic(err)
	}

	myABI = a

	c, e := ethclient.Dial(rpcUrl)
	if e != nil {
		panic(e)
	}
	rc, e := rpc.Dial(rpcUrl)
	if e != nil {
		panic(e)
	}

	caller, e := aabi.NewErc721gpCaller(common.HexToAddress(contract), c)
	if e != nil {
		panic(e)
	}

	abiCaller = caller

	rows, err := myDB.Query("SELECT number FROM number")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var start int64

	for rows.Next() {
		if err := rows.Scan(&start); err != nil {
			panic(err)
		}
	}

	recordBlock = start

	if _, err := myDB.Exec("DELETE FROM transfers WHERE block_number >= ?", start); err != nil {
		panic(err)
	}

	bn := uint64(endSearchBlock)
	end := int64(bn)

	for start < end {
		if end > start+step {
			end = start + step
		}
		findAB(c, rc, big.NewInt(start), big.NewInt(end))
		start = end + 1
		end = int64(bn)
	}

	for k, v := range fraudToken {
		fmt.Printf("fraud token id: %d, hash, owner: %s\n", k, v)
	}
}

func findInscriber(c *ethclient.Client, from, to *big.Int) {
	q := ethereum.FilterQuery{
		FromBlock: from,
		ToBlock:   to,
		Addresses: []common.Address{common.HexToAddress("0x8c578a6e31fc94b1facd58202be53a8385bacbf7")},
		Topics: [][]common.Hash{
			{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
			{common.HexToHash("0x0000000000000000000000000000000000000000")},
		},
	}
	logs, e := c.FilterLogs(context.Background(), q)
	if e != nil {
		panic(e)
	}

	for _, l := range logs {
		t, b := new(big.Int).SetString(l.Topics[3].Hex(), 0)
		if b {
			if _, f := inscriber[strings.TrimLeft(l.Topics[2].Hex(), trimLeft)]; !f {
				inscriber[strings.TrimLeft(l.Topics[2].Hex(), "0x000000000000000000000000")] = make([]int64, 0)
			}
			inscriber[strings.TrimLeft(l.Topics[2].Hex(), trimLeft)] = append(inscriber[strings.TrimLeft(l.Topics[2].Hex(), trimLeft)], t.Int64())
			// fmt.Printf("0x%s inscribe %d \n", strings.TrimLeft(l.Topics[2].Hex(), "0x000000000000000000000000"), t.Int64())
		}
	}

	file, err := os.Create("./inscriber.csv")
	if err != nil {
		panic(err)
	}
	cw := csv.NewWriter(file)

	for k, v := range inscriber {
		fmt.Printf("%s inscribe %v \n", k, v)
		var vv []string
		for i := range v {
			vv = append(vv, fmt.Sprintf("%d", v[i]))
		}
		if err := cw.Write([]string{k, strings.Join(vv, "&")}); err != nil {
			panic(err)
		}
	}
}

func findFraud(c *ethclient.Client, from, to *big.Int) {
	q := ethereum.FilterQuery{
		FromBlock: from,
		ToBlock:   to,
		Addresses: []common.Address{common.HexToAddress(contract)},
		Topics: [][]common.Hash{
			{common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")},
		},
	}
	logs, e := c.FilterLogs(context.Background(), q)
	if e != nil {
		panic(e)
	}

	for _, l := range logs {
		tx, _, err := c.TransactionByHash(context.Background(), l.TxHash)
		if err != nil {
			panic(err)
		}
		sender, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
		if err != nil {
			panic(err)
		}
		if strings.ToLower(l.Topics[1].Hex()) != strings.ToLower(sender.Hex()) {
			fmt.Printf("fraud block number %d, transaction hash %s \n", l.BlockNumber, l.TxHash.Hex())
		}
	}
}

func findAB(c *ethclient.Client, rc *rpc.Client, from, to *big.Int) {
	q := ethereum.FilterQuery{
		FromBlock: from,
		ToBlock:   to,
		Addresses: []common.Address{common.HexToAddress("0x8c578a6e31fc94b1facd58202be53a8385bacbf7")},
		Topics: [][]common.Hash{
			{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
		},
	}
	logs, e := c.FilterLogs(context.Background(), q)
	if e != nil {
		panic(e)
	}

	transactions := make(map[string]*types.Transaction)
	tracesStore := make(map[string][]WrapperAction)
	for _, l := range logs {
		if int64(l.BlockNumber) > recordBlock {
			if _, err := myDB.Exec("UPDATE number SET number = ?", l.BlockNumber); err != nil {
				panic(err)
			}
			recordBlock = int64(l.BlockNumber)
		}
		fa := "0x" + strings.TrimPrefix(l.Topics[1].Hex(), trimLeft)
		if fa == "0x0000000000000000000000000000000000000000" {
			continue
		}
		ta := "0x" + strings.TrimPrefix(l.Topics[2].Hex(), trimLeft)
		tokenId, ok := new(big.Int).SetString(l.Topics[3].Hex(), 0)
		if !ok {
			panic("wrong token id " + l.TxHash.Hex())
		}

		if _, err := myDB.Exec(`INSERT INTO transfers VALUES (?,?,?,?,?,?,?)`, l.TxHash.Hex(), l.BlockNumber, fa, ta, tokenId.Int64(), "", false); err != nil {
			panic(err)
		}
		fmt.Println("handle token ", tokenId.Int64(), "hash ", l.TxHash.Hex(), "block number ", l.BlockNumber)

		if _, f := fraudToken[tokenId.Int64()]; f {
			continue
		}

		time.Sleep(50 * time.Millisecond)

		if _, found := transactions[l.TxHash.Hex()]; !found {
			tx, _, err := c.TransactionByHash(context.Background(), l.TxHash)
			if err != nil {
				panic(err)
			}
			transactions[l.TxHash.Hex()] = tx
		}

		tx := transactions[l.TxHash.Hex()]

		sender, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
		if err != nil {
			panic(err)
		}

		owner := getOwner(tokenId, big.NewInt(int64(l.BlockNumber-1)))

		// 1
		if strings.ToLower(sender.Hex()) == strings.ToLower(owner) {
			continue
		}

		// 3 - trace
		if _, found := tracesStore[tx.Hash().Hex()]; !found {
			tracesStore[tx.Hash().Hex()] = getTrace(rc, tx.Hash())
		}

		// 3.1 is from exchange market
		var isFromExchangeMarket bool

		traces := tracesStore[tx.Hash().Hex()]

		for i := range traces {
			if strings.ToLower(traces[i].Action.To) != strings.ToLower(contract) {
				continue
			}
			var tid int64 = -1
			if strings.HasPrefix(traces[i].Action.Input, "0x23b872dd") {
				inputs, err := myABI.Methods["transferFrom"].Inputs.Unpack(hexutil.MustDecode("0x" + strings.TrimPrefix(traces[i].Action.Input, "0x23b872dd")))
				if err != nil {
					panic(err)
				}

				id, ok := inputs[2].(*big.Int)
				if !ok {
					panic(inputs)
				}

				tid = id.Int64()
			}
			if strings.HasPrefix(traces[i].Action.Input, "0x42842e0e") {
				inputs, err := myABI.Methods["safeTransferFrom"].Inputs.UnpackValues(hexutil.MustDecode("0x" + strings.TrimPrefix(traces[i].Action.Input, "0x42842e0e")))
				if err != nil {
					panic(err)
				}
				id, ok := inputs[2].(*big.Int)
				if !ok {
					panic(inputs)
				}
				tid = id.Int64()
			}
			if strings.HasPrefix(traces[i].Action.Input, "0xb88d4fde") {
				inputs, err := myABI.Methods["safeTransferFrom"].Inputs.UnpackValues(hexutil.MustDecode("0x" + strings.TrimPrefix(traces[i].Action.Input, "0xb88d4fde")))
				if err != nil {
					panic(err)
				}
				id, ok := inputs[2].(*big.Int)
				if !ok {
					panic(inputs)
				}
				tid = id.Int64()
			}

			if tid != tokenId.Int64() {
				continue
			}

			if strings.ToLower(traces[i].Action.From) == strings.ToLower(owner) {
				isFromExchangeMarket = true
				break
			}

			if isExchange(traces[i].Action.From) {
				isFromExchangeMarket = true
				break
			} else {
				// 3.2
				isOperator, err := abiCaller.IsApprovedForAll(&bind.CallOpts{
					BlockNumber: big.NewInt(int64(l.BlockNumber - 1)),
				}, common.HexToAddress(owner), common.HexToAddress(traces[i].Action.From))
				if err != nil {
					panic(err)
				}
				if isOperator {
					isFromExchangeMarket = true
					break
				}

				// 3.3 latest approve owner
				t := int64(l.BlockNumber)
				s := t - step
				if s < createContractBlock {
					s = createContractBlock
				}

				for {
					approveHash, ownerAddress := getLatestApprove(c, big.NewInt(s), big.NewInt(t), tokenId, l.TxHash.Hex())
					var trigger string
					if approveHash != "" {
						if _, found := tracesStore[approveHash]; !found {
							tracesStore[approveHash] = getTrace(rc, common.HexToHash(approveHash))
						}
						newTraces := tracesStore[approveHash]
						for i := range newTraces {
							if strings.ToLower(newTraces[i].Action.To) != strings.ToLower(contract) {
								continue
							}
							if strings.HasPrefix(newTraces[i].Action.Input, "0x095ea7b3") {
								inputs, err := myABI.Methods["approve"].Inputs.UnpackValues(hexutil.MustDecode("0x" + strings.TrimPrefix(newTraces[i].Action.Input, "0x095ea7b3")))
								if err != nil {
									panic(err)
								}

								id, ok := inputs[1].(*big.Int)
								if !ok {
									panic(inputs)
								}
								if id.Int64() != tokenId.Int64() {
									continue
								}

								trigger = newTraces[i].Action.From
								break
							}
						}

						if strings.ToLower(trigger) != strings.ToLower(ownerAddress) {
							if _, err := myDB.Exec("UPDATE transfers SET approveHash = ?,isFraud = ? WHERE hash = ? AND tokenId = ?", approveHash, true, l.TxHash.Hex(), tokenId.Int64()); err != nil {
								panic(err)
							}
							fmt.Printf("fraud transaction %s,block number %d, from %s to %s token %d, approve hash %s\n", l.TxHash.Hex(), l.BlockNumber, fa, ta, tokenId, approveHash)
							fraudToken[tokenId.Int64()] = tx.Hash().Hex() + "--" + owner
						}
						break
					}
					t = s - 1
					if t <= createContractBlock {
						break
					}
					s = t - step
					if s < createContractBlock {
						s = createContractBlock
					}
				}

				if s == createContractBlock {
					fmt.Printf("not found approved address transaction %s, token %d", l.TxHash.Hex(), tokenId)
				}
			}
		}
		if isFromExchangeMarket {
			continue
		}
	}

}

var exchanges = []string{"0x20f780a973856b93f63670377900c1d2a50a77c4", "0x0000000000e655fae4d56241588680f86e3b2377", "0x00000000000000adc04c56bf30ac9d3c0aaf14dc", "0xb2ecfe4e4d61f8790bbb9de2d1259b9e2410cea5"}

func isExchange(address string) bool {
	for i := range exchanges {
		if strings.ToLower(exchanges[i]) == strings.ToLower(address) {
			return true
		}
	}
	return false
}

type WrapperAction struct {
	Action Action `json:"action"`
}

type Action struct {
	From  string `json:"from"`
	Input string `json:"input"`
	To    string `json:"to"`
}

var t = time.NewTicker(300 * time.Millisecond)

func getTrace(c *rpc.Client, hash common.Hash) []WrapperAction {

	<-t.C

	var result interface{}

	for {
		err := c.CallContext(context.Background(), &result, "trace_transaction", hash.Hex())
		if err != nil {
			fmt.Println("Error ", err.Error())
			time.Sleep(time.Second)
		}
		break
	}

	data1, err1 := json.Marshal(result)
	if err1 != nil {
		panic(err1)
	}

	actions := make([]WrapperAction, 0)
	if err := json.Unmarshal(data1, &actions); err != nil {
		panic(err)
	}

	return actions
}

func getLatestApprove(c *ethclient.Client, start, end, tokenId *big.Int, transferHash string) (string, string) {
	fmt.Println("getLatestApprove", tokenId, start, end)
	q := ethereum.FilterQuery{
		FromBlock: start,
		ToBlock:   end,
		Addresses: []common.Address{common.HexToAddress(contract)},
		Topics: [][]common.Hash{
			{common.HexToHash("0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925")},
		},
	}
	logs, e := c.FilterLogs(context.Background(), q)
	if e != nil {
		panic(e)
	}

	for i := len(logs) - 1; i >= 0; i-- {
		if strings.ToLower(logs[i].TxHash.Hex()) == strings.ToLower(transferHash) && logs[i].Topics[2].Hex() == "0x0000000000000000000000000000000000000000000000000000000000000000" {
			continue
		}
		tid, ok := new(big.Int).SetString(logs[i].Topics[3].Hex(), 0)
		if !ok {
			panic(logs[i].TxHash.Hex())
		}
		if tokenId.Int64() == tid.Int64() {
			return logs[i].TxHash.Hex(), "0x" + strings.TrimPrefix(logs[i].Topics[1].Hex(), trimLeft)
		}
	}

	return "", ""
}

func getOwner(tokenID, bn *big.Int) string {
	owner, err := abiCaller.OwnerOf(&bind.CallOpts{BlockNumber: bn}, tokenID)
	if err != nil {
		fmt.Println(tokenID.Uint64())
		panic(err)
	}
	return owner.Hex()
}

func findLastApprove(c *ethclient.Client, rc *rpc.Client, bn *big.Int) {

}
