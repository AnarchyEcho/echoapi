package randusers

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func Randusers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	results, converr := strconv.Atoi(r.URL.Query().Get("results"))
	json, jsonerr := Json(results)
	if converr != nil || jsonerr != nil {
		w.WriteHeader(500)
		w.Write([]byte(""))
	}
	w.Write([]byte(json))
}

func Json(amount int) (string, error) {
	switch {
	case (amount > 5):
		amount = 5
	case amount < 1:
		amount = 1
	}
	arr := []string{"{\"gender\":\"male\",\"name\":{\"title\":\"Mr\",\"first\":\"Víctor\",\"last\":\"Gutiérrez\"},\"location\":{\"street\":{\"number\":4218,\"name\":\"Viaducto Roldán\"},\"city\":\"Tamán\",\"state\":\"Tlaxcala\",\"country\":\"Mexico\",\"postcode\":92507,\"coordinates\":{\"latitude\":\"25.5289\",\"longitude\":\"-172.1906\"},\"timezone\":{\"offset\":\"+9:00\",\"description\":\"Tokyo, Seoul, Osaka, Sapporo, Yakutsk\"}},\"email\":\"victor.gutierrez@example.com\",\"login\":{\"uuid\":\"f6d6fdee-3802-497b-a81a-0afdaeaeba4b\",\"username\":\"beautifulmouse125\",\"password\":\"luan\",\"salt\":\"lGbkLETz\",\"md5\":\"f65265cb3c28f71c83a35245c148544a\",\"sha1\":\"5af36f435696a331c8d7766de6bdd3c9cf3e706d\",\"sha256\":\"7c1a93b80f6779eeff34566c5bb2324975047d1c3e9d04f768ba848fb26f58ee\"},\"dob\":{\"date\":\"1980-12-01T01:23:44.431Z\",\"age\":44},\"registered\":{\"date\":\"2021-07-18T08:01:10.827Z\",\"age\":3},\"phone\":\"(669) 697 6267\",\"cell\":\"(696) 620 2363\",\"id\":{\"name\":\"NSS\",\"value\":\"83 61 28 9479 4\"},\"picture\":{\"large\":\"https://randomuser.me/api/portraits/men/82.jpg\",\"medium\":\"https://randomuser.me/api/portraits/med/men/82.jpg\",\"thumbnail\":\"https://randomuser.me/api/portraits/thumb/men/82.jpg\"},\"nat\":\"MX\"}", "{\"gender\":\"female\",\"name\":{\"title\":\"Mrs\",\"first\":\"Johanne\",\"last\":\"Nielsen\"},\"location\":{\"street\":{\"number\":3830,\"name\":\"Uglevej\"},\"city\":\"Randers Nø\",\"state\":\"Hovedstaden\",\"country\":\"Denmark\",\"postcode\":24793,\"coordinates\":{\"latitude\":\"52.7726\",\"longitude\":\"158.4007\"},\"timezone\":{\"offset\":\"+6:00\",\"description\":\"Almaty, Dhaka, Colombo\"}},\"email\":\"johanne.nielsen@example.com\",\"login\":{\"uuid\":\"aa4036a5-d691-43df-ace7-48699c6ea328\",\"username\":\"organicsnake996\",\"password\":\"adam25\",\"salt\":\"FhiA2JeJ\",\"md5\":\"817c5d3ed0f01622c7d0256f6f1a04dc\",\"sha1\":\"33afd224397167a03626b5f45cac34c73c4263f4\",\"sha256\":\"8a0b77b0bd0a4bea114df720d5ab831b68db53aedf791a1ced32dacd7e662cc1\"},\"dob\":{\"date\":\"1966-07-17T13:43:36.051Z\",\"age\":58},\"registered\":{\"date\":\"2007-03-07T09:52:50.078Z\",\"age\":18},\"phone\":\"36193048\",\"cell\":\"01598436\",\"id\":{\"name\":\"CPR\",\"value\":\"170766-4245\"},\"picture\":{\"large\":\"https://randomuser.me/api/portraits/women/0.jpg\",\"medium\":\"https://randomuser.me/api/portraits/med/women/0.jpg\",\"thumbnail\":\"https://randomuser.me/api/portraits/thumb/women/0.jpg\"},\"nat\":\"DK\"}", "{\"gender\":\"female\",\"name\":{\"title\":\"Miss\",\"first\":\"آیناز\",\"last\":\"یاسمی\"},\"location\":{\"street\":{\"number\":8059,\"name\":\"میدان 15خرداد\"},\"city\":\"اراک\",\"state\":\"هرمزگان\",\"country\":\"Iran\",\"postcode\":12735,\"coordinates\":{\"latitude\":\"-67.0106\",\"longitude\":\"-130.4555\"},\"timezone\":{\"offset\":\"+4:00\",\"description\":\"Abu Dhabi, Muscat, Baku, Tbilisi\"}},\"email\":\"aynz.ysmy@example.com\",\"login\":{\"uuid\":\"6d91919d-f630-45a2-b416-e029fd65a737\",\"username\":\"ticklishgorilla611\",\"password\":\"think\",\"salt\":\"RuGw4l6o\",\"md5\":\"b561a539f3be9526aa153937ef88fb9f\",\"sha1\":\"499f2bad7ffad58394a362977ba2ba29b3132907\",\"sha256\":\"cbb2f2e9fbfe3e2e79f18d15fe09e3defae5f3c5e3a9a87a4c07af2a0558746f\"},\"dob\":{\"date\":\"1993-12-25T06:09:07.885Z\",\"age\":31},\"registered\":{\"date\":\"2016-08-03T12:34:44.447Z\",\"age\":8},\"phone\":\"089-11391949\",\"cell\":\"0921-886-5847\",\"id\":{\"name\":\"\",\"value\":null},\"picture\":{\"large\":\"https://randomuser.me/api/portraits/women/46.jpg\",\"medium\":\"https://randomuser.me/api/portraits/med/women/46.jpg\",\"thumbnail\":\"https://randomuser.me/api/portraits/thumb/women/46.jpg\"},\"nat\":\"IR\"}", "{\"gender\":\"female\",\"name\":{\"title\":\"Mrs\",\"first\":\"Andrea\",\"last\":\"Katić\"},\"location\":{\"street\":{\"number\":1749,\"name\":\"Pivljanska\"},\"city\":\"Niš\",\"state\":\"Moravica\",\"country\":\"Serbia\",\"postcode\":42068,\"coordinates\":{\"latitude\":\"17.9022\",\"longitude\":\"37.4114\"},\"timezone\":{\"offset\":\"-3:00\",\"description\":\"Brazil, Buenos Aires, Georgetown\"}},\"email\":\"andrea.katic@example.com\",\"login\":{\"uuid\":\"92f49ca0-a0b5-485e-be3e-d69db7df225e\",\"username\":\"bluegoose358\",\"password\":\"hookup\",\"salt\":\"F2c0wMvW\",\"md5\":\"a42fcf1a300ac778113b25b33353cdfa\",\"sha1\":\"4aa2dde307f2a6ff42fc4d19c5348a08d2e3b0ae\",\"sha256\":\"b6da72fd4bf1f37d55dc949c1a6f996e7b606c8b3cb4da18abbafe9970cbe1c7\"},\"dob\":{\"date\":\"1962-05-25T12:58:54.812Z\",\"age\":62},\"registered\":{\"date\":\"2016-04-03T09:04:29.187Z\",\"age\":9},\"phone\":\"030-3697-544\",\"cell\":\"060-2481-929\",\"id\":{\"name\":\"SID\",\"value\":\"442581226\"},\"picture\":{\"large\":\"https://randomuser.me/api/portraits/women/57.jpg\",\"medium\":\"https://randomuser.me/api/portraits/med/women/57.jpg\",\"thumbnail\":\"https://randomuser.me/api/portraits/thumb/women/57.jpg\"},\"nat\":\"RS\"}", "{\"gender\":\"female\",\"name\":{\"title\":\"Mrs\",\"first\":\"Brittany\",\"last\":\"Shelton\"},\"location\":{\"street\":{\"number\":2169,\"name\":\"O'Connell Avenue\"},\"city\":\"Tullow\",\"state\":\"Limerick\",\"country\":\"Ireland\",\"postcode\":66315,\"coordinates\":{\"latitude\":\"-69.9053\",\"longitude\":\"151.4035\"},\"timezone\":{\"offset\":\"+4:00\",\"description\":\"Abu Dhabi, Muscat, Baku, Tbilisi\"}},\"email\":\"brittany.shelton@example.com\",\"login\":{\"uuid\":\"6e42eda9-e6cd-4c04-a52a-8558d0f595c7\",\"username\":\"goldentiger389\",\"password\":\"sites\",\"salt\":\"MKyFUJed\",\"md5\":\"46bcb9ab9e4b5e3010c171eb08ebfaf8\",\"sha1\":\"20e7d2efc1a4e42719975e948cc931a7e467320c\",\"sha256\":\"7858e02b4753a0ba6f5373e1c0ff73f7da7a800dfa0b3ac338bd19413c2b3cda\"},\"dob\":{\"date\":\"1961-02-08T04:39:24.158Z\",\"age\":64},\"registered\":{\"date\":\"2014-09-30T10:28:57.174Z\",\"age\":10},\"phone\":\"051-393-9805\",\"cell\":\"081-184-0227\",\"id\":{\"name\":\"PPS\",\"value\":\"3284810T\"},\"picture\":{\"large\":\"https://randomuser.me/api/portraits/women/55.jpg\",\"medium\":\"https://randomuser.me/api/portraits/med/women/55.jpg\",\"thumbnail\":\"https://randomuser.me/api/portraits/thumb/women/55.jpg\"},\"nat\":\"IE\"}"}

	tmplt := template.New("tmplt")
	data := struct {
		Arr    []string
		Amount int
	}{arr[:amount], amount}

	tmplt, err := tmplt.Parse(
		`{"results":[{{- range $i, $e := .Arr}}
		{{- if gt $i 0 -}}
		,
		{{- end -}}
		{{- . -}}
		{{- end -}}
		],"info":{"seed":"nurgle","results":{{.Amount}},"page":1,"version":"1.4"}}`,
	)
	if err != nil {
		log.Print(err)
		return "", err
	}
	var str bytes.Buffer
	if err2 := tmplt.Execute(&str, data); err2 != nil {
		log.Print(err2)
		return "", err2
	}
	return str.String(), nil
}
