package parser

import (
	"reflect"
	"testing"
)


func TestShlex_Parse(t *testing.T) {

crl := `curl --request POST \
  --url https://api.example.com/api/v1/products/3e12d388-dd9b-422b-862a-52463ec305f8/structures/23 \
  --header 'accept: application/json, text/plain, */*' \
  --header 'accept-language: en' \
  --header 'authority: api.bookletix.com' \
  --header 'authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7IklkIjoiN2FkMzYxNGQtMjg3YS00OTE4LTllYzctMWI5YmQ0NmZjN2IxIiwiQW5vbnltb3VzIjpmYWxzZX0sImV4cCI6MTYwNjU4ODM5MSwiaXNzIjoibWFpbiJ9.vNe-E6iBxm45ej0dyCWK-h2sNVKE563r0e4ZTi58eDE' \
  --header 'content-type: application/json;charset=UTF-8' \
  --header 'currency: USD' \
  --header 'origin: https://bookletix.com' \
  --header 'referer: https://bookletix.com/' \
  --header 'sec-fetch-dest: empty' \
  --header 'sec-fetch-mode: cors' \
  --header 'sec-fetch-site: same-site' \
  --header 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36'`

type fields struct {
		text string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Path
		wantErr bool
	}{

		{ name: "parse", fields: fields{
			text: crl,
		},
		want: &Path{
			SourceURL: "https://api.example.com/api/v1/products/3e12d388-dd9b-422b-862a-52463ec305f8",
			Method: "POST",
			TemplatePath: "/api/v1/products/{ProductID}",
		},
		wantErr: false},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Curl{
				text: tt.fields.text,
			}
			got, err := s.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShlex_BODYParse(t *testing.T) {

	crl := `
curl 'https://api.example.com/api/v1/products/b7fdda2b-4ddb-4cc8-bf38-73d5009a222e/reviews' \
  -H 'authority: api.pharmaspace.ru' \
  -H 'accept: application/json, text/plain, */*' \
  -H 'authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiNWNmZTY0MmIyZWNkYTM0OThjZmQ5N2RkIiwiYWN0aXZlIjp0cnVlLCJmaXJzdE5hbWUiOiLQodGC0LDQvdC40YHQu9Cw0LIiLCJsYXN0TmFtZSI6ItCh0LXQvNGR0L3QvtCyIiwicGhvbmUiOiIrNzkyOTU5NTg3NTgiLCJncm91cHMiOlt7ImlkIjoiNWNmZTY0MmEyZWNkYTM0OThjZmQ5N2Q4IiwibmFtZSI6ImFkbWluIn1dLCJlbWFpbENvbmZpcm0iOnRydWUsInBlcm1pc3Npb25zIjp7ImFuYWx5dGljcyI6eyJjcmVhdGUiOnRydWUsInJlYWQiOnRydWUsInVwZGF0ZSI6dHJ1ZSwiZGVsZXRlIjp0cnVlfSwiYmFuayI6eyJjcmVhdGUiOnRydWUsInJlYWQiOnRydWUsInVwZGF0ZSI6dHJ1ZSwiZGVsZXRlIjp0cnVlfSwiY2FzaGllciI6eyJjcmVhdGUiOnRydWUsInJlYWQiOnRydWUsInVwZGF0ZSI6dHJ1ZSwiZGVsZXRlIjp0cnVlfSwiY29udHJhY3QiOnsiY3JlYXRlIjp0cnVlLCJyZWFkIjp0cnVlLCJ1cGRhdGUiOnRydWUsImRlbGV0ZSI6dHJ1ZX0sIm1lcmNoYW50Ijp7ImNyZWF0ZSI6dHJ1ZSwicmVhZCI6dHJ1ZSwidXBkYXRlIjp0cnVlLCJkZWxldGUiOnRydWV9LCJub3RpZmljYXRpb24iOnsiY3JlYXRlIjp0cnVlLCJyZWFkIjp0cnVlLCJ1cGRhdGUiOnRydWUsImRlbGV0ZSI6dHJ1ZX0sInByb2R1Y3RzIjp7ImNyZWF0ZSI6dHJ1ZSwicmVhZCI6dHJ1ZSwidXBkYXRlIjp0cnVlLCJkZWxldGUiOnRydWV9LCJwcm9ncmVzcyI6eyJjcmVhdGUiOnRydWUsInJlYWQiOnRydWUsInVwZGF0ZSI6dHJ1ZSwiZGVsZXRlIjp0cnVlfSwicHJvdmlkZXIiOnsiY3JlYXRlIjp0cnVlLCJyZWFkIjp0cnVlLCJ1cGRhdGUiOnRydWUsImRlbGV0ZSI6dHJ1ZX0sInJlY2llcHQiOnsiY3JlYXRlIjp0cnVlLCJyZWFkIjp0cnVlLCJ1cGRhdGUiOnRydWUsImRlbGV0ZSI6dHJ1ZX0sInNldHRpbmdzIjp7ImNyZWF0ZSI6dHJ1ZSwicmVhZCI6dHJ1ZSwidXBkYXRlIjp0cnVlLCJkZWxldGUiOnRydWV9LCJzcGlkZXIiOnsiY3JlYXRlIjp0cnVlLCJyZWFkIjp0cnVlLCJ1cGRhdGUiOnRydWUsImRlbGV0ZSI6dHJ1ZX0sInN0b3JhZ2UiOnsiY3JlYXRlIjp0cnVlLCJyZWFkIjp0cnVlLCJ1cGRhdGUiOnRydWUsImRlbGV0ZSI6dHJ1ZX0sInVzZXIiOnsiY3JlYXRlIjp0cnVlLCJyZWFkIjp0cnVlLCJ1cGRhdGUiOnRydWUsImRlbGV0ZSI6dHJ1ZX19LCJlbWFpbCI6ImljanAyNjAwQGdtYWlsLmNvbSIsInR5cGUiOjQsImNyZWF0ZWRBdCI6eyJzZWNvbmRzIjotNjIxMzU1OTY3NDl9LCJ1cGRhdGVkQXQiOnsic2Vjb25kcyI6LTYyMTM1NTk2NzQ5fSwicGF5bG9hZCI6eyJtZXJjaGFudElkIjoiNWRhNmZlMjU2Y2IyYmUwOGIxZjgxYTBmIn19LCJleHAiOjE2MDQzMjk4MzksImlzcyI6ImdvLm1pY3JvLnNydi51c2VyIn0.bbA7j-jMLDYku_DQgOTvs29Ai2bPfFzwbHcXgy-Exs0' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.111 Safari/537.36' \
  -H 'origin: https://pharmaspace.ru' \
  -H 'sec-fetch-site: same-site' \
  -H 'sec-fetch-mode: cors' \
  -H 'sec-fetch-dest: empty' \
  -H 'referer: https://pharmaspace.ru/' \
  -H 'accept-language: ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7,zh-CN;q=0.6,zh;q=0.5,de;q=0.4' \
  --compressed`


type fields struct {
		text string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Path
		wantErr bool
	}{
		{ name: "parse body", fields: fields{
			text: crl,
		},
		want: &Path{
			SourceURL: "https://api.example.com/api/v1/products/b7fdda2b-4ddb-4cc8-bf38-73d5009a222e/reviews",
			Method: "GET",
			TemplatePath: "/api/v1/products/<ProductID>/reviews",
		},
		wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Curl{
				text: tt.fields.text,
			}
			got, err := s.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
