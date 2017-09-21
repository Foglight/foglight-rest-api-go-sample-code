package rest

import (
	"time"
	"encoding/json"
	"io/ioutil"
	"net/url"
	"strings"
	"net/http"
	"fmt"
)

//Topology Interface support data submission
type Topology interface {
	Type() string
	ToJson() string
}

//Client client for foglight rest api
type Client struct {
	config *Config
	token *string
}

//NewClient create a new client
func NewClient(conf *Config) *Client{
	client:=&Client{conf, nil}
	client.Login()
	return client
}

//Now time in ms
func Now() uint64 {
	return uint64(time.Now().UnixNano()) / uint64(time.Millisecond)
}

//Submit submit topology data
func (c *Client)Submit(data Topology, lastSubmissionMs uint64, now uint64){
	tpl := `
	{
		"data":[{
			"typeName": "%s",
			"properties": {
				%s
			}
	    }],
		"startTime": %d,
		"endTime": %d,
		"agentName": "%s"
		}
	`
	buf:=fmt.Sprintf(tpl, data.Type(), data.ToJson(), lastSubmissionMs, now, c.config.AgentName)
	Log("Submit Topology:\n", buf)
	req, _ := http.NewRequest("POST", c.serverURL("/api/v1/topology/pushData"), strings.NewReader(buf))
	req.Header.Add("Access-Token", *c.token)
	req.Header.Add("Content-Type", "application/json")
	resp, _ := http.DefaultClient.Do(req)
	logger.Println("push-data: ", resp.Status)
}

//Login login and get the token
func  (c *Client) Login() {
	form := &url.Values{"authToken": {c.config.AuthToken}}
	resp, err := http.PostForm(c.serverURL("/api/v1/security/login"), *form)
	if err != nil {
		logger.Fatalln("Fail to login:", err)
	}
	if resp.StatusCode != 200 {
		logger.Fatalln("Login failed:", resp.Status)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var data interface{}
	json.Unmarshal(body, &data)
	m := data.(map[string]interface{})
	m2 := m["data"].(map[string]interface{})
	t := m2["access-token"].(string)
	c.token = &t
	logger.Println("login: ", resp.Status, ", access-token=", t)
}

func (c *Client)serverURL(path string) string {
	return fmt.Sprintf("%s%s", c.config.ServerURL, path)
}

//Logout logout
func (c *Client)Logout() {
	req, _ := http.NewRequest("GET", c.serverURL("/api/v1/security/logout"), nil)
	req.Header.Add("Access-Token", *c.token)
	req.Header.Add("Accept", "application/json")
	resp, _ := http.DefaultClient.Do(req)
	logger.Println("logout: ", resp.Status)
}

