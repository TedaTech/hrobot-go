package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	neturl "net/url"
	"strconv"
	"strings"

	"github.com/syself/hrobot-go/models"
)

// VSwitchGetList returns all vSwitches.
func (c *Client) VSwitchGetList() ([]models.VSwitch, error) {
	url := c.baseURL + "/vswitch"
	bytes, err := c.doGetRequest(url)
	if err != nil {
		return nil, err
	}

	var vswitches []models.VSwitch
	err = json.Unmarshal(bytes, &vswitches)
	if err != nil {
		return nil, err
	}

	return vswitches, nil
}

// VSwitchGet returns a vSwitch by ID.
func (c *Client) VSwitchGet(id int) (*models.VSwitch, error) {
	url := fmt.Sprintf(c.baseURL+"/vswitch/%d", id)
	bytes, err := c.doGetRequest(url)
	if err != nil {
		return nil, err
	}

	var vswitch models.VSwitch
	err = json.Unmarshal(bytes, &vswitch)
	if err != nil {
		return nil, err
	}

	return &vswitch, nil
}

// VSwitchCreate creates a new vSwitch.
func (c *Client) VSwitchCreate(input *models.VSwitchCreateInput) (*models.VSwitch, error) {
	url := c.baseURL + "/vswitch"

	formData := neturl.Values{}
	formData.Set("name", input.Name)
	formData.Set("vlan", strconv.Itoa(input.VlanID))

	bytes, err := c.doPostFormRequest(url, formData)
	if err != nil {
		return nil, err
	}

	var vswitch models.VSwitch
	err = json.Unmarshal(bytes, &vswitch)
	if err != nil {
		return nil, err
	}

	return &vswitch, nil
}

// VSwitchAddServer adds a server to a vSwitch.
func (c *Client) VSwitchAddServer(vswitchID, serverNumber int) error {
	url := fmt.Sprintf(c.baseURL+"/vswitch/%d/server", vswitchID)

	formData := neturl.Values{}
	formData.Add("server[]", strconv.Itoa(serverNumber))

	_, err := c.doPostFormRequest(url, formData)
	return err
}

// VSwitchRemoveServer removes a server from a vSwitch.
func (c *Client) VSwitchRemoveServer(vswitchID, serverNumber int) error {
	url := fmt.Sprintf(c.baseURL+"/vswitch/%d/server", vswitchID)

	formData := neturl.Values{}
	formData.Add("server[]", strconv.Itoa(serverNumber))

	req, err := http.NewRequest(http.MethodDelete, url, strings.NewReader(formData.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	_, err = c.doRequest(req)
	return err
}
