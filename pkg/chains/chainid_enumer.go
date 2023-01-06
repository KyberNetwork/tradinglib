// Code generated by "enumer -type=ChainID -linecomment -json=true -sql=true -yaml"; DO NOT EDIT.

package chains

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _ChainIDName = "ethereumoptimismcronosbscetctomognosispolygonbttcfantomkccmoonbeamkavacantoklaytnfusionarbitrumceloavaxaurora"
const _ChainIDLowerName = "ethereumoptimismcronosbscetctomognosispolygonbttcfantomkccmoonbeamkavacantoklaytnfusionarbitrumceloavaxaurora"

var _ChainIDMap = map[ChainID]string{
	1:          _ChainIDName[0:8],
	10:         _ChainIDName[8:16],
	25:         _ChainIDName[16:22],
	56:         _ChainIDName[22:25],
	61:         _ChainIDName[25:28],
	88:         _ChainIDName[28:32],
	100:        _ChainIDName[32:38],
	137:        _ChainIDName[38:45],
	199:        _ChainIDName[45:49],
	250:        _ChainIDName[49:55],
	321:        _ChainIDName[55:58],
	1284:       _ChainIDName[58:66],
	2222:       _ChainIDName[66:70],
	7700:       _ChainIDName[70:75],
	8217:       _ChainIDName[75:81],
	32659:      _ChainIDName[81:87],
	42161:      _ChainIDName[87:95],
	42220:      _ChainIDName[95:99],
	43114:      _ChainIDName[99:103],
	1313161554: _ChainIDName[103:109],
}

func (i ChainID) String() string {
	if str, ok := _ChainIDMap[i]; ok {
		return str
	}
	return fmt.Sprintf("ChainID(%d)", i)
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ChainIDNoOp() {
	var x [1]struct{}
	_ = x[Ethereum-(1)]
	_ = x[Optimism-(10)]
	_ = x[Cronos-(25)]
	_ = x[BSC-(56)]
	_ = x[ETC-(61)]
	_ = x[Tomo-(88)]
	_ = x[Gnosis-(100)]
	_ = x[Polygon-(137)]
	_ = x[BTTC-(199)]
	_ = x[Fantom-(250)]
	_ = x[KCC-(321)]
	_ = x[Moonbeam-(1284)]
	_ = x[Kava-(2222)]
	_ = x[Canto-(7700)]
	_ = x[Klaytn-(8217)]
	_ = x[Fusion-(32659)]
	_ = x[Arbitrum-(42161)]
	_ = x[Celo-(42220)]
	_ = x[Avax-(43114)]
	_ = x[Aurora-(1313161554)]
}

var _ChainIDValues = []ChainID{Ethereum, Optimism, Cronos, BSC, ETC, Tomo, Gnosis, Polygon, BTTC, Fantom, KCC, Moonbeam, Kava, Canto, Klaytn, Fusion, Arbitrum, Celo, Avax, Aurora}

var _ChainIDNameToValueMap = map[string]ChainID{
	_ChainIDName[0:8]:          Ethereum,
	_ChainIDLowerName[0:8]:     Ethereum,
	_ChainIDName[8:16]:         Optimism,
	_ChainIDLowerName[8:16]:    Optimism,
	_ChainIDName[16:22]:        Cronos,
	_ChainIDLowerName[16:22]:   Cronos,
	_ChainIDName[22:25]:        BSC,
	_ChainIDLowerName[22:25]:   BSC,
	_ChainIDName[25:28]:        ETC,
	_ChainIDLowerName[25:28]:   ETC,
	_ChainIDName[28:32]:        Tomo,
	_ChainIDLowerName[28:32]:   Tomo,
	_ChainIDName[32:38]:        Gnosis,
	_ChainIDLowerName[32:38]:   Gnosis,
	_ChainIDName[38:45]:        Polygon,
	_ChainIDLowerName[38:45]:   Polygon,
	_ChainIDName[45:49]:        BTTC,
	_ChainIDLowerName[45:49]:   BTTC,
	_ChainIDName[49:55]:        Fantom,
	_ChainIDLowerName[49:55]:   Fantom,
	_ChainIDName[55:58]:        KCC,
	_ChainIDLowerName[55:58]:   KCC,
	_ChainIDName[58:66]:        Moonbeam,
	_ChainIDLowerName[58:66]:   Moonbeam,
	_ChainIDName[66:70]:        Kava,
	_ChainIDLowerName[66:70]:   Kava,
	_ChainIDName[70:75]:        Canto,
	_ChainIDLowerName[70:75]:   Canto,
	_ChainIDName[75:81]:        Klaytn,
	_ChainIDLowerName[75:81]:   Klaytn,
	_ChainIDName[81:87]:        Fusion,
	_ChainIDLowerName[81:87]:   Fusion,
	_ChainIDName[87:95]:        Arbitrum,
	_ChainIDLowerName[87:95]:   Arbitrum,
	_ChainIDName[95:99]:        Celo,
	_ChainIDLowerName[95:99]:   Celo,
	_ChainIDName[99:103]:       Avax,
	_ChainIDLowerName[99:103]:  Avax,
	_ChainIDName[103:109]:      Aurora,
	_ChainIDLowerName[103:109]: Aurora,
}

var _ChainIDNames = []string{
	_ChainIDName[0:8],
	_ChainIDName[8:16],
	_ChainIDName[16:22],
	_ChainIDName[22:25],
	_ChainIDName[25:28],
	_ChainIDName[28:32],
	_ChainIDName[32:38],
	_ChainIDName[38:45],
	_ChainIDName[45:49],
	_ChainIDName[49:55],
	_ChainIDName[55:58],
	_ChainIDName[58:66],
	_ChainIDName[66:70],
	_ChainIDName[70:75],
	_ChainIDName[75:81],
	_ChainIDName[81:87],
	_ChainIDName[87:95],
	_ChainIDName[95:99],
	_ChainIDName[99:103],
	_ChainIDName[103:109],
}

// ChainIDString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ChainIDString(s string) (ChainID, error) {
	if val, ok := _ChainIDNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ChainIDNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ChainID values", s)
}

// ChainIDValues returns all values of the enum
func ChainIDValues() []ChainID {
	return _ChainIDValues
}

// ChainIDStrings returns a slice of all String values of the enum
func ChainIDStrings() []string {
	strs := make([]string, len(_ChainIDNames))
	copy(strs, _ChainIDNames)
	return strs
}

// IsAChainID returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ChainID) IsAChainID() bool {
	_, ok := _ChainIDMap[i]
	return ok
}

// MarshalJSON implements the json.Marshaler interface for ChainID
func (i ChainID) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ChainID
func (i *ChainID) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ChainID should be a string, got %s", data)
	}

	var err error
	*i, err = ChainIDString(s)
	return err
}

// MarshalYAML implements a YAML Marshaler for ChainID
func (i ChainID) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for ChainID
func (i *ChainID) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = ChainIDString(s)
	return err
}

func (i ChainID) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *ChainID) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of ChainID: %[1]T(%[1]v)", value)
	}

	val, err := ChainIDString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
