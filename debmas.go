package main

// Customer master data distribution
type DEBMASDEBMAS07 struct {
	BEGIN    string                `xml:"BEGIN,attr" json:"-BEGIN" yaml:"BEGIN" doc:""`
	EDI_DC40 *EDI_DCDEBMASDEBMAS07 `xml:"EDI_DC40" json:"EDI_DC40" yaml:"EDI_DC40" doc:""`
	E1KNA1M  *DEBMASE1KNA1M        `xml:"E1KNA1M" json:"E1KNA1M" yaml:"E1KNA1M" doc:"Master customer master basic data (KNA1)"`
}

// DEBMASDEBMAS071 was auto-generated from WSDL.
type DEBMASDEBMAS071 struct {
	Idoc *DEBMAS07 `xml:"idoc" json:"idoc" yaml:"idoc" doc:""`
}

// DEBMAS07 was auto-generated from WSDL.
type DEBMAS07 struct {
	IDOC *DEBMASDEBMAS07 `xml:"IDOC" json:"IDOC" yaml:"IDOC" doc:"Customer master data distribution"`
}

// Customer Master: Additional General Fields  (KNA1)
type DEBMASE1KNA11 struct {
	SEGMENT    string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	KNURL      *string `xml:"KNURL,omitempty" json:"KNURL,omitempty" yaml:"KNURL,omitempty" doc:"Uniform Resource Locator"`
	J_1KFREPRE *string `xml:"J_1KFREPRE,omitempty" json:"J_1KFREPRE,omitempty" yaml:"J_1KFREPRE,omitempty" doc:"Name of Representative"`
	J_1KFTBUS  *string `xml:"J_1KFTBUS,omitempty" json:"J_1KFTBUS,omitempty" yaml:"J_1KFTBUS,omitempty" doc:"Type of Business"`
	J_1KFTIND  *string `xml:"J_1KFTIND,omitempty" json:"J_1KFTIND,omitempty" yaml:"J_1KFTIND,omitempty" doc:"Type of Industry"`
	PSOIS      *string `xml:"PSOIS,omitempty" json:"PSOIS,omitempty" yaml:"PSOIS,omitempty" doc:"Subledger acct preprocessing procedure"`
	PSON1      *string `xml:"PSON1,omitempty" json:"PSON1,omitempty" yaml:"PSON1,omitempty" doc:"Name 1"`
	PSON2      *string `xml:"PSON2,omitempty" json:"PSON2,omitempty" yaml:"PSON2,omitempty" doc:"Name 1"`
	PSON3      *string `xml:"PSON3,omitempty" json:"PSON3,omitempty" yaml:"PSON3,omitempty" doc:"Name 1"`
	PSOVN      *string `xml:"PSOVN,omitempty" json:"PSOVN,omitempty" yaml:"PSOVN,omitempty" doc:"First Name"`
	PSOTL      *string `xml:"PSOTL,omitempty" json:"PSOTL,omitempty" yaml:"PSOTL,omitempty" doc:"Title"`
	PSOO1      *string `xml:"PSOO1,omitempty" json:"PSOO1,omitempty" yaml:"PSOO1,omitempty" doc:"Description"`
	PSOO2      *string `xml:"PSOO2,omitempty" json:"PSOO2,omitempty" yaml:"PSOO2,omitempty" doc:"Description"`
	PSOO3      *string `xml:"PSOO3,omitempty" json:"PSOO3,omitempty" yaml:"PSOO3,omitempty" doc:"Description"`
	PSOO4      *string `xml:"PSOO4,omitempty" json:"PSOO4,omitempty" yaml:"PSOO4,omitempty" doc:"Description"`
	PSOO5      *string `xml:"PSOO5,omitempty" json:"PSOO5,omitempty" yaml:"PSOO5,omitempty" doc:"Description"`
	STCD5      *string `xml:"STCD5,omitempty" json:"STCD5,omitempty" yaml:"STCD5,omitempty" doc:"Tax Number 5"`
}

// Master customer master basic data: Texts, header
type DEBMASE1KNA1H struct {
	SEGMENT    string           `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN      *string          `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	TDOBJECT   *string          `xml:"TDOBJECT,omitempty" json:"TDOBJECT,omitempty" yaml:"TDOBJECT,omitempty" doc:"Texts: Application Object"`
	TDNAME     *string          `xml:"TDNAME,omitempty" json:"TDNAME,omitempty" yaml:"TDNAME,omitempty" doc:"Name"`
	TDID       *string          `xml:"TDID,omitempty" json:"TDID,omitempty" yaml:"TDID,omitempty" doc:"Text ID"`
	TDSPRAS    *string          `xml:"TDSPRAS,omitempty" json:"TDSPRAS,omitempty" yaml:"TDSPRAS,omitempty" doc:"Language Key"`
	TDTEXTTYPE *string          `xml:"TDTEXTTYPE,omitempty" json:"TDTEXTTYPE,omitempty" yaml:"TDTEXTTYPE,omitempty" doc:"SAPscript: Format of Text"`
	TDSPRASISO *string          `xml:"TDSPRASISO,omitempty" json:"TDSPRASISO,omitempty" yaml:"TDSPRASISO,omitempty" doc:"Language according to ISO 639"`
	E1KNA1L    []*DEBMASE1KNA1L `xml:"E1KNA1L,omitempty" json:"E1KNA1L,omitempty" yaml:"E1KNA1L,omitempty" doc:""`
}

// Master customer master basic data: Text line
type DEBMASE1KNA1L struct {
	SEGMENT  string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN    *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	TDFORMAT *string `xml:"TDFORMAT,omitempty" json:"TDFORMAT,omitempty" yaml:"TDFORMAT,omitempty" doc:"Tag column"`
	TDLINE   *string `xml:"TDLINE,omitempty" json:"TDLINE,omitempty" yaml:"TDLINE,omitempty" doc:"Text Line"`
}

// Master customer master basic data (KNA1)
type DEBMASE1KNA1M struct {
	SEGMENT   string           `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN     *string          `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	KUNNR     *string          `xml:"KUNNR,omitempty" json:"KUNNR,omitempty" yaml:"KUNNR,omitempty" doc:"Customer Number"`
	ANRED     *string          `xml:"ANRED,omitempty" json:"ANRED,omitempty" yaml:"ANRED,omitempty" doc:"Title"`
	AUFSD     *string          `xml:"AUFSD,omitempty" json:"AUFSD,omitempty" yaml:"AUFSD,omitempty" doc:"Central order block for customer"`
	BAHNE     *string          `xml:"BAHNE,omitempty" json:"BAHNE,omitempty" yaml:"BAHNE,omitempty" doc:"Express train station"`
	BAHNS     *string          `xml:"BAHNS,omitempty" json:"BAHNS,omitempty" yaml:"BAHNS,omitempty" doc:"Train station"`
	BBBNR     *string          `xml:"BBBNR,omitempty" json:"BBBNR,omitempty" yaml:"BBBNR,omitempty" doc:"International location number  (part 1)"`
	BBSNR     *string          `xml:"BBSNR,omitempty" json:"BBSNR,omitempty" yaml:"BBSNR,omitempty" doc:"International location number (Part 2)"`
	BEGRU     *string          `xml:"BEGRU,omitempty" json:"BEGRU,omitempty" yaml:"BEGRU,omitempty" doc:"Authorization Group"`
	BRSCH     *string          `xml:"BRSCH,omitempty" json:"BRSCH,omitempty" yaml:"BRSCH,omitempty" doc:"Industry key"`
	BUBKZ     *string          `xml:"BUBKZ,omitempty" json:"BUBKZ,omitempty" yaml:"BUBKZ,omitempty" doc:"Check digit for the international location number"`
	DATLT     *string          `xml:"DATLT,omitempty" json:"DATLT,omitempty" yaml:"DATLT,omitempty" doc:"Data communication line no."`
	FAKSD     *string          `xml:"FAKSD,omitempty" json:"FAKSD,omitempty" yaml:"FAKSD,omitempty" doc:"Central billing block for customer"`
	FISKN     *string          `xml:"FISKN,omitempty" json:"FISKN,omitempty" yaml:"FISKN,omitempty" doc:"Account number of the master record with the fiscal address"`
	KNRZA     *string          `xml:"KNRZA,omitempty" json:"KNRZA,omitempty" yaml:"KNRZA,omitempty" doc:"Account number of an alternative payer"`
	KONZS     *string          `xml:"KONZS,omitempty" json:"KONZS,omitempty" yaml:"KONZS,omitempty" doc:"Group key"`
	KTOKD     *string          `xml:"KTOKD,omitempty" json:"KTOKD,omitempty" yaml:"KTOKD,omitempty" doc:"Customer Account Group"`
	KUKLA     *string          `xml:"KUKLA,omitempty" json:"KUKLA,omitempty" yaml:"KUKLA,omitempty" doc:"Customer classification"`
	LAND1     *string          `xml:"LAND1,omitempty" json:"LAND1,omitempty" yaml:"LAND1,omitempty" doc:"Country Key"`
	LIFNR     *string          `xml:"LIFNR,omitempty" json:"LIFNR,omitempty" yaml:"LIFNR,omitempty" doc:"Account Number of Vendor or Creditor"`
	LIFSD     *string          `xml:"LIFSD,omitempty" json:"LIFSD,omitempty" yaml:"LIFSD,omitempty" doc:"Central delivery block for the customer"`
	LOCCO     *string          `xml:"LOCCO,omitempty" json:"LOCCO,omitempty" yaml:"LOCCO,omitempty" doc:"City Coordinates"`
	LOEVM     *string          `xml:"LOEVM,omitempty" json:"LOEVM,omitempty" yaml:"LOEVM,omitempty" doc:"Central Deletion Flag for Master Record"`
	NAME1     *string          `xml:"NAME1,omitempty" json:"NAME1,omitempty" yaml:"NAME1,omitempty" doc:"Name 1"`
	NAME2     *string          `xml:"NAME2,omitempty" json:"NAME2,omitempty" yaml:"NAME2,omitempty" doc:"Name 2"`
	NAME3     *string          `xml:"NAME3,omitempty" json:"NAME3,omitempty" yaml:"NAME3,omitempty" doc:"Name 3"`
	NAME4     *string          `xml:"NAME4,omitempty" json:"NAME4,omitempty" yaml:"NAME4,omitempty" doc:"Name 4"`
	NIELS     *string          `xml:"NIELS,omitempty" json:"NIELS,omitempty" yaml:"NIELS,omitempty" doc:"Nielsen ID"`
	ORT01     *string          `xml:"ORT01,omitempty" json:"ORT01,omitempty" yaml:"ORT01,omitempty" doc:"City"`
	ORT02     *string          `xml:"ORT02,omitempty" json:"ORT02,omitempty" yaml:"ORT02,omitempty" doc:"District"`
	PFACH     *string          `xml:"PFACH,omitempty" json:"PFACH,omitempty" yaml:"PFACH,omitempty" doc:"PO Box"`
	PSTL2     *string          `xml:"PSTL2,omitempty" json:"PSTL2,omitempty" yaml:"PSTL2,omitempty" doc:"P.O. Box Postal Code"`
	PSTLZ     *string          `xml:"PSTLZ,omitempty" json:"PSTLZ,omitempty" yaml:"PSTLZ,omitempty" doc:"Postal Code"`
	REGIO     *string          `xml:"REGIO,omitempty" json:"REGIO,omitempty" yaml:"REGIO,omitempty" doc:"Region (State, Province, County)"`
	COUNC     *string          `xml:"COUNC,omitempty" json:"COUNC,omitempty" yaml:"COUNC,omitempty" doc:"County Code"`
	CITYC     *string          `xml:"CITYC,omitempty" json:"CITYC,omitempty" yaml:"CITYC,omitempty" doc:"City Code"`
	RPMKR     *string          `xml:"RPMKR,omitempty" json:"RPMKR,omitempty" yaml:"RPMKR,omitempty" doc:"Regional Market"`
	SORTL     *string          `xml:"SORTL,omitempty" json:"SORTL,omitempty" yaml:"SORTL,omitempty" doc:"Sort field"`
	SPERR     *string          `xml:"SPERR,omitempty" json:"SPERR,omitempty" yaml:"SPERR,omitempty" doc:"Central posting block"`
	SPRAS     *string          `xml:"SPRAS,omitempty" json:"SPRAS,omitempty" yaml:"SPRAS,omitempty" doc:"Language Key"`
	STCD1     *string          `xml:"STCD1,omitempty" json:"STCD1,omitempty" yaml:"STCD1,omitempty" doc:"Tax Number 1"`
	STCD2     *string          `xml:"STCD2,omitempty" json:"STCD2,omitempty" yaml:"STCD2,omitempty" doc:"Tax Number 2"`
	STKZA     *string          `xml:"STKZA,omitempty" json:"STKZA,omitempty" yaml:"STKZA,omitempty" doc:"Indicator: Business Partner Subject to Equalization Tax?"`
	STKZU     *string          `xml:"STKZU,omitempty" json:"STKZU,omitempty" yaml:"STKZU,omitempty" doc:"Liable for VAT"`
	STRAS     *string          `xml:"STRAS,omitempty" json:"STRAS,omitempty" yaml:"STRAS,omitempty" doc:"House number and street"`
	TELBX     *string          `xml:"TELBX,omitempty" json:"TELBX,omitempty" yaml:"TELBX,omitempty" doc:"Telebox number"`
	TELF1     *string          `xml:"TELF1,omitempty" json:"TELF1,omitempty" yaml:"TELF1,omitempty" doc:"First telephone number"`
	TELF2     *string          `xml:"TELF2,omitempty" json:"TELF2,omitempty" yaml:"TELF2,omitempty" doc:"Second telephone number"`
	TELFX     *string          `xml:"TELFX,omitempty" json:"TELFX,omitempty" yaml:"TELFX,omitempty" doc:"Fax Number"`
	TELTX     *string          `xml:"TELTX,omitempty" json:"TELTX,omitempty" yaml:"TELTX,omitempty" doc:"Teletex number"`
	TELX1     *string          `xml:"TELX1,omitempty" json:"TELX1,omitempty" yaml:"TELX1,omitempty" doc:"Telex number"`
	LZONE     *string          `xml:"LZONE,omitempty" json:"LZONE,omitempty" yaml:"LZONE,omitempty" doc:"Transportation zone to or from which the goods are delivered"`
	XZEMP     *string          `xml:"XZEMP,omitempty" json:"XZEMP,omitempty" yaml:"XZEMP,omitempty" doc:"Indicator: Alternative payee in document allowed ?"`
	VBUND     *string          `xml:"VBUND,omitempty" json:"VBUND,omitempty" yaml:"VBUND,omitempty" doc:"Company ID of Trading Partner"`
	STCEG     *string          `xml:"STCEG,omitempty" json:"STCEG,omitempty" yaml:"STCEG,omitempty" doc:"VAT Registration Number"`
	GFORM     *string          `xml:"GFORM,omitempty" json:"GFORM,omitempty" yaml:"GFORM,omitempty" doc:"Legal status"`
	BRAN1     *string          `xml:"BRAN1,omitempty" json:"BRAN1,omitempty" yaml:"BRAN1,omitempty" doc:"Industry Code 1"`
	BRAN2     *string          `xml:"BRAN2,omitempty" json:"BRAN2,omitempty" yaml:"BRAN2,omitempty" doc:"Industry code 2"`
	BRAN3     *string          `xml:"BRAN3,omitempty" json:"BRAN3,omitempty" yaml:"BRAN3,omitempty" doc:"Industry code 3"`
	BRAN4     *string          `xml:"BRAN4,omitempty" json:"BRAN4,omitempty" yaml:"BRAN4,omitempty" doc:"Industry code 4"`
	BRAN5     *string          `xml:"BRAN5,omitempty" json:"BRAN5,omitempty" yaml:"BRAN5,omitempty" doc:"Industry code 5"`
	UMJAH     *string          `xml:"UMJAH,omitempty" json:"UMJAH,omitempty" yaml:"UMJAH,omitempty" doc:"Year For Which Sales are Given"`
	UWAER     *string          `xml:"UWAER,omitempty" json:"UWAER,omitempty" yaml:"UWAER,omitempty" doc:"Currency of sales figure"`
	JMZAH     *string          `xml:"JMZAH,omitempty" json:"JMZAH,omitempty" yaml:"JMZAH,omitempty" doc:"Yearly number of employees"`
	JMJAH     *string          `xml:"JMJAH,omitempty" json:"JMJAH,omitempty" yaml:"JMJAH,omitempty" doc:"Year for which the number of employees is given"`
	KATR1     *string          `xml:"KATR1,omitempty" json:"KATR1,omitempty" yaml:"KATR1,omitempty" doc:"e-Shop"`
	KATR2     *string          `xml:"KATR2,omitempty" json:"KATR2,omitempty" yaml:"KATR2,omitempty" doc:"Attribute 2"`
	KATR3     *string          `xml:"KATR3,omitempty" json:"KATR3,omitempty" yaml:"KATR3,omitempty" doc:"Attribute 3"`
	KATR4     *string          `xml:"KATR4,omitempty" json:"KATR4,omitempty" yaml:"KATR4,omitempty" doc:"Attribute 4"`
	KATR5     *string          `xml:"KATR5,omitempty" json:"KATR5,omitempty" yaml:"KATR5,omitempty" doc:"Attribute 5"`
	KATR6     *string          `xml:"KATR6,omitempty" json:"KATR6,omitempty" yaml:"KATR6,omitempty" doc:"Attribute 6"`
	KATR7     *string          `xml:"KATR7,omitempty" json:"KATR7,omitempty" yaml:"KATR7,omitempty" doc:"Attribute 7"`
	KATR8     *string          `xml:"KATR8,omitempty" json:"KATR8,omitempty" yaml:"KATR8,omitempty" doc:"Attribute 8"`
	KATR9     *string          `xml:"KATR9,omitempty" json:"KATR9,omitempty" yaml:"KATR9,omitempty" doc:"Attribute 9"`
	KATR10    *string          `xml:"KATR10,omitempty" json:"KATR10,omitempty" yaml:"KATR10,omitempty" doc:"Attribute 10"`
	STKZN     *string          `xml:"STKZN,omitempty" json:"STKZN,omitempty" yaml:"STKZN,omitempty" doc:"Natural Person"`
	UMSA1     *string          `xml:"UMSA1,omitempty" json:"UMSA1,omitempty" yaml:"UMSA1,omitempty" doc:"Field of length 16"`
	TXJCD     *string          `xml:"TXJCD,omitempty" json:"TXJCD,omitempty" yaml:"TXJCD,omitempty" doc:"Tax Jurisdiction"`
	PERIV     *string          `xml:"PERIV,omitempty" json:"PERIV,omitempty" yaml:"PERIV,omitempty" doc:"Fiscal Year Variant"`
	KTOCD     *string          `xml:"KTOCD,omitempty" json:"KTOCD,omitempty" yaml:"KTOCD,omitempty" doc:"Reference Account Group for One-Time Account (Customer)"`
	PFORT     *string          `xml:"PFORT,omitempty" json:"PFORT,omitempty" yaml:"PFORT,omitempty" doc:"PO Box city"`
	DTAMS     *string          `xml:"DTAMS,omitempty" json:"DTAMS,omitempty" yaml:"DTAMS,omitempty" doc:"Indicator for Data Medium Exchange"`
	DTAWS     *string          `xml:"DTAWS,omitempty" json:"DTAWS,omitempty" yaml:"DTAWS,omitempty" doc:"Instruction key for data medium exchange"`
	HZUOR     *string          `xml:"HZUOR,omitempty" json:"HZUOR,omitempty" yaml:"HZUOR,omitempty" doc:"Hierarchy assignment (batch input)"`
	CIVVE     *string          `xml:"CIVVE,omitempty" json:"CIVVE,omitempty" yaml:"CIVVE,omitempty" doc:"ID for mainly non-military use"`
	MILVE     *string          `xml:"MILVE,omitempty" json:"MILVE,omitempty" yaml:"MILVE,omitempty" doc:"ID for mainly military use"`
	SPRAS_ISO *string          `xml:"SPRAS_ISO,omitempty" json:"SPRAS_ISO,omitempty" yaml:"SPRAS_ISO,omitempty" doc:"Language according to ISO 639"`
	FITYP     *string          `xml:"FITYP,omitempty" json:"FITYP,omitempty" yaml:"FITYP,omitempty" doc:"Tax type"`
	STCDT     *string          `xml:"STCDT,omitempty" json:"STCDT,omitempty" yaml:"STCDT,omitempty" doc:"Tax Number Type"`
	STCD3     *string          `xml:"STCD3,omitempty" json:"STCD3,omitempty" yaml:"STCD3,omitempty" doc:"Tax Number 3"`
	STCD4     *string          `xml:"STCD4,omitempty" json:"STCD4,omitempty" yaml:"STCD4,omitempty" doc:"Tax Number 4"`
	XICMS     *string          `xml:"XICMS,omitempty" json:"XICMS,omitempty" yaml:"XICMS,omitempty" doc:"Customer is ICMS-exempt"`
	XXIPI     *string          `xml:"XXIPI,omitempty" json:"XXIPI,omitempty" yaml:"XXIPI,omitempty" doc:"Customer is IPI-exempt"`
	XSUBT     *string          `xml:"XSUBT,omitempty" json:"XSUBT,omitempty" yaml:"XSUBT,omitempty" doc:"Customer group of Substituiçao Tributária calculation- old!!"`
	CFOPC     *string          `xml:"CFOPC,omitempty" json:"CFOPC,omitempty" yaml:"CFOPC,omitempty" doc:"Customer's CFOP category"`
	TXLW1     *string          `xml:"TXLW1,omitempty" json:"TXLW1,omitempty" yaml:"TXLW1,omitempty" doc:"Tax law: ICMS"`
	TXLW2     *string          `xml:"TXLW2,omitempty" json:"TXLW2,omitempty" yaml:"TXLW2,omitempty" doc:"Tax law: IPI"`
	CCC01     *string          `xml:"CCC01,omitempty" json:"CCC01,omitempty" yaml:"CCC01,omitempty" doc:"Indicator for biochemical warfare for legal control"`
	CCC02     *string          `xml:"CCC02,omitempty" json:"CCC02,omitempty" yaml:"CCC02,omitempty" doc:"Indicator for nuclear nonproliferation for legal control"`
	CCC03     *string          `xml:"CCC03,omitempty" json:"CCC03,omitempty" yaml:"CCC03,omitempty" doc:"Indicator for national security for legal control"`
	CCC04     *string          `xml:"CCC04,omitempty" json:"CCC04,omitempty" yaml:"CCC04,omitempty" doc:"Indicator for missile technology for legal control"`
	CASSD     *string          `xml:"CASSD,omitempty" json:"CASSD,omitempty" yaml:"CASSD,omitempty" doc:"Central sales block for customer"`
	KDKG1     *string          `xml:"KDKG1,omitempty" json:"KDKG1,omitempty" yaml:"KDKG1,omitempty" doc:"Customer condition group 1"`
	KDKG2     *string          `xml:"KDKG2,omitempty" json:"KDKG2,omitempty" yaml:"KDKG2,omitempty" doc:"Customer condition group 2"`
	KDKG3     *string          `xml:"KDKG3,omitempty" json:"KDKG3,omitempty" yaml:"KDKG3,omitempty" doc:"Customer condition group 3"`
	KDKG4     *string          `xml:"KDKG4,omitempty" json:"KDKG4,omitempty" yaml:"KDKG4,omitempty" doc:"Customer condition group 4"`
	KDKG5     *string          `xml:"KDKG5,omitempty" json:"KDKG5,omitempty" yaml:"KDKG5,omitempty" doc:"Customer condition group 5"`
	NODEL     *string          `xml:"NODEL,omitempty" json:"NODEL,omitempty" yaml:"NODEL,omitempty" doc:"Central deletion block for master record"`
	XSUB2     *string          `xml:"XSUB2,omitempty" json:"XSUB2,omitempty" yaml:"XSUB2,omitempty" doc:"Customer group for Substituiçao Tributária calculation"`
	WERKS     *string          `xml:"WERKS,omitempty" json:"WERKS,omitempty" yaml:"WERKS,omitempty" doc:"Plant"`
	E1KNA11   *DEBMASE1KNA11   `xml:"E1KNA11,omitempty" json:"E1KNA11,omitempty" yaml:"E1KNA11,omitempty" doc:"Customer Master: Additional General Fields  (KNA1)"`
	E1KNA1H   []*DEBMASE1KNA1H `xml:"E1KNA1H,omitempty" json:"E1KNA1H,omitempty" yaml:"E1KNA1H,omitempty" doc:""`
	E1KNVVM   []*DEBMASE1KNVVM `xml:"E1KNVVM,omitempty" json:"E1KNVVM,omitempty" yaml:"E1KNVVM,omitempty" doc:""`
	E1KNB1M   []*DEBMASE1KNB1M `xml:"E1KNB1M,omitempty" json:"E1KNB1M,omitempty" yaml:"E1KNB1M,omitempty" doc:""`
	E1KNBKM   []*DEBMASE1KNBKM `xml:"E1KNBKM,omitempty" json:"E1KNBKM,omitempty" yaml:"E1KNBKM,omitempty" doc:""`
	E1KNVAM   []*DEBMASE1KNVAM `xml:"E1KNVAM,omitempty" json:"E1KNVAM,omitempty" yaml:"E1KNVAM,omitempty" doc:""`
	E1WRF12   []*DEBMASE1WRF12 `xml:"E1WRF12,omitempty" json:"E1WRF12,omitempty" yaml:"E1WRF12,omitempty" doc:""`
	E1WRF4M   []*DEBMASE1WRF4M `xml:"E1WRF4M,omitempty" json:"E1WRF4M,omitempty" yaml:"E1WRF4M,omitempty" doc:""`
	E1KNVKM   []*DEBMASE1KNVKM `xml:"E1KNVKM,omitempty" json:"E1KNVKM,omitempty" yaml:"E1KNVKM,omitempty" doc:""`
	E1KNEXM   []*DEBMASE1KNEXM `xml:"E1KNEXM,omitempty" json:"E1KNEXM,omitempty" yaml:"E1KNEXM,omitempty" doc:""`
	E1KNASM   []*DEBMASE1KNASM `xml:"E1KNASM,omitempty" json:"E1KNASM,omitempty" yaml:"E1KNASM,omitempty" doc:""`
	E1KNKAM   *DEBMASE1KNKAM   `xml:"E1KNKAM,omitempty" json:"E1KNKAM,omitempty" yaml:"E1KNKAM,omitempty" doc:"Master customer master credit management central data (KNKA)"`
	E1KNATM   []*DEBMASE1KNATM `xml:"E1KNATM,omitempty" json:"E1KNATM,omitempty" yaml:"E1KNATM,omitempty" doc:""`
	E1KNKKM   []*DEBMASE1KNKKM `xml:"E1KNKKM,omitempty" json:"E1KNKKM,omitempty" yaml:"E1KNKKM,omitempty" doc:""`
	E1VCKUN   []*DEBMASE1VCKUN `xml:"E1VCKUN,omitempty" json:"E1VCKUN,omitempty" yaml:"E1VCKUN,omitempty" doc:""`
	E1WRF1M   []*DEBMASE1WRF1M `xml:"E1WRF1M,omitempty" json:"E1WRF1M,omitempty" yaml:"E1WRF1M,omitempty" doc:""`
	E1WRF3M   []*DEBMASE1WRF3M `xml:"E1WRF3M,omitempty" json:"E1WRF3M,omitempty" yaml:"E1WRF3M,omitempty" doc:""`
	E1WRF5M   []*DEBMASE1WRF5M `xml:"E1WRF5M,omitempty" json:"E1WRF5M,omitempty" yaml:"E1WRF5M,omitempty" doc:""`
	E1WRF6M   []*DEBMASE1WRF6M `xml:"E1WRF6M,omitempty" json:"E1WRF6M,omitempty" yaml:"E1WRF6M,omitempty" doc:""`
	E1T023W   []*DEBMASE1T023W `xml:"E1T023W,omitempty" json:"E1T023W,omitempty" yaml:"E1T023W,omitempty" doc:""`
	E1T023X   []*DEBMASE1T023X `xml:"E1T023X,omitempty" json:"E1T023X,omitempty" yaml:"E1T023X,omitempty" doc:""`
}

// Master customer master additional EU tax number
type DEBMASE1KNASM struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	LAND1   *string `xml:"LAND1,omitempty" json:"LAND1,omitempty" yaml:"LAND1,omitempty" doc:"Country Key"`
	STCEG   *string `xml:"STCEG,omitempty" json:"STCEG,omitempty" yaml:"STCEG,omitempty" doc:"VAT Registration Number"`
}

// Master customer master tax categories (KNAT)
type DEBMASE1KNATM struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	TAXGR   *string `xml:"TAXGR,omitempty" json:"TAXGR,omitempty" yaml:"TAXGR,omitempty" doc:"Category indicator for tax codes"`
	SBJDF   *string `xml:"SBJDF,omitempty" json:"SBJDF,omitempty" yaml:"SBJDF,omitempty" doc:"Character field, 8 characters long"`
	SBJDT   *string `xml:"SBJDT,omitempty" json:"SBJDT,omitempty" yaml:"SBJDT,omitempty" doc:"Character field, 8 characters long"`
	EXNR    *string `xml:"EXNR,omitempty" json:"EXNR,omitempty" yaml:"EXNR,omitempty" doc:"Number of exemption certificate"`
	EXRT    *string `xml:"EXRT,omitempty" json:"EXRT,omitempty" yaml:"EXRT,omitempty" doc:"R/2 table"`
	EXDF    *string `xml:"EXDF,omitempty" json:"EXDF,omitempty" yaml:"EXDF,omitempty" doc:"Character field, 8 characters long"`
	EXDT    *string `xml:"EXDT,omitempty" json:"EXDT,omitempty" yaml:"EXDT,omitempty" doc:"Character field, 8 characters long"`
}

// Master customer master company code: Texts, header
type DEBMASE1KNB1H struct {
	SEGMENT    string           `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN      *string          `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	TDOBJECT   *string          `xml:"TDOBJECT,omitempty" json:"TDOBJECT,omitempty" yaml:"TDOBJECT,omitempty" doc:"Texts: Application Object"`
	TDNAME     *string          `xml:"TDNAME,omitempty" json:"TDNAME,omitempty" yaml:"TDNAME,omitempty" doc:"Name"`
	TDID       *string          `xml:"TDID,omitempty" json:"TDID,omitempty" yaml:"TDID,omitempty" doc:"Text ID"`
	TDSPRAS    *string          `xml:"TDSPRAS,omitempty" json:"TDSPRAS,omitempty" yaml:"TDSPRAS,omitempty" doc:"Language Key"`
	TDTEXTTYPE *string          `xml:"TDTEXTTYPE,omitempty" json:"TDTEXTTYPE,omitempty" yaml:"TDTEXTTYPE,omitempty" doc:"SAPscript: Format of Text"`
	TDSPRASISO *string          `xml:"TDSPRASISO,omitempty" json:"TDSPRASISO,omitempty" yaml:"TDSPRASISO,omitempty" doc:"Language according to ISO 639"`
	E1KNB1L    []*DEBMASE1KNB1L `xml:"E1KNB1L,omitempty" json:"E1KNB1L,omitempty" yaml:"E1KNB1L,omitempty" doc:""`
}

// Master customer master company code: text line
type DEBMASE1KNB1L struct {
	SEGMENT  string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN    *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	TDFORMAT *string `xml:"TDFORMAT,omitempty" json:"TDFORMAT,omitempty" yaml:"TDFORMAT,omitempty" doc:"Tag column"`
	TDLINE   *string `xml:"TDLINE,omitempty" json:"TDLINE,omitempty" yaml:"TDLINE,omitempty" doc:"Text Line"`
}

// Master Customer Master Company Code (KNB1)
type DEBMASE1KNB1M struct {
	SEGMENT    string           `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN      *string          `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	BUKRS      *string          `xml:"BUKRS,omitempty" json:"BUKRS,omitempty" yaml:"BUKRS,omitempty" doc:"Company Code"`
	SPERR      *string          `xml:"SPERR,omitempty" json:"SPERR,omitempty" yaml:"SPERR,omitempty" doc:"Posting block for company code"`
	LOEVM      *string          `xml:"LOEVM,omitempty" json:"LOEVM,omitempty" yaml:"LOEVM,omitempty" doc:"Deletion Flag for Master Record (Company Code Level)"`
	ZUAWA      *string          `xml:"ZUAWA,omitempty" json:"ZUAWA,omitempty" yaml:"ZUAWA,omitempty" doc:"Key for sorting according to assignment numbers"`
	BUSAB      *string          `xml:"BUSAB,omitempty" json:"BUSAB,omitempty" yaml:"BUSAB,omitempty" doc:"Accounting clerk"`
	AKONT      *string          `xml:"AKONT,omitempty" json:"AKONT,omitempty" yaml:"AKONT,omitempty" doc:"Reconciliation Account in General Ledger"`
	BEGRU      *string          `xml:"BEGRU,omitempty" json:"BEGRU,omitempty" yaml:"BEGRU,omitempty" doc:"Authorization Group"`
	KNRZE      *string          `xml:"KNRZE,omitempty" json:"KNRZE,omitempty" yaml:"KNRZE,omitempty" doc:"Head office account number (in branch accounts)"`
	KNRZB      *string          `xml:"KNRZB,omitempty" json:"KNRZB,omitempty" yaml:"KNRZB,omitempty" doc:"Account number of an alternative payer"`
	ZAMIM      *string          `xml:"ZAMIM,omitempty" json:"ZAMIM,omitempty" yaml:"ZAMIM,omitempty" doc:"Indicator: Payment notice to customer (with cleared items)?"`
	ZAMIV      *string          `xml:"ZAMIV,omitempty" json:"ZAMIV,omitempty" yaml:"ZAMIV,omitempty" doc:"Indicator: payment notice to sales department?"`
	ZAMIR      *string          `xml:"ZAMIR,omitempty" json:"ZAMIR,omitempty" yaml:"ZAMIR,omitempty" doc:"Indicator: payment notice to legal department?"`
	ZAMIB      *string          `xml:"ZAMIB,omitempty" json:"ZAMIB,omitempty" yaml:"ZAMIB,omitempty" doc:"Indicator: Payment notice to the accounting department ?"`
	ZAMIO      *string          `xml:"ZAMIO,omitempty" json:"ZAMIO,omitempty" yaml:"ZAMIO,omitempty" doc:"Indicator: payment notice to customer (w/o cleared items)?"`
	ZWELS      *string          `xml:"ZWELS,omitempty" json:"ZWELS,omitempty" yaml:"ZWELS,omitempty" doc:"List of the Payment Methods to be Considered"`
	XVERR      *string          `xml:"XVERR,omitempty" json:"XVERR,omitempty" yaml:"XVERR,omitempty" doc:"Indicator: Clearing between customer and vendor ?"`
	ZAHLS      *string          `xml:"ZAHLS,omitempty" json:"ZAHLS,omitempty" yaml:"ZAHLS,omitempty" doc:"Block Key for Payment"`
	ZTERM      *string          `xml:"ZTERM,omitempty" json:"ZTERM,omitempty" yaml:"ZTERM,omitempty" doc:"Terms of Payment Key"`
	WAKON      *string          `xml:"WAKON,omitempty" json:"WAKON,omitempty" yaml:"WAKON,omitempty" doc:"Terms of payment key for bill of exchange charges"`
	VZSKZ      *string          `xml:"VZSKZ,omitempty" json:"VZSKZ,omitempty" yaml:"VZSKZ,omitempty" doc:"Interest calculation indicator"`
	ZINDT      *string          `xml:"ZINDT,omitempty" json:"ZINDT,omitempty" yaml:"ZINDT,omitempty" doc:"Key date of the last interest calculation"`
	ZINRT      *string          `xml:"ZINRT,omitempty" json:"ZINRT,omitempty" yaml:"ZINRT,omitempty" doc:"Interest calculation frequency in months"`
	EIKTO      *string          `xml:"EIKTO,omitempty" json:"EIKTO,omitempty" yaml:"EIKTO,omitempty" doc:"Our account number at customer"`
	ZSABE      *string          `xml:"ZSABE,omitempty" json:"ZSABE,omitempty" yaml:"ZSABE,omitempty" doc:"User at customer"`
	KVERM      *string          `xml:"KVERM,omitempty" json:"KVERM,omitempty" yaml:"KVERM,omitempty" doc:"Memo"`
	FDGRV      *string          `xml:"FDGRV,omitempty" json:"FDGRV,omitempty" yaml:"FDGRV,omitempty" doc:"Planning group"`
	VRBKZ      *string          `xml:"VRBKZ,omitempty" json:"VRBKZ,omitempty" yaml:"VRBKZ,omitempty" doc:"Export credit insurance institution number"`
	VLIBB      *string          `xml:"VLIBB,omitempty" json:"VLIBB,omitempty" yaml:"VLIBB,omitempty" doc:"Amount Insured"`
	VRSZL      *string          `xml:"VRSZL,omitempty" json:"VRSZL,omitempty" yaml:"VRSZL,omitempty" doc:"Insurance lead months"`
	VRSPR      *string          `xml:"VRSPR,omitempty" json:"VRSPR,omitempty" yaml:"VRSPR,omitempty" doc:"Deductible percentage rate"`
	VRSNR      *string          `xml:"VRSNR,omitempty" json:"VRSNR,omitempty" yaml:"VRSNR,omitempty" doc:"Insurance number"`
	VERDT      *string          `xml:"VERDT,omitempty" json:"VERDT,omitempty" yaml:"VERDT,omitempty" doc:"Insurance validity date"`
	PERKZ      *string          `xml:"PERKZ,omitempty" json:"PERKZ,omitempty" yaml:"PERKZ,omitempty" doc:"Collective invoice variant"`
	XDEZV      *string          `xml:"XDEZV,omitempty" json:"XDEZV,omitempty" yaml:"XDEZV,omitempty" doc:"Indicator: Local processing?"`
	XAUSZ      *string          `xml:"XAUSZ,omitempty" json:"XAUSZ,omitempty" yaml:"XAUSZ,omitempty" doc:"Indicator for periodic account statements"`
	WEBTR      *string          `xml:"WEBTR,omitempty" json:"WEBTR,omitempty" yaml:"WEBTR,omitempty" doc:"Bill of exchange limit (in local currency)"`
	REMIT      *string          `xml:"REMIT,omitempty" json:"REMIT,omitempty" yaml:"REMIT,omitempty" doc:"Next payee"`
	DATLZ      *string          `xml:"DATLZ,omitempty" json:"DATLZ,omitempty" yaml:"DATLZ,omitempty" doc:"Date of the last interest calculation run"`
	XZVER      *string          `xml:"XZVER,omitempty" json:"XZVER,omitempty" yaml:"XZVER,omitempty" doc:"Indicator: Record Payment History ?"`
	TOGRU      *string          `xml:"TOGRU,omitempty" json:"TOGRU,omitempty" yaml:"TOGRU,omitempty" doc:"Tolerance group for the business partner/G/L account"`
	KULTG      *string          `xml:"KULTG,omitempty" json:"KULTG,omitempty" yaml:"KULTG,omitempty" doc:"Probable time until check is paid"`
	HBKID      *string          `xml:"HBKID,omitempty" json:"HBKID,omitempty" yaml:"HBKID,omitempty" doc:"Short Key for a House Bank"`
	XPORE      *string          `xml:"XPORE,omitempty" json:"XPORE,omitempty" yaml:"XPORE,omitempty" doc:"Indicator: Pay all items separately ?"`
	BLNKZ      *string          `xml:"BLNKZ,omitempty" json:"BLNKZ,omitempty" yaml:"BLNKZ,omitempty" doc:"Subsidy Indicator for Determining the Reduction Rates"`
	ALTKN      *string          `xml:"ALTKN,omitempty" json:"ALTKN,omitempty" yaml:"ALTKN,omitempty" doc:"Previous Master Record Number"`
	ZGRUP      *string          `xml:"ZGRUP,omitempty" json:"ZGRUP,omitempty" yaml:"ZGRUP,omitempty" doc:"Key for Payment Grouping"`
	URLID      *string          `xml:"URLID,omitempty" json:"URLID,omitempty" yaml:"URLID,omitempty" doc:"Short Key for Known/Negotiated Leave"`
	MGRUP      *string          `xml:"MGRUP,omitempty" json:"MGRUP,omitempty" yaml:"MGRUP,omitempty" doc:"Key for dunning notice grouping"`
	LOCKB      *string          `xml:"LOCKB,omitempty" json:"LOCKB,omitempty" yaml:"LOCKB,omitempty" doc:"Key of the Lockbox to Which the Customer Is To Pay"`
	UZAWE      *string          `xml:"UZAWE,omitempty" json:"UZAWE,omitempty" yaml:"UZAWE,omitempty" doc:"Payment Method Supplement"`
	EKVBD      *string          `xml:"EKVBD,omitempty" json:"EKVBD,omitempty" yaml:"EKVBD,omitempty" doc:"Account Number of Buying Group"`
	SREGL      *string          `xml:"SREGL,omitempty" json:"SREGL,omitempty" yaml:"SREGL,omitempty" doc:"Selection Rule for Payment Advices"`
	XEDIP      *string          `xml:"XEDIP,omitempty" json:"XEDIP,omitempty" yaml:"XEDIP,omitempty" doc:"Indicator: Send Payment Advices by EDI"`
	FRGRP      *string          `xml:"FRGRP,omitempty" json:"FRGRP,omitempty" yaml:"FRGRP,omitempty" doc:"Release Approval Group"`
	VRSDG      *string          `xml:"VRSDG,omitempty" json:"VRSDG,omitempty" yaml:"VRSDG,omitempty" doc:"Reason Code Conversion Version"`
	TLFXS      *string          `xml:"TLFXS,omitempty" json:"TLFXS,omitempty" yaml:"TLFXS,omitempty" doc:"Accounting clerk's fax number at the customer/vendor"`
	PERNR      *string          `xml:"PERNR,omitempty" json:"PERNR,omitempty" yaml:"PERNR,omitempty" doc:"Personnel Number"`
	INTAD      *string          `xml:"INTAD,omitempty" json:"INTAD,omitempty" yaml:"INTAD,omitempty" doc:"Internet address of partner company clerk"`
	GUZTE      *string          `xml:"GUZTE,omitempty" json:"GUZTE,omitempty" yaml:"GUZTE,omitempty" doc:"Payment Terms Key for Credit Memos"`
	GRICD      *string          `xml:"GRICD,omitempty" json:"GRICD,omitempty" yaml:"GRICD,omitempty" doc:"Activity Code for Gross Income Tax"`
	GRIDT      *string          `xml:"GRIDT,omitempty" json:"GRIDT,omitempty" yaml:"GRIDT,omitempty" doc:"Distribution Type for Employment Tax"`
	WBRSL      *string          `xml:"WBRSL,omitempty" json:"WBRSL,omitempty" yaml:"WBRSL,omitempty" doc:"Value Adjustment Key"`
	NODEL      *string          `xml:"NODEL,omitempty" json:"NODEL,omitempty" yaml:"NODEL,omitempty" doc:"Deletion bock for master record (company code level)"`
	TLFNS      *string          `xml:"TLFNS,omitempty" json:"TLFNS,omitempty" yaml:"TLFNS,omitempty" doc:"Accounting clerk's telephone number at business partner"`
	CESSION_KZ *string          `xml:"CESSION_KZ,omitempty" json:"CESSION_KZ,omitempty" yaml:"CESSION_KZ,omitempty" doc:"Accounts Receivable Pledging Indicator"`
	GMVKZD     *string          `xml:"GMVKZD,omitempty" json:"GMVKZD,omitempty" yaml:"GMVKZD,omitempty" doc:"Customer is in execution"`
	AVSND      *string          `xml:"AVSND,omitempty" json:"AVSND,omitempty" yaml:"AVSND,omitempty" doc:"Indicator: Send Payment Advice by XML"`
	SMTP_ADDR  *string          `xml:"SMTP_ADDR,omitempty" json:"SMTP_ADDR,omitempty" yaml:"SMTP_ADDR,omitempty" doc:"E-mail Address"`
	E1KNBWM    []*DEBMASE1KNBWM `xml:"E1KNBWM,omitempty" json:"E1KNBWM,omitempty" yaml:"E1KNBWM,omitempty" doc:""`
	E1KNB5M    []*DEBMASE1KNB5M `xml:"E1KNB5M,omitempty" json:"E1KNB5M,omitempty" yaml:"E1KNB5M,omitempty" doc:""`
	E1KNB1H    []*DEBMASE1KNB1H `xml:"E1KNB1H,omitempty" json:"E1KNB1H,omitempty" yaml:"E1KNB1H,omitempty" doc:""`
}

// Master customer master reminder data (KNB5)
type DEBMASE1KNB5M struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	MABER   *string `xml:"MABER,omitempty" json:"MABER,omitempty" yaml:"MABER,omitempty" doc:"Dunning Area"`
	MAHNA   *string `xml:"MAHNA,omitempty" json:"MAHNA,omitempty" yaml:"MAHNA,omitempty" doc:"Dunning Procedure"`
	MANSP   *string `xml:"MANSP,omitempty" json:"MANSP,omitempty" yaml:"MANSP,omitempty" doc:"Dunning Block"`
	MADAT   *string `xml:"MADAT,omitempty" json:"MADAT,omitempty" yaml:"MADAT,omitempty" doc:"Date of Last Dunning Notice"`
	MAHNS   *string `xml:"MAHNS,omitempty" json:"MAHNS,omitempty" yaml:"MAHNS,omitempty" doc:"Dunning Level"`
	KNRMA   *string `xml:"KNRMA,omitempty" json:"KNRMA,omitempty" yaml:"KNRMA,omitempty" doc:"Account number of the dunning recipient"`
	GMVDT   *string `xml:"GMVDT,omitempty" json:"GMVDT,omitempty" yaml:"GMVDT,omitempty" doc:"Date of the legal dunning proceedings"`
	BUSAB   *string `xml:"BUSAB,omitempty" json:"BUSAB,omitempty" yaml:"BUSAB,omitempty" doc:"Dunning clerk"`
}

// Master customer master bank details and bank master
type DEBMASE1KNBKM struct {
	SEGMENT    string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN      *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	BANKS      *string `xml:"BANKS,omitempty" json:"BANKS,omitempty" yaml:"BANKS,omitempty" doc:"Bank country key"`
	BANKL      *string `xml:"BANKL,omitempty" json:"BANKL,omitempty" yaml:"BANKL,omitempty" doc:"Bank number"`
	BANKN      *string `xml:"BANKN,omitempty" json:"BANKN,omitempty" yaml:"BANKN,omitempty" doc:"Bank account number"`
	BKONT      *string `xml:"BKONT,omitempty" json:"BKONT,omitempty" yaml:"BKONT,omitempty" doc:"Bank Control Key"`
	BVTYP      *string `xml:"BVTYP,omitempty" json:"BVTYP,omitempty" yaml:"BVTYP,omitempty" doc:"Partner Bank Type"`
	XEZER      *string `xml:"XEZER,omitempty" json:"XEZER,omitempty" yaml:"XEZER,omitempty" doc:"Indicator: Is there collection authorization ?"`
	BKREF      *string `xml:"BKREF,omitempty" json:"BKREF,omitempty" yaml:"BKREF,omitempty" doc:"Reference specifications for bank details"`
	BANKA      *string `xml:"BANKA,omitempty" json:"BANKA,omitempty" yaml:"BANKA,omitempty" doc:"Name of bank"`
	STRAS      *string `xml:"STRAS,omitempty" json:"STRAS,omitempty" yaml:"STRAS,omitempty" doc:"House number and street"`
	ORT01      *string `xml:"ORT01,omitempty" json:"ORT01,omitempty" yaml:"ORT01,omitempty" doc:"City"`
	SWIFT      *string `xml:"SWIFT,omitempty" json:"SWIFT,omitempty" yaml:"SWIFT,omitempty" doc:"SWIFT Code for International Payments"`
	BGRUP      *string `xml:"BGRUP,omitempty" json:"BGRUP,omitempty" yaml:"BGRUP,omitempty" doc:"Bank group (bank network)"`
	XPGRO      *string `xml:"XPGRO,omitempty" json:"XPGRO,omitempty" yaml:"XPGRO,omitempty" doc:"Post Office Bank Current Account"`
	BNKLZ      *string `xml:"BNKLZ,omitempty" json:"BNKLZ,omitempty" yaml:"BNKLZ,omitempty" doc:"Bank number"`
	PSKTO      *string `xml:"PSKTO,omitempty" json:"PSKTO,omitempty" yaml:"PSKTO,omitempty" doc:"Account Number of Bank Account At Post Office"`
	BRNCH      *string `xml:"BRNCH,omitempty" json:"BRNCH,omitempty" yaml:"BRNCH,omitempty" doc:"Bank Branch"`
	PROVZ      *string `xml:"PROVZ,omitempty" json:"PROVZ,omitempty" yaml:"PROVZ,omitempty" doc:"Region (State, Province, County)"`
	KOINH      *string `xml:"KOINH,omitempty" json:"KOINH,omitempty" yaml:"KOINH,omitempty" doc:"Account Holder Name"`
	KOINH_N    *string `xml:"KOINH_N,omitempty" json:"KOINH_N,omitempty" yaml:"KOINH_N,omitempty" doc:"Account Holder Name"`
	KOVON      *string `xml:"KOVON,omitempty" json:"KOVON,omitempty" yaml:"KOVON,omitempty" doc:"Date (batch input)"`
	KOBIS      *string `xml:"KOBIS,omitempty" json:"KOBIS,omitempty" yaml:"KOBIS,omitempty" doc:"Date (batch input)"`
	IBAN_BANKS *string `xml:"IBAN_BANKS,omitempty" json:"IBAN_BANKS,omitempty" yaml:"IBAN_BANKS,omitempty" doc:"Bank country key"`
	IBAN_BANKL *string `xml:"IBAN_BANKL,omitempty" json:"IBAN_BANKL,omitempty" yaml:"IBAN_BANKL,omitempty" doc:"Bank Keys"`
	IBAN_BIC   *string `xml:"IBAN_BIC,omitempty" json:"IBAN_BIC,omitempty" yaml:"IBAN_BIC,omitempty" doc:"SWIFT Code for International Payments"`
	IBAN       *string `xml:"IBAN,omitempty" json:"IBAN,omitempty" yaml:"IBAN,omitempty" doc:"IBAN (International Bank Account Number)"`
}

// Segment for Witholding Tax Categories in the Customer Master
type DEBMASE1KNBWM struct {
	SEGMENT   string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN     *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	WITHT     *string `xml:"WITHT,omitempty" json:"WITHT,omitempty" yaml:"WITHT,omitempty" doc:"Indicator for withholding tax type"`
	WT_WITHCD *string `xml:"WT_WITHCD,omitempty" json:"WT_WITHCD,omitempty" yaml:"WT_WITHCD,omitempty" doc:"Withholding tax code"`
	WT_AGENT  *string `xml:"WT_AGENT,omitempty" json:"WT_AGENT,omitempty" yaml:"WT_AGENT,omitempty" doc:"Indicator: Withholding tax agent?"`
	WT_AGTDF  *string `xml:"WT_AGTDF,omitempty" json:"WT_AGTDF,omitempty" yaml:"WT_AGTDF,omitempty" doc:"Date (batch input)"`
	WT_AGTDT  *string `xml:"WT_AGTDT,omitempty" json:"WT_AGTDT,omitempty" yaml:"WT_AGTDT,omitempty" doc:"Date (batch input)"`
	WT_WTSTCD *string `xml:"WT_WTSTCD,omitempty" json:"WT_WTSTCD,omitempty" yaml:"WT_WTSTCD,omitempty" doc:"Withholding tax identification number"`
}

// Master customer master export data
type DEBMASE1KNEXM struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	LNDEX   *string `xml:"LNDEX,omitempty" json:"LNDEX,omitempty" yaml:"LNDEX,omitempty" doc:"Country key for export control in customer master"`
	TDOCO   *string `xml:"TDOCO,omitempty" json:"TDOCO,omitempty" yaml:"TDOCO,omitempty" doc:"ID: TDO boycott list for export control"`
	TDODA   *string `xml:"TDODA,omitempty" json:"TDODA,omitempty" yaml:"TDODA,omitempty" doc:"Date of last check of TDO list for export control"`
	SDNCO   *string `xml:"SDNCO,omitempty" json:"SDNCO,omitempty" yaml:"SDNCO,omitempty" doc:"ID: SDN boycott list for export control"`
	SDNDA   *string `xml:"SDNDA,omitempty" json:"SDNDA,omitempty" yaml:"SDNDA,omitempty" doc:"Date of last check of SDN list for export control"`
	DHRCO   *string `xml:"DHRCO,omitempty" json:"DHRCO,omitempty" yaml:"DHRCO,omitempty" doc:"ID: Customer boycott list for export control"`
	DHRDA   *string `xml:"DHRDA,omitempty" json:"DHRDA,omitempty" yaml:"DHRDA,omitempty" doc:"Date of last check in inter. boycott list for exp. control"`
}

// Master customer master credit management central data (KNKA)
type DEBMASE1KNKAM struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	KLIMG   *string `xml:"KLIMG,omitempty" json:"KLIMG,omitempty" yaml:"KLIMG,omitempty" doc:"Field of length 16"`
	KLIME   *string `xml:"KLIME,omitempty" json:"KLIME,omitempty" yaml:"KLIME,omitempty" doc:"Field of length 16"`
	WAERS   *string `xml:"WAERS,omitempty" json:"WAERS,omitempty" yaml:"WAERS,omitempty" doc:"Currency Key"`
	DLAUS   *string `xml:"DLAUS,omitempty" json:"DLAUS,omitempty" yaml:"DLAUS,omitempty" doc:"Date of the last general information"`
}

// Master customer master credit management: Text, header
type DEBMASE1KNKKH struct {
	SEGMENT    string           `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN      *string          `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	TDOBJECT   *string          `xml:"TDOBJECT,omitempty" json:"TDOBJECT,omitempty" yaml:"TDOBJECT,omitempty" doc:"Texts: Application Object"`
	TDNAME     *string          `xml:"TDNAME,omitempty" json:"TDNAME,omitempty" yaml:"TDNAME,omitempty" doc:"Name"`
	TDID       *string          `xml:"TDID,omitempty" json:"TDID,omitempty" yaml:"TDID,omitempty" doc:"Text ID"`
	TDSPRAS    *string          `xml:"TDSPRAS,omitempty" json:"TDSPRAS,omitempty" yaml:"TDSPRAS,omitempty" doc:"Language Key"`
	TDTEXTTYPE *string          `xml:"TDTEXTTYPE,omitempty" json:"TDTEXTTYPE,omitempty" yaml:"TDTEXTTYPE,omitempty" doc:"SAPscript: Format of Text"`
	TDSPRASISO *string          `xml:"TDSPRASISO,omitempty" json:"TDSPRASISO,omitempty" yaml:"TDSPRASISO,omitempty" doc:"Language according to ISO 639"`
	E1KNKKL    []*DEBMASE1KNKKL `xml:"E1KNKKL,omitempty" json:"E1KNKKL,omitempty" yaml:"E1KNKKL,omitempty" doc:""`
}

// Master customer master credit management: Text, line
type DEBMASE1KNKKL struct {
	SEGMENT  string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN    *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	TDFORMAT *string `xml:"TDFORMAT,omitempty" json:"TDFORMAT,omitempty" yaml:"TDFORMAT,omitempty" doc:"Tag column"`
	TDLINE   *string `xml:"TDLINE,omitempty" json:"TDLINE,omitempty" yaml:"TDLINE,omitempty" doc:"Text Line"`
}

// Master customer master credit mgmt control area data (KNKK)
type DEBMASE1KNKKM struct {
	SEGMENT string           `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string          `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	KKBER   *string          `xml:"KKBER,omitempty" json:"KKBER,omitempty" yaml:"KKBER,omitempty" doc:"Credit Control Area"`
	KLIMK   *string          `xml:"KLIMK,omitempty" json:"KLIMK,omitempty" yaml:"KLIMK,omitempty" doc:"Field of length 16"`
	KNKLI   *string          `xml:"KNKLI,omitempty" json:"KNKLI,omitempty" yaml:"KNKLI,omitempty" doc:"Customer's account number with credit limit reference"`
	CTLPC   *string          `xml:"CTLPC,omitempty" json:"CTLPC,omitempty" yaml:"CTLPC,omitempty" doc:"Credit management: Risk category"`
	DTREV   *string          `xml:"DTREV,omitempty" json:"DTREV,omitempty" yaml:"DTREV,omitempty" doc:"Last internal review"`
	CRBLB   *string          `xml:"CRBLB,omitempty" json:"CRBLB,omitempty" yaml:"CRBLB,omitempty" doc:"Indicator: Blocked by credit management ?"`
	SBGRP   *string          `xml:"SBGRP,omitempty" json:"SBGRP,omitempty" yaml:"SBGRP,omitempty" doc:"Credit representative group for credit management"`
	NXTRV   *string          `xml:"NXTRV,omitempty" json:"NXTRV,omitempty" yaml:"NXTRV,omitempty" doc:"Next internal review"`
	KRAUS   *string          `xml:"KRAUS,omitempty" json:"KRAUS,omitempty" yaml:"KRAUS,omitempty" doc:"Credit information number"`
	PAYDB   *string          `xml:"PAYDB,omitempty" json:"PAYDB,omitempty" yaml:"PAYDB,omitempty" doc:"do not use - replaced by DBPAY_CM"`
	DBRAT   *string          `xml:"DBRAT,omitempty" json:"DBRAT,omitempty" yaml:"DBRAT,omitempty" doc:"do not use - replaced by DBRTG_CM"`
	REVDB   *string          `xml:"REVDB,omitempty" json:"REVDB,omitempty" yaml:"REVDB,omitempty" doc:"Last review (external)"`
	GRUPP   *string          `xml:"GRUPP,omitempty" json:"GRUPP,omitempty" yaml:"GRUPP,omitempty" doc:"Customer Credit Group"`
	SBDAT   *string          `xml:"SBDAT,omitempty" json:"SBDAT,omitempty" yaml:"SBDAT,omitempty" doc:"Reference Date"`
	KDGRP   *string          `xml:"KDGRP,omitempty" json:"KDGRP,omitempty" yaml:"KDGRP,omitempty" doc:"Customer Group"`
	DBPAY   *string          `xml:"DBPAY,omitempty" json:"DBPAY,omitempty" yaml:"DBPAY,omitempty" doc:"Payment Index"`
	DBRTG   *string          `xml:"DBRTG,omitempty" json:"DBRTG,omitempty" yaml:"DBRTG,omitempty" doc:"Rating"`
	DBEKR   *string          `xml:"DBEKR,omitempty" json:"DBEKR,omitempty" yaml:"DBEKR,omitempty" doc:"Recommended credit limit"`
	DBWAE   *string          `xml:"DBWAE,omitempty" json:"DBWAE,omitempty" yaml:"DBWAE,omitempty" doc:"ISO currency code"`
	DBMON   *string          `xml:"DBMON,omitempty" json:"DBMON,omitempty" yaml:"DBMON,omitempty" doc:"Date Monitoring"`
	E1KNKKH []*DEBMASE1KNKKH `xml:"E1KNKKH,omitempty" json:"E1KNKKH,omitempty" yaml:"E1KNKKH,omitempty" doc:""`
}

// Master customer master unloading point (KNVA)
type DEBMASE1KNVAM struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	ABLAD   *string `xml:"ABLAD,omitempty" json:"ABLAD,omitempty" yaml:"ABLAD,omitempty" doc:"Unloading Point"`
	KNFAK   *string `xml:"KNFAK,omitempty" json:"KNFAK,omitempty" yaml:"KNFAK,omitempty" doc:"Customer's factory calendar"`
	WANID   *string `xml:"WANID,omitempty" json:"WANID,omitempty" yaml:"WANID,omitempty" doc:"Goods receiving hours ID (default value)"`
	MOAB1   *string `xml:"MOAB1,omitempty" json:"MOAB1,omitempty" yaml:"MOAB1,omitempty" doc:"Goods receipt times: Monday morning from ..."`
	MOBI1   *string `xml:"MOBI1,omitempty" json:"MOBI1,omitempty" yaml:"MOBI1,omitempty" doc:"Goods receiving hours: Monday morning until ..."`
	MOAB2   *string `xml:"MOAB2,omitempty" json:"MOAB2,omitempty" yaml:"MOAB2,omitempty" doc:"Goods receiving hours: Monday afternoon from ..."`
	MOBI2   *string `xml:"MOBI2,omitempty" json:"MOBI2,omitempty" yaml:"MOBI2,omitempty" doc:"Goods receiving hours: Monday afternoon until ..."`
	DIAB1   *string `xml:"DIAB1,omitempty" json:"DIAB1,omitempty" yaml:"DIAB1,omitempty" doc:"Goods receiving hours: Tuesday morning from..."`
	DIBI1   *string `xml:"DIBI1,omitempty" json:"DIBI1,omitempty" yaml:"DIBI1,omitempty" doc:"Goods receipt times: Tuesday morning until ..."`
	DIAB2   *string `xml:"DIAB2,omitempty" json:"DIAB2,omitempty" yaml:"DIAB2,omitempty" doc:"Goods receiving hours: Tuesday afternoon from ..."`
	DIBI2   *string `xml:"DIBI2,omitempty" json:"DIBI2,omitempty" yaml:"DIBI2,omitempty" doc:"Goods receiving hours: Tuesday afternoon until ..."`
	MIAB1   *string `xml:"MIAB1,omitempty" json:"MIAB1,omitempty" yaml:"MIAB1,omitempty" doc:"Goods receiving hours: Wednesday morning from ..."`
	MIBI1   *string `xml:"MIBI1,omitempty" json:"MIBI1,omitempty" yaml:"MIBI1,omitempty" doc:"Goods receiving hours: Wednesday morning until ..."`
	MIAB2   *string `xml:"MIAB2,omitempty" json:"MIAB2,omitempty" yaml:"MIAB2,omitempty" doc:"Goods receiving hours: Wednesday afternoon from ..."`
	MIBI2   *string `xml:"MIBI2,omitempty" json:"MIBI2,omitempty" yaml:"MIBI2,omitempty" doc:"Goods receiving hours: Wednesday afternoon until ..."`
	DOAB1   *string `xml:"DOAB1,omitempty" json:"DOAB1,omitempty" yaml:"DOAB1,omitempty" doc:"Goods receiving hours: Thursday morning from ..."`
	DOBI1   *string `xml:"DOBI1,omitempty" json:"DOBI1,omitempty" yaml:"DOBI1,omitempty" doc:"Goods receiving hours: Thursday morning until ..."`
	DOAB2   *string `xml:"DOAB2,omitempty" json:"DOAB2,omitempty" yaml:"DOAB2,omitempty" doc:"Goods receiving hours: Thursday afternoon from ..."`
	DOBI2   *string `xml:"DOBI2,omitempty" json:"DOBI2,omitempty" yaml:"DOBI2,omitempty" doc:"Goods receiving hours: Thursday afternoon until..."`
	FRAB1   *string `xml:"FRAB1,omitempty" json:"FRAB1,omitempty" yaml:"FRAB1,omitempty" doc:"Goods receiving hours: Friday morning from ..."`
	FRBI1   *string `xml:"FRBI1,omitempty" json:"FRBI1,omitempty" yaml:"FRBI1,omitempty" doc:"Goods receiving hours: Friday morning until ..."`
	FRAB2   *string `xml:"FRAB2,omitempty" json:"FRAB2,omitempty" yaml:"FRAB2,omitempty" doc:"Goods receiving hours: Friday afternoon from ..."`
	FRBI2   *string `xml:"FRBI2,omitempty" json:"FRBI2,omitempty" yaml:"FRBI2,omitempty" doc:"Goods receiving hours: Friday afternoon until ..."`
	SAAB1   *string `xml:"SAAB1,omitempty" json:"SAAB1,omitempty" yaml:"SAAB1,omitempty" doc:"Goods receiving hours: Saturday morning from ..."`
	SABI1   *string `xml:"SABI1,omitempty" json:"SABI1,omitempty" yaml:"SABI1,omitempty" doc:"Goods receiving hours: Saturday morning from ..."`
	SAAB2   *string `xml:"SAAB2,omitempty" json:"SAAB2,omitempty" yaml:"SAAB2,omitempty" doc:"Goods receiving hours: Saturday afternoon from ..."`
	SABI2   *string `xml:"SABI2,omitempty" json:"SABI2,omitempty" yaml:"SABI2,omitempty" doc:"Goods receiving hours: Saturday afternoon until ..."`
	SOAB1   *string `xml:"SOAB1,omitempty" json:"SOAB1,omitempty" yaml:"SOAB1,omitempty" doc:"Goods receiving hours: Sunday morning from ..."`
	SOBI1   *string `xml:"SOBI1,omitempty" json:"SOBI1,omitempty" yaml:"SOBI1,omitempty" doc:"Goods receiving hours: Sunday morning until ..."`
	SOAB2   *string `xml:"SOAB2,omitempty" json:"SOAB2,omitempty" yaml:"SOAB2,omitempty" doc:"Goods receiving hours: Sunday afternoon from ..."`
	SOBI2   *string `xml:"SOBI2,omitempty" json:"SOBI2,omitempty" yaml:"SOBI2,omitempty" doc:"Goods receiving hours: Sunday afternoon until ..."`
	DEFAB   *string `xml:"DEFAB,omitempty" json:"DEFAB,omitempty" yaml:"DEFAB,omitempty" doc:"Default unloading point"`
}

// Master customer master document request (KNVD)
type DEBMASE1KNVDM struct {
	SEGMENT   string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN     *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	DOCTP     *string `xml:"DOCTP,omitempty" json:"DOCTP,omitempty" yaml:"DOCTP,omitempty" doc:"Output Type"`
	SPRAS     *string `xml:"SPRAS,omitempty" json:"SPRAS,omitempty" yaml:"SPRAS,omitempty" doc:"Message language"`
	DOANZ     *string `xml:"DOANZ,omitempty" json:"DOANZ,omitempty" yaml:"DOANZ,omitempty" doc:"3-Byte field"`
	DOVER     *string `xml:"DOVER,omitempty" json:"DOVER,omitempty" yaml:"DOVER,omitempty" doc:"Dispatch time"`
	NACHA     *string `xml:"NACHA,omitempty" json:"NACHA,omitempty" yaml:"NACHA,omitempty" doc:"Message transmission medium"`
	SPRAS_ISO *string `xml:"SPRAS_ISO,omitempty" json:"SPRAS_ISO,omitempty" yaml:"SPRAS_ISO,omitempty" doc:"Language according to ISO 639"`
}

// Master customer master tax indicators (KNVI)
type DEBMASE1KNVIM struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	ALAND   *string `xml:"ALAND,omitempty" json:"ALAND,omitempty" yaml:"ALAND,omitempty" doc:"Departure country (country from which the goods are sent)"`
	TATYP   *string `xml:"TATYP,omitempty" json:"TATYP,omitempty" yaml:"TATYP,omitempty" doc:"Tax category (sales tax, federal sales tax,...)"`
	TAXKD   *string `xml:"TAXKD,omitempty" json:"TAXKD,omitempty" yaml:"TAXKD,omitempty" doc:"Tax classification for customer"`
}

// Master customer master contact person: Texts, header
type DEBMASE1KNVKH struct {
	SEGMENT    string           `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN      *string          `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	TDOBJECT   *string          `xml:"TDOBJECT,omitempty" json:"TDOBJECT,omitempty" yaml:"TDOBJECT,omitempty" doc:"Texts: Application Object"`
	TDNAME     *string          `xml:"TDNAME,omitempty" json:"TDNAME,omitempty" yaml:"TDNAME,omitempty" doc:"Name"`
	TDID       *string          `xml:"TDID,omitempty" json:"TDID,omitempty" yaml:"TDID,omitempty" doc:"Text ID"`
	TDSPRAS    *string          `xml:"TDSPRAS,omitempty" json:"TDSPRAS,omitempty" yaml:"TDSPRAS,omitempty" doc:"Language Key"`
	TDTEXTTYPE *string          `xml:"TDTEXTTYPE,omitempty" json:"TDTEXTTYPE,omitempty" yaml:"TDTEXTTYPE,omitempty" doc:"SAPscript: Format of Text"`
	TDSPRASISO *string          `xml:"TDSPRASISO,omitempty" json:"TDSPRASISO,omitempty" yaml:"TDSPRASISO,omitempty" doc:"Language according to ISO 639"`
	E1KNVKL    []*DEBMASE1KNVKL `xml:"E1KNVKL,omitempty" json:"E1KNVKL,omitempty" yaml:"E1KNVKL,omitempty" doc:""`
}

// Master customer master contact person: text line
type DEBMASE1KNVKL struct {
	SEGMENT  string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN    *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	TDFORMAT *string `xml:"TDFORMAT,omitempty" json:"TDFORMAT,omitempty" yaml:"TDFORMAT,omitempty" doc:"Tag column"`
	TDLINE   *string `xml:"TDLINE,omitempty" json:"TDLINE,omitempty" yaml:"TDLINE,omitempty" doc:"Text Line"`
}

// Master customer master contact person (KNVK)
type DEBMASE1KNVKM struct {
	SEGMENT   string           `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN     *string          `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	PARNR     *string          `xml:"PARNR,omitempty" json:"PARNR,omitempty" yaml:"PARNR,omitempty" doc:"Number of contact person"`
	NAMEV     *string          `xml:"NAMEV,omitempty" json:"NAMEV,omitempty" yaml:"NAMEV,omitempty" doc:"First name"`
	NAME1     *string          `xml:"NAME1,omitempty" json:"NAME1,omitempty" yaml:"NAME1,omitempty" doc:"Name 1"`
	ABTPA     *string          `xml:"ABTPA,omitempty" json:"ABTPA,omitempty" yaml:"ABTPA,omitempty" doc:"Contact person's department at customer"`
	ABTNR     *string          `xml:"ABTNR,omitempty" json:"ABTNR,omitempty" yaml:"ABTNR,omitempty" doc:"Contact person department"`
	UEPAR     *string          `xml:"UEPAR,omitempty" json:"UEPAR,omitempty" yaml:"UEPAR,omitempty" doc:"Higher-level partner"`
	TELF1     *string          `xml:"TELF1,omitempty" json:"TELF1,omitempty" yaml:"TELF1,omitempty" doc:"First telephone number"`
	ANRED     *string          `xml:"ANRED,omitempty" json:"ANRED,omitempty" yaml:"ANRED,omitempty" doc:"Form of address for contact person (Mr, Mrs...etc)"`
	PAFKT     *string          `xml:"PAFKT,omitempty" json:"PAFKT,omitempty" yaml:"PAFKT,omitempty" doc:"Contact person function"`
	PARVO     *string          `xml:"PARVO,omitempty" json:"PARVO,omitempty" yaml:"PARVO,omitempty" doc:"Partner's Authority"`
	PAVIP     *string          `xml:"PAVIP,omitempty" json:"PAVIP,omitempty" yaml:"PAVIP,omitempty" doc:"VIP Partner"`
	PARGE     *string          `xml:"PARGE,omitempty" json:"PARGE,omitempty" yaml:"PARGE,omitempty" doc:"Partner's gender"`
	PARLA     *string          `xml:"PARLA,omitempty" json:"PARLA,omitempty" yaml:"PARLA,omitempty" doc:"Partner language"`
	GBDAT     *string          `xml:"GBDAT,omitempty" json:"GBDAT,omitempty" yaml:"GBDAT,omitempty" doc:"Date of Birth"`
	VRTNR     *string          `xml:"VRTNR,omitempty" json:"VRTNR,omitempty" yaml:"VRTNR,omitempty" doc:"Representative number"`
	BRYTH     *string          `xml:"BRYTH,omitempty" json:"BRYTH,omitempty" yaml:"BRYTH,omitempty" doc:"Call frequency"`
	AKVER     *string          `xml:"AKVER,omitempty" json:"AKVER,omitempty" yaml:"AKVER,omitempty" doc:"Buying habits"`
	NMAIL     *string          `xml:"NMAIL,omitempty" json:"NMAIL,omitempty" yaml:"NMAIL,omitempty" doc:"Advertising material indicator"`
	PARAU     *string          `xml:"PARAU,omitempty" json:"PARAU,omitempty" yaml:"PARAU,omitempty" doc:"Notes about contact person"`
	PARH1     *string          `xml:"PARH1,omitempty" json:"PARH1,omitempty" yaml:"PARH1,omitempty" doc:"Field of Interest 1"`
	PARH2     *string          `xml:"PARH2,omitempty" json:"PARH2,omitempty" yaml:"PARH2,omitempty" doc:"Field of Interest 2"`
	PARH3     *string          `xml:"PARH3,omitempty" json:"PARH3,omitempty" yaml:"PARH3,omitempty" doc:"Field of Interest 3"`
	PARH4     *string          `xml:"PARH4,omitempty" json:"PARH4,omitempty" yaml:"PARH4,omitempty" doc:"Field of Interest 4"`
	PARH5     *string          `xml:"PARH5,omitempty" json:"PARH5,omitempty" yaml:"PARH5,omitempty" doc:"Field of Interest 5"`
	MOAB1     *string          `xml:"MOAB1,omitempty" json:"MOAB1,omitempty" yaml:"MOAB1,omitempty" doc:"Contact person's visiting hours: Monday morning from ..."`
	MOBI1     *string          `xml:"MOBI1,omitempty" json:"MOBI1,omitempty" yaml:"MOBI1,omitempty" doc:"Contact person's visiting hours: Monday morning until ..."`
	MOAB2     *string          `xml:"MOAB2,omitempty" json:"MOAB2,omitempty" yaml:"MOAB2,omitempty" doc:"Contact person's visiting hours: Monday afternoon from ..."`
	MOBI2     *string          `xml:"MOBI2,omitempty" json:"MOBI2,omitempty" yaml:"MOBI2,omitempty" doc:"Contact person's visiting hours: Monday afternoon until ..."`
	DIAB1     *string          `xml:"DIAB1,omitempty" json:"DIAB1,omitempty" yaml:"DIAB1,omitempty" doc:"Contact person's visiting hours: Tuesday morning from..."`
	DIBI1     *string          `xml:"DIBI1,omitempty" json:"DIBI1,omitempty" yaml:"DIBI1,omitempty" doc:"Contact person's visiting hours: Tuesday morning until ..."`
	DIAB2     *string          `xml:"DIAB2,omitempty" json:"DIAB2,omitempty" yaml:"DIAB2,omitempty" doc:"Contact person's visiting hours: Tuesday afternoon from.."`
	DIBI2     *string          `xml:"DIBI2,omitempty" json:"DIBI2,omitempty" yaml:"DIBI2,omitempty" doc:"Contact person's visiting hours: Tuesday afternoon until ..."`
	MIAB1     *string          `xml:"MIAB1,omitempty" json:"MIAB1,omitempty" yaml:"MIAB1,omitempty" doc:"Contact person's visiting hours: Wednesday morning from..."`
	MIBI1     *string          `xml:"MIBI1,omitempty" json:"MIBI1,omitempty" yaml:"MIBI1,omitempty" doc:"Contact person's visiting hours: Wednesday morning until ..."`
	MIAB2     *string          `xml:"MIAB2,omitempty" json:"MIAB2,omitempty" yaml:"MIAB2,omitempty" doc:"Contact person's visiting hours: Wednesday afternoon from .."`
	MIBI2     *string          `xml:"MIBI2,omitempty" json:"MIBI2,omitempty" yaml:"MIBI2,omitempty" doc:"Contact person's visiting hours: Wednesday afternoon until.."`
	DOAB1     *string          `xml:"DOAB1,omitempty" json:"DOAB1,omitempty" yaml:"DOAB1,omitempty" doc:"Contact person's visiting hours: Thursday morning from ...."`
	DOBI1     *string          `xml:"DOBI1,omitempty" json:"DOBI1,omitempty" yaml:"DOBI1,omitempty" doc:"Contact person's visiting hours: Thursday morning until ...."`
	DOAB2     *string          `xml:"DOAB2,omitempty" json:"DOAB2,omitempty" yaml:"DOAB2,omitempty" doc:"Contact person's visiting hours: Thursday afternoon from..."`
	DOBI2     *string          `xml:"DOBI2,omitempty" json:"DOBI2,omitempty" yaml:"DOBI2,omitempty" doc:"Contact person's visiting hours: Thursday afternoon until .."`
	FRAB1     *string          `xml:"FRAB1,omitempty" json:"FRAB1,omitempty" yaml:"FRAB1,omitempty" doc:"Contact person's visiting hours: Friday morning from ..."`
	FRBI1     *string          `xml:"FRBI1,omitempty" json:"FRBI1,omitempty" yaml:"FRBI1,omitempty" doc:"Contact person's visiting hours: Friday morning until ..."`
	FRAB2     *string          `xml:"FRAB2,omitempty" json:"FRAB2,omitempty" yaml:"FRAB2,omitempty" doc:"Contact person's visiting hours: Friday afternoon from ..."`
	FRBI2     *string          `xml:"FRBI2,omitempty" json:"FRBI2,omitempty" yaml:"FRBI2,omitempty" doc:"Contact person's visiting hours: Friday afternoon until ..."`
	SAAB1     *string          `xml:"SAAB1,omitempty" json:"SAAB1,omitempty" yaml:"SAAB1,omitempty" doc:"Contact person's visiting hours: Saturday morning from ..."`
	SABI1     *string          `xml:"SABI1,omitempty" json:"SABI1,omitempty" yaml:"SABI1,omitempty" doc:"Contact person's visiting hours: Saturday morning until ..."`
	SAAB2     *string          `xml:"SAAB2,omitempty" json:"SAAB2,omitempty" yaml:"SAAB2,omitempty" doc:"Contact person's visiting hours: Saturday afternoon from ..."`
	SABI2     *string          `xml:"SABI2,omitempty" json:"SABI2,omitempty" yaml:"SABI2,omitempty" doc:"Contact person's visiting hours: Saturday afternoon until .."`
	SOAB1     *string          `xml:"SOAB1,omitempty" json:"SOAB1,omitempty" yaml:"SOAB1,omitempty" doc:"Contact person's visiting hours: Sunday morning from ..."`
	SOBI1     *string          `xml:"SOBI1,omitempty" json:"SOBI1,omitempty" yaml:"SOBI1,omitempty" doc:"Contact person's visiting hours: Sunday morning until ..."`
	SOAB2     *string          `xml:"SOAB2,omitempty" json:"SOAB2,omitempty" yaml:"SOAB2,omitempty" doc:"Contact person's visiting hours: Sunday afternoon from ..."`
	SOBI2     *string          `xml:"SOBI2,omitempty" json:"SOBI2,omitempty" yaml:"SOBI2,omitempty" doc:"Contact person's visiting hours: Sunday afternoon until ..."`
	PAKN1     *string          `xml:"PAKN1,omitempty" json:"PAKN1,omitempty" yaml:"PAKN1,omitempty" doc:"Field of Interest 6"`
	PAKN2     *string          `xml:"PAKN2,omitempty" json:"PAKN2,omitempty" yaml:"PAKN2,omitempty" doc:"Field of Interest 7"`
	PAKN3     *string          `xml:"PAKN3,omitempty" json:"PAKN3,omitempty" yaml:"PAKN3,omitempty" doc:"Field of Interest 8"`
	PAKN4     *string          `xml:"PAKN4,omitempty" json:"PAKN4,omitempty" yaml:"PAKN4,omitempty" doc:"Contact person: Attribute 9"`
	PAKN5     *string          `xml:"PAKN5,omitempty" json:"PAKN5,omitempty" yaml:"PAKN5,omitempty" doc:"Contact person: Attribute 10"`
	SORTL     *string          `xml:"SORTL,omitempty" json:"SORTL,omitempty" yaml:"SORTL,omitempty" doc:"Sort field"`
	FAMST     *string          `xml:"FAMST,omitempty" json:"FAMST,omitempty" yaml:"FAMST,omitempty" doc:"Marital Status Key"`
	SPNAM     *string          `xml:"SPNAM,omitempty" json:"SPNAM,omitempty" yaml:"SPNAM,omitempty" doc:"Nickname"`
	TITEL_AP  *string          `xml:"TITEL_AP,omitempty" json:"TITEL_AP,omitempty" yaml:"TITEL_AP,omitempty" doc:"Title of contact person (description of function)"`
	PARLA_ISO *string          `xml:"PARLA_ISO,omitempty" json:"PARLA_ISO,omitempty" yaml:"PARLA_ISO,omitempty" doc:"Language according to ISO 639"`
	E1KNVKH   []*DEBMASE1KNVKH `xml:"E1KNVKH,omitempty" json:"E1KNVKH,omitempty" yaml:"E1KNVKH,omitempty" doc:""`
}

// Master customer master licenses (KNVL)
type DEBMASE1KNVLM struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	ALAND   *string `xml:"ALAND,omitempty" json:"ALAND,omitempty" yaml:"ALAND,omitempty" doc:"Departure country (country from which the goods are sent)"`
	TATYP   *string `xml:"TATYP,omitempty" json:"TATYP,omitempty" yaml:"TATYP,omitempty" doc:"Tax category (sales tax, federal sales tax,...)"`
	LICNR   *string `xml:"LICNR,omitempty" json:"LICNR,omitempty" yaml:"LICNR,omitempty" doc:"License number"`
	DATAB   *string `xml:"DATAB,omitempty" json:"DATAB,omitempty" yaml:"DATAB,omitempty" doc:"Valid-From Date"`
	DATBI   *string `xml:"DATBI,omitempty" json:"DATBI,omitempty" yaml:"DATBI,omitempty" doc:"Valid To Date"`
	BELIC   *string `xml:"BELIC,omitempty" json:"BELIC,omitempty" yaml:"BELIC,omitempty" doc:"Confirmation for licenses"`
}

// Master customer master partner roles (KNVP)
type DEBMASE1KNVPM struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	PARVW   *string `xml:"PARVW,omitempty" json:"PARVW,omitempty" yaml:"PARVW,omitempty" doc:"Partner Function"`
	KUNN2   *string `xml:"KUNN2,omitempty" json:"KUNN2,omitempty" yaml:"KUNN2,omitempty" doc:"Customer number of business partner"`
	DEFPA   *string `xml:"DEFPA,omitempty" json:"DEFPA,omitempty" yaml:"DEFPA,omitempty" doc:"Default Partner"`
	KNREF   *string `xml:"KNREF,omitempty" json:"KNREF,omitempty" yaml:"KNREF,omitempty" doc:"Customer description of partner (plant, storage location)"`
	PARZA   *string `xml:"PARZA,omitempty" json:"PARZA,omitempty" yaml:"PARZA,omitempty" doc:"Partner counter"`
}

// Master customer master sales data: Texts, header
type DEBMASE1KNVVH struct {
	SEGMENT    string           `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN      *string          `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	TDOBJECT   *string          `xml:"TDOBJECT,omitempty" json:"TDOBJECT,omitempty" yaml:"TDOBJECT,omitempty" doc:"Texts: Application Object"`
	TDNAME     *string          `xml:"TDNAME,omitempty" json:"TDNAME,omitempty" yaml:"TDNAME,omitempty" doc:"Name"`
	TDID       *string          `xml:"TDID,omitempty" json:"TDID,omitempty" yaml:"TDID,omitempty" doc:"Text ID"`
	TDSPRAS    *string          `xml:"TDSPRAS,omitempty" json:"TDSPRAS,omitempty" yaml:"TDSPRAS,omitempty" doc:"Language Key"`
	TDTEXTTYPE *string          `xml:"TDTEXTTYPE,omitempty" json:"TDTEXTTYPE,omitempty" yaml:"TDTEXTTYPE,omitempty" doc:"SAPscript: Format of Text"`
	TDSPRASISO *string          `xml:"TDSPRASISO,omitempty" json:"TDSPRASISO,omitempty" yaml:"TDSPRASISO,omitempty" doc:"Language according to ISO 639"`
	E1KNVVL    []*DEBMASE1KNVVL `xml:"E1KNVVL,omitempty" json:"E1KNVVL,omitempty" yaml:"E1KNVVL,omitempty" doc:""`
}

// Master customer master sales data: text line
type DEBMASE1KNVVL struct {
	SEGMENT  string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN    *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	TDFORMAT *string `xml:"TDFORMAT,omitempty" json:"TDFORMAT,omitempty" yaml:"TDFORMAT,omitempty" doc:"Tag column"`
	TDLINE   *string `xml:"TDLINE,omitempty" json:"TDLINE,omitempty" yaml:"TDLINE,omitempty" doc:"Text Line"`
}

// Master customer master sales data (KNVV)
type DEBMASE1KNVVM struct {
	SEGMENT       string           `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN         *string          `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	VKORG         *string          `xml:"VKORG,omitempty" json:"VKORG,omitempty" yaml:"VKORG,omitempty" doc:"Sales Organization"`
	VTWEG         *string          `xml:"VTWEG,omitempty" json:"VTWEG,omitempty" yaml:"VTWEG,omitempty" doc:"Distribution Channel"`
	SPART         *string          `xml:"SPART,omitempty" json:"SPART,omitempty" yaml:"SPART,omitempty" doc:"Division"`
	BEGRU         *string          `xml:"BEGRU,omitempty" json:"BEGRU,omitempty" yaml:"BEGRU,omitempty" doc:"Authorization Group"`
	LOEVM         *string          `xml:"LOEVM,omitempty" json:"LOEVM,omitempty" yaml:"LOEVM,omitempty" doc:"Deletion flag for customer (sales level)"`
	VERSG         *string          `xml:"VERSG,omitempty" json:"VERSG,omitempty" yaml:"VERSG,omitempty" doc:"Customer Statistics Group"`
	AUFSD         *string          `xml:"AUFSD,omitempty" json:"AUFSD,omitempty" yaml:"AUFSD,omitempty" doc:"Customer order block (sales area)"`
	KALKS         *string          `xml:"KALKS,omitempty" json:"KALKS,omitempty" yaml:"KALKS,omitempty" doc:"Pricing procedure assigned to this customer"`
	KDGRP         *string          `xml:"KDGRP,omitempty" json:"KDGRP,omitempty" yaml:"KDGRP,omitempty" doc:"Customer group"`
	BZIRK         *string          `xml:"BZIRK,omitempty" json:"BZIRK,omitempty" yaml:"BZIRK,omitempty" doc:"Sales district"`
	KONDA         *string          `xml:"KONDA,omitempty" json:"KONDA,omitempty" yaml:"KONDA,omitempty" doc:"Price group (customer)"`
	PLTYP         *string          `xml:"PLTYP,omitempty" json:"PLTYP,omitempty" yaml:"PLTYP,omitempty" doc:"Price list type"`
	AWAHR         *string          `xml:"AWAHR,omitempty" json:"AWAHR,omitempty" yaml:"AWAHR,omitempty" doc:"Order probability of the item"`
	INCO1         *string          `xml:"INCO1,omitempty" json:"INCO1,omitempty" yaml:"INCO1,omitempty" doc:"Incoterms (Part 1)"`
	INCO2         *string          `xml:"INCO2,omitempty" json:"INCO2,omitempty" yaml:"INCO2,omitempty" doc:"Incoterms (Part 2)"`
	LIFSD         *string          `xml:"LIFSD,omitempty" json:"LIFSD,omitempty" yaml:"LIFSD,omitempty" doc:"Customer delivery block (sales area)"`
	AUTLF         *string          `xml:"AUTLF,omitempty" json:"AUTLF,omitempty" yaml:"AUTLF,omitempty" doc:"Complete delivery defined for each sales order?"`
	ANTLF         *string          `xml:"ANTLF,omitempty" json:"ANTLF,omitempty" yaml:"ANTLF,omitempty" doc:"Maximum Number of Partial Deliveries Allowed Per Item"`
	KZTLF         *string          `xml:"KZTLF,omitempty" json:"KZTLF,omitempty" yaml:"KZTLF,omitempty" doc:"Partial delivery at item level"`
	KZAZU         *string          `xml:"KZAZU,omitempty" json:"KZAZU,omitempty" yaml:"KZAZU,omitempty" doc:"Order Combination Indicator"`
	CHSPL         *string          `xml:"CHSPL,omitempty" json:"CHSPL,omitempty" yaml:"CHSPL,omitempty" doc:"Batch split allowed"`
	LPRIO         *string          `xml:"LPRIO,omitempty" json:"LPRIO,omitempty" yaml:"LPRIO,omitempty" doc:"Delivery Priority"`
	EIKTO         *string          `xml:"EIKTO,omitempty" json:"EIKTO,omitempty" yaml:"EIKTO,omitempty" doc:"Shipper's (Our) Account Number at the Customer or Vendor"`
	VSBED         *string          `xml:"VSBED,omitempty" json:"VSBED,omitempty" yaml:"VSBED,omitempty" doc:"Shipping Conditions"`
	FAKSD         *string          `xml:"FAKSD,omitempty" json:"FAKSD,omitempty" yaml:"FAKSD,omitempty" doc:"Billing block for customer (sales and distribution)"`
	MRNKZ         *string          `xml:"MRNKZ,omitempty" json:"MRNKZ,omitempty" yaml:"MRNKZ,omitempty" doc:"Manual invoice maintenance"`
	PERFK         *string          `xml:"PERFK,omitempty" json:"PERFK,omitempty" yaml:"PERFK,omitempty" doc:"Invoice dates (calendar identification)"`
	PERRL         *string          `xml:"PERRL,omitempty" json:"PERRL,omitempty" yaml:"PERRL,omitempty" doc:"Invoice list schedule (calendar identification)"`
	WAERS         *string          `xml:"WAERS,omitempty" json:"WAERS,omitempty" yaml:"WAERS,omitempty" doc:"Currency"`
	KTGRD         *string          `xml:"KTGRD,omitempty" json:"KTGRD,omitempty" yaml:"KTGRD,omitempty" doc:"Account assignment group for this customer"`
	ZTERM         *string          `xml:"ZTERM,omitempty" json:"ZTERM,omitempty" yaml:"ZTERM,omitempty" doc:"Terms of Payment Key"`
	VWERK         *string          `xml:"VWERK,omitempty" json:"VWERK,omitempty" yaml:"VWERK,omitempty" doc:"Delivering Plant"`
	VKGRP         *string          `xml:"VKGRP,omitempty" json:"VKGRP,omitempty" yaml:"VKGRP,omitempty" doc:"Sales Group"`
	VKBUR         *string          `xml:"VKBUR,omitempty" json:"VKBUR,omitempty" yaml:"VKBUR,omitempty" doc:"Sales Office"`
	VSORT         *string          `xml:"VSORT,omitempty" json:"VSORT,omitempty" yaml:"VSORT,omitempty" doc:"Item proposal"`
	KVGR1         *string          `xml:"KVGR1,omitempty" json:"KVGR1,omitempty" yaml:"KVGR1,omitempty" doc:"Service Flatrate"`
	KVGR2         *string          `xml:"KVGR2,omitempty" json:"KVGR2,omitempty" yaml:"KVGR2,omitempty" doc:"Customer Category"`
	KVGR3         *string          `xml:"KVGR3,omitempty" json:"KVGR3,omitempty" yaml:"KVGR3,omitempty" doc:"Customer group 3"`
	KVGR4         *string          `xml:"KVGR4,omitempty" json:"KVGR4,omitempty" yaml:"KVGR4,omitempty" doc:"Customer group 4"`
	KVGR5         *string          `xml:"KVGR5,omitempty" json:"KVGR5,omitempty" yaml:"KVGR5,omitempty" doc:"Customer group 5"`
	BOKRE         *string          `xml:"BOKRE,omitempty" json:"BOKRE,omitempty" yaml:"BOKRE,omitempty" doc:"Indicator: Customer Is Rebate-Relevant"`
	KURST         *string          `xml:"KURST,omitempty" json:"KURST,omitempty" yaml:"KURST,omitempty" doc:"Exchange Rate Type"`
	PRFRE         *string          `xml:"PRFRE,omitempty" json:"PRFRE,omitempty" yaml:"PRFRE,omitempty" doc:"Relevant for price determination ID"`
	KLABC         *string          `xml:"KLABC,omitempty" json:"KLABC,omitempty" yaml:"KLABC,omitempty" doc:"Customer classification (ABC analysis)"`
	KABSS         *string          `xml:"KABSS,omitempty" json:"KABSS,omitempty" yaml:"KABSS,omitempty" doc:"Customer payment guarantee procedure"`
	KKBER         *string          `xml:"KKBER,omitempty" json:"KKBER,omitempty" yaml:"KKBER,omitempty" doc:"Credit Control Area"`
	CASSD         *string          `xml:"CASSD,omitempty" json:"CASSD,omitempty" yaml:"CASSD,omitempty" doc:"Sales block for customer (sales area)"`
	RDOFF         *string          `xml:"RDOFF,omitempty" json:"RDOFF,omitempty" yaml:"RDOFF,omitempty" doc:"Switch off rounding?"`
	AGREL         *string          `xml:"AGREL,omitempty" json:"AGREL,omitempty" yaml:"AGREL,omitempty" doc:"Indicator: Relevant for agency business"`
	MEGRU         *string          `xml:"MEGRU,omitempty" json:"MEGRU,omitempty" yaml:"MEGRU,omitempty" doc:"Unit of Measure Group"`
	UEBTO         *string          `xml:"UEBTO,omitempty" json:"UEBTO,omitempty" yaml:"UEBTO,omitempty" doc:"Overdelivery tolerance limit (BTCI)"`
	UNTTO         *string          `xml:"UNTTO,omitempty" json:"UNTTO,omitempty" yaml:"UNTTO,omitempty" doc:"Underdelivery tolerance (BTCI)"`
	UEBTK         *string          `xml:"UEBTK,omitempty" json:"UEBTK,omitempty" yaml:"UEBTK,omitempty" doc:"Unlimited overdelivery allowed"`
	PVKSM         *string          `xml:"PVKSM,omitempty" json:"PVKSM,omitempty" yaml:"PVKSM,omitempty" doc:"Customer procedure for product proposal"`
	PODKZ         *string          `xml:"PODKZ,omitempty" json:"PODKZ,omitempty" yaml:"PODKZ,omitempty" doc:"Relevant for POD processing"`
	PODTG         *string          `xml:"PODTG,omitempty" json:"PODTG,omitempty" yaml:"PODTG,omitempty" doc:"Timeframe for Confirmation of POD (BI)"`
	BLIND         *string          `xml:"BLIND,omitempty" json:"BLIND,omitempty" yaml:"BLIND,omitempty" doc:"Indicator: Doc. index compilation active for purchase orders"`
	CARRIER_NOTIF *string          `xml:"CARRIER_NOTIF,omitempty" json:"CARRIER_NOTIF,omitempty" yaml:"CARRIER_NOTIF,omitempty" doc:"Carrier is to be notified"`
	E1KNVPM       []*DEBMASE1KNVPM `xml:"E1KNVPM,omitempty" json:"E1KNVPM,omitempty" yaml:"E1KNVPM,omitempty" doc:""`
	E1KNVDM       []*DEBMASE1KNVDM `xml:"E1KNVDM,omitempty" json:"E1KNVDM,omitempty" yaml:"E1KNVDM,omitempty" doc:""`
	E1KNVIM       []*DEBMASE1KNVIM `xml:"E1KNVIM,omitempty" json:"E1KNVIM,omitempty" yaml:"E1KNVIM,omitempty" doc:""`
	E1KNVLM       []*DEBMASE1KNVLM `xml:"E1KNVLM,omitempty" json:"E1KNVLM,omitempty" yaml:"E1KNVLM,omitempty" doc:""`
	E1KNVVH       []*DEBMASE1KNVVH `xml:"E1KNVVH,omitempty" json:"E1KNVVH,omitempty" yaml:"E1KNVVH,omitempty" doc:""`
}

// Segment for value-only material determination
type DEBMASE1T023W struct {
	SEGMENT        string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN          *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	LOCNR          *string `xml:"LOCNR,omitempty" json:"LOCNR,omitempty" yaml:"LOCNR,omitempty" doc:"Customer Number for Plant"`
	MATKL          *string `xml:"MATKL,omitempty" json:"MATKL,omitempty" yaml:"MATKL,omitempty" doc:"Material Group"`
	WWGPA          *string `xml:"WWGPA,omitempty" json:"WWGPA,omitempty" yaml:"WWGPA,omitempty" doc:"Material group material"`
	KEDET          *string `xml:"KEDET,omitempty" json:"KEDET,omitempty" yaml:"KEDET,omitempty" doc:"Indicates exceptions to type of Inventory Management"`
	WWGPA_EXTERNAL *string `xml:"WWGPA_EXTERNAL,omitempty" json:"WWGPA_EXTERNAL,omitempty" yaml:"WWGPA_EXTERNAL,omitempty" doc:"Material Group Material"`
	WWGPA_VERSION  *string `xml:"WWGPA_VERSION,omitempty" json:"WWGPA_VERSION,omitempty" yaml:"WWGPA_VERSION,omitempty" doc:"Version Number (Future Development) for Field WWGPA"`
	WWGPA_GUID     *string `xml:"WWGPA_GUID,omitempty" json:"WWGPA_GUID,omitempty" yaml:"WWGPA_GUID,omitempty" doc:"External GUID (Future Development) for Field WWGPA"`
}

// Segment for exceptions in VO-material determination
type DEBMASE1T023X struct {
	SEGMENT        string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN          *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	LOCNR          *string `xml:"LOCNR,omitempty" json:"LOCNR,omitempty" yaml:"LOCNR,omitempty" doc:"Customer Number for Plant"`
	MATNR          *string `xml:"MATNR,omitempty" json:"MATNR,omitempty" yaml:"MATNR,omitempty" doc:"Material Number"`
	WMATN          *string `xml:"WMATN,omitempty" json:"WMATN,omitempty" yaml:"WMATN,omitempty" doc:"Posting material number of value-only or individual material"`
	MATKL          *string `xml:"MATKL,omitempty" json:"MATKL,omitempty" yaml:"MATKL,omitempty" doc:"Material Group"`
	MATNR_EXTERNAL *string `xml:"MATNR_EXTERNAL,omitempty" json:"MATNR_EXTERNAL,omitempty" yaml:"MATNR_EXTERNAL,omitempty" doc:"Material Number"`
	MATNR_VERSION  *string `xml:"MATNR_VERSION,omitempty" json:"MATNR_VERSION,omitempty" yaml:"MATNR_VERSION,omitempty" doc:"Version Number for MATNR Field"`
	MATNR_GUID     *string `xml:"MATNR_GUID,omitempty" json:"MATNR_GUID,omitempty" yaml:"MATNR_GUID,omitempty" doc:"External GUID for MATNR Field"`
	WMATN_EXTERNAL *string `xml:"WMATN_EXTERNAL,omitempty" json:"WMATN_EXTERNAL,omitempty" yaml:"WMATN_EXTERNAL,omitempty" doc:"Length of Material No. (Future Development) for Field WMATN"`
	WMATN_VERSION  *string `xml:"WMATN_VERSION,omitempty" json:"WMATN_VERSION,omitempty" yaml:"WMATN_VERSION,omitempty" doc:"Version Number (Future Development) for Field WMATN"`
	WMATN_GUID     *string `xml:"WMATN_GUID,omitempty" json:"WMATN_GUID,omitempty" yaml:"WMATN_GUID,omitempty" doc:"External GUID (Future Development) for Field WMATN"`
}

// Customer Master: Credit Card Data
type DEBMASE1VCKUN struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	CCINS   *string `xml:"CCINS,omitempty" json:"CCINS,omitempty" yaml:"CCINS,omitempty" doc:"Payment cards: Card type"`
	CCNUM   *string `xml:"CCNUM,omitempty" json:"CCNUM,omitempty" yaml:"CCNUM,omitempty" doc:"Payment cards: Card number"`
	CCDEF   *string `xml:"CCDEF,omitempty" json:"CCDEF,omitempty" yaml:"CCDEF,omitempty" doc:"Payment cards: Indicator for default payment card"`
	CCNAME  *string `xml:"CCNAME,omitempty" json:"CCNAME,omitempty" yaml:"CCNAME,omitempty" doc:"Payment cards: Name of cardholder"`
	DATAB   *string `xml:"DATAB,omitempty" json:"DATAB,omitempty" yaml:"DATAB,omitempty" doc:"Payment cards: Valid from"`
	DATBI   *string `xml:"DATBI,omitempty" json:"DATBI,omitempty" yaml:"DATBI,omitempty" doc:"Payment Cards: Valid To"`
	CCTYP   *string `xml:"CCTYP,omitempty" json:"CCTYP,omitempty" yaml:"CCTYP,omitempty" doc:"Payment cards: Card category"`
	CCLOCK  *string `xml:"CCLOCK,omitempty" json:"CCLOCK,omitempty" yaml:"CCLOCK,omitempty" doc:"Payment cards: Reason for payment card block"`
	CCGUID  *string `xml:"CCGUID,omitempty" json:"CCGUID,omitempty" yaml:"CCGUID,omitempty" doc:"GUID of a Payment Card"`
}

// Segment for plant/receiving points
type DEBMASE1WRF12 struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	LOCNR   *string `xml:"LOCNR,omitempty" json:"LOCNR,omitempty" yaml:"LOCNR,omitempty" doc:"Customer Number for Plant"`
	EMPST   *string `xml:"EMPST,omitempty" json:"EMPST,omitempty" yaml:"EMPST,omitempty" doc:"Receiving point"`
	KUNN2   *string `xml:"KUNN2,omitempty" json:"KUNN2,omitempty" yaml:"KUNN2,omitempty" doc:"Customer number of business partner"`
	ABLAD   *string `xml:"ABLAD,omitempty" json:"ABLAD,omitempty" yaml:"ABLAD,omitempty" doc:"Unloading Point"`
}

// Segment for plant master
type DEBMASE1WRF1M struct {
	SEGMENT   string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN     *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	LOCNR     *string `xml:"LOCNR,omitempty" json:"LOCNR,omitempty" yaml:"LOCNR,omitempty" doc:"Customer Number for Plant"`
	EROED     *string `xml:"EROED,omitempty" json:"EROED,omitempty" yaml:"EROED,omitempty" doc:"Opening date"`
	SCHLD     *string `xml:"SCHLD,omitempty" json:"SCHLD,omitempty" yaml:"SCHLD,omitempty" doc:"Closing date"`
	SPDAB     *string `xml:"SPDAB,omitempty" json:"SPDAB,omitempty" yaml:"SPDAB,omitempty" doc:"Block from"`
	SPDBI     *string `xml:"SPDBI,omitempty" json:"SPDBI,omitempty" yaml:"SPDBI,omitempty" doc:"Block to"`
	AUTOB     *string `xml:"AUTOB,omitempty" json:"AUTOB,omitempty" yaml:"AUTOB,omitempty" doc:"Automatic purchase order"`
	KOPRO     *string `xml:"KOPRO,omitempty" json:"KOPRO,omitempty" yaml:"KOPRO,omitempty" doc:"POS outbound profile"`
	LAYVR     *string `xml:"LAYVR,omitempty" json:"LAYVR,omitempty" yaml:"LAYVR,omitempty" doc:"Layout"`
	FLVAR     *string `xml:"FLVAR,omitempty" json:"FLVAR,omitempty" yaml:"FLVAR,omitempty" doc:"Area schema"`
	STFAK     *string `xml:"STFAK,omitempty" json:"STFAK,omitempty" yaml:"STFAK,omitempty" doc:"Calendar"`
	WANID     *string `xml:"WANID,omitempty" json:"WANID,omitempty" yaml:"WANID,omitempty" doc:"Goods receiving hours ID (default value)"`
	MOAB1     *string `xml:"MOAB1,omitempty" json:"MOAB1,omitempty" yaml:"MOAB1,omitempty" doc:"Goods receipt times: Monday morning from ..."`
	MOBI1     *string `xml:"MOBI1,omitempty" json:"MOBI1,omitempty" yaml:"MOBI1,omitempty" doc:"Goods receiving hours: Monday morning until ..."`
	MOAB2     *string `xml:"MOAB2,omitempty" json:"MOAB2,omitempty" yaml:"MOAB2,omitempty" doc:"Goods receiving hours: Monday afternoon from ..."`
	MOBI2     *string `xml:"MOBI2,omitempty" json:"MOBI2,omitempty" yaml:"MOBI2,omitempty" doc:"Goods receiving hours: Monday afternoon until ..."`
	DIAB1     *string `xml:"DIAB1,omitempty" json:"DIAB1,omitempty" yaml:"DIAB1,omitempty" doc:"Goods receiving hours: Tuesday morning from..."`
	DIBI1     *string `xml:"DIBI1,omitempty" json:"DIBI1,omitempty" yaml:"DIBI1,omitempty" doc:"Goods receipt times: Tuesday morning until ..."`
	DIAB2     *string `xml:"DIAB2,omitempty" json:"DIAB2,omitempty" yaml:"DIAB2,omitempty" doc:"Goods receiving hours: Tuesday afternoon from ..."`
	DIBI2     *string `xml:"DIBI2,omitempty" json:"DIBI2,omitempty" yaml:"DIBI2,omitempty" doc:"Goods receiving hours: Tuesday afternoon until ..."`
	MIAB1     *string `xml:"MIAB1,omitempty" json:"MIAB1,omitempty" yaml:"MIAB1,omitempty" doc:"Goods receiving hours: Wednesday morning from ..."`
	MIBI1     *string `xml:"MIBI1,omitempty" json:"MIBI1,omitempty" yaml:"MIBI1,omitempty" doc:"Goods receiving hours: Wednesday morning until ..."`
	MIAB2     *string `xml:"MIAB2,omitempty" json:"MIAB2,omitempty" yaml:"MIAB2,omitempty" doc:"Goods receiving hours: Wednesday morning until ..."`
	MIBI2     *string `xml:"MIBI2,omitempty" json:"MIBI2,omitempty" yaml:"MIBI2,omitempty" doc:"Goods receiving hours: Wednesday afternoon from ..."`
	DOAB1     *string `xml:"DOAB1,omitempty" json:"DOAB1,omitempty" yaml:"DOAB1,omitempty" doc:"Goods receiving hours: Wednesday afternoon until ..."`
	DOBI1     *string `xml:"DOBI1,omitempty" json:"DOBI1,omitempty" yaml:"DOBI1,omitempty" doc:"Goods receiving hours: Thursday morning from ..."`
	DOAB2     *string `xml:"DOAB2,omitempty" json:"DOAB2,omitempty" yaml:"DOAB2,omitempty" doc:"Goods receiving hours: Thursday morning until ..."`
	DOBI2     *string `xml:"DOBI2,omitempty" json:"DOBI2,omitempty" yaml:"DOBI2,omitempty" doc:"Goods receiving hours: Thursday afternoon from ..."`
	FRAB1     *string `xml:"FRAB1,omitempty" json:"FRAB1,omitempty" yaml:"FRAB1,omitempty" doc:"Goods receiving hours: Thursday afternoon until..."`
	FRBI1     *string `xml:"FRBI1,omitempty" json:"FRBI1,omitempty" yaml:"FRBI1,omitempty" doc:"Goods receiving hours: Friday morning from ..."`
	FRAB2     *string `xml:"FRAB2,omitempty" json:"FRAB2,omitempty" yaml:"FRAB2,omitempty" doc:"Goods receiving hours: Friday morning until ..."`
	FRBI2     *string `xml:"FRBI2,omitempty" json:"FRBI2,omitempty" yaml:"FRBI2,omitempty" doc:"Goods receiving hours: Friday afternoon from ..."`
	SAAB1     *string `xml:"SAAB1,omitempty" json:"SAAB1,omitempty" yaml:"SAAB1,omitempty" doc:"Goods receiving hours: Friday afternoon until ..."`
	SABI1     *string `xml:"SABI1,omitempty" json:"SABI1,omitempty" yaml:"SABI1,omitempty" doc:"Goods receiving hours: Saturday morning from ..."`
	SAAB2     *string `xml:"SAAB2,omitempty" json:"SAAB2,omitempty" yaml:"SAAB2,omitempty" doc:"Goods receiving hours: Saturday afternoon from ..."`
	SABI2     *string `xml:"SABI2,omitempty" json:"SABI2,omitempty" yaml:"SABI2,omitempty" doc:"Goods receiving hours: Saturday afternoon until ..."`
	SOAB1     *string `xml:"SOAB1,omitempty" json:"SOAB1,omitempty" yaml:"SOAB1,omitempty" doc:"Goods receiving hours: Sunday morning from ..."`
	SOBI1     *string `xml:"SOBI1,omitempty" json:"SOBI1,omitempty" yaml:"SOBI1,omitempty" doc:"Goods receiving hours: Sunday morning until ..."`
	SOAB2     *string `xml:"SOAB2,omitempty" json:"SOAB2,omitempty" yaml:"SOAB2,omitempty" doc:"Goods receiving hours: Sunday afternoon from ..."`
	SOBI2     *string `xml:"SOBI2,omitempty" json:"SOBI2,omitempty" yaml:"SOBI2,omitempty" doc:"Goods receiving hours: Sunday afternoon until ..."`
	VERFL     *string `xml:"VERFL,omitempty" json:"VERFL,omitempty" yaml:"VERFL,omitempty" doc:"Sales area (floor space)"`
	VERFE     *string `xml:"VERFE,omitempty" json:"VERFE,omitempty" yaml:"VERFE,omitempty" doc:"Sales area (floor space) unit"`
	SPGR1     *string `xml:"SPGR1,omitempty" json:"SPGR1,omitempty" yaml:"SPGR1,omitempty" doc:"Blocking reason"`
	INPRO     *string `xml:"INPRO,omitempty" json:"INPRO,omitempty" yaml:"INPRO,omitempty" doc:"POS inbound profile"`
	EKOAR     *string `xml:"EKOAR,omitempty" json:"EKOAR,omitempty" yaml:"EKOAR,omitempty" doc:"POS outbound: condition type group"`
	KZLIK     *string `xml:"KZLIK,omitempty" json:"KZLIK,omitempty" yaml:"KZLIK,omitempty" doc:"Listing conditions should be created per assortment"`
	BETRP     *string `xml:"BETRP,omitempty" json:"BETRP,omitempty" yaml:"BETRP,omitempty" doc:"Plant profile"`
	ERDAT     *string `xml:"ERDAT,omitempty" json:"ERDAT,omitempty" yaml:"ERDAT,omitempty" doc:"Date on Which Record Was Created"`
	ERNAM     *string `xml:"ERNAM,omitempty" json:"ERNAM,omitempty" yaml:"ERNAM,omitempty" doc:"Name of Person who Created the Object"`
	NLMATFB   *string `xml:"NLMATFB,omitempty" json:"NLMATFB,omitempty" yaml:"NLMATFB,omitempty" doc:"ID: Carry out subsequent listing"`
	BWWRK     *string `xml:"BWWRK,omitempty" json:"BWWRK,omitempty" yaml:"BWWRK,omitempty" doc:"Plant for retail price determination"`
	BWVKO     *string `xml:"BWVKO,omitempty" json:"BWVKO,omitempty" yaml:"BWVKO,omitempty" doc:"Sales organization for retail price determination"`
	BWVTW     *string `xml:"BWVTW,omitempty" json:"BWVTW,omitempty" yaml:"BWVTW,omitempty" doc:"Distribution channel for retail price determination"`
	BBPRO     *string `xml:"BBPRO,omitempty" json:"BBPRO,omitempty" yaml:"BBPRO,omitempty" doc:"Assortment list profile"`
	VKBUR_WRK *string `xml:"VKBUR_WRK,omitempty" json:"VKBUR_WRK,omitempty" yaml:"VKBUR_WRK,omitempty" doc:"Sales office"`
	VLFKZ     *string `xml:"VLFKZ,omitempty" json:"VLFKZ,omitempty" yaml:"VLFKZ,omitempty" doc:"Plant category"`
	LSTFL     *string `xml:"LSTFL,omitempty" json:"LSTFL,omitempty" yaml:"LSTFL,omitempty" doc:"Listing procedure for store or other assortment categories"`
	LIGRD     *string `xml:"LIGRD,omitempty" json:"LIGRD,omitempty" yaml:"LIGRD,omitempty" doc:"Basic listing rule for assortments"`
	VKORG     *string `xml:"VKORG,omitempty" json:"VKORG,omitempty" yaml:"VKORG,omitempty" doc:"Sales organization for intercompany billing"`
	VTWEG     *string `xml:"VTWEG,omitempty" json:"VTWEG,omitempty" yaml:"VTWEG,omitempty" doc:"Distribution channel for intercompany billing"`
	DESROI    *string `xml:"DESROI,omitempty" json:"DESROI,omitempty" yaml:"DESROI,omitempty" doc:"Required ROI (for ALE)"`
	TIMINC    *string `xml:"TIMINC,omitempty" json:"TIMINC,omitempty" yaml:"TIMINC,omitempty" doc:"Time Increment for Investment Buying Algorithms (for ALE)"`
	POSWS     *string `xml:"POSWS,omitempty" json:"POSWS,omitempty" yaml:"POSWS,omitempty" doc:"Currency of POS systems"`
	SSOPT_PRO *string `xml:"SSOPT_PRO,omitempty" json:"SSOPT_PRO,omitempty" yaml:"SSOPT_PRO,omitempty" doc:"Space management profile"`
	WBPRO     *string `xml:"WBPRO,omitempty" json:"WBPRO,omitempty" yaml:"WBPRO,omitempty" doc:"Profile for value-based inventory management"`
	ORGPRICE  *string `xml:"ORGPRICE,omitempty" json:"ORGPRICE,omitempty" yaml:"ORGPRICE,omitempty" doc:"Retail: Original Price for Segment Definition E2WRF1M"`
	PRCTR     *string `xml:"PRCTR,omitempty" json:"PRCTR,omitempty" yaml:"PRCTR,omitempty" doc:"Profit Center"`
	RMA_PROF  *string `xml:"RMA_PROF,omitempty" json:"RMA_PROF,omitempty" yaml:"RMA_PROF,omitempty" doc:"Retail: RMA Profile for Segment Definition E2WRF1M"`
	RMA_VKORG *string `xml:"RMA_VKORG,omitempty" json:"RMA_VKORG,omitempty" yaml:"RMA_VKORG,omitempty" doc:"Cost allocation sales organization"`
	RMA_VTWEG *string `xml:"RMA_VTWEG,omitempty" json:"RMA_VTWEG,omitempty" yaml:"RMA_VTWEG,omitempty" doc:"Cost allocation distribution channel"`
}

// Segment for supplying plant, time-dependent
type DEBMASE1WRF3M struct {
	SEGMENT         string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN           *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	LOCNR           *string `xml:"LOCNR,omitempty" json:"LOCNR,omitempty" yaml:"LOCNR,omitempty" doc:"Customer Number for Plant"`
	MATKL           *string `xml:"MATKL,omitempty" json:"MATKL,omitempty" yaml:"MATKL,omitempty" doc:"Material Group"`
	DATBI           *string `xml:"DATBI,omitempty" json:"DATBI,omitempty" yaml:"DATBI,omitempty" doc:"Valid to date"`
	DATAB           *string `xml:"DATAB,omitempty" json:"DATAB,omitempty" yaml:"DATAB,omitempty" doc:"Valid-from date"`
	LOCLB           *string `xml:"LOCLB,omitempty" json:"LOCLB,omitempty" yaml:"LOCLB,omitempty" doc:"Supplying plant (source of supply)"`
	PRIORITAET      *string `xml:"PRIORITAET,omitempty" json:"PRIORITAET,omitempty" yaml:"PRIORITAET,omitempty" doc:"Supplying Plant Priority (BI)"`
	TRANSPORT_CHAIN *string `xml:"TRANSPORT_CHAIN,omitempty" json:"TRANSPORT_CHAIN,omitempty" yaml:"TRANSPORT_CHAIN,omitempty" doc:"Transportation Chain"`
}

// Segment for plant/departments
type DEBMASE1WRF4M struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	LOCNR   *string `xml:"LOCNR,omitempty" json:"LOCNR,omitempty" yaml:"LOCNR,omitempty" doc:"Customer Number for Plant"`
	ABTNR   *string `xml:"ABTNR,omitempty" json:"ABTNR,omitempty" yaml:"ABTNR,omitempty" doc:"Department number"`
	EMPST   *string `xml:"EMPST,omitempty" json:"EMPST,omitempty" yaml:"EMPST,omitempty" doc:"Receiving point"`
	VERFL   *string `xml:"VERFL,omitempty" json:"VERFL,omitempty" yaml:"VERFL,omitempty" doc:"Sales area (floor space)"`
	VERFE   *string `xml:"VERFE,omitempty" json:"VERFE,omitempty" yaml:"VERFE,omitempty" doc:"Sales area (floor space) unit"`
	LAYVR   *string `xml:"LAYVR,omitempty" json:"LAYVR,omitempty" yaml:"LAYVR,omitempty" doc:"Layout"`
	FLVAR   *string `xml:"FLVAR,omitempty" json:"FLVAR,omitempty" yaml:"FLVAR,omitempty" doc:"Area schema"`
}

// Segment for Plant/Merchant ID for Credit Card Company
type DEBMASE1WRF5M struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	LOCNR   *string `xml:"LOCNR,omitempty" json:"LOCNR,omitempty" yaml:"LOCNR,omitempty" doc:"Customer Number"`
	CCINS   *string `xml:"CCINS,omitempty" json:"CCINS,omitempty" yaml:"CCINS,omitempty" doc:"Payment cards: Card type"`
	MERCH   *string `xml:"MERCH,omitempty" json:"MERCH,omitempty" yaml:"MERCH,omitempty" doc:"Payment cards: Merchant ID at the clearing house"`
	BEZEI   *string `xml:"BEZEI,omitempty" json:"BEZEI,omitempty" yaml:"BEZEI,omitempty" doc:"Description of merchant ID of credit card company"`
}

// Segment for plant/material groups
type DEBMASE1WRF6M struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	MSGFN   *string `xml:"MSGFN,omitempty" json:"MSGFN,omitempty" yaml:"MSGFN,omitempty" doc:"Function"`
	LOCNR   *string `xml:"LOCNR,omitempty" json:"LOCNR,omitempty" yaml:"LOCNR,omitempty" doc:"Customer Number for Plant"`
	MATKL   *string `xml:"MATKL,omitempty" json:"MATKL,omitempty" yaml:"MATKL,omitempty" doc:"Material Group"`
	SSTUF   *string `xml:"SSTUF,omitempty" json:"SSTUF,omitempty" yaml:"SSTUF,omitempty" doc:"Assortment grade"`
	SORBR   *string `xml:"SORBR,omitempty" json:"SORBR,omitempty" yaml:"SORBR,omitempty" doc:"Assortment length"`
	ABTNR   *string `xml:"ABTNR,omitempty" json:"ABTNR,omitempty" yaml:"ABTNR,omitempty" doc:"Department number"`
	PRIOT   *string `xml:"PRIOT,omitempty" json:"PRIOT,omitempty" yaml:"PRIOT,omitempty" doc:"Assortment Priority"`
	DISPR   *string `xml:"DISPR,omitempty" json:"DISPR,omitempty" yaml:"DISPR,omitempty" doc:"Material: MRP profile"`
	PROPR   *string `xml:"PROPR,omitempty" json:"PROPR,omitempty" yaml:"PROPR,omitempty" doc:"Forecast profile"`
	PRIMW   *string `xml:"PRIMW,omitempty" json:"PRIMW,omitempty" yaml:"PRIMW,omitempty" doc:"Price including sales tax"`
	UPROF   *string `xml:"UPROF,omitempty" json:"UPROF,omitempty" yaml:"UPROF,omitempty" doc:"Retail revalution profile"`
	WDAUS   *string `xml:"WDAUS,omitempty" json:"WDAUS,omitempty" yaml:"WDAUS,omitempty" doc:"Exclude material group from POS outbound processing"`
	RQGRP   *string `xml:"RQGRP,omitempty" json:"RQGRP,omitempty" yaml:"RQGRP,omitempty" doc:"Replenishment: requirement group"`
	FBPRO   *string `xml:"FBPRO,omitempty" json:"FBPRO,omitempty" yaml:"FBPRO,omitempty" doc:"Store procurement profile for store order, replenishment"`
	PLTYP_P *string `xml:"PLTYP_P,omitempty" json:"PLTYP_P,omitempty" yaml:"PLTYP_P,omitempty" doc:"Price determination: Item-based price list type"`
	PSTRA   *string `xml:"PSTRA,omitempty" json:"PSTRA,omitempty" yaml:"PSTRA,omitempty" doc:"Sales Pricing: Pricing Strategy"`
	MGINT   *string `xml:"MGINT,omitempty" json:"MGINT,omitempty" yaml:"MGINT,omitempty" doc:"Internal Class Number of a Competitor Group (for ALE)"`
}

// EDI_DCDEBMASDEBMAS07 was auto-generated from WSDL.
type EDI_DCDEBMASDEBMAS07 struct {
	SEGMENT string  `xml:"SEGMENT,attr" json:"-SEGMENT" yaml:"SEGMENT" doc:""`
	TABNAM  string  `xml:"TABNAM" json:"TABNAM" yaml:"TABNAM" doc:"TABNAM"`
	MANDT   *string `xml:"MANDT,omitempty" json:"MANDT,omitempty" yaml:"MANDT,omitempty" doc:"MANDT"`
	DOCNUM  *string `xml:"DOCNUM,omitempty" json:"DOCNUM,omitempty" yaml:"DOCNUM,omitempty" doc:"DOCNUM"`
	DOCREL  *string `xml:"DOCREL,omitempty" json:"DOCREL,omitempty" yaml:"DOCREL,omitempty" doc:"DOCREL"`
	STATUS  *string `xml:"STATUS,omitempty" json:"STATUS,omitempty" yaml:"STATUS,omitempty" doc:"STATUS"`
	DIRECT  int64   `xml:"DIRECT" json:"DIRECT" yaml:"DIRECT" doc:""`
	OUTMOD  *string `xml:"OUTMOD,omitempty" json:"OUTMOD,omitempty" yaml:"OUTMOD,omitempty" doc:"OUTMOD"`
	EXPRSS  *string `xml:"EXPRSS,omitempty" json:"EXPRSS,omitempty" yaml:"EXPRSS,omitempty" doc:"EXPRSS"`
	TEST    *string `xml:"TEST,omitempty" json:"TEST,omitempty" yaml:"TEST,omitempty" doc:"TEST"`
	IDOCTYP string  `xml:"IDOCTYP" json:"IDOCTYP" yaml:"IDOCTYP" doc:"IDOCTYP"`
	CIMTYP  *string `xml:"CIMTYP,omitempty" json:"CIMTYP,omitempty" yaml:"CIMTYP,omitempty" doc:"CIMTYP"`
	MESTYP  string  `xml:"MESTYP" json:"MESTYP" yaml:"MESTYP" doc:"MESTYP"`
	MESCOD  *string `xml:"MESCOD,omitempty" json:"MESCOD,omitempty" yaml:"MESCOD,omitempty" doc:"MESCOD"`
	MESFCT  *string `xml:"MESFCT,omitempty" json:"MESFCT,omitempty" yaml:"MESFCT,omitempty" doc:"MESFCT"`
	STD     *string `xml:"STD,omitempty" json:"STD,omitempty" yaml:"STD,omitempty" doc:"STD"`
	STDVRS  *string `xml:"STDVRS,omitempty" json:"STDVRS,omitempty" yaml:"STDVRS,omitempty" doc:"STDVRS"`
	STDMES  *string `xml:"STDMES,omitempty" json:"STDMES,omitempty" yaml:"STDMES,omitempty" doc:"STDMES"`
	SNDPOR  string  `xml:"SNDPOR" json:"SNDPOR" yaml:"SNDPOR" doc:"SNDPOR"`
	SNDPRT  string  `xml:"SNDPRT" json:"SNDPRT" yaml:"SNDPRT" doc:"SNDPRT"`
	SNDPFC  *string `xml:"SNDPFC,omitempty" json:"SNDPFC,omitempty" yaml:"SNDPFC,omitempty" doc:"SNDPFC"`
	SNDPRN  string  `xml:"SNDPRN" json:"SNDPRN" yaml:"SNDPRN" doc:"SNDPRN"`
	SNDSAD  *string `xml:"SNDSAD,omitempty" json:"SNDSAD,omitempty" yaml:"SNDSAD,omitempty" doc:"SNDSAD"`
	SNDLAD  *string `xml:"SNDLAD,omitempty" json:"SNDLAD,omitempty" yaml:"SNDLAD,omitempty" doc:"SNDLAD"`
	RCVPOR  string  `xml:"RCVPOR" json:"RCVPOR" yaml:"RCVPOR" doc:"RCVPOR"`
	RCVPRT  *string `xml:"RCVPRT,omitempty" json:"RCVPRT,omitempty" yaml:"RCVPRT,omitempty" doc:"RCVPRT"`
	RCVPFC  *string `xml:"RCVPFC,omitempty" json:"RCVPFC,omitempty" yaml:"RCVPFC,omitempty" doc:"RCVPFC"`
	RCVPRN  string  `xml:"RCVPRN" json:"RCVPRN" yaml:"RCVPRN" doc:"RCVPRN"`
	RCVSAD  *string `xml:"RCVSAD,omitempty" json:"RCVSAD,omitempty" yaml:"RCVSAD,omitempty" doc:"RCVSAD"`
	RCVLAD  *string `xml:"RCVLAD,omitempty" json:"RCVLAD,omitempty" yaml:"RCVLAD,omitempty" doc:"RCVLAD"`
	CREDAT  *string `xml:"CREDAT,omitempty" json:"CREDAT,omitempty" yaml:"CREDAT,omitempty" doc:"CREDAT"`
	CRETIM  *string `xml:"CRETIM,omitempty" json:"CRETIM,omitempty" yaml:"CRETIM,omitempty" doc:"CRETIM"`
	REFINT  *string `xml:"REFINT,omitempty" json:"REFINT,omitempty" yaml:"REFINT,omitempty" doc:"REFINT"`
	REFGRP  *string `xml:"REFGRP,omitempty" json:"REFGRP,omitempty" yaml:"REFGRP,omitempty" doc:"REFGRP"`
	REFMES  *string `xml:"REFMES,omitempty" json:"REFMES,omitempty" yaml:"REFMES,omitempty" doc:"REFMES"`
	ARCKEY  *string `xml:"ARCKEY,omitempty" json:"ARCKEY,omitempty" yaml:"ARCKEY,omitempty" doc:"ARCKEY"`
	SERIAL  *string `xml:"SERIAL,omitempty" json:"SERIAL,omitempty" yaml:"SERIAL,omitempty" doc:"SERIAL"`
}
