package containers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/together-coding/runtime-bridge/db"
	"log"
	"net/http"
	"time"
)

type AgentStatus struct {
	StatusCode int    `json:"-"`
	Ping       string `json:"ping"`
	Init       bool   `json:"init"`
	KeyStatus  bool   `json:"key_status"`
}

// GetContainerFromKey returns RuntimeAllocation associated with the
// passed apiKey. If there is no allocation, not complete RuntimeAllocation
// would be returned.
func GetContainerFromKey(apiKey string) *RuntimeAllocation {
	alloc := RuntimeAllocation{
		ContAPIKey: apiKey,
	}
	db.DB.Where(&alloc, "ContAPIKey").First(&alloc)

	return &alloc
}

// PingAgent sends a ping request to the agent server.
// Caller should check AgentStatus.StatusCode and err to see
// whether the agent is in an expected status.
func PingAgent(url string, port uint16) (agentStatus AgentStatus, err error) {
	agentUrl := fmt.Sprintf("http://%v:%v/ping", url, port)
	httpClient := http.Client{
		Timeout: time.Second * 1,
	}

	resp, err := httpClient.Get(agentUrl)

	// Server or container might not be launched yet.
	if err != nil {
		return
	}

	agentStatus.StatusCode = resp.StatusCode
	if agentStatus.StatusCode != 200 {
		return
	}

	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&agentStatus)
	if err != nil {
		agentStatus.StatusCode = http.StatusUnprocessableEntity
	}

	return
}

// MustPingAgent sends a ping request to the agent server until it responds.
// Because it must be sure that the agent server is up, this function retries maximum `retry` times.
// The maximum time it can take is 500 * `retry` + 50 * `retry` * (`retry` - 1)  milliseconds.
// e.g. retry = 10, then 9.5 seconds.
func MustPingAgent(url string, port uint16, retry int) (agentStatus AgentStatus) {
	for i := 0; i < retry; i++ {
		time.Sleep(time.Millisecond * time.Duration(500+100*i))
		agentStatus, err := PingAgent(url, port)

		if agentStatus.StatusCode != 200 || err != nil {
			continue
		}
		break
	}
	return
}

// InitAgent initialized Runtime Agent server. Bridge send API key for
// identified communication between Bridge and Agent server.
// The agent server not only initialize itself but also return SSH
// credentials for later usage.
func InitAgent(url string, port uint16, contAPIKey string, retry int) map[string]interface{} {
	agentUrl := fmt.Sprintf("http://%v:%v/init", url, port)
	agentReqBody, err := json.Marshal(map[string]string{
		"api_key": contAPIKey,
	})
	if err != nil {
		log.Fatalln(err)
	}

	httpClient := http.Client{
		Timeout: time.Second * 1,
	}
	resp, err := httpClient.Post(agentUrl, "application/json", bytes.NewBuffer(agentReqBody))
	if err != nil {
		if retry > 0 {
			return InitAgent(url, port, contAPIKey, retry-1)
		}
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		log.Fatalln(err)
	}

	return result
}
