package repository

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"hackaton/internal/app/config"
	"hackaton/internal/models"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

type egov struct {
	client *http.Client
	token  string
	mutex  sync.Mutex
}

func NewEgov(cfg *config.Config) Egov {
	transport := http.DefaultTransport.(*http.Transport).Clone()
	//transport.Proxy = http.ProxyFromEnvironment
	transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} // nolint:gosec
	httpClient := &http.Client{
		Timeout:   60 * time.Second,
		Transport: transport,
	}

	r := &egov{
		client: httpClient,
		mutex:  sync.Mutex{},
	}

	token, err := r.GetToken(context.Background())
	if err != nil {
		panic(err)
	}

	r.token = token
	return r
}

func (r *egov) GetToken(ctx context.Context) (token string, err error) {
	url := "http://hakaton-idp.gov4c.kz/auth/realms/con-web/protocol/openid-connect/token"
	method := "POST"

	payload := strings.NewReader("username=test-operator&password=DjrsmA9RMXRl&client_id=cw-queue-service&grant_type=password")

	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	res, err := r.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	var auth models.Auth
	err = json.Unmarshal(body, &auth)
	if err != nil {
		return
	}

	r.mutex.Lock()
	r.token = auth.AccessToken
	token = auth.AccessToken
	r.mutex.Unlock()

	return
}

func (r *egov) GetPersonData(ctx context.Context, iin string) (person models.Person, err error) {
	url := "http://hakaton-fl.gov4c.kz/api/persons/" + iin
	method := "GET"

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return
	}

REQ:
	req.Header.Add("Authorization", "Bearer "+r.token)

	res, err := r.client.Do(req)
	if err != nil {
		return
	}

	if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
		_, getTokenErr := r.GetToken(ctx)
		if getTokenErr != nil {
			return
		}

		goto REQ
	} else if res.StatusCode != http.StatusOK {
		return person, errors.New(res.Status)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &person)
	if err != nil {
		return
	}

	return
}

func (r *egov) SendSMS(ctx context.Context, msg models.SendSMSRequest) (err error) {
	url := "http://hak-sms123.gov4c.kz/api/smsgateway/send"
	method := "POST"

	byteData, err := json.Marshal(msg)
	if err != nil {
		return
	}

	reader := bytes.NewReader(byteData)
	req, err := http.NewRequest(method, url, reader)

	if err != nil {
		return
	}
REQ:
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+r.token)

	res, err := r.client.Do(req)
	if err != nil {
		return
	}
	if res.StatusCode == http.StatusUnauthorized || res.StatusCode == http.StatusForbidden {
		_, getTokenErr := r.GetToken(ctx)
		if getTokenErr != nil {
			return
		}

		goto REQ
	} else if res.StatusCode != http.StatusOK {
		return errors.New(res.Status)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	resp := models.SendSMSResponse{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return err
	}

	if resp.Status == "200" {
		return nil
	}

	return errors.New(resp.StatusMessage)
}
