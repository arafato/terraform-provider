package cdn

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

// DomainConfigs is a nested struct in cdn response
type DomainConfigs struct {
	CcConfig                CcConfig                `json:"CcConfig" xml:"CcConfig"`
	ErrorPageConfig         ErrorPageConfig         `json:"ErrorPageConfig" xml:"ErrorPageConfig"`
	OptimizeConfig          OptimizeConfig          `json:"OptimizeConfig" xml:"OptimizeConfig"`
	PageCompressConfig      PageCompressConfig      `json:"PageCompressConfig" xml:"PageCompressConfig"`
	IgnoreQueryStringConfig IgnoreQueryStringConfig `json:"IgnoreQueryStringConfig" xml:"IgnoreQueryStringConfig"`
	RangeConfig             RangeConfig             `json:"RangeConfig" xml:"RangeConfig"`
	RefererConfig           RefererConfig           `json:"RefererConfig" xml:"RefererConfig"`
	ReqAuthConfig           ReqAuthConfig           `json:"ReqAuthConfig" xml:"ReqAuthConfig"`
	SrcHostConfig           SrcHostConfig           `json:"SrcHostConfig" xml:"SrcHostConfig"`
	VideoSeekConfig         VideoSeekConfig         `json:"VideoSeekConfig" xml:"VideoSeekConfig"`
	WafConfig               WafConfig               `json:"WafConfig" xml:"WafConfig"`
	NotifyUrlConfig         NotifyUrlConfig         `json:"NotifyUrlConfig" xml:"NotifyUrlConfig"`
	RedirectTypeConfig      RedirectTypeConfig      `json:"RedirectTypeConfig" xml:"RedirectTypeConfig"`
	ForwardSchemeConfig     ForwardSchemeConfig     `json:"ForwardSchemeConfig" xml:"ForwardSchemeConfig"`
	RemoveQueryStringConfig RemoveQueryStringConfig `json:"RemoveQueryStringConfig" xml:"RemoveQueryStringConfig"`
	L2OssKeyConfig          L2OssKeyConfig          `json:"L2OssKeyConfig" xml:"L2OssKeyConfig"`
	MacServiceConfig        MacServiceConfig        `json:"MacServiceConfig" xml:"MacServiceConfig"`
	GreenManagerConfig      GreenManagerConfig      `json:"GreenManagerConfig" xml:"GreenManagerConfig"`
	HttpsOptionConfig       HttpsOptionConfig       `json:"HttpsOptionConfig" xml:"HttpsOptionConfig"`
	AliBusinessConfig       AliBusinessConfig       `json:"AliBusinessConfig" xml:"AliBusinessConfig"`
	IpAllowListConfig       IpAllowListConfig       `json:"IpAllowListConfig" xml:"IpAllowListConfig"`
	CacheExpiredConfigs     CacheExpiredConfigs     `json:"CacheExpiredConfigs" xml:"CacheExpiredConfigs"`
	HttpErrorPageConfigs    HttpErrorPageConfigs    `json:"HttpErrorPageConfigs" xml:"HttpErrorPageConfigs"`
	HttpHeaderConfigs       HttpHeaderConfigs       `json:"HttpHeaderConfigs" xml:"HttpHeaderConfigs"`
	DynamicConfigs          DynamicConfigs          `json:"DynamicConfigs" xml:"DynamicConfigs"`
	ReqHeaderConfigs        ReqHeaderConfigs        `json:"ReqHeaderConfigs" xml:"ReqHeaderConfigs"`
	SetVarsConfigs          SetVarsConfigs          `json:"SetVarsConfigs" xml:"SetVarsConfigs"`
}
