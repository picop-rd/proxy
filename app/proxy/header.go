package proxy

import "go.opentelemetry.io/otel/baggage"

const EnvIDKey = "env-id"

func GetEnvIDFromBaggage(v string) (string, error) {
	bag, err := baggage.Parse(v)
	if err != nil {
		return "", err
	}
	return bag.Member(EnvIDKey).Value(), nil
}
