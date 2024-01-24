/*
Copyright © 2021 ax-i-om <addressaxiom@pm.me>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package email contains the types and functions used for querying for
// information regarding an email address
package email

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/url"
	"strings"

	"github.com/ax-i-om/bitcrook/internal/http"
)

/* ///////////////////////////////////////////////////////

   MELISSA

*/ ///////////////////////////////////////////////////////

type emailResponse struct {
	Version               string `json:"Version"`
	Transmissionreference string `json:"TransmissionReference"`
	Transmissionresults   string `json:"TransmissionResults"`
	Totalrecords          string `json:"TotalRecords"`
	Records               []struct {
		Recordid                      string `json:"RecordID"`
		Deliverabilityconfidencescore string `json:"DeliverabilityConfidenceScore"`
		Results                       string `json:"Results"`
		Emailaddress                  string `json:"EmailAddress"`
		Mailboxname                   string `json:"MailboxName"`
		Domainname                    string `json:"DomainName"`
		Topleveldomain                string `json:"TopLevelDomain"`
		Topleveldomainname            string `json:"TopLevelDomainName"`
		Datechecked                   string `json:"DateChecked"`
		Emailageestimated             string `json:"EmailAgeEstimated"`
		Domainageestimated            string `json:"DomainAgeEstimated"`
		Domainexpirationdate          string `json:"DomainExpirationDate"`
		Domaincreateddate             string `json:"DomainCreatedDate"`
		Domainupdateddate             string `json:"DomainUpdatedDate"`
		Domainemail                   string `json:"DomainEmail"`
		Domainorganization            string `json:"DomainOrganization"`
		Domainaddress1                string `json:"DomainAddress1"`
		Domainlocality                string `json:"DomainLocality"`
		Domainadministrativearea      string `json:"DomainAdministrativeArea"`
		Domainpostalcode              string `json:"DomainPostalCode"`
		Domaincountry                 string `json:"DomainCountry"`
		Domainavailability            string `json:"DomainAvailability"`
		Domaincountrycode             string `json:"DomainCountryCode"`
		Domainprivateproxy            string `json:"DomainPrivateProxy"`
		Privacyflag                   string `json:"PrivacyFlag"`
		Mxserver                      string `json:"MXServer"`
	} `json:"Records"`
}

// MelissaEmail is a type that represents the results of Melissa's GlobalEmail API.
type MelissaEmail struct {
	Recordid                      string
	Deliverabilityconfidencescore string
	Results                       string
	Emailaddress                  string
	Mailboxname                   string
	Domainname                    string
	Topleveldomain                string
	Topleveldomainname            string
	Datechecked                   string
	Emailageestimated             string
	Domainageestimated            string
	Domainexpirationdate          string
	Domaincreateddate             string
	Domainupdateddate             string
	Domainemail                   string
	Domainorganization            string
	Domainaddress1                string
	Domainlocality                string
	Domainadministrativearea      string
	Domainpostalcode              string
	Domaincountry                 string
	Domainavailability            string
	Domaincountrycode             string
	Domainprivateproxy            string
	Privacyflag                   string
	Mxserver                      string
}

// MelissaLookup takes a Melissa API key and an email as its parameters which are passed through the Melissa Global Email API whose response is then represented by a *MelissaEmail type.
func MelissaLookup(key, email string) (*MelissaEmail, error) {
	resp, err := http.GetReq("https://globalemail.melissadata.net/v4/WEB/GlobalEmail/doGlobalEmail?id=" + key + "&email=" + email + "&format=json")
	if err != nil {
		return nil, err
	}
	response := new(emailResponse)
	err = json.Unmarshal([]byte(resp), &response)
	if err != nil {
		return nil, err
	}
	newRes := new(MelissaEmail)
	for _, v := range response.Records {
		newRes.Recordid = v.Recordid
		newRes.Deliverabilityconfidencescore = v.Deliverabilityconfidencescore
		newRes.Results = v.Results
		newRes.Emailaddress = v.Emailaddress
		newRes.Mailboxname = v.Mailboxname
		newRes.Domainname = v.Domainname
		newRes.Topleveldomain = v.Topleveldomain
		newRes.Topleveldomainname = v.Topleveldomainname
		newRes.Datechecked = v.Datechecked
		newRes.Emailageestimated = v.Emailageestimated
		newRes.Domainageestimated = v.Domainageestimated
		newRes.Domainexpirationdate = v.Domainexpirationdate
		newRes.Domaincreateddate = v.Domaincreateddate
		newRes.Domainupdateddate = v.Domainupdateddate
		newRes.Domainemail = v.Domainemail
		newRes.Domainorganization = v.Domainorganization
		newRes.Domainaddress1 = v.Domainaddress1
		newRes.Domainlocality = v.Domainlocality
		newRes.Domainadministrativearea = v.Domainadministrativearea
		newRes.Domainpostalcode = v.Domainpostalcode
		newRes.Domaincountry = v.Domaincountry
		newRes.Domainavailability = v.Domainavailability
		newRes.Domaincountrycode = v.Domaincountrycode
		newRes.Domainprivateproxy = v.Domainprivateproxy
		newRes.Privacyflag = v.Privacyflag
		newRes.Mxserver = v.Mxserver

		break
	}
	return newRes, nil
}

/* ///////////////////////////////////////////////////////

   HaveIBeenPwned

*/ ///////////////////////////////////////////////////////

type breachList []struct {
	Name string `json:"Name"`
}

// AccBreaches is a representation of the names of site breaches returned by the HIBPLookup() function.
type AccBreaches struct {
	Name string `json:"Name"`
}

// HIBPLookup takes a Have I Been Pwned API key and an email as its parameters and returns a list of related breaches (if any).
func HIBPLookup(key, email string) ([]AccBreaches, error) {
	headers := []string{"hibp-api-key", key}
	body, err := http.CustomGet("https://haveibeenpwned.com/api/v3/breachedaccount/"+email, headers)
	if err != nil {
		return nil, err
	}
	response := new(breachList)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, err
	}

	var results []AccBreaches
	for _, v := range *response {
		results = append(results, v)
	}
	return results, nil
}

/* ///////////////////////////////////////////////////////

   Pinterest

*/ ///////////////////////////////////////////////////////

// PinterestResponse is a representation of the information returned by the Pinterest API.
type PinterestResponse struct {
	ResourceResponse struct {
		Status                    string `json:"status"`
		Code                      int    `json:"code"`
		Message                   string `json:"message"`
		EndpointName              string `json:"endpoint_name"`
		Data                      bool   `json:"data"`
		XPinterestSliEndpointName string `json:"x_pinterest_sli_endpoint_name"`
		HTTPStatus                int    `json:"http_status"`
	} `json:"resource_response"`
	ClientContext struct {
		AnalysisUa struct {
			AppType        int    `json:"app_type"`
			BrowserName    string `json:"browser_name"`
			BrowserVersion string `json:"browser_version"`
			DeviceType     any    `json:"device_type"`
			Device         string `json:"device"`
			OsName         string `json:"os_name"`
			OsVersion      string `json:"os_version"`
		} `json:"analysis_ua"`
		AppTypeDetailed            int      `json:"app_type_detailed"`
		AppVersion                 string   `json:"app_version"`
		BatchExp                   bool     `json:"batch_exp"`
		BrowserLocale              string   `json:"browser_locale"`
		BrowserName                string   `json:"browser_name"`
		BrowserType                any      `json:"browser_type"`
		BrowserVersion             string   `json:"browser_version"`
		Country                    string   `json:"country"`
		CountryFromHostname        string   `json:"country_from_hostname"`
		CountryFromIP              string   `json:"country_from_ip"`
		CspNonce                   string   `json:"csp_nonce"`
		CurrentURL                 string   `json:"current_url"`
		Debug                      bool     `json:"debug"`
		DeepLink                   string   `json:"deep_link"`
		EnabledAdvertiserCountries []string `json:"enabled_advertiser_countries"`
		FacebookToken              any      `json:"facebook_token"`
		FullPath                   string   `json:"full_path"`
		HTTPReferrer               string   `json:"http_referrer"`
		ImpersonatorUserID         any      `json:"impersonator_user_id"`
		InviteCode                 string   `json:"invite_code"`
		InviteSenderID             string   `json:"invite_sender_id"`
		IsAuthenticated            bool     `json:"is_authenticated"`
		IsBot                      string   `json:"is_bot"`
		IsInternalIP               bool     `json:"is_internal_ip"`
		IsFullPage                 bool     `json:"is_full_page"`
		IsMobileAgent              bool     `json:"is_mobile_agent"`
		IsSterlingOnSteroids       bool     `json:"is_sterling_on_steroids"`
		IsTabletAgent              bool     `json:"is_tablet_agent"`
		Language                   string   `json:"language"`
		Locale                     string   `json:"locale"`
		Origin                     string   `json:"origin"`
		Path                       string   `json:"path"`
		PlacedExperiences          any      `json:"placed_experiences"`
		Referrer                   any      `json:"referrer"`
		RegionFromIP               string   `json:"region_from_ip"`
		RequestHost                string   `json:"request_host"`
		RequestIdentifier          string   `json:"request_identifier"`
		SocialBot                  string   `json:"social_bot"`
		Stage                      string   `json:"stage"`
		SterlingOnSteroidsLdap     any      `json:"sterling_on_steroids_ldap"`
		SterlingOnSteroidsUserType any      `json:"sterling_on_steroids_user_type"`
		UnauthID                   string   `json:"unauth_id"`
		SeoDebug                   bool     `json:"seo_debug"`
		UserAgentCanUseNativeApp   bool     `json:"user_agent_can_use_native_app"`
		UserAgentPlatform          string   `json:"user_agent_platform"`
		UserAgentPlatformVersion   any      `json:"user_agent_platform_version"`
		UserAgent                  string   `json:"user_agent"`
		User                       struct {
			UnauthID  string `json:"unauth_id"`
			IPCountry string `json:"ip_country"`
			IPRegion  string `json:"ip_region"`
		} `json:"user"`
		UtmCampaign any    `json:"utm_campaign"`
		VisibleURL  string `json:"visible_url"`
	} `json:"client_context"`
	Resource struct {
		Name    string `json:"name"`
		Options struct {
			Bookmarks []string `json:"bookmarks"`
			Email     string   `json:"email"`
		} `json:"options"`
	} `json:"resource"`
	RequestIdentifier string `json:"request_identifier"`
}

// PinterestLookup takes an email as an argument and returns information about any existed Pinterest account association.
func PinterestLookup(email string) (*PinterestResponse, error) {
	body, err := http.GetReq(`https://www.pinterest.com/_ngjs/resource/EmailExistsResource/get/?source_url=%2F&data=%7B%0A%20%20%20%22options%22%3A%7B%0A%20%20%20%20%20%20%22email%22%3A%22` + url.QueryEscape(email) + `%22%0A%20%20%20%7D,%0A%20%20%20%22context%22%3A%7B%0A%20%20%20%20%20%20%0A%20%20%20%7D%0A%7D`)
	if err != nil {
		return nil, err
	}
	response := new(PinterestResponse)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

/* ///////////////////////////////////////////////////////

   Twitter

*/ ///////////////////////////////////////////////////////

// TwitterValidityResponse is a representation of the information returned by the TwitterLookup() function.
type TwitterValidityResponse struct {
	Valid bool   `json:"valid"`
	Msg   string `json:"msg"`
	Taken bool   `json:"taken"`
}

// TwitterValidityLookup takes an email as an argument and returns a type TwitterValidityResponse, indicating whether or not it is taken.
func TwitterValidityLookup(email string) (*TwitterValidityResponse, error) {
	body, err := http.GetReq("https://api.twitter.com/i/users/email_available.json?email=" + email)
	if err != nil {
		return nil, err
	}
	response := new(TwitterValidityResponse)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/* ///////////////////////////////////////////////////////

   Gravatar

*/ ///////////////////////////////////////////////////////

// GravatarResponse is a representation of the information returned by the GravatarLookup() function.
type GravatarResponse struct {
	Entry []struct {
		Hash              string `json:"hash"`
		RequestHash       string `json:"requestHash"`
		ProfileURL        string `json:"profileUrl"`
		PreferredUsername string `json:"preferredUsername"`
		ThumbnailURL      string `json:"thumbnailUrl"`
		Photos            []struct {
			Value string `json:"value"`
			Type  string `json:"type"`
		} `json:"photos"`
		LastProfileEdit   string `json:"last_profile_edit"`
		ProfileBackground struct {
			Color string `json:"color"`
			URL   string `json:"url"`
		} `json:"profileBackground"`
		Name struct {
			GivenName  string `json:"givenName"`
			FamilyName string `json:"familyName"`
			Formatted  string `json:"formatted"`
		} `json:"name"`
		DisplayName     string `json:"displayName"`
		AboutMe         string `json:"aboutMe"`
		CurrentLocation string `json:"currentLocation"`
		PhoneNumbers    []struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"phoneNumbers"`
		Emails []struct {
			Primary string `json:"primary"`
			Value   string `json:"value"`
		} `json:"emails"`
		Urls []struct {
			Value string `json:"value"`
			Title string `json:"title"`
		} `json:"urls"`
		ShareFlags struct {
			SearchEngines bool `json:"search_engines"`
		} `json:"share_flags"`
	} `json:"entry"`
}

// GravatarLookup takes an email as an argument and returns a type GravatarResponse, indicating whether or not it is taken.
func GravatarLookup(email string) (*GravatarResponse, error) {
	chkd := md5.Sum([]byte(email))
	body, err := http.GetReq("https://gravatar.com/" + hex.EncodeToString(chkd[:]) + ".json")
	if err != nil {
		return nil, err
	}
	response := new(GravatarResponse)
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

/* ///////////////////////////////////////////////////////

   Wordpress

*/ ///////////////////////////////////////////////////////

// WordpressLookup takes an email as an argument and returns a bool, indicating whether or not it is taken.
func WordpressLookup(email string) (bool, error) {
	body, err := http.GetReq("https://public-api.wordpress.com/rest/v1.1/users/" + email + "/auth-options")
	if err != nil {
		return false, err
	}

	if strings.Contains(body, "User does not exist.") {
		return false, nil
	}
	return true, nil
}
