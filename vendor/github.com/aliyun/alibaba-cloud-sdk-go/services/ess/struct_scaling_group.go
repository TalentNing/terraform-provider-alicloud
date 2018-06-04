package ess

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

// ScalingGroup is a nested struct in ess response
type ScalingGroup struct {
	DefaultCooldown              int             `json:"DefaultCooldown" xml:"DefaultCooldown"`
	MaxSize                      int             `json:"MaxSize" xml:"MaxSize"`
	PendingCapacity              int             `json:"PendingCapacity" xml:"PendingCapacity"`
	RemovingCapacity             int             `json:"RemovingCapacity" xml:"RemovingCapacity"`
	ScalingGroupName             string          `json:"ScalingGroupName" xml:"ScalingGroupName"`
	ActiveCapacity               int             `json:"ActiveCapacity" xml:"ActiveCapacity"`
	StandbyCapacity              int             `json:"StandbyCapacity" xml:"StandbyCapacity"`
	ProtectedCapacity            int             `json:"ProtectedCapacity" xml:"ProtectedCapacity"`
	ActiveScalingConfigurationId string          `json:"ActiveScalingConfigurationId" xml:"ActiveScalingConfigurationId"`
	ScalingGroupId               string          `json:"ScalingGroupId" xml:"ScalingGroupId"`
	RegionId                     string          `json:"RegionId" xml:"RegionId"`
	TotalCapacity                int             `json:"TotalCapacity" xml:"TotalCapacity"`
	MinSize                      int             `json:"MinSize" xml:"MinSize"`
	LifecycleState               string          `json:"LifecycleState" xml:"LifecycleState"`
	CreationTime                 string          `json:"CreationTime" xml:"CreationTime"`
	ModificationTime             string          `json:"ModificationTime" xml:"ModificationTime"`
	VpcId                        string          `json:"VpcId" xml:"VpcId"`
	VSwitchId                    string          `json:"VSwitchId" xml:"VSwitchId"`
	MultiAZPolicy                string          `json:"MultiAZPolicy" xml:"MultiAZPolicy"`
	VSwitchIds                   VSwitchIds      `json:"VSwitchIds" xml:"VSwitchIds"`
	RemovalPolicies              RemovalPolicies `json:"RemovalPolicies" xml:"RemovalPolicies"`
	DBInstanceIds                DBInstanceIds   `json:"DBInstanceIds" xml:"DBInstanceIds"`
	LoadBalancerIds              LoadBalancerIds `json:"LoadBalancerIds" xml:"LoadBalancerIds"`
}
