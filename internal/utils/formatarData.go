package utils

import "time"

func FormatarData(dataStr string) (string, error) {
	// Remove as barras e converte para o formato desejado
	parsedTime, err := time.Parse("02/01/2006", dataStr)
	if err != nil {
		return "", err
	}
	return parsedTime.Format("20060102"), nil
}

func FormatarDataExibicao(dataStr string) (string, error) {
	layouts := []string{
		"02/01/2006", // BR: DD/MM/YYYY
		"01/02/2006", // US: MM/DD/YYYY
		"2006-01-02", // ISO: YYYY-MM-DD
		"2006/01/02", // YYYY/MM/DD
		time.RFC3339, // RFC3339
	}

	var lastErr error
	for _, layout := range layouts {
		if t, err := time.Parse(layout, dataStr); err == nil {
			return t.Format("02/01/2006"), nil
		} else {
			lastErr = err
		}
	}

	return "", lastErr
}
