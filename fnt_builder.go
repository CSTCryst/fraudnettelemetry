package fraudnettelemetry

import (
	"sync"

	"github.com/CSTCryst/fraudnettelemetry/internal"
	"github.com/CSTCryst/fraudnettelemetry/internal/fnt"
)

type (
	Version interface {
		String() string
	}
	v0_1_9 struct{}
	v2_0_1 struct{}
	v2_0_4 struct{}
)

var (
	fntBuilderPool sync.Pool = sync.Pool{
		New: func() interface{} {
			return new(fntBuilder)
		},
	}
	V0_1_9 v0_1_9
	V2_0_1 v2_0_1
	V2_0_4 v2_0_4
)

func (v0_1_9) String() string {
	return "0.1.9"
}

func (v2_0_1) String() string {
	return "2.0.1"
}

func (v2_0_4) String() string {
	return "2.0.4"
}

type fntBuilder struct {
	d                    map[string]string
	dKeys                []string
	scVersion            string
	syncStatus           string
	f                    string
	s                    string
	dc                   string
	wv                   bool
	web_integration_type string
	cookie_enabled       bool
	chk                  *fnt.Chk
}

func newFNT(builder *fntBuilder) fnt.IFNT {
	if builder == nil {
		return nil
	}

	m := fnt.NewFNTBaseBuilder()

	m.SCVersion = builder.scVersion
	m.SyncStatus = builder.syncStatus
	m.F = builder.f
	m.S = builder.s
	*m.Chk = *builder.chk
	m.DC = builder.dc
	m.D = builder.d
	m.WV = builder.wv                                   // Set wv from builder
	m.WebIntegrationType = builder.web_integration_type // Set web_integration_type from builder
	m.CookieEnabled = builder.cookie_enabled            // Set cookie_enabled from builder

	return m
}

// Creates a new fntBuilder struct instance in sync.Pool
func (v0_1_9) NewFNTBuilder() *fntBuilder {
	obj := fntBuilderPool.Get().(*fntBuilder)
	obj.scVersion = "0.1.9"
	obj.chk = fnt.NewChkBuilder()
	return obj
}

// Creates a new fntBuilder struct instance in sync.Pool
func (v2_0_1) NewFNTBuilder() *fntBuilder {
	obj := fntBuilderPool.Get().(*fntBuilder)
	obj.scVersion = "2.0.1"
	obj.chk = fnt.NewChkBuilder()
	return obj
}

// Creates a new fntBuilder struct instance in sync.Pool
func (v2_0_4) NewFNTBuilder() *fntBuilder {
	obj := fntBuilderPool.Get().(*fntBuilder)
	obj.scVersion = "2.0.4"
	obj.chk = fnt.NewChkBuilder()
	return obj
}

// Check wether if version is supported or not
func validateVersion(version Version) error {
	switch version.String() {
	case "0.1.9", "2.0.1", "2.0.4":
		return nil
	default:
		return internal.UndefinedVersionError(version.String())
	}
}

// Get script version
func (m *fntBuilder) GetSCVersion() string {
	if m == nil {
		return ""
	}
	return m.scVersion
}

// Set script version
func (m *fntBuilder) SetSCVersion(version Version) (*fntBuilder, error) {
	if m == nil {
		return nil, internal.NullStructError("fntBuilder")
	}
	err := validateVersion(version)
	if err != nil {
		m.Reset(false)
		return nil, err
	}
	m.scVersion = version.String()
	return m, nil
}

// Get FraudNet Telemetry synchronization status
func (m *fntBuilder) GetSyncStatus() string {
	if m == nil {
		return ""
	}
	return m.syncStatus
}

// Set FraudNet Telemetry synchronization status
func (m *fntBuilder) SetSyncStatus() (*fntBuilder, error) {
	if m == nil {
		return nil, internal.NullStructError("fntBuilder")
	}
	m.syncStatus = "data"
	return m, nil
}

// Get fnSessionId or order token
func (m *fntBuilder) GetF() string {
	if m == nil {
		return ""
	}
	return m.f
}

// Set fnSessionId or order token
func (m *fntBuilder) SetF(fnSessionId string) *fntBuilder {
	if m == nil {
		return nil
	}
	m.f = fnSessionId
	return m
}

// Get source id
func (m *fntBuilder) GetS() string {
	if m == nil {
		return ""
	}
	return m.s
}

// Set source id
func (m *fntBuilder) SetS(source string) *fntBuilder {
	if m == nil {
		return nil
	}
	m.s = source
	return m
}

func (m *fntBuilder) GetChk() *fnt.Chk {
	if m == nil {
		return nil
	}
	return m.chk
}
func (m *fntBuilder) SetChk() (*fntBuilder, error) {
	if m == nil {
		return nil, internal.NullStructError("fntBuilder")
	}
	out, err := m.generateChk()
	if err != nil {
		m.Reset(false)
		return nil, err
	}
	return out, nil
}

func (m *fntBuilder) GetDC() string {
	if m == nil {
		return ""
	}
	return m.dc
}
func (m *fntBuilder) SetDC(userAgent string) (*fntBuilder, error) {
	if m == nil {
		return nil, internal.NullStructError("fntBuilder")
	}
	out, err := m.generateDC(userAgent)
	if err != nil {
		m.Reset(false)
		return nil, err
	}
	return out, nil
}

// Get simulated input traces data
func (m *fntBuilder) GetD() map[string]string {
	if m == nil {
		return nil
	}
	return m.d
}

// Set simulated input traces data from inputArr string input array
func (m *fntBuilder) SetD(inputArr []string, extractKeys bool) (*fntBuilder, []string, error) {
	if m == nil {
		return nil, nil, internal.NullStructError("fntBuilder")
	}
	out, err := m.generateD(inputArr)
	if err != nil {
		m.Reset(false)
		return nil, nil, err
	}
	keys := make([]string, 0, len(inputArr))
	if extractKeys {
		keys = append(keys, internal.ExtractStrMapStrKeys(m.d)...)
	}
	return out, keys, nil
}

// Set wv field
func (m *fntBuilder) SetWV(wv bool) *fntBuilder {
	if m == nil {
		return nil
	}
	m.wv = wv
	return m
}

// Set web_integration_type field
func (m *fntBuilder) SetWebIntegrationType(integrationType string) *fntBuilder {
	if m == nil {
		return nil
	}
	m.web_integration_type = integrationType
	return m
}

// Set cookie_enabled field
func (m *fntBuilder) SetCookieEnabled(enabled bool) *fntBuilder {
	if m == nil {
		return nil
	}
	m.cookie_enabled = enabled
	return m
}

func (m *fntBuilder) Generate(inputArr, sourceIdArr []string, fnSessionId, userAgent string, urlEncode bool) ([]string, error) {
	if m == nil {
		return nil, internal.NullStructError("fntBuilder")
	}

	inputArrLen, sourceIdArrLen := len(inputArr), len(sourceIdArr)
	if inputArr == nil || sourceIdArr == nil || inputArrLen == 0 || sourceIdArrLen == 0 || fnSessionId == "" || userAgent == "" {
		return nil, internal.EmptyOrInvalidInputError("Generate")
	}

	var (
		dKeys []string = make([]string, 0, inputArrLen)
		err   error    = nil
	)

	// Set additional fields here
	m.SetWV(true).
		SetWebIntegrationType("WEB_REDIRECT").
		SetCookieEnabled(false)

	if _, err = m.SetSyncStatus(); err != nil {
		return nil, err
	}

	if _, err = m.SetF(fnSessionId).SetDC(userAgent); err != nil {
		return nil, err
	}

	if _, err = m.SetChk(); err != nil {
		return nil, err
	}

	if _, dKeys, err = m.SetD(inputArr, true); err != nil {
		return nil, err
	}

	dCopy := make(map[string]string, inputArrLen)
	for i := 0; i < inputArrLen; i++ {
		dCopy[dKeys[i]] = m.d[dKeys[i]]
	}

	output := make([]string, sourceIdArrLen)
	for i, sourceIdLen := 0, sourceIdArrLen; i < sourceIdLen; i++ {
		m.SetS(sourceIdArr[i])
		clear(m.d)
		for j := 0; j <= i && j < inputArrLen; j++ {
			m.d[dKeys[j]] = dCopy[dKeys[j]]
		}
		if output[i], err = m.String(urlEncode); err != nil {
			return nil, err
		}
	}

	return output, nil
}

// Finalize the current fntBuilder instance to a new fntBase struct instance in sync.Pool behind as an interface
func (m *fntBuilder) Build() fnt.IFNT {
	if m == nil {
		return nil
	}
	out := newFNT(m)
	return out
}

func (m *fntBuilder) String(urlEncoded bool) (string, error) {
	if m == nil {
		return "", internal.NullStructError("FNTBuilder")
	}
	out := m.Build()
	defer out.Reset(true)
	str, err := out.String(urlEncoded)
	if err != nil {
		return "", err
	}
	return str, err
}

func (m *fntBuilder) Reset(put bool) {
	if m == nil {
		return
	}
	m.chk.Reset(put)
	m.d = nil
	m.dKeys = nil
	m.dc = ""
	m.f = ""
	m.s = ""
	m.syncStatus = ""
	if put {
		fntBuilderPool.Put(m)
	}
}
