package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// VpnGateway is a nested struct in vpc response
type VpnGateway struct {
	VpnType           string                    `json:"VpnType" xml:"VpnType"`
	Status            string                    `json:"Status" xml:"Status"`
	VpcId             string                    `json:"VpcId" xml:"VpcId"`
	SslMaxConnections int64                     `json:"SslMaxConnections" xml:"SslMaxConnections"`
	Spec              string                    `json:"Spec" xml:"Spec"`
	InternetIp        string                    `json:"InternetIp" xml:"InternetIp"`
	CreateTime        int64                     `json:"CreateTime" xml:"CreateTime"`
	AutoPropagate     bool                      `json:"AutoPropagate" xml:"AutoPropagate"`
	ChargeType        string                    `json:"ChargeType" xml:"ChargeType"`
	VpnGatewayId      string                    `json:"VpnGatewayId" xml:"VpnGatewayId"`
	Tag               string                    `json:"Tag" xml:"Tag"`
	IpsecVpn          string                    `json:"IpsecVpn" xml:"IpsecVpn"`
	EndTime           int64                     `json:"EndTime" xml:"EndTime"`
	VSwitchId         string                    `json:"VSwitchId" xml:"VSwitchId"`
	Description       string                    `json:"Description" xml:"Description"`
	EnableBgp         bool                      `json:"EnableBgp" xml:"EnableBgp"`
	BusinessStatus    string                    `json:"BusinessStatus" xml:"BusinessStatus"`
	SslVpn            string                    `json:"SslVpn" xml:"SslVpn"`
	Name              string                    `json:"Name" xml:"Name"`
	ReservationData   ReservationData           `json:"ReservationData" xml:"ReservationData"`
	Tags              TagsInDescribeVpnGateways `json:"Tags" xml:"Tags"`
}