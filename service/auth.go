package service

type Auth struct {
	Code  string `json:"code"`
	Group []string
	Uid   string
}

func (a *Auth) Authorized() error {
	//client := &http.Client{}
	//request, err := http.NewRequest("GET", "http://localhost/api/auth", nil)
	//request.Header.Add("code", a.Code)
	//
	//if err != nil {
	//	return err
	//}
	//
	//response, err := client.Do(request)
	//defer response.Body.Close()
	//
	//a.Uid = response.Header.Get("uid")

	a.Uid = a.Code
	a.Group = []string{"teacher", "student"}
	return nil
}
