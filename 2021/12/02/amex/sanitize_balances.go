// SanitizeGenesisBalances sorts addresses and coin sets.
func SanitizeGenesisBalances(balances []Balance) []Balance {
	sort.Slice(balances, func(i, j int) bool {
		addr1, _ := sdk.AccAddressFromBech32(balances[i].Address)
		addr2, _ := sdk.AccAddressFromBech32(balances[j].Address)
		return bytes.Compare(addr1.Bytes(), addr2.Bytes()) < 0
	})

	for _, balance := range balances {
		balance.Coins = balance.Coins.Sort()
	}

	return balances
}