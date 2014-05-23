// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingSerialKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_SERIAL_BAUD:
		t = ktypeUint32
	case NM_SETTING_SERIAL_BITS:
		t = ktypeUint32
	case NM_SETTING_SERIAL_PARITY:
		t = ktypeByte
	case NM_SETTING_SERIAL_STOPBITS:
		t = ktypeUint32
	case NM_SETTING_SERIAL_SEND_DELAY:
		t = ktypeUint64
	}
	return
}

// Check is key in current setting field
func isKeyInSettingSerial(key string) bool {
	switch key {
	case NM_SETTING_SERIAL_BAUD:
		return true
	case NM_SETTING_SERIAL_BITS:
		return true
	case NM_SETTING_SERIAL_PARITY:
		return true
	case NM_SETTING_SERIAL_STOPBITS:
		return true
	case NM_SETTING_SERIAL_SEND_DELAY:
		return true
	}
	return false
}

// Get key's default value
func getSettingSerialDefaultValue(key string) (value interface{}) {
	switch key {
	default:
		logger.Error("invalid key:", key)
	case NM_SETTING_SERIAL_BAUD:
		value = 57600
	case NM_SETTING_SERIAL_BITS:
		value = 8
	case NM_SETTING_SERIAL_PARITY:
		value = 110
	case NM_SETTING_SERIAL_STOPBITS:
		value = 1
	case NM_SETTING_SERIAL_SEND_DELAY:
		value = 0
	}
	return
}

// Get JSON value generally
func generalGetSettingSerialKeyJSON(data connectionData, key string) (value string) {
	switch key {
	default:
		logger.Error("generalGetSettingSerialKeyJSON: invalide key", key)
	case NM_SETTING_SERIAL_BAUD:
		value = getSettingSerialBaudJSON(data)
	case NM_SETTING_SERIAL_BITS:
		value = getSettingSerialBitsJSON(data)
	case NM_SETTING_SERIAL_PARITY:
		value = getSettingSerialParityJSON(data)
	case NM_SETTING_SERIAL_STOPBITS:
		value = getSettingSerialStopbitsJSON(data)
	case NM_SETTING_SERIAL_SEND_DELAY:
		value = getSettingSerialSendDelayJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingSerialKeyJSON(data connectionData, key, valueJSON string) (err error) {
	switch key {
	default:
		logger.Error("generalSetSettingSerialKeyJSON: invalide key", key)
	case NM_SETTING_SERIAL_BAUD:
		err = setSettingSerialBaudJSON(data, valueJSON)
	case NM_SETTING_SERIAL_BITS:
		err = setSettingSerialBitsJSON(data, valueJSON)
	case NM_SETTING_SERIAL_PARITY:
		err = setSettingSerialParityJSON(data, valueJSON)
	case NM_SETTING_SERIAL_STOPBITS:
		err = setSettingSerialStopbitsJSON(data, valueJSON)
	case NM_SETTING_SERIAL_SEND_DELAY:
		err = setSettingSerialSendDelayJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingSerialBaudExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BAUD)
}
func isSettingSerialBitsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BITS)
}
func isSettingSerialParityExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_PARITY)
}
func isSettingSerialStopbitsExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_STOPBITS)
}
func isSettingSerialSendDelayExists(data connectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_SEND_DELAY)
}

// Ensure field and key exists and not empty
func ensureFieldSettingSerialExists(data connectionData, errs fieldErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_SERIAL_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_SERIAL_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_SERIAL_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_SERIAL_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_SERIAL_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_SERIAL_SETTING_NAME))
	}
}
func ensureSettingSerialBaudNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingSerialBaudExists(data) {
		rememberError(errs, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BAUD, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingSerialBitsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingSerialBitsExists(data) {
		rememberError(errs, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BITS, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingSerialParityNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingSerialParityExists(data) {
		rememberError(errs, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_PARITY, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingSerialStopbitsNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingSerialStopbitsExists(data) {
		rememberError(errs, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_STOPBITS, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingSerialSendDelayNoEmpty(data connectionData, errs fieldErrors) {
	if !isSettingSerialSendDelayExists(data) {
		rememberError(errs, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_SEND_DELAY, NM_KEY_ERROR_MISSING_VALUE)
	}
}

// Getter
func getSettingSerialBaud(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BAUD).(uint32)
	return
}
func getSettingSerialBits(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BITS).(uint32)
	return
}
func getSettingSerialParity(data connectionData) (value byte) {
	value, _ = getSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_PARITY).(byte)
	return
}
func getSettingSerialStopbits(data connectionData) (value uint32) {
	value, _ = getSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_STOPBITS).(uint32)
	return
}
func getSettingSerialSendDelay(data connectionData) (value uint64) {
	value, _ = getSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_SEND_DELAY).(uint64)
	return
}

// Setter
func setSettingSerialBaud(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BAUD, value)
}
func setSettingSerialBits(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BITS, value)
}
func setSettingSerialParity(data connectionData, value byte) {
	setSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_PARITY, value)
}
func setSettingSerialStopbits(data connectionData, value uint32) {
	setSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_STOPBITS, value)
}
func setSettingSerialSendDelay(data connectionData, value uint64) {
	setSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_SEND_DELAY, value)
}

// JSON Getter
func getSettingSerialBaudJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BAUD, getSettingSerialKeyType(NM_SETTING_SERIAL_BAUD))
	return
}
func getSettingSerialBitsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BITS, getSettingSerialKeyType(NM_SETTING_SERIAL_BITS))
	return
}
func getSettingSerialParityJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_PARITY, getSettingSerialKeyType(NM_SETTING_SERIAL_PARITY))
	return
}
func getSettingSerialStopbitsJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_STOPBITS, getSettingSerialKeyType(NM_SETTING_SERIAL_STOPBITS))
	return
}
func getSettingSerialSendDelayJSON(data connectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_SEND_DELAY, getSettingSerialKeyType(NM_SETTING_SERIAL_SEND_DELAY))
	return
}

// JSON Setter
func setSettingSerialBaudJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BAUD, valueJSON, getSettingSerialKeyType(NM_SETTING_SERIAL_BAUD))
}
func setSettingSerialBitsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BITS, valueJSON, getSettingSerialKeyType(NM_SETTING_SERIAL_BITS))
}
func setSettingSerialParityJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_PARITY, valueJSON, getSettingSerialKeyType(NM_SETTING_SERIAL_PARITY))
}
func setSettingSerialStopbitsJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_STOPBITS, valueJSON, getSettingSerialKeyType(NM_SETTING_SERIAL_STOPBITS))
}
func setSettingSerialSendDelayJSON(data connectionData, valueJSON string) (err error) {
	return setSettingKeyJSON(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_SEND_DELAY, valueJSON, getSettingSerialKeyType(NM_SETTING_SERIAL_SEND_DELAY))
}

// Logic JSON Setter

// Remover
func removeSettingSerialBaud(data connectionData) {
	removeSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BAUD)
}
func removeSettingSerialBits(data connectionData) {
	removeSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_BITS)
}
func removeSettingSerialParity(data connectionData) {
	removeSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_PARITY)
}
func removeSettingSerialStopbits(data connectionData) {
	removeSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_STOPBITS)
}
func removeSettingSerialSendDelay(data connectionData) {
	removeSettingKey(data, NM_SETTING_SERIAL_SETTING_NAME, NM_SETTING_SERIAL_SEND_DELAY)
}
