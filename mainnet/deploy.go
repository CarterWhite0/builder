const (
	badPLC        = "bad.plc"
	goodPLC       = "good.plc"
	badPLR        = "bad.plr"
	goodPLR       = "good.plr"
	moduledPLR    = "moduled.plr"
	testModulePDT = "test-module.pdt"

const goodRecipeLiteral = `
{
    "cookbookId": "cookbookLoudTest",
    #id_name LOUDGetCharacter LOUD-Get-Character-Recipe,
    "description": "Creates a basic character in LOUD",
    "version": "v0.0.1",
    #no_input,
    "entries": {
        #no_coin_or_item_modify_output,
        "itemOutputs": [
            {}

    "cookbookID": "cookbookLoudTest",
    "ID": "LOUDGetCharacter",
    "name": "LOUD-Get-Character-Recipe",
    "description": "Creates a basic character in LOUD (but don't b/c it doesn't work)",
    "version": "v0.0.1",
    "coinInputs": [],
	"beef": "edible",
    "itemInputs": [],
    "entries": {}
