type Cfg struct {
	GoPrefix string
	// evaluated recursively, defaults to ["."]
	SrcDirs []string

  	VendorMultipleBuildFiles bool
	// whether to manage kubernetes' pkg/generated/openapi.
	K8sOpenAPIGen bool


  func NewDefaultGenesisState(cdc codec.JSONCodec) GenesisState {
	gs := ModuleBasics.DefaultGenesis(cdc)
	// add default SendEnabled params for paymentProcessors tokens
	bankModuleName := banktypes.ModuleName
	bankModule, ok := gs[bankModuleName]
	if ok {
		var bankGenesisState banktypes.GenesisState
		cdc.MustUnmarshalJSON(bankModule, &bankGenesisState)
		for _, token := range types.DefaultPaymentProcessorsTokensBankParams {
			bankGenesisState.Params = bankGenesisState.Params.SetSendEnabledParam(token.Denom, token.Enabled)
		}
		gs[bankModuleName] = cdc.MustMarshalJSON(&bankGenesisState)
	}
	return gs
}
}
