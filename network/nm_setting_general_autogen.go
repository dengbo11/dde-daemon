// This file is automatically generated, please don't edit manully.
package main

func generalIsKeyInSettingField(field, key string) bool {
	if isVirtualKey(field, key) {
		return true
	}
	switch field {
	default:
		logger.Warning("invalid field name", field)
	case NM_SETTING_802_1X_SETTING_NAME:
		return isKeyInSetting8021x(key)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		return isKeyInSettingConnection(key)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		return isKeyInSettingIp4Config(key)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		return isKeyInSettingIp6Config(key)
	case NM_SETTING_PPP_SETTING_NAME:
		return isKeyInSettingPpp(key)
	case NM_SETTING_PPPOE_SETTING_NAME:
		return isKeyInSettingPppoe(key)
	case NM_SETTING_VPN_SETTING_NAME:
		return isKeyInSettingVpn(key)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		return isKeyInSettingVpnL2tp(key)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		return isKeyInSettingVpnL2tpPpp(key)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		return isKeyInSettingVpnL2tpIpsec(key)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		return isKeyInSettingVpnOpenconnect(key)
	case NM_SETTING_VF_VPN_OPENVPN_SETTING_NAME:
		return isKeyInSettingVpnOpenvpn(key)
	case NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME:
		return isKeyInSettingVpnOpenvpnAdvanced(key)
	case NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME:
		return isKeyInSettingVpnOpenvpnSecurity(key)
	case NM_SETTING_VF_VPN_OPENVPN_TLSAUTH_SETTING_NAME:
		return isKeyInSettingVpnOpenvpnTlsauth(key)
	case NM_SETTING_VF_VPN_OPENVPN_PROXIES_SETTING_NAME:
		return isKeyInSettingVpnOpenvpnProxies(key)
	case NM_SETTING_VF_VPN_PPTP_SETTING_NAME:
		return isKeyInSettingVpnPptp(key)
	case NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME:
		return isKeyInSettingVpnPptpPpp(key)
	case NM_SETTING_VF_VPN_VPNC_SETTING_NAME:
		return isKeyInSettingVpnVpnc(key)
	case NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME:
		return isKeyInSettingVpnVpncAdvanced(key)
	case NM_SETTING_WIRED_SETTING_NAME:
		return isKeyInSettingWired(key)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		return isKeyInSettingWireless(key)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		return isKeyInSettingWirelessSecurity(key)
	}
	return false
}

func generalGetSettingKeyType(field, key string) (t ktype) {
	if isVirtualKey(field, key) {
		t = getSettingVkKeyType(field, key)
		return
	}
	switch field {
	default:
		logger.Warning("invalid field name", field)
	case NM_SETTING_802_1X_SETTING_NAME:
		t = getSetting8021xKeyType(key)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		t = getSettingConnectionKeyType(key)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		t = getSettingIp4ConfigKeyType(key)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		t = getSettingIp6ConfigKeyType(key)
	case NM_SETTING_PPP_SETTING_NAME:
		t = getSettingPppKeyType(key)
	case NM_SETTING_PPPOE_SETTING_NAME:
		t = getSettingPppoeKeyType(key)
	case NM_SETTING_VPN_SETTING_NAME:
		t = getSettingVpnKeyType(key)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		t = getSettingVpnL2tpKeyType(key)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		t = getSettingVpnL2tpPppKeyType(key)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		t = getSettingVpnL2tpIpsecKeyType(key)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		t = getSettingVpnOpenconnectKeyType(key)
	case NM_SETTING_VF_VPN_OPENVPN_SETTING_NAME:
		t = getSettingVpnOpenvpnKeyType(key)
	case NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME:
		t = getSettingVpnOpenvpnAdvancedKeyType(key)
	case NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME:
		t = getSettingVpnOpenvpnSecurityKeyType(key)
	case NM_SETTING_VF_VPN_OPENVPN_TLSAUTH_SETTING_NAME:
		t = getSettingVpnOpenvpnTlsauthKeyType(key)
	case NM_SETTING_VF_VPN_OPENVPN_PROXIES_SETTING_NAME:
		t = getSettingVpnOpenvpnProxiesKeyType(key)
	case NM_SETTING_VF_VPN_PPTP_SETTING_NAME:
		t = getSettingVpnPptpKeyType(key)
	case NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME:
		t = getSettingVpnPptpPppKeyType(key)
	case NM_SETTING_VF_VPN_VPNC_SETTING_NAME:
		t = getSettingVpnVpncKeyType(key)
	case NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME:
		t = getSettingVpnVpncAdvancedKeyType(key)
	case NM_SETTING_WIRED_SETTING_NAME:
		t = getSettingWiredKeyType(key)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		t = getSettingWirelessKeyType(key)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		t = getSettingWirelessSecurityKeyType(key)
	}
	return
}

func generalGetSettingAvailableKeys(data connectionData, field string) (keys []string) {
	switch field {
	case NM_SETTING_802_1X_SETTING_NAME:
		keys = getSetting8021xAvailableKeys(data)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		keys = getSettingConnectionAvailableKeys(data)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		keys = getSettingIp4ConfigAvailableKeys(data)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		keys = getSettingIp6ConfigAvailableKeys(data)
	case NM_SETTING_PPP_SETTING_NAME:
		keys = getSettingPppAvailableKeys(data)
	case NM_SETTING_PPPOE_SETTING_NAME:
		keys = getSettingPppoeAvailableKeys(data)
	case NM_SETTING_VPN_SETTING_NAME:
		keys = getSettingVpnAvailableKeys(data)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		keys = getSettingVpnL2tpAvailableKeys(data)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		keys = getSettingVpnL2tpPppAvailableKeys(data)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		keys = getSettingVpnL2tpIpsecAvailableKeys(data)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		keys = getSettingVpnOpenconnectAvailableKeys(data)
	case NM_SETTING_VF_VPN_OPENVPN_SETTING_NAME:
		keys = getSettingVpnOpenvpnAvailableKeys(data)
	case NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME:
		keys = getSettingVpnOpenvpnAdvancedAvailableKeys(data)
	case NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME:
		keys = getSettingVpnOpenvpnSecurityAvailableKeys(data)
	case NM_SETTING_VF_VPN_OPENVPN_TLSAUTH_SETTING_NAME:
		keys = getSettingVpnOpenvpnTlsauthAvailableKeys(data)
	case NM_SETTING_VF_VPN_OPENVPN_PROXIES_SETTING_NAME:
		keys = getSettingVpnOpenvpnProxiesAvailableKeys(data)
	case NM_SETTING_VF_VPN_PPTP_SETTING_NAME:
		keys = getSettingVpnPptpAvailableKeys(data)
	case NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME:
		keys = getSettingVpnPptpPppAvailableKeys(data)
	case NM_SETTING_VF_VPN_VPNC_SETTING_NAME:
		keys = getSettingVpnVpncAvailableKeys(data)
	case NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME:
		keys = getSettingVpnVpncAdvancedAvailableKeys(data)
	case NM_SETTING_WIRED_SETTING_NAME:
		keys = getSettingWiredAvailableKeys(data)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		keys = getSettingWirelessAvailableKeys(data)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		keys = getSettingWirelessSecurityAvailableKeys(data)
	}
	return
}

func generalGetSettingAvailableValues(data connectionData, field, key string) (values []kvalue) {
	if isVirtualKey(field, key) {
		values = generalGetSettingVkAvailableValues(data, field, key)
		return
	}
	switch field {
	case NM_SETTING_802_1X_SETTING_NAME:
		values = getSetting8021xAvailableValues(data, key)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		values = getSettingConnectionAvailableValues(data, key)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		values = getSettingIp4ConfigAvailableValues(data, key)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		values = getSettingIp6ConfigAvailableValues(data, key)
	case NM_SETTING_PPP_SETTING_NAME:
		values = getSettingPppAvailableValues(data, key)
	case NM_SETTING_PPPOE_SETTING_NAME:
		values = getSettingPppoeAvailableValues(data, key)
	case NM_SETTING_VPN_SETTING_NAME:
		values = getSettingVpnAvailableValues(data, key)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		values = getSettingVpnL2tpAvailableValues(data, key)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		values = getSettingVpnL2tpPppAvailableValues(data, key)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		values = getSettingVpnL2tpIpsecAvailableValues(data, key)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		values = getSettingVpnOpenconnectAvailableValues(data, key)
	case NM_SETTING_VF_VPN_OPENVPN_SETTING_NAME:
		values = getSettingVpnOpenvpnAvailableValues(data, key)
	case NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME:
		values = getSettingVpnOpenvpnAdvancedAvailableValues(data, key)
	case NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME:
		values = getSettingVpnOpenvpnSecurityAvailableValues(data, key)
	case NM_SETTING_VF_VPN_OPENVPN_TLSAUTH_SETTING_NAME:
		values = getSettingVpnOpenvpnTlsauthAvailableValues(data, key)
	case NM_SETTING_VF_VPN_OPENVPN_PROXIES_SETTING_NAME:
		values = getSettingVpnOpenvpnProxiesAvailableValues(data, key)
	case NM_SETTING_VF_VPN_PPTP_SETTING_NAME:
		values = getSettingVpnPptpAvailableValues(data, key)
	case NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME:
		values = getSettingVpnPptpPppAvailableValues(data, key)
	case NM_SETTING_VF_VPN_VPNC_SETTING_NAME:
		values = getSettingVpnVpncAvailableValues(data, key)
	case NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME:
		values = getSettingVpnVpncAdvancedAvailableValues(data, key)
	case NM_SETTING_WIRED_SETTING_NAME:
		values = getSettingWiredAvailableValues(data, key)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		values = getSettingWirelessAvailableValues(data, key)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		values = getSettingWirelessSecurityAvailableValues(data, key)
	}
	return
}

func generalCheckSettingValues(data connectionData, field string) (errs fieldErrors) {
	switch field {
	default:
		logger.Error("invalid field name", field)
	case NM_SETTING_802_1X_SETTING_NAME:
		errs = checkSetting8021xValues(data)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		errs = checkSettingConnectionValues(data)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		errs = checkSettingIp4ConfigValues(data)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		errs = checkSettingIp6ConfigValues(data)
	case NM_SETTING_PPP_SETTING_NAME:
		errs = checkSettingPppValues(data)
	case NM_SETTING_PPPOE_SETTING_NAME:
		errs = checkSettingPppoeValues(data)
	case NM_SETTING_VPN_SETTING_NAME:
		errs = checkSettingVpnValues(data)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		errs = checkSettingVpnL2tpValues(data)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		errs = checkSettingVpnL2tpPppValues(data)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		errs = checkSettingVpnL2tpIpsecValues(data)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		errs = checkSettingVpnOpenconnectValues(data)
	case NM_SETTING_VF_VPN_OPENVPN_SETTING_NAME:
		errs = checkSettingVpnOpenvpnValues(data)
	case NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME:
		errs = checkSettingVpnOpenvpnAdvancedValues(data)
	case NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME:
		errs = checkSettingVpnOpenvpnSecurityValues(data)
	case NM_SETTING_VF_VPN_OPENVPN_TLSAUTH_SETTING_NAME:
		errs = checkSettingVpnOpenvpnTlsauthValues(data)
	case NM_SETTING_VF_VPN_OPENVPN_PROXIES_SETTING_NAME:
		errs = checkSettingVpnOpenvpnProxiesValues(data)
	case NM_SETTING_VF_VPN_PPTP_SETTING_NAME:
		errs = checkSettingVpnPptpValues(data)
	case NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME:
		errs = checkSettingVpnPptpPppValues(data)
	case NM_SETTING_VF_VPN_VPNC_SETTING_NAME:
		errs = checkSettingVpnVpncValues(data)
	case NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME:
		errs = checkSettingVpnVpncAdvancedValues(data)
	case NM_SETTING_WIRED_SETTING_NAME:
		errs = checkSettingWiredValues(data)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		errs = checkSettingWirelessValues(data)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		errs = checkSettingWirelessSecurityValues(data)
	}
	return
}

func generalGetSettingKeyJSON(data connectionData, field, key string) (valueJSON string) {
	if isVirtualKey(field, key) {
		valueJSON = generalGetVirtualKeyJSON(data, field, key)
		return
	}
	switch field {
	default:
		logger.Warning("invalid field name", field)
	case NM_SETTING_802_1X_SETTING_NAME:
		valueJSON = generalGetSetting8021xKeyJSON(data, key)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		valueJSON = generalGetSettingConnectionKeyJSON(data, key)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		valueJSON = generalGetSettingIp4ConfigKeyJSON(data, key)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		valueJSON = generalGetSettingIp6ConfigKeyJSON(data, key)
	case NM_SETTING_PPP_SETTING_NAME:
		valueJSON = generalGetSettingPppKeyJSON(data, key)
	case NM_SETTING_PPPOE_SETTING_NAME:
		valueJSON = generalGetSettingPppoeKeyJSON(data, key)
	case NM_SETTING_VPN_SETTING_NAME:
		valueJSON = generalGetSettingVpnKeyJSON(data, key)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		valueJSON = generalGetSettingVpnL2tpKeyJSON(data, key)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		valueJSON = generalGetSettingVpnL2tpPppKeyJSON(data, key)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		valueJSON = generalGetSettingVpnL2tpIpsecKeyJSON(data, key)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		valueJSON = generalGetSettingVpnOpenconnectKeyJSON(data, key)
	case NM_SETTING_VF_VPN_OPENVPN_SETTING_NAME:
		valueJSON = generalGetSettingVpnOpenvpnKeyJSON(data, key)
	case NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME:
		valueJSON = generalGetSettingVpnOpenvpnAdvancedKeyJSON(data, key)
	case NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME:
		valueJSON = generalGetSettingVpnOpenvpnSecurityKeyJSON(data, key)
	case NM_SETTING_VF_VPN_OPENVPN_TLSAUTH_SETTING_NAME:
		valueJSON = generalGetSettingVpnOpenvpnTlsauthKeyJSON(data, key)
	case NM_SETTING_VF_VPN_OPENVPN_PROXIES_SETTING_NAME:
		valueJSON = generalGetSettingVpnOpenvpnProxiesKeyJSON(data, key)
	case NM_SETTING_VF_VPN_PPTP_SETTING_NAME:
		valueJSON = generalGetSettingVpnPptpKeyJSON(data, key)
	case NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME:
		valueJSON = generalGetSettingVpnPptpPppKeyJSON(data, key)
	case NM_SETTING_VF_VPN_VPNC_SETTING_NAME:
		valueJSON = generalGetSettingVpnVpncKeyJSON(data, key)
	case NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME:
		valueJSON = generalGetSettingVpnVpncAdvancedKeyJSON(data, key)
	case NM_SETTING_WIRED_SETTING_NAME:
		valueJSON = generalGetSettingWiredKeyJSON(data, key)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		valueJSON = generalGetSettingWirelessKeyJSON(data, key)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		valueJSON = generalGetSettingWirelessSecurityKeyJSON(data, key)
	}
	return
}

func generalSetSettingKeyJSON(data connectionData, field, key, valueJSON string) (ok bool, errMsg string) {
	ok = true
	if isVirtualKey(field, key) {
		ok, errMsg = generalSetVirtualKeyJSON(data, field, key, valueJSON)
		return
	}
	switch field {
	default:
		logger.Warning("invalid field name", field)
	case NM_SETTING_802_1X_SETTING_NAME:
		ok, errMsg = generalSetSetting8021xKeyJSON(data, key, valueJSON)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		ok, errMsg = generalSetSettingConnectionKeyJSON(data, key, valueJSON)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		ok, errMsg = generalSetSettingIp4ConfigKeyJSON(data, key, valueJSON)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		ok, errMsg = generalSetSettingIp6ConfigKeyJSON(data, key, valueJSON)
	case NM_SETTING_PPP_SETTING_NAME:
		ok, errMsg = generalSetSettingPppKeyJSON(data, key, valueJSON)
	case NM_SETTING_PPPOE_SETTING_NAME:
		ok, errMsg = generalSetSettingPppoeKeyJSON(data, key, valueJSON)
	case NM_SETTING_VPN_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnL2tpKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnL2tpPppKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnL2tpIpsecKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnOpenconnectKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_OPENVPN_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnOpenvpnKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnOpenvpnAdvancedKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnOpenvpnSecurityKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_OPENVPN_TLSAUTH_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnOpenvpnTlsauthKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_OPENVPN_PROXIES_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnOpenvpnProxiesKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_PPTP_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnPptpKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnPptpPppKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_VPNC_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnVpncKeyJSON(data, key, valueJSON)
	case NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME:
		ok, errMsg = generalSetSettingVpnVpncAdvancedKeyJSON(data, key, valueJSON)
	case NM_SETTING_WIRED_SETTING_NAME:
		ok, errMsg = generalSetSettingWiredKeyJSON(data, key, valueJSON)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		ok, errMsg = generalSetSettingWirelessKeyJSON(data, key, valueJSON)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		ok, errMsg = generalSetSettingWirelessSecurityKeyJSON(data, key, valueJSON)
	}
	return
}

func getSettingKeyDefaultValueJSON(field, key string) (valueJSON string) {
	switch field {
	default:
		logger.Warning("invalid field name", field)
	case NM_SETTING_802_1X_SETTING_NAME:
		valueJSON = getSetting8021xKeyDefaultValueJSON(key)
	case NM_SETTING_CONNECTION_SETTING_NAME:
		valueJSON = getSettingConnectionKeyDefaultValueJSON(key)
	case NM_SETTING_IP4_CONFIG_SETTING_NAME:
		valueJSON = getSettingIp4ConfigKeyDefaultValueJSON(key)
	case NM_SETTING_IP6_CONFIG_SETTING_NAME:
		valueJSON = getSettingIp6ConfigKeyDefaultValueJSON(key)
	case NM_SETTING_PPP_SETTING_NAME:
		valueJSON = getSettingPppKeyDefaultValueJSON(key)
	case NM_SETTING_PPPOE_SETTING_NAME:
		valueJSON = getSettingPppoeKeyDefaultValueJSON(key)
	case NM_SETTING_VPN_SETTING_NAME:
		valueJSON = getSettingVpnKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_L2TP_SETTING_NAME:
		valueJSON = getSettingVpnL2tpKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_L2TP_PPP_SETTING_NAME:
		valueJSON = getSettingVpnL2tpPppKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_L2TP_IPSEC_SETTING_NAME:
		valueJSON = getSettingVpnL2tpIpsecKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME:
		valueJSON = getSettingVpnOpenconnectKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_OPENVPN_SETTING_NAME:
		valueJSON = getSettingVpnOpenvpnKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_OPENVPN_ADVANCED_SETTING_NAME:
		valueJSON = getSettingVpnOpenvpnAdvancedKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_OPENVPN_SECURITY_SETTING_NAME:
		valueJSON = getSettingVpnOpenvpnSecurityKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_OPENVPN_TLSAUTH_SETTING_NAME:
		valueJSON = getSettingVpnOpenvpnTlsauthKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_OPENVPN_PROXIES_SETTING_NAME:
		valueJSON = getSettingVpnOpenvpnProxiesKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_PPTP_SETTING_NAME:
		valueJSON = getSettingVpnPptpKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_PPTP_PPP_SETTING_NAME:
		valueJSON = getSettingVpnPptpPppKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_VPNC_SETTING_NAME:
		valueJSON = getSettingVpnVpncKeyDefaultValueJSON(key)
	case NM_SETTING_VF_VPN_VPNC_ADVANCED_SETTING_NAME:
		valueJSON = getSettingVpnVpncAdvancedKeyDefaultValueJSON(key)
	case NM_SETTING_WIRED_SETTING_NAME:
		valueJSON = getSettingWiredKeyDefaultValueJSON(key)
	case NM_SETTING_WIRELESS_SETTING_NAME:
		valueJSON = getSettingWirelessKeyDefaultValueJSON(key)
	case NM_SETTING_WIRELESS_SECURITY_SETTING_NAME:
		valueJSON = getSettingWirelessSecurityKeyDefaultValueJSON(key)
	}
	return
}
