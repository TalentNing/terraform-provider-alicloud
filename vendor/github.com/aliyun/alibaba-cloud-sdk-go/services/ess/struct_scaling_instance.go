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

// ScalingInstance is a nested struct in ess response
type ScalingInstance struct {
	InstanceId             string `json:"InstanceId" xml:"InstanceId"`
	ScalingConfigurationId string `json:"ScalingConfigurationId" xml:"ScalingConfigurationId"`
	ScalingGroupId         string `json:"ScalingGroupId" xml:"ScalingGroupId"`
	HealthStatus           string `json:"HealthStatus" xml:"HealthStatus"`
	LoadBalancerWeight     int    `json:"LoadBalancerWeight" xml:"LoadBalancerWeight"`
	LifecycleState         string `json:"LifecycleState" xml:"LifecycleState"`
	CreationTime           string `json:"CreationTime" xml:"CreationTime"`
	CreationType           string `json:"CreationType" xml:"CreationType"`
}
